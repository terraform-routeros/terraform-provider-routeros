package routeros

import (
	"fmt"
	"strconv"

	roscl "github.com/gnewbury1/terraform-provider-routeros/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIPFirewallFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceIPFirewallFilterCreate,
		Read:   resourceIPFirewallFilterRead,
		Update: resourceIPFirewallFilterUpdate,
		Delete: resourceIPFirewallFilterDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeString,
				Required: true,
			},
			"address-list-timeout": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bytes": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"chain": {
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"connection_bytes": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"connection_limit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"conneciton_mark": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"connection_nat_state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"connection_rate": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"connection_state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"connection_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"content": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"dscp": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dst_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dst_address_list": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dst_address_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dst_limit": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dst_port": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"fragment": {
				Type:     schema.TypeBool, //This is a yes/no bool rather than a true/false bool
				Optional: true,
				Computed: true,
			},
			"hotspot": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icmp_options": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"in_bridge_port": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"in_bridge_port_list": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"in_interface": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"in_interface_list": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ingress_priority": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"invalid": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ipsec_policy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_options": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"jump_target": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"layer7_protocol": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"limit": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"log_prefix": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nth": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"out_bridge_port": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"out_bridge_port_list": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"out_interface": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"out_interface_list": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"packet_mark": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"packet_size": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"per_connection_classifier": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"psd": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"random": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"reject_with": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"routing_table": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_address_list": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_address_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_port": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_mac_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_flags": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_mss": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tls_host": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ttl": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceIPFirewallFilterCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)

	firewall_filter := new(roscl.IPFirewallFilter)
	firewall_filter.Action = d.Get("action").(string)
	firewall_filter.AddressListTimeout = d.Get("address_list_timeout").(string)
	firewall_filter.Chain = d.Get("chain").(string)
	firewall_filter.Comment = d.Get("comment").(string)
	firewall_filter.ConnectionBytes = d.Get("connection_bytes").(string)
	firewall_filter.ConnectionLimit = d.Get("connection_limit").(string)
	firewall_filter.ConnectionMark = d.Get("connection_mark").(string)
	firewall_filter.ConnectionNatState = d.Get("connection_nat_state").(string)
	firewall_filter.ConnectionRate = strconv.Itoa(d.Get("connection_rate").(int))
	firewall_filter.ConnectionState = d.Get("connection_state").(string)
	firewall_filter.ConnectionType = d.Get("connection_type").(string)
	firewall_filter.Content = d.Get("content").(string)
	firewall_filter.Disabled = strconv.FormatBool(d.Get("disabled").(bool))
	firewall_filter.Dscp = strconv.Itoa(d.Get("dscp").(int))
	firewall_filter.DstAddress = d.Get("dst_address").(string)
	firewall_filter.DstAddressList = d.Get("dst_address_list").(string)
	firewall_filter.DstAddressType = d.Get("dst_address_type").(string)
	firewall_filter.DstLimit = d.Get("dst_limit").(string)
	firewall_filter.DstPort = d.Get("dst_port").(string)
	firewall_filter.Dynamic = strconv.FormatBool(d.Get("dynamic").(bool))
	firewall_filter.Fragment = BoolStringYesNo(strconv.FormatBool(d.Get("fragment").(bool)))
	firewall_filter.HotSpot = d.Get("hotspot").(string)
	firewall_filter.IcmpOptions = d.Get("icmp_options").(string)
	firewall_filter.InBridgePort = d.Get("in_bridge_port").(string)
	firewall_filter.InBridgePortList = d.Get("in_bridge_port_list").(string)
	firewall_filter.InInterface = d.Get("in_interface").(string)
	firewall_filter.InInterfaceList = d.Get("in_interface_list").(string)
	firewall_filter.IngressPriority = strconv.Itoa(d.Get("ingress_priority").(int))
	firewall_filter.IpsecPolicy = d.Get("ipsec_policy").(string)
	firewall_filter.Ipv4Options = d.Get("ipv4_options").(string)
	firewall_filter.JumpTarget = d.Get("jump_target").(string)
	firewall_filter.Layer7Protocol = d.Get("layer7_protocol").(string)
	firewall_filter.Limit = d.Get("limit").(string)
	firewall_filter.Log = strconv.FormatBool(d.Get("log").(bool))
	firewall_filter.LogPrefix = d.Get("log_prefix").(string)
	firewall_filter.Nth = d.Get("nth").(string)
	firewall_filter.OutBridgePort = d.Get("out_bridge_port").(string)
	firewall_filter.OutBridgePortList = d.Get("out_bridge_port_list").(string)
	firewall_filter.OutInterface = d.Get("out_interface").(string)
	firewall_filter.OutInterfaceList = d.Get("out_interface_list").(string)
	firewall_filter.PacketMark = d.Get("packet_mark").(string)
	firewall_filter.PacketSize = d.Get("packet_size").(string)
	firewall_filter.PerConnectionClassifier = d.Get("per_connection_classifier").(string)
	firewall_filter.Port = d.Get("port").(string)
	firewall_filter.Priority = strconv.Itoa(d.Get("priority").(int))
	firewall_filter.Protocol = d.Get("protocol").(string)
	firewall_filter.Psd = d.Get("psd").(string)
	firewall_filter.Random = strconv.Itoa(d.Get("random").(int))
	firewall_filter.RejectWith = d.Get("reject_with").(string)
	firewall_filter.RoutingTable = d.Get("routing_table").(string)
	firewall_filter.RoutingMark = d.Get("routing_mark").(string)
	firewall_filter.SrcAddress = d.Get("src_address").(string)
	firewall_filter.SrcAddressList = d.Get("src_address_list").(string)
	firewall_filter.SrcAddressType = d.Get("src_address_type").(string)
	firewall_filter.SrcPort = d.Get("src_port").(string)
	firewall_filter.SrcMacAddress = d.Get("src_mac_address").(string)
	firewall_filter.TcpFlags = d.Get("tcp_flags").(string)
	firewall_filter.TcpMss = d.Get("tcp_mss").(string)
	firewall_filter.Time = d.Get("time").(string)
	firewall_filter.TlsHost = d.Get("tls_host").(string)
	firewall_filter.Ttl = d.Get("ttl").(string)

	res, err := c.CreateIPFirewallFilter(firewall_filter)
	if err != nil {
		return fmt.Errorf("error creating firewall filter rule: %s", err.Error())
	}

	d.SetId(res.ID)
	return nil
}

func resourceIPFirewallFilterRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	firewall_filter, err := c.ReadIPFirewallFilter(d.Id())

	if err != nil {
		return fmt.Errorf("error fetching firewall filter rule: %s", err.Error())
	}

	d.SetId(firewall_filter.ID)

	return nil

}

func resourceIPFirewallFilterUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	firewall_filter := new(roscl.IPFirewallFilter)
	firewall_filter.Action = d.Get("action").(string)
	firewall_filter.AddressListTimeout = d.Get("address_list_timeout").(string)
	firewall_filter.Chain = d.Get("chain").(string)
	firewall_filter.Comment = d.Get("comment").(string)
	firewall_filter.ConnectionBytes = d.Get("connection_bytes").(string)
	firewall_filter.ConnectionLimit = d.Get("connection_limit").(string)
	firewall_filter.ConnectionMark = d.Get("connection_mark").(string)
	firewall_filter.ConnectionNatState = d.Get("connection_nat_state").(string)
	firewall_filter.ConnectionRate = strconv.Itoa(d.Get("connection_rate").(int))
	firewall_filter.ConnectionState = d.Get("connection_state").(string)
	firewall_filter.ConnectionType = d.Get("connection_type").(string)
	firewall_filter.Content = d.Get("content").(string)
	firewall_filter.Disabled = strconv.FormatBool(d.Get("disabled").(bool))
	firewall_filter.Dscp = strconv.Itoa(d.Get("dscp").(int))
	firewall_filter.DstAddress = d.Get("dst_address").(string)
	firewall_filter.DstAddressList = d.Get("dst_address_list").(string)
	firewall_filter.DstAddressType = d.Get("dst_address_type").(string)
	firewall_filter.DstLimit = d.Get("dst_limit").(string)
	firewall_filter.DstPort = d.Get("dst_port").(string)
	firewall_filter.Dynamic = strconv.FormatBool(d.Get("dynamic").(bool))
	firewall_filter.Fragment = BoolStringYesNo(strconv.FormatBool(d.Get("fragment").(bool)))
	firewall_filter.HotSpot = d.Get("hotspot").(string)
	firewall_filter.IcmpOptions = d.Get("icmp_options").(string)
	firewall_filter.InBridgePort = d.Get("in_bridge_port").(string)
	firewall_filter.InBridgePortList = d.Get("in_bridge_port_list").(string)
	firewall_filter.InInterface = d.Get("in_interface").(string)
	firewall_filter.InInterfaceList = d.Get("in_interface_list").(string)
	firewall_filter.IngressPriority = strconv.Itoa(d.Get("ingress_priority").(int))
	firewall_filter.IpsecPolicy = d.Get("ipsec_policy").(string)
	firewall_filter.Ipv4Options = d.Get("ipv4_options").(string)
	firewall_filter.JumpTarget = d.Get("jump_target").(string)
	firewall_filter.Layer7Protocol = d.Get("layer7_protocol").(string)
	firewall_filter.Limit = d.Get("limit").(string)
	firewall_filter.Log = strconv.FormatBool(d.Get("log").(bool))
	firewall_filter.LogPrefix = d.Get("log_prefix").(string)
	firewall_filter.Nth = d.Get("nth").(string)
	firewall_filter.OutBridgePort = d.Get("out_bridge_port").(string)
	firewall_filter.OutBridgePortList = d.Get("out_bridge_port_list").(string)
	firewall_filter.OutInterface = d.Get("out_interface").(string)
	firewall_filter.OutInterfaceList = d.Get("out_interface_list").(string)
	firewall_filter.PacketMark = d.Get("packet_mark").(string)
	firewall_filter.PacketSize = d.Get("packet_size").(string)
	firewall_filter.PerConnectionClassifier = d.Get("per_connection_classifier").(string)
	firewall_filter.Port = d.Get("port").(string)
	firewall_filter.Priority = strconv.Itoa(d.Get("priority").(int))
	firewall_filter.Protocol = d.Get("protocol").(string)
	firewall_filter.Psd = d.Get("psd").(string)
	firewall_filter.Random = strconv.Itoa(d.Get("random").(int))
	firewall_filter.RejectWith = d.Get("reject_with").(string)
	firewall_filter.RoutingTable = d.Get("routing_table").(string)
	firewall_filter.RoutingMark = d.Get("routing_mark").(string)
	firewall_filter.SrcAddress = d.Get("src_address").(string)
	firewall_filter.SrcAddressList = d.Get("src_address_list").(string)
	firewall_filter.SrcAddressType = d.Get("src_address_type").(string)
	firewall_filter.SrcPort = d.Get("src_port").(string)
	firewall_filter.SrcMacAddress = d.Get("src_mac_address").(string)
	firewall_filter.TcpFlags = d.Get("tcp_flags").(string)
	firewall_filter.TcpMss = d.Get("tcp_mss").(string)
	firewall_filter.Time = d.Get("time").(string)
	firewall_filter.TlsHost = d.Get("tls_host").(string)
	firewall_filter.Ttl = d.Get("ttl").(string)

	res, err := c.UpdateIPFirewallFilter(d.Id(), firewall_filter)

	if err != nil {
		return fmt.Errorf("error updating firewall filter rule: %s", err.Error())
	}

	d.SetId(res.ID)

	return nil
}

func resourceIPFirewallFilterDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	firewall_filter, _ := c.ReadIPFirewallFilter(d.Id())
	err := c.DeleteIPFirewallFilter(firewall_filter)
	if err != nil {
		return fmt.Errorf("error deleting firewall filter rule: %s", err.Error())
	}
	d.SetId("")
	return nil
}
