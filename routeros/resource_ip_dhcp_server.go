package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ResourceDhcpServer https://help.mikrotik.com/docs/display/ROS/DHCP#DHCP-Leases
func ResourceDhcpServer() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/dhcp-server"),
		MetaId:           PropId(Id),

		"address_pool": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "static-only",
			Description: "IP pool, from which to take IP addresses for the clients. If set to static-only, then only " +
				"the clients that have a static lease (added in lease submenu) will be allowed.",
		},
		"authoritative": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "yes",
			Description:  "Option changes the way how a server responds to DHCP requests.",
			ValidateFunc: validation.StringInSlice([]string{"after-10sec-delay", "after-2sec-delay", "yes", "no"}, false),
		},
		KeyComment:   PropCommentRw,
		KeyDisabled:  PropDisabledRw,
		KeyDynamic:   PropDynamicRo,
		KeyInterface: PropInterfaceRw,
		KeyInvalid:   PropInvalidRo,
		"lease_script": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "A script that will be executed after a lease is assigned or de-assigned.",
		},
		"lease_time": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "10m",
			Description: "The time that a client may use the assigned address. The client will try to renew this " +
				"address after half of this time and will request a new address after the time limit expires.",
		},
		KeyName: PropNameRw,
		"use_radius": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		// N.B. Some options have not been added!
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
