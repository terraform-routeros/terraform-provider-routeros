package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/* /caps-man/manager
{
  "ca-certificate": "none",
  "certificate": "none",
  "enabled": "false",
  "generated-ca-certificate":"CAPsMAN-CA-000000000000",
  "generated-certificate":"CAPsMAN-000000000000",
  "package-path": "",
  "require-peer-certificate": "false",
  "upgrade-policy": "none"
}
*/

// https://help.mikrotik.com/docs/display/ROS/CAPsMAN
func ResourceCapsManManager() *schema.Resource {

	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/caps-man/manager"),
		MetaId:           PropId(Name),

		"ca_certificate": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "none",
			Description: "Device CA certificate.",
		},
		"certificate": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "none",
			Description: "Device certificate.",
		},
		KeyEnabled: PropEnabled("Disable or enable CAPsMAN functionality."),
		"generated_ca_certificate": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Generated CA certificate.",
		},
		"generated_certificate": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Generated CAPsMAN certificate.",
		},
		"package_path": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Folder location for the RouterOS packages. For example, use '/upgrade' to specify the " +
				"upgrade folder from the files section. If empty string is set, CAPsMAN can use built-in RouterOS " +
				"packages, note that in this case only CAPs with the same architecture as CAPsMAN will be upgraded.",
		},
		"require_peer_certificate": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Require all connecting CAPs to have a valid certificate.",
		},
		"upgrade_policy": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "none",
			Description:  "Upgrade policy options.",
			ValidateFunc: validation.StringInSlice([]string{"none", "require-same-version", "suggest-same-version"}, false),
		},
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

/* /caps-man/manager/interface
{
  ".id": "*1",
  "default": "true",
  "disabled": "false",
  "dynamic": "false",
  "forbid": "false",
  "interface": "all"
}
*/

// https://wiki.mikrotik.com/wiki/Manual:Simple_CAPsMAN_setup
func ResourceCapsManManagerInterface() *schema.Resource {

	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/caps-man/manager/interface"),
		MetaId:           PropId(Id),

		KeyComment:  PropCommentRw,
		KeyDefault:  PropDefaultRo,
		KeyDisabled: PropDisabledRw,
		KeyDynamic:  PropDynamicRo,
		"forbid": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Disable interface listening.",
		},
		KeyInterface: PropInterfaceRw,
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

/* /caps-man/aaa
{
  "called-format": "mac:ssid",
  "interim-update": "disabled",
  "mac-caching": "disabled",
  "mac-format": "XX:XX:XX:XX:XX:XX",
  "mac-mode": "as-username"
}
*/

// https://help.mikrotik.com/docs/display/ROS/CAPsMAN
func ResourceCapsManAaa() *schema.Resource {

	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/caps-man/aaa"),
		MetaId:           PropId(Name),

		"called_format": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "mac:ssid",
			Description: "Format of how the 'called-id' identifier will be passed to RADIUS. When configuring radius " +
				"server clients, you can specify 'called-id' in order to separate multiple entires.",
			ValidateFunc: validation.StringInSlice([]string{"mac", "mac:ssid", "ssid"}, false),
		},
		"interim_update": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "When RADIUS accounting is used, Access Point periodically sends accounting information " +
				"updates to the RADIUS server. This property specifies the default update interval that can be " +
				"overridden by the RADIUS server using the Acct-Interim-Interval attribute.",
			//DiffSuppressFunc: TimeEquall, // "interim-update": "disabled"
		},
		"mac_caching": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "If this value is set to a time interval, the Access Point will cache RADIUS MAC authentication " +
				"responses for a specified time, and will not contact the RADIUS server if matching cache entry already " +
				"exists. The value disabled will disable the cache, Access Point will always contact the RADIUS server.",
			//DiffSuppressFunc: TimeEquall, // "mac-caching": "disabled"
		},
		"mac_format": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "XX:XX:XX:XX:XX:XX",
			Description: "Controls how the MAC address of the client is encoded by Access Point in the User-Name " +
				"attribute of the MAC authentication and MAC accounting RADIUS requests.",
		},
		"mac_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "as-username",
			Description: "By default Access Point uses an empty password, when sending Access-Request during MAC " +
				"authentication. When this property is set to as-username-and-password, Access Point will use the same " +
				"value for the User-Password attribute as for the User-Name attribute.",
			ValidateFunc: validation.StringInSlice([]string{"as-username", "as-username-and-password"}, false),
		},
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
