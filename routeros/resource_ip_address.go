package routeros

import (
	"fmt"
	"strconv"

	roscl "github.com/gnewbury1/terraform-provider-routeros/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIPAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceIPAddressCreate,
		Read:   resourceIPAddressRead,
		Update: resourceIPAddressUpdate,
		Delete: resourceIPAddressDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"interface": {
				Type:     schema.TypeString,
				Required: true,
			},
			"network": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceIPAddressCreate(d *schema.ResourceData, m interface{}) error {
	address := d.Get("address").(string)
	comment := d.Get("comment").(string)
	disabled := d.Get("disabled").(bool)
	network_interface := d.Get("interface").(string)
	network := d.Get("network").(string)

	c := m.(*roscl.Client)
	ip_addr := new(roscl.IPAddress)

	ip_addr.Address = address
	ip_addr.Comment = comment
	ip_addr.Disabled = strconv.FormatBool(disabled)
	ip_addr.Interface = network_interface
	ip_addr.Network = network

	res, err := c.CreateIPAddress(ip_addr)
	if err != nil {
		return fmt.Errorf("error creating ip address: %s", err.Error())
	}

	d.SetId(res.ID)
	return nil
}

func resourceIPAddressRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	ipaddr, err := c.GetIPAddresses(d.Id())

	if err != nil {
		return fmt.Errorf("error fetching ip address: %s", err.Error())
	}

	d.SetId(ipaddr.ID)

	return nil

}

func resourceIPAddressUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(roscl.Client)
	ipaddr, err := c.GetIPAddresses(d.Id())

	if err != nil {
		return fmt.Errorf("error updating ip address: %s", err.Error())
	}
	ipaddr.Address = d.Get("address").(string)
	ipaddr.Comment = d.Get("comment").(string)
	ipaddr.Interface = d.Get("interface").(string)
	ipaddr.Disabled = strconv.FormatBool(d.Get("disabled").(bool))
	ipaddr.Network = d.Get("network").(string)

	res, err := c.UpdateIPAddress(d.Id(), ipaddr)

	if err != nil {
		return fmt.Errorf("error updating ip address: %s", err.Error())
	}

	d.SetId(res.ID)

	return nil
}

func resourceIPAddressDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(roscl.Client)
	err := c.DeleteIPAddress(d.Id())
	if err != nil {
		return fmt.Errorf("error deleting ip address: %s", err.Error())
	}
	d.SetId("")
	return nil
}
