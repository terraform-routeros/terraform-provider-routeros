package routeros

import (
	"fmt"
	"strconv"

	roscl "github.com/gnewbury1/terraform-provider-routeros/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceVlan() *schema.Resource {
	return &schema.Resource{
		Create: resourceInterfaceVlanCreate,
		Read:   resourceInterfaceVlanRead,
		Update: resourceInterfaceVlanUpdate,
		Delete: resourceInterfaceVlanDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"interface": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vlan_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"arp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"arp_timeout": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"l2mtu": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"loop_protect": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"loop_protect_disable_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"loop_protect_send_interval": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"loop_protect_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mac_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mtu": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1500,
			},
			"running": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"use_service_tag": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceInterfaceVlanCreate(d *schema.ResourceData, m interface{}) error {

	c := m.(*roscl.Client)
	vlan_obj := new(roscl.VLAN)

	vlan_obj.Name = d.Get("name").(string)
	vlan_obj.Disabled = strconv.FormatBool(d.Get("disabled").(bool))
	vlan_obj.Interface = d.Get("interface").(string)
	vlan_obj.VlanID = strconv.Itoa(d.Get("vlan_id").(int))

	res, err := c.CreateVLAN(vlan_obj)
	if err != nil {
		return fmt.Errorf("error creating vlan: %s", err.Error())
	}

	d.SetId(res.ID)
	return resourceInterfaceVlanRead(d, m)
}

func resourceInterfaceVlanRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	vlan, err := c.ReadVLAN(d.Id())

	if err != nil {
		return fmt.Errorf("error fetching vlan: %s", err.Error())
	}

	vlan_id, _ := strconv.Atoi(vlan.VlanID)

	disabled, _ := strconv.ParseBool(vlan.Disabled)

	mtu, err := strconv.Atoi(vlan.Mtu)
	if err != nil {
		return err
	}
	use_service_tag, _ := strconv.ParseBool(vlan.UseServiceTag)

	running, _ := strconv.ParseBool(vlan.Running)

	d.SetId(vlan.ID)
	d.Set("name", vlan.Name)
	d.Set("interface", vlan.Interface)
	d.Set("vlan_id", vlan_id)
	d.Set("disabled", disabled)
	d.Set("arp", vlan.Arp)
	d.Set("arp_timeout", vlan.ArpTimeout)
	d.Set("l2mtu", vlan.L2Mtu)
	d.Set("loop_protect", vlan.LoopProtect)
	d.Set("loop_protect_disable_time", vlan.LoopProtectDisableTime)
	d.Set("loop_protect_send_interval", vlan.LoopProtectSendInterval)
	d.Set("loop_protect_status", vlan.LoopProtectStatus)
	d.Set("mac_address", vlan.MacAddress)
	d.Set("mtu", mtu)
	d.Set("use_service_tag", use_service_tag)
	d.Set("running", running)

	return nil

}

func resourceInterfaceVlanUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	vlan_obj := new(roscl.VLAN)
	vlan_obj.Name = d.Get("name").(string)
	vlan_obj.Disabled = strconv.FormatBool(d.Get("disabled").(bool))
	vlan_obj.Interface = d.Get("interface").(string)
	vlan_obj.VlanID = strconv.Itoa(d.Get("vlan_id").(int))

	res, err := c.UpdateVLAN(d.Id(), vlan_obj)

	if err != nil {
		return fmt.Errorf("error updating vlan: %s", err.Error())
	}

	d.SetId(res.ID)

	return nil
}

func resourceInterfaceVlanDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	err := c.DeleteVLAN(d.Id())
	if err != nil {
		return fmt.Errorf("error deleting vlan: %s", err.Error())
	}
	d.SetId("")
	return nil
}
