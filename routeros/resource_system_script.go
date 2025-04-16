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
    "dont-require-permissions": "false",
    "invalid": "false",
    "last-started": "jan/13/2023 00:16:01",
    "name": "unreg_died",
    "owner": "admin",
    "policy": "read,write,policy,password,sensitive",
    "run-count": "47",
    "source": ":log info \"TEST\";\r\n"
  }
*/

// ResourceSystemScript https://help.mikrotik.com/docs/display/ROS/Scripting#Scripting-Scriptrepository
func ResourceSystemScript() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/script"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("launch_trigger"),

		KeyComment: PropCommentRw,
		"dont_require_permissions": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Bypass permissions check when the script is being executed, useful when scripts are being executed " +
				"from services that have limited permissions, such as Netwatch.",
		},
		KeyInvalid: PropInvalidRo,
		"last_started": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Date and time when the script was last invoked.",
		},
		KeyName: PropName("Name of the script."),
		"owner": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"policy": {
			Type:     schema.TypeSet,
			Computed: true,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"ftp", "reboot", "read", "write", "policy", "test",
					"password", "sniff", "sensitive"}, false),
			},
			Description: `List of applicable policies:
	* ftp - Policy that grants full rights to log in remotely via FTP, to read/write/erase files and to transfer files from/to the router. Should be used together with read/write policies.  
	* password - Policy that grants rights to change the password.  
	* policy - Policy that grants user management rights. Should be used together with the write policy. Allows also to see global variables created by other users (requires also 'test' policy).  
	* read - Policy that grants read access to the router's configuration. All console commands that do not alter router's configuration are allowed. Doesn't affect FTP.  
	* reboot - Policy that allows rebooting the router.  
	* sensitive - Policy that grants rights to change "hide sensitive" option, if this policy is disabled sensitive information is not displayed.  
	* sniff - Policy that grants rights to use packet sniffer tool.  
	* test - Policy that grants rights to run ping, traceroute, bandwidth-test, wireless scan, snooper, and other test commands.  
	* write - Policy that grants write access to the router's configuration, except for user management. This policy does not allow to read the configuration, so make sure to enable read policy as well.  
policy = ["ftp", "read", "write"]
`,
		},
		"launch_trigger": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Changing the attribute value causes the script to run.",
		},
		"run_count": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "This counter is incremented each time the script is executed.",
		},
		"source": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Script source code.",
		},
	}

	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			diags := ResourceCreate(ctx, resSchema, d, m)
			if diags.HasError() {
				return diags
			}

			if d.Get("launch_trigger").(string) != "" {
				startScript(ctx, resSchema, d, m)
			}

			return ResourceRead(ctx, resSchema, d, m)
		},

		ReadContext: DefaultRead(resSchema),
		UpdateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			diags := ResourceUpdate(ctx, resSchema, d, m)
			if diags.HasError() {
				return diags
			}

			if d.Get("launch_trigger").(string) != "" {
				startScript(ctx, resSchema, d, m)
			}

			return ResourceRead(ctx, resSchema, d, m)
		},

		DeleteContext: DefaultDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}

func startScript(ctx context.Context, s map[string]*schema.Schema, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Start the script.
	var resUrl = &URL{
		Path: s[MetaResourcePath].Default.(string),
	}
	if m.(Client).GetTransport() == TransportREST {
		resUrl.Path += "/run"
	}

	err := m.(Client).SendRequest(crudStart, resUrl, MikrotikItem{Id.String(): d.Id()}, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
