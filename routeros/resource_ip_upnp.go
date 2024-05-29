package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceUPNPSettings https://help.mikrotik.com/docs/display/ROS/UPnP
func ResourceUPNPSettings() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/upnp"),
		MetaId:           PropId(Id),

		"allow_disable_external_interface": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether or not should the users are allowed to disable the router's external interface. " +
				"This functionality (for users to be able to turn the router's external interface off without any " +
				"authentication procedure) is required by the standard, but as it is sometimes not expected or " +
				"unwanted in UPnP deployments which the standard was not designed for (it was designed mostly for " +
				"home users to establish their own local networks), you can disable this behavior",
		},
		KeyEnabled: PropEnabled("Enable UPnP service."),
		"show_dummy_rule": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "nable a workaround for some broken implementations, which are handling the absence of " +
				"UPnP rules incorrectly (for example, popping up error messages). This option will instruct the " +
				"server to install a dummy (meaningless) UPnP rule that can be observed by the clients, which refuse " +
				"to work correctly otherwise",
		},
	}
	return &schema.Resource{
		Description: "*<span style=\"color:red\">If you do not disable the `allow-disable-external-interface`, any " +
			"user from the local network will be able (without any authentication procedures) to disable the router's " +
			"external interface.</span>*",
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
