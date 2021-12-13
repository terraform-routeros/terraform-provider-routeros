package routeros

import (
	"fmt"
	"strconv"

	roscl "github.com/gnewbury1/terraform-provider-routeros/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceInterfaces() *schema.Resource {
	return &schema.Resource{
		Read: datasourceInterfacesRead,
		Schema: map[string]*schema.Schema{
			"interfaces": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"actual_mtu": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"default_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"disabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"fp_rx_byte": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"fp_rx_packet": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"fp_tx_byte": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"fp_tx_packet": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"l2mtu": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"last_link_down_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_link_up_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"link_downs": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"mac_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"max_l2mtu": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"mtu": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"running": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"rx_byte": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"rx_packet": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"slave": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"tx_byte": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"tx_packet": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"tx_queue_drop": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func datasourceInterfacesRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	res, err := c.ReadInterfaces()

	if err != nil {
		return fmt.Errorf("error fetching interfaces: %s", err.Error())
	}

	interfaces := make([]map[string]interface{}, len(res))
	ros_interface := make(map[string]interface{})
	for k, v := range res {
		ros_interface["id"] = v.ID
		ros_interface["actual_mtu"], _ = strconv.Atoi(v.ActualMtu)
		ros_interface["default_name"] = v.DefaultName
		ros_interface["disabled"], _ = strconv.ParseBool(v.Disabled)
		ros_interface["fp_rx_byte"], _ = strconv.Atoi(v.FpRxByte)
		ros_interface["fp_rx_packet"], _ = strconv.Atoi(v.FpRxPacket)
		ros_interface["fp_tx_byte"], _ = strconv.Atoi(v.FpTxByte)
		ros_interface["fp_tx_packet"], _ = strconv.Atoi(v.FpTxPacket)
		ros_interface["l2mtu"], _ = strconv.Atoi(v.L2Mtu)
		ros_interface["last_link_down_time"] = v.LastLinkDownTime
		ros_interface["last_link_up_time"] = v.LastLinkUpTime
		ros_interface["link_downs"], _ = strconv.Atoi(v.LinkDowns)
		ros_interface["mac_address"] = v.MacAddress
		ros_interface["max_l2mtu"], _ = strconv.Atoi(v.MaxL2Mtu)
		ros_interface["mtu"], _ = strconv.Atoi(v.Mtu)
		ros_interface["name"] = v.Name
		ros_interface["running"], _ = strconv.ParseBool(v.Running)
		ros_interface["rx_byte"], _ = strconv.Atoi(v.RxByte)
		ros_interface["rx_packet"], _ = strconv.Atoi(v.RxPacket)
		ros_interface["slave"], _ = strconv.ParseBool(v.Slave)
		ros_interface["tx_byte"], _ = strconv.Atoi(v.TxByte)
		ros_interface["tx_packet"], _ = strconv.Atoi(v.TxPacket)
		ros_interface["tx_queue_drop"], _ = strconv.Atoi(v.TxQueueDrop)
		ros_interface["type"] = v.Type
		interfaces[k] = ros_interface
	}

	d.SetId(resource.UniqueId())
	if err := d.Set("interfaces", interfaces); err != nil {
		return err
	}

	return nil

}
