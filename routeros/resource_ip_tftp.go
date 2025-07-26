package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*1",
    "allow": "true",
    "allow-rollover": "false",
    "disabled": "false",
	"hits": 0,
    "ip-addresses": "10.0.0.0/24",
    "read-only": "true",
    "real-filename": "/usb1/file.txt",
    "req-filename": "file.txt"
  }
*/

// ResourceIPTFTP https://wiki.mikrotik.com/Manual:IP/TFTP
func ResourceIpTFTP() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/tftp"),
		MetaId:           PropId(Id),

		"allow": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Allow connection.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"allow_overwrite": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "If `true`, overwriting the file is allowed.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"allow_rollover": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "If set, server will allow sequence number to roll over when " +
				"maximum value is reached. This is used to enable large " +
				"downloads using TFTP server.",
			Default: false,
		},
		KeyDisabled: PropDisabledRw,
		"hits": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "How many times this access rule entry has been used.",
		},
		"ip_addresses": {
			Type: schema.TypeSet,
			Elem: &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.Any(
					validation.IsCIDR,
					validation.IsIPAddress,
				),
			},
			Optional: true,
			Description: "Range of IP addresses accepted as clients. " +
				"If empty `0.0.0.0/0` will be used.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"read_only": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Sets if file can be written to. If set to `false` write " +
				"attempts will fail with an error.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"reading_window_size": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "TFTP Windowsize option value.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"real_filename": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "If `req-filename` and `real-filename` values are set and " +
				"valid, the requested filename will be replaced with " +
				"matched file. This field has to be set. If multiple regex " +
				"are specified in `req-filename`, with this field you can " +
				"set which ones should match, so this rule is validated. " +
				"`real-filename` format for using multiple regex is " +
				"`filename\\0\\5\\6`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"req_filename": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Requested filename as regular expression (regex) if field " +
				"is left empty it defaults to `.*`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
	}
	return &schema.Resource{
		CreateContext: DefaultCreate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultUpdate(resSchema),
		DeleteContext: DefaultDelete(resSchema),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
