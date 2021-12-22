package routeros

import (
	"log"
	"strconv"

	roscl "github.com/gnewbury1/terraform-provider-routeros/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceBridge() *schema.Resource {
	return &schema.Resource{
		Create: resourceInterfaceBridgeCreate,
		Read:   resourceInterfaceBridgeRead,
		Update: resourceInterfaceBridgeUpdate,
		Delete: resourceInterfaceBridgeDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"actual_mtu": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"admin_mac": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ageing_time": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"arp": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"arp_timeout": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"auto_mac": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_snooping": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ether_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fast_forward": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"forward_delay": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"frame_types": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"igmp_snooping": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ingress_filtering": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"l2mtu": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"mac_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_message_age": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mtu": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"priority": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protocol_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pvid": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"running": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"transmit_hold_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  6,
			},
			"vlan_filtering": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceInterfaceBridgeCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)

	bridge := new(roscl.InterfaceBridge)
	bridge.AdminMac = d.Get("admin_mac").(string)
	bridge.AgeingTime = d.Get("ageing_time").(string)
	bridge.Arp = d.Get("arp").(string)
	bridge.Comment = d.Get("comment").(string)
	bridge.DhcpSnooping = strconv.FormatBool(d.Get("dhcp_snooping").(bool))
	bridge.Disabled = strconv.FormatBool(d.Get("disabled").(bool))
	bridge.FastForward = strconv.FormatBool(d.Get("fast_forward").(bool))
	bridge.ForwardDelay = d.Get("forward_delay").(string)
	bridge.IgmpSnooping = strconv.FormatBool(d.Get("igmp_snooping").(bool))
	bridge.IngressFiltering = strconv.FormatBool(d.Get("ingress_filtering").(bool))
	bridge.MaxMessageAge = d.Get("max_message_age").(string)
	bridge.Mtu = d.Get("mtu").(string)
	bridge.Name = d.Get("name").(string)
	bridge.Priority = d.Get("priority").(string)
	bridge.ProtocolMode = d.Get("protocol_mode").(string)
	bridge.TransmitHoldCount = strconv.Itoa(d.Get("transmit_hold_count").(int))
	bridge.VlanFiltering = strconv.FormatBool(d.Get("vlan_filtering").(bool))

	if d.Get("vlan_filtering").(bool) {
		bridge.Pvid = strconv.Itoa(d.Get("pvid").(int))
		bridge.FrameTypes = d.Get("frame_types").(string)
	}

	res, err := c.CreateInterfaceBridge(bridge)
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PUT request to the API")
		log.Fatal(err.Error())
		return err
	}

	actual_mtu, _ := strconv.Atoi(res.ActualMtu)
	auto_mac, _ := strconv.ParseBool(res.AutoMac)
	dhcp_snooping, _ := strconv.ParseBool(res.DhcpSnooping)
	disabled, _ := strconv.ParseBool(res.Disabled)
	fast_forward, _ := strconv.ParseBool(res.FastForward)
	igmp_snooping, _ := strconv.ParseBool(res.IgmpSnooping)
	ingress_filtering, _ := strconv.ParseBool(res.IngressFiltering)
	l2mtu, _ := strconv.Atoi(res.L2Mtu)
	pvid, _ := strconv.Atoi(res.Pvid)
	running, _ := strconv.ParseBool(res.Running)
	transmit_hold_count, _ := strconv.Atoi(res.TransmitHoldCount)
	vlan_filtering, _ := strconv.ParseBool(res.VlanFiltering)

	d.SetId(res.ID)
	d.Set("actual_mtu", actual_mtu)
	d.Set("admin_mac", res.AdminMac)
	d.Set("ageing_time", res.AgeingTime)
	d.Set("arp", res.Arp)
	d.Set("arp_timeout", res.ArpTimeout)
	d.Set("auto_mac", auto_mac)
	d.Set("comment", res.Comment)
	d.Set("dhcp_snooping", dhcp_snooping)
	d.Set("disabled", disabled)
	d.Set("ether_type", res.EtherType)
	d.Set("fast_forward", fast_forward)
	d.Set("forward_delay", res.ForwardDelay)
	d.Set("frame_types", res.FrameTypes)
	d.Set("igmp_snooping", igmp_snooping)
	d.Set("ingress_filtering", ingress_filtering)
	d.Set("l2mtu", l2mtu)
	d.Set("mac_address", res.MacAddress)
	d.Set("max_message_age", res.MaxMessageAge)
	d.Set("mtu", res.Mtu)
	d.Set("name", res.Name)
	d.Set("priority", res.Priority)
	d.Set("protocol_mode", res.ProtocolMode)
	d.Set("pvid", pvid)
	d.Set("running", running)
	d.Set("transmit_hold_count", transmit_hold_count)
	d.Set("vlan_filtering", vlan_filtering)

	return nil
}

func resourceInterfaceBridgeRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	res, err := c.ReadInterfaceBridge(d.Id())

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a GET request to the API")
		log.Fatal(err.Error())
		return err
	}

	actual_mtu, _ := strconv.Atoi(res.ActualMtu)
	auto_mac, _ := strconv.ParseBool(res.AutoMac)
	dhcp_snooping, _ := strconv.ParseBool(res.DhcpSnooping)
	disabled, _ := strconv.ParseBool(res.Disabled)
	fast_forward, _ := strconv.ParseBool(res.FastForward)
	igmp_snooping, _ := strconv.ParseBool(res.IgmpSnooping)
	ingress_filtering, _ := strconv.ParseBool(res.IngressFiltering)
	l2mtu, _ := strconv.Atoi(res.L2Mtu)
	pvid, _ := strconv.Atoi(res.Pvid)
	running, _ := strconv.ParseBool(res.Running)
	transmit_hold_count, _ := strconv.Atoi(res.TransmitHoldCount)
	vlan_filtering, _ := strconv.ParseBool(res.VlanFiltering)

	d.SetId(res.ID)
	d.Set("actual_mtu", actual_mtu)
	d.Set("admin_mac", res.AdminMac)
	d.Set("ageing_time", res.AgeingTime)
	d.Set("arp", res.Arp)
	d.Set("arp_timeout", res.ArpTimeout)
	d.Set("auto_mac", auto_mac)
	d.Set("comment", res.Comment)
	d.Set("dhcp_snooping", dhcp_snooping)
	d.Set("disabled", disabled)
	d.Set("ether_type", res.EtherType)
	d.Set("fast_forward", fast_forward)
	d.Set("forward_delay", res.ForwardDelay)
	d.Set("frame_types", res.FrameTypes)
	d.Set("igmp_snooping", igmp_snooping)
	d.Set("ingress_filtering", ingress_filtering)
	d.Set("l2mtu", l2mtu)
	d.Set("mac_address", res.MacAddress)
	d.Set("max_message_age", res.MaxMessageAge)
	d.Set("mtu", res.Mtu)
	d.Set("name", res.Name)
	d.Set("priority", res.Priority)
	d.Set("protocol_mode", res.ProtocolMode)
	d.Set("pvid", pvid)
	d.Set("running", running)
	d.Set("transmit_hold_count", transmit_hold_count)
	d.Set("vlan_filtering", vlan_filtering)

	return nil

}

func resourceInterfaceBridgeUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)

	bridge := new(roscl.InterfaceBridge)
	bridge.AdminMac = d.Get("admin_mac").(string)
	bridge.AgeingTime = d.Get("ageing_time").(string)
	bridge.Arp = d.Get("arp").(string)
	bridge.Comment = d.Get("comment").(string)
	bridge.DhcpSnooping = strconv.FormatBool(d.Get("dhcp_snooping").(bool))
	bridge.Disabled = strconv.FormatBool(d.Get("disabled").(bool))
	bridge.FastForward = strconv.FormatBool(d.Get("fast_forward").(bool))
	bridge.ForwardDelay = d.Get("forward_delay").(string)
	bridge.IgmpSnooping = strconv.FormatBool(d.Get("igmp_snooping").(bool))
	bridge.IngressFiltering = strconv.FormatBool(d.Get("ingress_filtering").(bool))
	bridge.MaxMessageAge = d.Get("max_message_age").(string)
	bridge.Mtu = d.Get("mtu").(string)
	bridge.Name = d.Get("name").(string)
	bridge.Priority = d.Get("priority").(string)
	bridge.ProtocolMode = d.Get("protocol_mode").(string)
	bridge.TransmitHoldCount = strconv.Itoa(d.Get("transmit_hold_count").(int))
	bridge.VlanFiltering = strconv.FormatBool(d.Get("vlan_filtering").(bool))

	if d.Get("vlan_filtering").(bool) {
		bridge.Pvid = strconv.Itoa(d.Get("pvid").(int))
		bridge.FrameTypes = d.Get("frame_types").(string)
	}

	res, err := c.UpdateInterfaceBridge(d.Id(), bridge)

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PATCH request to the API")
		log.Fatal(err.Error())
		return err
	}

	actual_mtu, _ := strconv.Atoi(res.ActualMtu)
	auto_mac, _ := strconv.ParseBool(res.AutoMac)
	dhcp_snooping, _ := strconv.ParseBool(res.DhcpSnooping)
	disabled, _ := strconv.ParseBool(res.Disabled)
	fast_forward, _ := strconv.ParseBool(res.FastForward)
	igmp_snooping, _ := strconv.ParseBool(res.IgmpSnooping)
	ingress_filtering, _ := strconv.ParseBool(res.IngressFiltering)
	l2mtu, _ := strconv.Atoi(res.L2Mtu)
	pvid, _ := strconv.Atoi(res.Pvid)
	running, _ := strconv.ParseBool(res.Running)
	transmit_hold_count, _ := strconv.Atoi(res.TransmitHoldCount)
	vlan_filtering, _ := strconv.ParseBool(res.VlanFiltering)

	d.SetId(res.ID)
	d.Set("actual_mtu", actual_mtu)
	d.Set("admin_mac", res.AdminMac)
	d.Set("ageing_time", res.AgeingTime)
	d.Set("arp", res.Arp)
	d.Set("arp_timeout", res.ArpTimeout)
	d.Set("auto_mac", auto_mac)
	d.Set("comment", res.Comment)
	d.Set("dhcp_snooping", dhcp_snooping)
	d.Set("disabled", disabled)
	d.Set("ether_type", res.EtherType)
	d.Set("fast_forward", fast_forward)
	d.Set("forward_delay", res.ForwardDelay)
	d.Set("frame_types", res.FrameTypes)
	d.Set("igmp_snooping", igmp_snooping)
	d.Set("ingress_filtering", ingress_filtering)
	d.Set("l2mtu", l2mtu)
	d.Set("mac_address", res.MacAddress)
	d.Set("max_message_age", res.MaxMessageAge)
	d.Set("mtu", res.Mtu)
	d.Set("name", res.Name)
	d.Set("priority", res.Priority)
	d.Set("protocol_mode", res.ProtocolMode)
	d.Set("pvid", pvid)
	d.Set("running", running)
	d.Set("transmit_hold_count", transmit_hold_count)
	d.Set("vlan_filtering", vlan_filtering)

	return nil
}

func resourceInterfaceBridgeDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	bridge, _ := c.ReadInterfaceBridge(d.Id())
	err := c.DeleteInterfaceBridge(bridge)
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a DELETE request to the API")
		log.Fatal(err.Error())
		return err
	}
	d.SetId("")
	return nil
}
