package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
    ".id": "*1",
    "called-format": "II-II-II-II-II-II:S",
    "calling-format": "AA-AA-AA-AA-AA-AA",
    "disabled": "false",
    "interim-update": "disabled",
    "mac-caching": "disabled",
    "name": "aaa1",
    "nas-identifier": "router",
    "password-format": "",
    "username-format": "AA:AA:AA:AA:AA:AA"
}
*/

// https://help.mikrotik.com/docs/display/ROS/WiFi#WiFi-AAAproperties
func ResourceWifiAaa() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/wifi/aaa"),
		MetaId:           PropId(Id),

		"called_format": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Format of the `Called-Station-Id` RADIUS attribute.",
		},
		"calling_format": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Format of the `Calling-Station-Id` RADIUS attribute.",
		},
		KeyComment: PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"interim_update": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Interval at which to send interim updates about traffic accounting to the RADIUS server.",
		},
		"mac_caching": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Time to cache RADIUS server replies when MAC address authentication is enabled.",
		},
		KeyName: PropName("Name of the AAA profile."),
		"nas_identifier": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Value of the `NAS-Identifier` RADIUS attribute.",
		},
		"password_format": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Format of the `User-Password` RADIUS attribute.",
		},
		"username_format": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Format of the `User-Name` RADIUS attribute.",
		},
	}

	return &schema.Resource{
		Description:   `*<span style="color:red">This resource requires a minimum version of RouterOS 7.13.</span>*`,
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
