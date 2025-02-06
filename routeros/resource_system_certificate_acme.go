package routeros

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// https://help.mikrotik.com/docs/display/ROS/Certificates
// https://wiki.mikrotik.com/wiki/Manual:System/Certificates
func ResourceSystemCertificateAcme() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/certificate"),
		MetaId:           PropId(Id),
		MetaTransformSet: PropTransformSet("challenge_type:type"),

		// ACME:
		"eab_hmac_key": {
			Type:             schema.TypeString,
			Optional:         true,
			ForceNew:         true,
			Sensitive:        true,
			Description:      "HMAC key for ACME External Account Binding (optional).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"eab_kid": {
			Type:             schema.TypeString,
			Optional:         true,
			ForceNew:         true,
			Description:      "Key identifier.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"directory_url": {
			Type:             schema.TypeString,
			Optional:         true,
			ForceNew:         true,
			Description:      "ACME directory url.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"dns_name": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
			Description: " If the dns-name is not specified, it will default to the automatically generated " +
				"`/ip cloud name` (ie. http://example.sn.mynetname.net).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"challenge_type": {
			Type:             schema.TypeString,
			Optional:         true,
			ForceNew:         true,
			Description:      "ACME challenge.",
			ValidateFunc:     validation.StringInSlice([]string{"cloud-dns", "http-01"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},

		// Certificate:
		"akid": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Authority Key Identifier.",
		},
		"common_name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Common Name (e.g. server FQDN or YOUR name).",
		},
		"country": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Country Name (2 letter code).",
		},
		"crl": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"days_valid": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Certificate lifetime.",
		},
		"digest_algorithm": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"expired": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Set to true if certificate is expired.",
		},
		"expires_after": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"fingerprint": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"invalid_after": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The date after which certificate wil be invalid.",
		},
		"invalid_before": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The date before which certificate is invalid.",
		},
		"issued": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"issuer": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"key_type": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"key_size": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"key_usage": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "Detailed key usage descriptions can be found in RFC 5280.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"locality": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Locality Name (eg, city).",
		},
		KeyName: {
			Type:     schema.TypeString,
			Computed: true,
		},
		"organization": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Organizational Unit Name (eg, section)",
		},
		"private_key": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"serial_number": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"skid": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Subject Key Identifier.",
		},
		"state": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "State or Province Name (full name).",
		},
		"subject_alt_name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "SANs (subject alternative names).",
		},
		"trusted": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "If set to yes certificate is included 'in trusted certificate chain'.",
		},
		"unit": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Organizational Unit Name (eg, section).",
		},
	}

	serviceCheck := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		var diags diag.Diagnostics

		if challengeType := d.Get("challenge_type").(string); challengeType == "" || challengeType == "http-01" {
			res, err := ReadItemsFiltered([]string{"name=www"}, "/ip/service", m.(Client))
			if err != nil {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  err.Error(),
				})
			} else if len(*res) != 1 {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Service `name=www` not found.",
				})
			} else if BoolFromMikrotikJSON((*res)[0]["disabled"]) {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "The WWW service must be enabled before communicating with the ACME server.",
				})
			}
		}

		if d.Get("dns_name").(string) == "" {
			res, err := ReadItems(nil, "/ip/cloud", m.(Client))
			if err != nil {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  err.Error(),
				})
			} else if len(*res) != 1 {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Cloud service not found.",
				})

				// Kebab case!
			} else if (*res)[0]["ddns-enabled"] != "yes" {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "DNS name is empty, you need to enable the cloud service.",
				})
			}
		}

		return diags
	}

	resCreate := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		if diags := serviceCheck(ctx, d, m); diags.HasError() {
			return diags
		}

		item, metadata := TerraformResourceDataToMikrotik(resSchema, d)

		var resUrl string
		if m.(Client).GetTransport() == TransportREST {
			resUrl = apiMethodName[crudEnableSslCertificate]
		}

		err := m.(Client).SendRequest(crudEnableSslCertificate, &URL{Path: metadata.Path + resUrl}, item, nil)
		if err != nil {
			return diag.FromErr(err)
		}

		d.SetId(d.Get("name").(string))
		id, err := dynamicIdLookup(Name, resSchema[MetaResourcePath].Default.(string), m.(Client), d)

		if err != nil {
			return diag.FromErr(err)
		}

		d.SetId(id)

		return ResourceRead(ctx, resSchema, d, m)
	}

	resDelete := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

		var resUrl = &URL{
			Path: resSchema[MetaResourcePath].Default.(string) + "",
		}

		if m.(Client).GetTransport() == TransportREST {
			resUrl.Path += "/remove"
		}

		// {"number":"*54"}
		err := m.(Client).SendRequest(crudRemove, resUrl, MikrotikItem{"numbers": d.Id()}, nil)
		if err != nil {
			return diag.FromErr(err)
		}

		d.SetId("")
		return nil
	}

	return &schema.Resource{
		Description:   `*<span style="color:red">Various acme servers are supported starting with RouterOS 7.15beta7.</span>*`,
		CreateContext: resCreate,
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: resCreate,
		DeleteContext: resDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
