package routeros

import (
	"log"
	"strconv"

	roscl "github.com/gnewbury1/terraform-provider-routeros/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceVrrp() *schema.Resource {
	return &schema.Resource{
		Create: resourceInterfaceVrrpCreate,
		Read:   resourceInterfaceVrrpRead,
		Update: resourceInterfaceVrrpUpdate,
		Delete: resourceInterfaceVrrpDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"arp": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "enabled",
			},
			"v3_protocol": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ipv4",
			},
			"authentication": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "none",
			},
			"interface": {
				Type:     schema.TypeString,
				Required: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"mtu": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1500,
			},
			"password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"on_fail": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"on_master": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"on_backup": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sync_connection_tracking": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
			"vrid": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"interval": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "1s",
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"preemption_mode": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"version": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3,
			},
			"arp_timeout": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "auto",
			},
		},
	}
}

func resourceInterfaceVrrpCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)

	interface_vrrp := new(roscl.InterfaceVrrp)
	interface_vrrp.Name = d.Get("name").(string)
	interface_vrrp.Arp = d.Get("arp").(string)
	interface_vrrp.ArpTimeout = d.Get("arp_timeout").(string)
	interface_vrrp.Mtu = strconv.Itoa(d.Get("mtu").(int))
	interface_vrrp.V3Protocol = d.Get("v3_protocol").(string)
	interface_vrrp.Authentication = d.Get("authentication").(string)
	interface_vrrp.Interface = d.Get("interface").(string)
	interface_vrrp.Password = d.Get("password").(string)
	interface_vrrp.Priority = strconv.Itoa(d.Get("priority").(int))
	interface_vrrp.OnFail = d.Get("on_fail").(string)
	interface_vrrp.OnMaster = d.Get("on_master").(string)
	interface_vrrp.OnBackup = d.Get("on_backup").(string)
	interface_vrrp.SyncConnectionTracking = d.Get("sync_connection_tracking").(string)
	interface_vrrp.Vrid = strconv.Itoa(d.Get("vrid").(int))
	interface_vrrp.Interval = d.Get("interval").(string)
	interface_vrrp.PreemptionMode = strconv.FormatBool(d.Get("preemption_mode").(bool))
	interface_vrrp.Version = strconv.Itoa(d.Get("version").(int))

	res, err := c.CreateInterfaceVrrp(interface_vrrp)
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PUT request to the API")
		log.Fatal(err.Error())
		return err
	}

	d.SetId(res.ID)
	return nil
}

func resourceInterfaceVrrpRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	res, err := c.ReadInterfaceVrrp(d.Id())

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a GET request to the API")
		log.Fatal(err.Error())
		return err
	}

	priority, _ := strconv.Atoi(res.Priority)
	mtu, _ := strconv.Atoi(res.Mtu)
	vrid, _ := strconv.Atoi(res.Vrid)
	version, _ := strconv.Atoi(res.Version)
	preemption_mode, _ := strconv.ParseBool(res.PreemptionMode)

	d.SetId(res.ID)
	d.Set("arp", res.Arp)
	d.Set("priority", priority)
	d.Set("mtu", mtu)
	d.Set("vrid", vrid)
	d.Set("version", version)
	d.Set("arp_timeout", res.ArpTimeout)
	d.Set("preemption_mode", preemption_mode)
	d.Set("v3_protocol", res.V3Protocol)
	d.Set("authentication", res.Authentication)
	d.Set("interface", res.Interface)
	d.Set("password", res.Password)
	d.Set("on_fail", res.OnFail)
	d.Set("on_master", res.OnMaster)
	d.Set("on_backup", res.OnBackup)
	d.Set("name", res.Name)
	d.Set("sync_connection_tracking", res.SyncConnectionTracking)
	d.Set("interval", res.Interval)

	return nil

}

func resourceInterfaceVrrpUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	interface_vrrp := new(roscl.InterfaceVrrp)
	interface_vrrp.Name = d.Get("name").(string)
	interface_vrrp.Arp = d.Get("arp").(string)
	interface_vrrp.ArpTimeout = d.Get("arp_timeout").(string)
	interface_vrrp.Mtu = strconv.Itoa(d.Get("mtu").(int))
	interface_vrrp.V3Protocol = d.Get("v3_protocol").(string)
	interface_vrrp.Authentication = d.Get("authentication").(string)
	interface_vrrp.Interface = d.Get("interface").(string)
	interface_vrrp.Password = d.Get("password").(string)
	interface_vrrp.Priority = strconv.Itoa(d.Get("priority").(int))
	interface_vrrp.OnFail = d.Get("on_fail").(string)
	interface_vrrp.OnMaster = d.Get("on_master").(string)
	interface_vrrp.OnBackup = d.Get("on_backup").(string)
	interface_vrrp.SyncConnectionTracking = d.Get("sync_connection_tracking").(string)
	interface_vrrp.Vrid = strconv.Itoa(d.Get("vrid").(int))
	interface_vrrp.Interval = d.Get("interval").(string)
	interface_vrrp.PreemptionMode = strconv.FormatBool(d.Get("preemption_mode").(bool))
	interface_vrrp.Version = strconv.Itoa(d.Get("version").(int))

	res, err := c.UpdateInterfaceVrrp(d.Id(), interface_vrrp)

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PATCH request to the API")
		log.Fatal(err.Error())
		return err
	}

	d.SetId(res.ID)

	return nil
}

func resourceInterfaceVrrpDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	interface_vrrp, _ := c.ReadInterfaceVrrp(d.Id())
	err := c.DeleteInterfaceVrrp(interface_vrrp)
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a DELETE request to the API")
		log.Fatal(err.Error())
		return err
	}
	d.SetId("")
	return nil
}
