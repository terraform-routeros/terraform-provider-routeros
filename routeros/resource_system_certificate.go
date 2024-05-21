package routeros

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*1",
    "akid": "",
    "authority": "true",
    "common-name": "MyRouter",
    "crl": "false",
    "days-valid": "3650",
    "digest-algorithm": "sha256",
    "expires-after": "519w4d15h28m25s",
    "fingerprint": "ad9b324e93dee5135d2f6292480b78a5e00ae7ab44bf082b10cb1947993793c7",
    "invalid-after": "mar/10/2033 17:57:09",
    "invalid-before": "mar/13/2023 17:57:09",
    "key-size": "2048",
    "key-type": "rsa",
    "key-usage": "key-cert-sign,crl-sign",
    "name": "root-cert",
    "private-key": "true",
    "serial-number": "21B9F571B54195D3",
    "skid": "c90ec1a6d381b97bfa6b2c2c5c3ee81cf80ea729",
    "smart-card-key": "false",
    "subject-alt-name": "",
    "trusted": "true"
  }
*/

// https://help.mikrotik.com/docs/display/ROS/Certificates
// https://wiki.mikrotik.com/wiki/Manual:System/Certificates
func ResourceSystemCertificate() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/certificate"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("import", "sign", "sign_via_scep"),

		"authority": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"akid": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Authority Key Identifier.",
		},
		"ca": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"ca_crl_host": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"ca_fingerprint": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"challenge_password": {
			Type:        schema.TypeString,
			Computed:    true,
			Sensitive:   true,
			Description: "A challenge password for scep client.",
		},
		"common_name": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "Common Name (e.g. server FQDN or YOUR name).",
		},
		"copy_from": {
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
			Description: "",
		},
		"country": {
			Type:         schema.TypeString,
			Optional:     true,
			ForceNew:     true,
			Description:  "Country Name (2 letter code).",
			ValidateFunc: validation.StringLenBetween(2, 2),
		},
		"crl": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"days_valid": {
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			ForceNew:    true,
			Description: "Certificate lifetime.",
		},
		"dsa": {
			Type:     schema.TypeBool,
			Computed: true,
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
		"import": {
			Type:          schema.TypeSet,
			Optional:      true,
			ForceNew:      true,
			ConflictsWith: []string{"sign", "sign_via_scep"},
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"cert_file_name": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "Certificate file name that will be imported.",
					},
					"key_file_name": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Key file name that will be imported.",
					},
					"passphrase": {
						Type:        schema.TypeString,
						Optional:    true,
						Sensitive:   true,
						Description: "File passphrase if there is such.",
					},
				},
			},
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
			Optional: true,
			Computed: true,
			ForceNew: true,
			ValidateFunc: validation.StringInSlice([]string{"1024", "1536", "2048", "4096", "8192",
				"prime256v1", "secp384r1", "secp521r1"}, false),
		},
		"key_usage": {
			Type:        schema.TypeSet,
			Optional:    true,
			Computed:    true,
			ForceNew:    true,
			Description: "Detailed key usage descriptions can be found in RFC 5280.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringInSlice(
					[]string{
						"digital-signature",
						"content-commitment",
						"key-encipherment",
						"data-encipherment",
						"key-agreement",
						"key-cert-sign",
						"crl-sign",
						"encipher-only",
						"decipher-only",
						"dvcs",
						"server-gated-crypto",
						"ocsp-sign",
						"timestamp",
						"ipsec-user",
						"ipsec-tunnel",
						"ipsec-end-system",
						"email-protect",
						"code-sign",
						"tls-server",
						"tls-client",
					}, false)},
		},
		"locality": {
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
			Description: "Locality Name (eg, city).",
		},
		KeyName: PropName("Name of the certificate. Name can be edited."),
		"organization": {
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
			Description: "Organizational Unit Name (eg, section)",
		},
		"private_key": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"req_fingerprint": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"revoked": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"scep_url": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"serial_number": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"sign": {
			Type:          schema.TypeSet,
			Optional:      true,
			ForceNew:      true,
			ConflictsWith: []string{"sign_via_scep"},
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"ca": {
						Type:        schema.TypeString,
						Optional:    true,
						ForceNew:    true,
						Description: "Which CA to use if signing issued certificates.",
					},
					"ca_crl_host": {
						Type:        schema.TypeString,
						Optional:    true,
						ForceNew:    true,
						Description: "CRL host if issuing CA certificate.",
					},
					// "ca_on_smart_card": {
					// 	Type:        schema.TypeString,
					// 	Optional:    true,
					// 	Description: "",
					// },
					// We do not change the name of the certificate after signing so that there is non empty plan.
					// "name": {
					// 	Type:        schema.TypeString,
					// 	Optional:    true,
					// 	Description: "What name to assign to issued certificate.",
					// },
				},
			},
		},
		"sign_via_scep": {
			Type:          schema.TypeSet,
			Optional:      true,
			ForceNew:      true,
			ConflictsWith: []string{"sign"},
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"scep_url": {
						Type:         schema.TypeString,
						Required:     true,
						ForceNew:     true,
						ValidateFunc: validation.IsURLWithScheme([]string{"http"}),
						Description:  "HTTP URL to the SCEP server.",
					},
					"challenge_password": {
						Type:        schema.TypeString,
						Optional:    true,
						Sensitive:   true,
						ForceNew:    true,
						Description: "A challenge password.",
					},
					"ca_identity": {
						Type:        schema.TypeString,
						Optional:    true,
						ForceNew:    true,
						Description: "SCEP CA identity.",
					},
					"on_smart_card": {
						Type:        schema.TypeBool,
						Optional:    true,
						ForceNew:    true,
						Description: "Whether to store a private key on smart card if hardware supports it.",
					},
					"refresh": {
						Type:        schema.TypeBool,
						Optional:    true,
						Default:     true,
						ForceNew:    true,
						Description: "Check certificate expiration and refresh it if expired.",
					},
				},
			},
		},
		"skid": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Subject Key Identifier.",
		},
		"smart_card_key": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"state": {
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
			Description: "State or Province Name (full name).",
		},
		"status": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Shows current status of scep client.",
		},
		"subject_alt_name": {
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
			Description: "SANs (subject alternative names).",
		},
		"trusted": {
			Type:             schema.TypeBool,
			Optional:         true,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			Description:      "If set to yes certificate is included 'in trusted certificate chain'.",
		},
		"unit": {
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
			Description: "Organizational Unit Name (eg, section).",
		},
	}

	certImport := func(ctx context.Context, attrBlock any, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		var certName string

		bl := attrBlock.(*schema.Set).List()[0].(map[string]interface{})
		var resUrl = &URL{Path: resSchema[MetaResourcePath].Default.(string)}
		if m.(Client).GetTransport() == TransportREST {
			resUrl.Path += "/import"
		}

		params := MikrotikItem{KeyName: d.Get(KeyName).(string), "file-name": bl["cert_file_name"].(string)}
		if passwd, ok := bl["passphrase"]; ok {
			params["passphrase"] = passwd.(string)
		}

		// Import certificate
		err := m.(Client).SendRequest(crudImport, resUrl, params, nil)
		if err != nil {
			return diag.FromErr(err)
		}

		if keyFile, ok := bl["key_file_name"]; ok {
			params = MikrotikItem{KeyName: d.Get(KeyName).(string), "file-name": keyFile.(string)}
			if passwd, ok := bl["passphrase"]; ok {
				params["passphrase"] = passwd.(string)
			}

			// Import key
			err := m.(Client).SendRequest(crudImport, resUrl, params, nil)
			if err != nil {
				return diag.FromErr(err)
			}
		}

		res, err := ReadItemsFiltered([]string{"name=" + d.Get(KeyName).(string)},
			resSchema[MetaResourcePath].Default.(string), m.(Client))
		if err != nil {
			return diag.FromErr(err)
		}

		switch len(*res) {
		case 0:
			return diag.Errorf("resource not found: name=%v", certName)
		case 1:
			retId, ok := (*res)[0][Id.String()]
			if !ok {
				return diag.Errorf("attribute %v not found in the response", Id.String())
			}
			d.SetId(retId)
		default:
			return diag.Errorf("more than one resource found: name=%v", certName)
		}

		return ResourceRead(ctx, resSchema, d, m)
	}

	resCreate := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		var diags diag.Diagnostics
		var cmdBlock any // User config for certificate signing
		var crudMethod crudMethod
		var command string // MikroTik command to sign certificate
		var ok bool

		if _, ok = d.GetOk("import"); !ok {
			// Run DefaultCreate.
			diags = ResourceCreate(ctx, resSchema, d, m)
			if diags.HasError() {
				return diags
			}
		}

		var params MikrotikItem // Parameters for MikroTik command

		if cmdBlock, ok = d.GetOk("sign"); ok {
			// {"number":"*54", ca: "Test-CA"}
			params = MikrotikItem{"number": d.Id()}
			crudMethod = crudSign
			// https://router/rest/certificate/sign
			command = "/sign"
		} else if cmdBlock, ok = d.GetOk("sign_via_scep"); ok {
			params = MikrotikItem{"template": d.Get("name").(string)}
			crudMethod = crudSignViaScep
			// https://router/rest/certificate/add-scep
			command = "/add-scep"
		} else if cmdBlock, ok = d.GetOk("import"); ok {
			return certImport(ctx, cmdBlock, d, m)
		} else {
			return diags
		}

		// []interface{map[string]interface{...}}
		for k, v := range cmdBlock.(*schema.Set).List()[0].(map[string]interface{}) {
			k = SnakeToKebab(k)
			switch v := v.(type) {
			case string:
				if v == "" {
					continue
				}
				params[k] = v
			case bool:
				params[k] = BoolToMikrotikJSON(v)
			default:
				panic("ResourceSystemCertificate resCreate: unhandled type switch")
			}
		}

		var resUrl = &URL{
			Path: resSchema[MetaResourcePath].Default.(string),
		}
		if m.(Client).GetTransport() == TransportREST {
			resUrl.Path += command
		}

		err := m.(Client).SendRequest(crudMethod, resUrl, params, nil)
		if err != nil {
			return diag.FromErr(err)
		}

		return ResourceRead(ctx, resSchema, d, m)
	}

	resDelete := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		// {"number":"*54"}
		item := MikrotikItem{"numbers": d.Id()}

		var resUrl = &URL{
			Path: resSchema[MetaResourcePath].Default.(string),
		}
		var method crudMethod = crudRemove
		if _, ok := d.State().Attributes["ca"]; ok {
			// Not Root CA.
			method = crudRevoke
		}

		if m.(Client).GetTransport() == TransportREST {
			if _, ok := d.State().Attributes["ca"]; ok {
				// Not Root CA.
				resUrl.Path += "/issued-revoke"
			} else {
				// Root CA.
				resUrl.Path += "/remove"
			}
		}

		err := m.(Client).SendRequest(method, resUrl, item, nil)
		if err != nil {
			return diag.FromErr(err)
		}

		d.SetId("")
		return nil
	}

	return &schema.Resource{
		CreateContext: resCreate,
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultUpdate(resSchema),
		DeleteContext: resDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
