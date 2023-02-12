package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
[
  {
    ".id": "*0",
    ".nextid": "*1",
    "auto-isolate": "false",
    "bpdu-guard": "false",
    "bridge": "bridge",
    "broadcast-flood": "true",
    "comment": "defconf",
    "debug-info": " prio 0x0 num 2\n
role:Root (1) learn 1 forward 1 infoIs Rcvd edge 0 sendRSTP 1\n
proposing 0 agreed 0 agree 1 synced 1 isolate 0 newInfo 0\n
migration:SENSING tc:ACTIVE\n
ptimes: Msg:0 Max: 5120 FD: 3840 HT: 512\n
pprio: RBI: 8000:744D288FCA7D RPC: 0 BI: 8000:744D288FCA7D tP: 0x1 rP: 0x2\n
dtimes: Msg:256 Max: 5120 FD: 3840 HT: 512\n
dprio: RBI: 8000:744D288FCA7D RPC: 10 BI: 8000:C4AD3407AD79 tP: 0x2 rP: 0x2\n
",
    "designated-bridge": "0x8000.74:4D:28:8F:CA:7D",
    "designated-cost": "0",
    "designated-port-number": "1",
    "disabled": "false",
    "dynamic": "false",
    "edge": "auto",
    "edge-port": "false",
    "edge-port-discovery": "true",
    "external-fdb-status": "false",
    "fast-leave": "false",
    "forwarding": "true",
    "frame-types": "admit-only-vlan-tagged",
    "horizon": "none",
    "hw": "true",
    "hw-offload": "false",
    "inactive": "false",
    "ingress-filtering": "true",
    "interface": "ether2",
    "internal-path-cost": "10",
    "learn": "auto",
    "learning": "true",
    "multicast-router": "temporary-query",
    "path-cost": "10",
    "point-to-point": "auto",
    "point-to-point-port": "true",
    "port-number": "2",
    "priority": "0x80",
    "pvid": "1",
    "restricted-role": "false",
    "restricted-tcn": "false",
    "role": "root-port",
    "root-path-cost": "10",
    "sending-rstp": "true",
    "status": "in-bridge",
    "tag-stacking": "false",
    "trusted": "false",
    "unknown-multicast-flood": "true",
    "unknown-unicast-flood": "true"
  },
  {...}
*/

// ResourceInterfaceBridgePort https://wiki.mikrotik.com/wiki/Manual:Interface/Bridge#Port_Settings
func ResourceInterfaceBridgePort() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/bridge/port"),
		MetaId:           PropId(Id),

		"nextid": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"auto_isolate": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
			Description: "When enabled, prevents a port moving from discarding into forwarding state if no BPDUs " +
				"are received from the neighboring bridge. The port will change into a forwarding state only when " +
				"a BPDU is received. This property only has an effect when protocol-mode is set to rstp or mstp and " +
				"edge is set to no.",
		},
		"bpdu_guard": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "This property has no effect when protocol-mode is set to none.",
		},
		"bridge": {
			Type:     schema.TypeString,
			Required: true,
		},
		"broadcast_flood": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
			Description: "When enabled, bridge floods broadcast traffic to all bridge egress ports. " +
				"When disabled, drops broadcast traffic on egress ports. ",
		},
		KeyComment: PropCommentRw,
		"debug_info": {
			Type:     schema.TypeString,
			Computed: true,
		},
		KeyDisabled: PropDisabledRw,
		KeyDynamic:  PropDynamicRo,
		"edge": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "auto",
			Description: "Set port as edge port or non-edge port, or enable edge discovery. " +
				"Edge ports are connected to a LAN that has no other bridges attached. ",
			ValidateFunc: validation.StringInSlice([]string{"auto", "no", "no-discover", "yes", "yes-discover"}, false),
		},
		"edge_port": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether port is an edge port or not.",
		},
		"edge_port_discovery": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether port is set to automatically detect edge ports.",
		},
		// Where external-fdb (auto | no | yes; Default: auto) ???
		"external_fdb_status": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether registration table is used instead of forwarding data base.",
		},
		"fast_leave": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Enables IGMP Fast leave feature on the port.",
		},
		"forwarding": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Shows if the port is not blocked by (R/M)STP.",
		},
		"frame_types": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "admit-all",
			Description: "Specifies allowed ingress frame types on a bridge port. " +
				"This property only has effect when vlan-filtering is set to yes.",
			ValidateFunc: validation.StringInSlice([]string{"admit-all",
				"admit-only-untagged-and-priority-tagged",
				"admit-only-vlan-tagged"}, false),
		},
		"horizon": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "none",
			Description: "Use split horizon bridging to prevent bridging loops. Set the same value for group of ports, " +
				"to prevent them from sending data to ports with the same horizon value. Split horizon is a software " +
				"feature that disables hardware offloading. This value is integer '0'..'429496729' or 'none'.",
		},
		"hw": {
			Type:     schema.TypeBool,
			Computed: true,
			Optional: true,
		},
		"hw_offload": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"hw_offload_group": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Switch chip used by the port.",
		},
		"inactive": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"ingress_filtering": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
			Description: "Enables or disables VLAN ingress filtering, which checks if the ingress port is a member of " +
				"the received VLAN ID in the bridge VLAN table. Should be used with frame-types to specify if the " +
				"ingress traffic should be tagged or untagged. This property only has effect when vlan-filtering " +
				"is set to yes.",
		},
		KeyInterface: PropInterfaceRw,
		"internal_path_cost": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  10,
			Description: "Path cost to the interface for MSTI0 inside a region. This property only has effect when " +
				"protocol-mode is set to mstp.",
		},
		"learn": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "auto",
			Description:  "Changes MAC learning behaviour on a bridge port ",
			ValidateFunc: ValidationAutoYesNo,
		},
		"learning": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Shows whether the port is capable of learning MAC addresses.",
		},
		"multicast_router": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "temporary-query",
			Description: "Changes the state of a bridge port whether IGMP membership reports are going to be " +
				"forwarded to this port.",
			ValidateFunc: validation.StringInSlice([]string{"disabled", "permanent", "temporary-query"}, false),
		},
		// This field has a string value because on the x86 architecture there is no good way to validate
		// values up to 4294967295. And in this case, an overflow occurs with an errors:
		// "Cannot use 4294967295 (untyped int constant) as int value in argument to validation.IntBetween (overflows)"
		// or "Attribute must be a whole number, got 4.294967295e+09".
		"path_cost": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "10",
			Description: `Path cost to the interface, used by STP to determine the "best" path, used by MSTP to` +
				`determine "best" path between regions. This property has no effect when protocol-mode is set to none.`,
		},
		"point_to_point": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "auto",
			Description: "Specifies if a bridge port is connected to a bridge using a point-to-point link for faster " +
				"convergence in case of failure. This property has no effect when protocol-mode is set to none.",
			ValidateFunc: ValidationAutoYesNo,
		},
		"point_to_point_port": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the port is connected to a bridge port using full-duplex (true) or half-duplex (false).",
		},
		"port_number": {
			Type:     schema.TypeInt,
			Computed: true,
			Description: "Port number will be assigned in the order that ports got added to the bridge, " +
				"but this is only true until reboot. After reboot internal numbering will be used.",
		},
		"priority": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  128,
			Description: "The priority of the interface, used by STP to determine the root port, " +
				"used by MSTP to determine root port between regions.",
			ValidateFunc:     validation.IntBetween(0, 240),
			DiffSuppressFunc: HexEqual,
		},
		"pvid": {
			Type:     schema.TypeInt,
			Required: true,
			Description: "ort VLAN ID (pvid) specifies which VLAN the untagged ingress traffic is assigned to. " +
				"This property only has effect when vlan-filtering is set to yes.",
			ValidateFunc: validation.IntBetween(1, 4096),
		},
		"restricted_role": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
			Description: "Enable the restricted role on a port, used by STP to forbid a port becoming a root port. " +
				"This property only has effect when protocol-mode is set to mstp.",
		},
		"restricted_tcn": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
			Description: "Disable topology change notification (TCN) sending on a port, used by STP to forbid network " +
				"topology changes to propagate. This property only has effect when protocol-mode is set to mstp.",
		},
		"role": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "(R/M)STP algorithm assigned role of the port",
		},
		"sending_rstp": {
			Type:     schema.TypeString,
			Computed: true,
			Description: "Whether the port is sending RSTP or MSTP BPDU types. A port will transit to STP type " +
				"when RSTP/MSTP enabled port receives a STP BPDU",
		},
		"status": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Port status ('in-bridge' - port is enabled).",
		},
		"tag_stacking": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
			Description: "Forces all packets to be treated as untagged packets. Packets on ingress port will be tagged " +
				"with another VLAN tag regardless if a VLAN tag already exists, packets will be tagged with a VLAN ID " +
				"that matches the pvid value and will use EtherType that is specified in ether-type. " +
				"This property only has effect when vlan-filtering is set to yes.",
		},
		"trusted": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
			Description: "When enabled, it allows to forward DHCP packets towards DHCP server through this port. " +
				"Mainly used to limit unauthorized servers to provide malicious information for users. " +
				"This property only has effect when dhcp-snooping is set to yes.",
		},
		"unknown_multicast_flood": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "When enabled, bridge floods unknown multicast traffic to all bridge egress ports.",
		},
		"unknown_unicast_flood": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "When enabled, bridge floods unknown unicast traffic to all bridge egress ports.",
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
