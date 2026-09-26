package routeros

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  ".id": "*1",
  "name": "qkd1",
  "type": "qkd",
  "key-size": "32",
  "qkd-address": "10.2.3.4:8020",
  "qkd-kme-id": "server-kme-id",
  "qkd-certificate": "none",
  "qkd-peer-sae-id": "client-sae-id",
  "qkd-cache-size": "1",
  "qkd-cache-state": "0",
  "qkd-total-keys-received": "0"
}
*/

// ResourceSystemKeymatProvider manages the generic key material providers used
// by IPsec (QKD / HKDF). Available since RouterOS 7.24.
// https://manual.mikrotik.com/docs/cli-reference/system/keymat-provider/
func ResourceSystemKeymatProvider() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/keymat-provider"),
		MetaId:           PropId(Name),

		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the key material provider.",
		},
		KeyDisabled: PropDisabledRw,
		"type": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			Description:  "Type of the key material provider: qkd or hkdf.",
			ValidateFunc: validation.StringInSlice([]string{"qkd", "hkdf"}, false),
		},
		"key_size": {
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Requested key size in bits.",
		},
		"qkd_address": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "KME device address (IP:port).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"qkd_kme_id": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "KME ID that must match the KME ID in the received TLS certificate.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"qkd_certificate": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Certificate used to authenticate to the KME; also specifies the local SAE ID.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"qkd_peer_sae_id": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Peer (master or slave) SAE ID.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"qkd_cache_size": {
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Number of unused keys to keep in the cache.",
		},
		"hkdf_secret": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "Secret used by the HKDF provider (hexadecimal value).",
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				return strings.EqualFold(old, new)
			},
		},
		"hkdf_hash": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			Description:  "Hash algorithm used by the HKDF provider.",
			ValidateFunc: validation.StringInSlice([]string{"sha256", "sha512"}, false),
		},
		"qkd_cache_state": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Current number of unused keys in the cache.",
		},
		"qkd_total_keys_received": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Total number of keys received from the KME.",
		},
	}

	return &schema.Resource{
		CreateContext: DefaultCreate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultUpdate(resSchema),
		DeleteContext: DefaultDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: ImportStateCustomContext(resSchema),
		},

		Schema: resSchema,
	}
}