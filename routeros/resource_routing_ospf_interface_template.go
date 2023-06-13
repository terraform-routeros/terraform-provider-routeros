package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*54",
    "interfaces": "",
    "network": "",
    "area": "",
    "auth": "",
    "auth-id": "",
    "authentication-key": "",
    "comment": "",
    "cost": "",
    "dead-interval": "",
    "disabled": "",
    "hello-interval": "",
    "instance-id": "",
    "passive": "",
    "prefix-list": "",
    "priority": "",
    "retransmit-interval": "",
    "transmit-delay": "",
    "type": "",
    "vlink-neighbor-id": "",
    "vlink-transit-area": "",
  }
*/

// ResourceRoutingOspfInterfaceTemplate https://help.mikrotik.com/docs/display/ROS/OSPF
func ResourceRoutingOspfInterfaceTemplate() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/routing/ospf/interface-template"),
		MetaId:           PropId(Id),

		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"interfaces": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "none",
			Description: "Interfaces to match.",
		},
		"network": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The network prefix associated with the area.",
		},
		"area": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The OSPF area to which the matching interface will be associated.",
		},
		"auth": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Specifies authentication method for OSPF protocol messages.",
			ValidateFunc: validation.StringInSlice([]string{"simple", "md5", "sha1", "sha256", "sha384", "sha512"}, true),
		},
		"auth_id": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "The key id is used to calculate message digest (used when MD5 or SHA authentication is enabled).",
		},
		"authentication_key": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The authentication key to be used, should match on all the neighbors of the network segment.",
		},
		"cost": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      1,
			Description:  "Interface cost expressed as link state metric.",
			ValidateFunc: validation.IntBetween(0, 65535),
		},
		"dead_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "40s",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEquall,
			Description:      "Specifies the interval after which a neighbor is declared dead.",
		},
		"hello_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "10s",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEquall,
			Description:      "The interval between HELLO packets that the router sends out this interface.",
		},
		"instance_id": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "Interface cost expressed as link state metric.",
			Default:      0,
			ValidateFunc: validation.IntBetween(0, 255),
		},
		"passive": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "If enabled, then do not send or receive OSPF traffic on the matching interfaces",
		},
		"prefix_list": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the address list containing networks that should be advertised to the v3 interface.",
		},
		"priority": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "Router's priority. Used to determine the designated router in a broadcast network.",
			Default:      128,
			ValidateFunc: validation.IntBetween(0, 255),
		},
		"retransmit_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "5s",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEquall,
			Description:      "Time interval the lost link state advertisement will be resent.",
		},
		"transmit_delay": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "1s",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEquall,
			Description:      "Link-state transmit delay is the estimated time it takes to transmit a link-state update packet on the interface.",
		},
		"type": {
			Type:         schema.TypeString,
			Description:  "The OSPF network type on this interface.",
			Optional:     true,
			Default:      "broadcast",
			ValidateFunc: validation.StringInSlice([]string{"broadcast", "nbma", "ptp", "ptmp", "ptp-unnumbered", "virtual-link"}, true),
		},
		"vlink_neighbor_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Specifies the router-id of the neighbor which should be connected over the virtual link.",
		},
		"vlink_transit_area": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "A non-backbone area the two routers have in common over which the virtual link will be established.",
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
