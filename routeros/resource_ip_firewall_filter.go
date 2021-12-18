package routeros

import (
	"fmt"

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
