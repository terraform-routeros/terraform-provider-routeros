package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*7",
    "allow-fast-path": "true",
    "arp": "enabled",
    "arp-timeout": "auto",
    "comment": "Comment",
    "disabled": "false",
    "dont-fragment": "disabled",
    "l2mtu": "65535",
    "loop-protect": "default",
    "loop-protect-disable-time": "5m",
    "loop-protect-send-interval": "5s",
    "loop-protect-status": "off",
    "mac-address": "8A:36:70:DA:F7:FA",
    "max-fdb-size": "4096",
    "mtu": "1500",
    "name": "vxlan1",
    "port": "8472",
    "running": "true",
    "vni": "1",
    "vrf": "main",
    "vteps-ip-version": "ipv4"
  }
*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/100007937/VXLAN
func ResourceInterfaceVxlan() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/vxlan"),
		MetaId:           PropId(Id),

		"allow_fast_path": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether to allow Fast Path processing. Fragmented and flooded packets over VXLAN are redirected " +
				"via a slow path. Fast Path is disabled for VXLAN interface that uses IPv6 VTEP version or VRF. The setting " +
				"is available since RouterOS version 7.8.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyArp:        PropArpRw,
		KeyArpTimeout: PropArpTimeoutRw,
		KeyComment:    PropCommentRw,
		KeyDisabled:   PropDisabledRw,
		"dont_fragment": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The Don't Fragment (DF) flag controls whether a packet can be broken into smaller packets, " +
				"called fragments, before being sent over a network. When configuring VXLAN, this setting determines " +
				"the presence of the DF flag on the outer IPv4 header and can control packet fragmentation if the encapsulated " +
				"packet exceeds the outgoing interface MTU. This setting has three options:\n  * disabled - the DF flag is " +
				"not set on the outer IPv4 header, which means that packets can be fragmented if they are too large to " +
				"be sent over the outgoing interface. This also allows packet fragmentation when VXLAN uses IPv6 underlay. " +
				"\n  * enabled - the DF flag is always set on the outer IPv4 header, which means that packets will not be fragmented " +
				"and will be dropped if they exceed the outgoing interface's MTU. This also avoids packet fragmentation " +
				"when VXLAN uses IPv6 underlay.\n  * inherit - The DF flag on the outer IPv4 header is based on the inner " +
				"IPv4 DF flag. If the inner IPv4 header has the DF flag set, the outer IPv4 header will also have it " +
				"set. If the packet exceeds the outgoing interface's MTU and DF is set, it will be dropped. If the inner " +
				"packet is non-IP, the outer IPv4 header will not have the DF flag set and packets can be fragmented. " +
				"If the inner packet is IPv6, the outer IPv4 header will always set the DF flag and packets cannot be " +
				"fragmented. Note that when VXLAN uses IPv6 underlay, this setting does not have any effect and is treated " +
				"the same as disabled. The setting is available since RouterOS version 7.8.",
			ValidateFunc:     validation.StringInSlice([]string{"disabled", "enabled", "inherit"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"group": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "When specified, a multicast group address can be used to forward broadcast, unknown-unicast, " +
				"and multicast traffic between VTEPs. This property requires specifying the interface setting. The interface " +
				"will use IGMP or MLD to join the specified multicast group, make sure to add the necessary PIM and IGMP/MDL " +
				"configuration. When this property is set, the vteps-ip-version automatically gets updated to the used " +
				"multicast IP version.",
			ValidateFunc: validation.StringInSlice([]string{"IPv4", "IPv6"}, false),
		},
		"interface": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Interface name used for multicast forwarding. This property requires specifying the group " +
				"setting.",
		},
		"local_address": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Specifies the local source address for the VXLAN interface. If not set, one IP address of " +
				"the egress interface will be selected as a source address for VXLAN packets. When the property is set, " +
				"the vteps-ip-version automatically gets updated to the used local IP version. The setting is available " +
				"since RouterOS version 7.7.",
		},
		"mac_address": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Static MAC address of the interface. A randomly generated MAC address will be assigned when " +
				"not specified.",
			ValidateFunc:     validation.IsMACAddress,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"max_fdb_size": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Limits the maximum number of MAC addresses that VXLAN can store in the forwarding database " +
				"(FDB).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyMtu:  PropMtuRw(),
		KeyName: PropName("Name of the interface."),
		"port": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Used UDP port number.",
			ValidateFunc:     Validation64k,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"vni": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "VXLAN Network Identifier (VNI).",
			ValidateFunc:     validation.IntAtLeast(1),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyVrf: PropVrfRw,
		"vteps_ip_version": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Used IP protocol version for statically configured VTEPs. The RouterOS VXLAN interface does " +
				"not support dual-stack, any configured remote VTEPs with the opposite IP version will be ignored. When " +
				"multicast group or local-address properties are set, the vteps-ip-version automatically gets updated " +
				"to the used IP version. The setting is available since RouterOS version 7.6.",
			ValidateFunc:     validation.StringInSlice([]string{"ipv4", "ipv6"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
	}

	return &schema.Resource{
		CreateContext: DefaultCreate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultUpdate(resSchema),
		DeleteContext: DefaultDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: ImportStateCustomContext(resSchema),
		},

		Schema: resSchema,
	}
}
