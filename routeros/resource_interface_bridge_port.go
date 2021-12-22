package routeros

import (
	"log"
	"strconv"

	roscl "github.com/gnewbury1/terraform-provider-routeros/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceBridgePort() *schema.Resource {
	return &schema.Resource{
		Create: resourceInterfaceBridgePortCreate,
		Read:   resourceInterfaceBridgePortRead,
		Update: resourceInterfaceBridgePortUpdate,
		Delete: resourceInterfaceBridgePortDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"nextid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"auto_isolate": {
				Type:     schema.TypeBool,
				Computed: true,
				Optional: true,
			},
			"bpdu_guard": {
				Type:     schema.TypeBool,
				Computed: true,
				Optional: true,
			},
			"bridge": {
				Type:     schema.TypeString,
				Required: true,
			},
			"broadcast_flood": {
				Type:     schema.TypeBool,
				Computed: true,
				Optional: true,
			},
			"comment": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  "",
			},
			"debug_info": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"edge": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"edge_port": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"edge_port_discovery": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"external_fdb_status": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"fast_leave": {
				Type:     schema.TypeBool,
				Computed: true,
				Optional: true,
			},
			"forwarding": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"frame_types": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"horizon": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
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
				Type:     schema.TypeString,
				Computed: true,
			},
			"inactive": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ingress_filtering": {
				Type:     schema.TypeBool,
				Computed: true,
				Optional: true,
			},
			"interface": {
				Type:     schema.TypeString,
				Required: true,
			},
			"internal_path_cost": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"learn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"learning": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"multicast_router": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"path_cost": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"point_to_point": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"point_to_point_port": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"port_number": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"priority": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pvid": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"restricted_role": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"restricted_tcn": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"role": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sending_rstp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tag_stacking": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"trusted": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"unknown_multicast_flood": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"unknown_unicast_flood": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func resourceInterfaceBridgePortCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)

	bridge_port := new(roscl.InterfaceBridgePort)
	bridge_port.AutoIsolate = strconv.FormatBool(d.Get("auto_isolate").(bool))
	bridge_port.BpduGuard = strconv.FormatBool(d.Get("bpdu_guard").(bool))
	bridge_port.Bridge = d.Get("bridge").(string)
	bridge_port.BroadcastFlood = strconv.FormatBool(d.Get("broadcast_flood").(bool))
	bridge_port.Disabled = strconv.FormatBool(d.Get("disabled").(bool))
	bridge_port.Edge = d.Get("edge").(string)
	bridge_port.FastLeave = strconv.FormatBool(d.Get("fast_leave").(bool))
	bridge_port.FrameTypes = d.Get("frame_types").(string)
	bridge_port.Horizon = d.Get("horizon").(string)
	bridge_port.Hw = strconv.FormatBool(d.Get("hw").(bool))
	bridge_port.IngressFiltering = strconv.FormatBool(d.Get("ingress_filtering").(bool))
	bridge_port.Interface = d.Get("interface").(string)
	bridge_port.InternalPathCost = strconv.Itoa(d.Get("internal_path_cost").(int))
	bridge_port.Learn = d.Get("learn").(string)
	bridge_port.PathCost = strconv.Itoa(d.Get("path_cost").(int))
	bridge_port.PointToPoint = d.Get("point_to_point").(string)
	bridge_port.Priority = d.Get("priority").(string)
	bridge_port.Pvid = strconv.Itoa(d.Get("pvid").(int))
	bridge_port.RestrictedRole = strconv.FormatBool(d.Get("restricted_role").(bool))
	bridge_port.RestrictedTcn = strconv.FormatBool(d.Get("restricted_tcn").(bool))
	bridge_port.TagStacking = strconv.FormatBool(d.Get("tag_stacking").(bool))
	bridge_port.Trusted = strconv.FormatBool(d.Get("trusted").(bool))

	res, err := c.CreateInterfaceBridgePort(bridge_port)
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PUT request to the API")
		log.Fatal(err.Error())
		return err
	}

	auto_isolate, _ := strconv.ParseBool(res.AutoIsolate)
	bpdu_guard, _ := strconv.ParseBool(res.BpduGuard)
	broadcast_flood, _ := strconv.ParseBool(res.BroadcastFlood)
	disabled, _ := strconv.ParseBool(res.Disabled)
	edge_port, _ := strconv.ParseBool(res.EdgePort)
	edge_port_discovery, _ := strconv.ParseBool(res.EdgePortDiscovery)
	external_fdb_status, _ := strconv.ParseBool(res.ExternalFdbStatus)
	fast_leave, _ := strconv.ParseBool(res.FastLeave)
	forwarding, _ := strconv.ParseBool(res.Forwarding)
	hw, _ := strconv.ParseBool(res.Hw)
	hw_offload, _ := strconv.ParseBool(res.HwOffload)
	inactive, _ := strconv.ParseBool(res.Inactive)
	ingress_filtering, _ := strconv.ParseBool(res.IngressFiltering)
	internal_path_cost, _ := strconv.Atoi(res.InternalPathCost)
	learning, _ := strconv.ParseBool(res.Learning)
	path_cost, _ := strconv.Atoi(res.PathCost)
	point_to_point_port, _ := strconv.ParseBool(res.PointToPointPort)
	port_number, _ := strconv.Atoi(res.PortNumber)
	pvid, _ := strconv.Atoi(res.Pvid)
	restricted_role, _ := strconv.ParseBool(res.RestrictedRole)
	restricted_tcn, _ := strconv.ParseBool(res.RestrictedTcn)
	tag_stacking, _ := strconv.ParseBool(res.TagStacking)
	trusted, _ := strconv.ParseBool(res.Trusted)
	unknown_multicast_flood, _ := strconv.ParseBool(res.UnknownMulticastFlood)
	unknown_unicast_flood, _ := strconv.ParseBool(res.UnknownUnicastFlood)

	d.SetId(res.ID)
	d.Set("nextid", res.Nextid)
	d.Set("auto_isolate", auto_isolate)
	d.Set("bpdu_guard", bpdu_guard)
	d.Set("bridge", res.Bridge)
	d.Set("broadcast_flood", broadcast_flood)
	d.Set("debug_info", res.DebugInfo)
	d.Set("disabled", disabled)
	d.Set("edge", res.Edge)
	d.Set("edge_port", edge_port)
	d.Set("edge_port_discovery", edge_port_discovery)
	d.Set("external_fdb_status", external_fdb_status)
	d.Set("fast_leave", fast_leave)
	d.Set("forwarding", forwarding)
	d.Set("frame_types", res.FrameTypes)
	d.Set("horizon", res.Horizon)
	d.Set("hw", hw)
	d.Set("hw_offload", hw_offload)
	d.Set("hw_offload_group", res.HwOffloadGroup)
	d.Set("inactive", inactive)
	d.Set("ingress_filtering", ingress_filtering)
	d.Set("interface", res.Interface)
	d.Set("internal_path_cost", internal_path_cost)
	d.Set("learn", res.Learn)
	d.Set("learning", learning)
	d.Set("multicast_router", res.MulticastRouter)
	d.Set("path_cost", path_cost)
	d.Set("point_to_point", res.PointToPoint)
	d.Set("point_to_point_port", point_to_point_port)
	d.Set("port_number", port_number)
	d.Set("priority", res.Priority)
	d.Set("pvid", pvid)
	d.Set("restricted_role", restricted_role)
	d.Set("restricted_tcn", restricted_tcn)
	d.Set("role", res.Role)
	d.Set("sending_rstp", res.SendingRstp)
	d.Set("status", res.Status)
	d.Set("tag_stacking", tag_stacking)
	d.Set("trusted", trusted)
	d.Set("unknown_multicast_flood", unknown_multicast_flood)
	d.Set("unknown_unicast_flood", unknown_unicast_flood)

	return nil
}

func resourceInterfaceBridgePortRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	res, err := c.ReadInterfaceBridgePort(d.Id())

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a GET request to the API")
		log.Fatal(err.Error())
		return err
	}

	auto_isolate, _ := strconv.ParseBool(res.AutoIsolate)
	bpdu_guard, _ := strconv.ParseBool(res.BpduGuard)
	broadcast_flood, _ := strconv.ParseBool(res.BroadcastFlood)
	disabled, _ := strconv.ParseBool(res.Disabled)
	edge_port, _ := strconv.ParseBool(res.EdgePort)
	edge_port_discovery, _ := strconv.ParseBool(res.EdgePortDiscovery)
	external_fdb_status, _ := strconv.ParseBool(res.ExternalFdbStatus)
	fast_leave, _ := strconv.ParseBool(res.FastLeave)
	forwarding, _ := strconv.ParseBool(res.Forwarding)
	hw, _ := strconv.ParseBool(res.Hw)
	hw_offload, _ := strconv.ParseBool(res.HwOffload)
	inactive, _ := strconv.ParseBool(res.Inactive)
	ingress_filtering, _ := strconv.ParseBool(res.IngressFiltering)
	internal_path_cost, _ := strconv.Atoi(res.InternalPathCost)
	learning, _ := strconv.ParseBool(res.Learning)
	path_cost, _ := strconv.Atoi(res.PathCost)
	point_to_point_port, _ := strconv.ParseBool(res.PointToPointPort)
	port_number, _ := strconv.Atoi(res.PortNumber)
	pvid, _ := strconv.Atoi(res.Pvid)
	restricted_role, _ := strconv.ParseBool(res.RestrictedRole)
	restricted_tcn, _ := strconv.ParseBool(res.RestrictedTcn)
	tag_stacking, _ := strconv.ParseBool(res.TagStacking)
	trusted, _ := strconv.ParseBool(res.Trusted)
	unknown_multicast_flood, _ := strconv.ParseBool(res.UnknownMulticastFlood)
	unknown_unicast_flood, _ := strconv.ParseBool(res.UnknownUnicastFlood)

	d.SetId(res.ID)
	d.Set("nextid", res.Nextid)
	d.Set("auto_isolate", auto_isolate)
	d.Set("bpdu_guard", bpdu_guard)
	d.Set("bridge", res.Bridge)
	d.Set("broadcast_flood", broadcast_flood)
	d.Set("debug_info", res.DebugInfo)
	d.Set("disabled", disabled)
	d.Set("edge", res.Edge)
	d.Set("edge_port", edge_port)
	d.Set("edge_port_discovery", edge_port_discovery)
	d.Set("external_fdb_status", external_fdb_status)
	d.Set("fast_leave", fast_leave)
	d.Set("forwarding", forwarding)
	d.Set("frame_types", res.FrameTypes)
	d.Set("horizon", res.Horizon)
	d.Set("hw", hw)
	d.Set("hw_offload", hw_offload)
	d.Set("hw_offload_group", res.HwOffloadGroup)
	d.Set("inactive", inactive)
	d.Set("ingress_filtering", ingress_filtering)
	d.Set("interface", res.Interface)
	d.Set("internal_path_cost", internal_path_cost)
	d.Set("learn", res.Learn)
	d.Set("learning", learning)
	d.Set("multicast_router", res.MulticastRouter)
	d.Set("path_cost", path_cost)
	d.Set("point_to_point", res.PointToPoint)
	d.Set("point_to_point_port", point_to_point_port)
	d.Set("port_number", port_number)
	d.Set("priority", res.Priority)
	d.Set("pvid", pvid)
	d.Set("restricted_role", restricted_role)
	d.Set("restricted_tcn", restricted_tcn)
	d.Set("role", res.Role)
	d.Set("sending_rstp", res.SendingRstp)
	d.Set("status", res.Status)
	d.Set("tag_stacking", tag_stacking)
	d.Set("trusted", trusted)
	d.Set("unknown_multicast_flood", unknown_multicast_flood)
	d.Set("unknown_unicast_flood", unknown_unicast_flood)

	return nil

}

func resourceInterfaceBridgePortUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)

	bridge_port := new(roscl.InterfaceBridgePort)
	bridge_port.AutoIsolate = strconv.FormatBool(d.Get("auto_isolate").(bool))
	bridge_port.BpduGuard = strconv.FormatBool(d.Get("bpdu_guard").(bool))
	bridge_port.Bridge = d.Get("bridge").(string)
	bridge_port.BroadcastFlood = strconv.FormatBool(d.Get("broadcast_flood").(bool))
	bridge_port.Disabled = strconv.FormatBool(d.Get("disabled").(bool))
	bridge_port.Edge = d.Get("edge").(string)
	bridge_port.FastLeave = strconv.FormatBool(d.Get("fast_leave").(bool))
	bridge_port.FrameTypes = d.Get("frame_types").(string)
	bridge_port.Horizon = d.Get("horizon").(string)
	bridge_port.Hw = strconv.FormatBool(d.Get("hw").(bool))
	bridge_port.IngressFiltering = strconv.FormatBool(d.Get("ingress_filtering").(bool))
	bridge_port.Interface = d.Get("interface").(string)
	bridge_port.InternalPathCost = strconv.Itoa(d.Get("internal_path_cost").(int))
	bridge_port.Learn = d.Get("learn").(string)
	bridge_port.PathCost = strconv.Itoa(d.Get("path_cost").(int))
	bridge_port.PointToPoint = d.Get("point_to_point").(string)
	bridge_port.Priority = d.Get("priority").(string)
	bridge_port.Pvid = strconv.Itoa(d.Get("pvid").(int))
	bridge_port.RestrictedRole = strconv.FormatBool(d.Get("restricted_role").(bool))
	bridge_port.RestrictedTcn = strconv.FormatBool(d.Get("restricted_tcn").(bool))
	bridge_port.TagStacking = strconv.FormatBool(d.Get("tag_stacking").(bool))
	bridge_port.Trusted = strconv.FormatBool(d.Get("trusted").(bool))

	res, err := c.UpdateInterfaceBridgePort(d.Id(), bridge_port)

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PUT request to the API")
		log.Fatal(err.Error())
		return err
	}

	auto_isolate, _ := strconv.ParseBool(res.AutoIsolate)
	bpdu_guard, _ := strconv.ParseBool(res.BpduGuard)
	broadcast_flood, _ := strconv.ParseBool(res.BroadcastFlood)
	disabled, _ := strconv.ParseBool(res.Disabled)
	edge_port, _ := strconv.ParseBool(res.EdgePort)
	edge_port_discovery, _ := strconv.ParseBool(res.EdgePortDiscovery)
	external_fdb_status, _ := strconv.ParseBool(res.ExternalFdbStatus)
	fast_leave, _ := strconv.ParseBool(res.FastLeave)
	forwarding, _ := strconv.ParseBool(res.Forwarding)
	hw, _ := strconv.ParseBool(res.Hw)
	hw_offload, _ := strconv.ParseBool(res.HwOffload)
	inactive, _ := strconv.ParseBool(res.Inactive)
	ingress_filtering, _ := strconv.ParseBool(res.IngressFiltering)
	internal_path_cost, _ := strconv.Atoi(res.InternalPathCost)
	learning, _ := strconv.ParseBool(res.Learning)
	path_cost, _ := strconv.Atoi(res.PathCost)
	point_to_point_port, _ := strconv.ParseBool(res.PointToPointPort)
	port_number, _ := strconv.Atoi(res.PortNumber)
	pvid, _ := strconv.Atoi(res.Pvid)
	restricted_role, _ := strconv.ParseBool(res.RestrictedRole)
	restricted_tcn, _ := strconv.ParseBool(res.RestrictedTcn)
	tag_stacking, _ := strconv.ParseBool(res.TagStacking)
	trusted, _ := strconv.ParseBool(res.Trusted)
	unknown_multicast_flood, _ := strconv.ParseBool(res.UnknownMulticastFlood)
	unknown_unicast_flood, _ := strconv.ParseBool(res.UnknownUnicastFlood)

	d.SetId(res.ID)
	d.Set("nextid", res.Nextid)
	d.Set("auto_isolate", auto_isolate)
	d.Set("bpdu_guard", bpdu_guard)
	d.Set("bridge", res.Bridge)
	d.Set("broadcast_flood", broadcast_flood)
	d.Set("debug_info", res.DebugInfo)
	d.Set("disabled", disabled)
	d.Set("edge", res.Edge)
	d.Set("edge_port", edge_port)
	d.Set("edge_port_discovery", edge_port_discovery)
	d.Set("external_fdb_status", external_fdb_status)
	d.Set("fast_leave", fast_leave)
	d.Set("forwarding", forwarding)
	d.Set("frame_types", res.FrameTypes)
	d.Set("horizon", res.Horizon)
	d.Set("hw", hw)
	d.Set("hw_offload", hw_offload)
	d.Set("hw_offload_group", res.HwOffloadGroup)
	d.Set("inactive", inactive)
	d.Set("ingress_filtering", ingress_filtering)
	d.Set("interface", res.Interface)
	d.Set("internal_path_cost", internal_path_cost)
	d.Set("learn", res.Learn)
	d.Set("learning", learning)
	d.Set("multicast_router", res.MulticastRouter)
	d.Set("path_cost", path_cost)
	d.Set("point_to_point", res.PointToPoint)
	d.Set("point_to_point_port", point_to_point_port)
	d.Set("port_number", port_number)
	d.Set("priority", res.Priority)
	d.Set("pvid", pvid)
	d.Set("restricted_role", restricted_role)
	d.Set("restricted_tcn", restricted_tcn)
	d.Set("role", res.Role)
	d.Set("sending_rstp", res.SendingRstp)
	d.Set("status", res.Status)
	d.Set("tag_stacking", tag_stacking)
	d.Set("trusted", trusted)
	d.Set("unknown_multicast_flood", unknown_multicast_flood)
	d.Set("unknown_unicast_flood", unknown_unicast_flood)

	return nil
}

func resourceInterfaceBridgePortDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	bridge_port, _ := c.ReadInterfaceBridgePort(d.Id())
	err := c.DeleteInterfaceBridgePort(bridge_port)
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a DELETE request to the API")
		log.Fatal(err.Error())
		return err
	}
	d.SetId("")
	return nil
}
