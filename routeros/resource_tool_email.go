package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  "from": "<>",
  "password": "",
  "port": "25",
  "server": "0.0.0.0",
  "tls": "no",
  "user": "",
  "vrf": "main"
}
*/

// https://help.mikrotik.com/docs/display/ROS/E-mail
func ResourceToolEmail() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/tool/e-mail"),
		MetaId:           PropId(Id),

		"from": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Name or email address that will be shown as a receiver.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"password": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "Password used for authenticating to an SMTP server.",
		},
		"port": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "SMTP server's port.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"server": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "SMTP server's IP address.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tls": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Whether to use TLS encryption:" +
				"\n  * yes - sends STARTTLS and drops the session if TLS is not available on the server" +
				"\n  * no - do not send STARTTLS" +
				"\n  * starttls - sends STARTTLS and continue without TLS if a server responds that TLS is not available",
			ValidateFunc:     validation.StringInSlice([]string{"yes", "no", "starttls"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"user": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The username used for authenticating to an SMTP server.",
		},
		"vrf": PropVrfRw,
	}

	return &schema.Resource{
		CreateContext: DefaultSystemCreate(resSchema),
		ReadContext:   DefaultSystemRead(resSchema),
		UpdateContext: DefaultSystemUpdate(resSchema),
		DeleteContext: DefaultSystemDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
