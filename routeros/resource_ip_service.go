package routeros

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*0",
    "address": "",
    "disabled": "false",
    "invalid": "false",
    "name": "telnet",
    "port": "23",
    "vrf": "main"
  },
  {
    ".id": "*6",
    "address": "",
    "certificate": "https-cert",
    "disabled": "false",
    "invalid": "false",
    "name": "www-ssl",
    "port": "443",
    "tls-version": "any",
    "vrf": "main"
  },
*/

// In RouterOS 7.20+ `/ip service print` surfaces a dynamic, per-connection
// entry for each active session alongside each configurable service row.
// All such entries share the same `name` field (e.g. two `name="ssh"` rows:
// one static service config, one for the running connection). The legacy
// name-based identifier therefore matches multiple rows and breaks import
// + read. We work around this by:
//
//   - Using the `.id` (`*N`) handle as the canonical identifier internally
//     (MetaId: PropId(Id)).
//   - Translating between the user-facing service name (`numbers`) and the
//     internal `*N` via the helper below, always filtering `dynamic=false`.
//
// See: https://github.com/terraform-routeros/terraform-provider-routeros/issues/905
func resolveIpServiceId(name string, c Client) (string, error) {
	res, err := ReadItemsFiltered(
		[]string{"name=" + name, "dynamic=false"},
		"/ip/service", c,
	)
	if err != nil {
		return "", err
	}
	if len(*res) == 0 {
		return "", fmt.Errorf("ip service not found: name=%s", name)
	}
	if len(*res) > 1 {
		return "", fmt.Errorf("ip service ambiguous: name=%s matched %d non-dynamic rows", name, len(*res))
	}
	id, ok := (*res)[0][".id"]
	if !ok {
		return "", fmt.Errorf("ip service lookup for %q returned no .id", name)
	}
	return id, nil
}

// https://help.mikrotik.com/docs/display/ROS/Services
func ResourceIpService() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/service"),
		MetaId:           PropId(Id),

		"address": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "",
			Description: "List of IP/IPv6 prefixes from which the service is accessible.",
			DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
				if oldValue == "" && newValue == "0.0.0.0/0" {
					return false
				}
				return oldValue == newValue
			},
		},
		"certificate": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The name of the certificate used by a particular service. Applicable only for services " +
				"that depend on certificates ( www-ssl, api-ssl ).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyDisabled: PropDisabledRw,
		KeyDynamic:  PropDynamicRo,
		KeyInvalid:  PropInvalidRo,
		"max_sessions": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Maximum number of concurrent connections to a particular service. This option is available in RouterOS starting from version 7.16.",
			ValidateFunc:     validation.IntAtLeast(1),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Service name.",
		},
		"numbers": {
			Type:     schema.TypeString,
			Required: true,
			Description: "The name of the service whose settings will be changed ( api, api-ssl, ftp, ssh, telnet, " +
				"winbox, www, www-ssl ).",
			ValidateDiagFunc: ValidationMultiValInSlice([]string{"api", "api-ssl", "ftp", "ssh", "telnet", "winbox",
				"www", "www-ssl"}, false, false),
		},
		"port": {
			Type:         schema.TypeInt,
			Required:     true,
			Description:  "The port particular service listens on.",
			ValidateFunc: validation.IntBetween(1, 65535),
		},
		"proto": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"tls_version": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Specifies which TLS versions to allow by a particular service.",
			ValidateFunc:     validation.StringInSlice([]string{"any", "only-1.2"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyVrf: PropVrfRw,
	}

	resCreateUpdate := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		item, metadata := TerraformResourceDataToMikrotik(resSchema, d)

		var resUrl string
		if m.(Client).GetTransport() == TransportREST {
			// https://router/rest/system/identity/set
			// https://router/rest/caps-man/manager/set
			resUrl = "/set"
		}

		err := m.(Client).SendRequest(crudPost, &URL{Path: metadata.Path + resUrl}, item, nil)
		if err != nil {
			return diag.FromErr(err)
		}

		// After SET, resolve the static (non-dynamic) row's `.id` and store
		// it as the resource id. Subsequent Read/Update/Delete then operate
		// on `.id`, which is unique even when dynamic per-connection rows
		// duplicate the service `name`.
		id, err := resolveIpServiceId(d.Get("numbers").(string), m.(Client))
		if err != nil {
			return diag.FromErr(err)
		}
		d.SetId(id)

		return resourceIpServiceRead(ctx, resSchema, d, m)
	}

	return &schema.Resource{
		CreateContext: resCreateUpdate,
		ReadContext:   func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return resourceIpServiceRead(ctx, resSchema, d, m)
		},
		UpdateContext: resCreateUpdate,
		DeleteContext: DefaultSystemDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: importIpServiceState,
		},

		Schema: resSchema,
	}
}

// resourceIpServiceRead Read filtered to the static (non-dynamic) row only,
// keyed by the resource's `.id` set during create/import.
func resourceIpServiceRead(ctx context.Context, s map[string]*schema.Schema, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	res, err := ReadItemsFiltered(
		[]string{".id=" + d.Id(), "dynamic=false"},
		"/ip/service", m.(Client),
	)
	if err != nil {
		return diag.FromErr(err)
	}
	if len(*res) == 0 {
		// Resource gone; let Terraform recreate.
		d.SetId("")
		return nil
	}
	if id, ok := (*res)[0][".id"]; ok {
		d.SetId(id)
	}
	return MikrotikResourceDataToTerraform((*res)[0], s, d)
}

// importIpServiceState accepts either a `*N` handle (passes through) or a
// service name like "ssh" (resolved via `name=<n> dynamic=false` to its `.id`).
// Backwards-compatible with the previous name-based import IDs that landed
// in user state.
func importIpServiceState(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	id := d.Id()
	if len(id) == 0 {
		return nil, fmt.Errorf("empty import id")
	}

	// `*N`-style handle — accept as-is, the read step will refresh attributes.
	if id[0] == '*' {
		return []*schema.ResourceData{d}, nil
	}

	// Allow `field=value` form for forward compatibility, default to name.
	field, value := "name", id
	if parts := strings.SplitN(id, "=", 2); len(parts) == 2 {
		field, value = parts[0], parts[1]
	}

	res, err := ReadItemsFiltered(
		[]string{SnakeToKebab(field) + "=" + value, "dynamic=false"},
		"/ip/service", m.(Client),
	)
	if err != nil {
		return nil, err
	}
	switch len(*res) {
	case 0:
		return nil, fmt.Errorf("ip service not found: %s=%s", field, value)
	case 1:
		resolved, ok := (*res)[0][".id"]
		if !ok {
			return nil, fmt.Errorf("ip service lookup returned no .id")
		}
		d.SetId(resolved)
	default:
		return nil, fmt.Errorf("ip service ambiguous: %s=%s matched %d rows", field, value, len(*res))
	}

	return []*schema.ResourceData{d}, nil
}
