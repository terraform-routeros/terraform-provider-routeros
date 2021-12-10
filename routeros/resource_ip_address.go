package routeros

import (
	"fmt"

	roscl "github.com/gnewbury1/terraform-provider-routeros/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIPAddress() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIPAddressCreate,
		ReadContext:   resourceIPAddressRead,
		UpdateContext: resourceIPAddressUpdate,
		DeleteContext: resourceIPAddressDelete,
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
	ip_addr.Disabled = disabled
	ip_addr.Interface = network_interface
	ip_addr.Network = network

	res, err := c.CreateIPAddress(ip_addr)
	if err != nil {
		return fmt.Errorf("Error creating IP address: %s", err.Error())
	}

	d.SetId(res.ID)
	return nil
}
