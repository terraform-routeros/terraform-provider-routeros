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
		MetaId:           PropId(Name),

		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"interfaces": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "Interfaces to match.",
		},
		"network": {
			Type:        schema.TypeString,
			Description: "The network prefix associated with the area.",
		},
		"area": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The OSPF area to which the matching interface will be associated.",
		},
		"auth": {
			Type:         schema.TypeString,
			Description:  "Specifies authentication method for OSPF protocol messages.",
			ValidateFunc: validation.StringInSlice([]string{"simple", "md5", "sha1", "sha256", "sha384", "sha512"}, true),
		},
		"auth-id": {
			Type:        schema.TypeInt,
			Description: "The key id is used to calculate message digest (used when MD5 or SHA authentication is enabled).",
		},
		"authentication-key": {
			Type:        schema.TypeString,
			Description: "The authentication key to be used, should match on all the neighbors of the network segment.",
		},
		"cost": {
			Type:         schema.TypeInt,
			Default:      1,
			Description:  "Interface cost expressed as link state metric.",
			ValidateFunc: validation.IntBetween(0, 65535),
		},
		"dead-interval": {
			Type:        schema.TypeString,
			Default:     "00:00:40",
			Description: "Specifies the interval after which a neighbor is declared dead.",
		},
		"hello-interval": {
			Type:        schema.TypeString,
			Default:     "00:00:10",
			Description: "The interval between HELLO packets that the router sends out this interface.",
		},
		"instance-id": {
			Type:         schema.TypeInt,
			Description:  "Interface cost expressed as link state metric.",
			Default:      0,
			ValidateFunc: validation.IntBetween(0, 255),
		},
		"passive": {
			Type:        schema.TypeBool,
			Default:     false,
			Description: "If enabled, then do not send or receive OSPF traffic on the matching interfaces",
		},
		"prefix-list": {
			Type:        schema.TypeString,
			Description: "Name of the address list containing networks that should be advertised to the v3 interface.",
		},
		"priority": {
			Type:         schema.TypeInt,
			Description:  "Router's priority. Used to determine the designated router in a broadcast network.",
			Default:      128,
			ValidateFunc: validation.IntBetween(0, 255),
		},
		"retransmit-interval": {
			Type:        schema.TypeString,
			Default:     "00:00:05",
			Description: "Time interval the lost link state advertisement will be resent.",
		},
		"transmit-delay": {
			Type:        schema.TypeString,
			Default:     "00:00:01",
			Description: "Link-state transmit delay is the estimated time it takes to transmit a link-state update packet on the interface.",
		},
		"type": {
			Type:         schema.TypeString,
			Description:  "The OSPF network type on this interface.",
			Default:      "broadcast",
			ValidateFunc: validation.StringInSlice([]string{"broadcast", "nbma", "ptp", "ptmp", "ptp-unnumbered", "virtual-link"}, true),
		},
		"vlink-neighbor-id": {
			Type:        schema.TypeString,
			Description: "Specifies the router-id of the neighbor which should be connected over the virtual link.",
		},
		"vlink-transit-area": {
			Type:        schema.TypeString,
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
