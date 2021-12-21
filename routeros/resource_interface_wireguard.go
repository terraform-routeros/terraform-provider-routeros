package routeros

import (
	"fmt"
	"strconv"

	roscl "github.com/gnewbury1/terraform-provider-routeros/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceWireguard() *schema.Resource {
	return &schema.Resource{
		Create: resourceInterfaceWireguardCreate,
		Read:   resourceInterfaceWireguardRead,
		Update: resourceInterfaceWireguardUpdate,
		Delete: resourceInterfaceWireguardDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"listen_port": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"mtu": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1420,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"private_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"running": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func resourceInterfaceWireguardCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)

	interface_wireguard := new(roscl.InterfaceWireguard)
	interface_wireguard.Name = d.Get("name").(string)
	interface_wireguard.Disabled = strconv.FormatBool(d.Get("disabled").(bool))

	res, err := c.CreateInterfaceWireguard(interface_wireguard)
	if err != nil {
		return fmt.Errorf("error creating ip pool: %s", err.Error())
	}

	d.SetId(res.ID)
	return nil
}

func resourceInterfaceWireguardRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	interface_wireguard, err := c.ReadInterfaceWireguard(d.Id())

	if err != nil {
		return fmt.Errorf("error fetching ip pool: %s", err.Error())
	}

	d.SetId(interface_wireguard.ID)
	d.Set("name", interface_wireguard.Name)
	d.Set("ranges", interface_wireguard.Ranges)

	return nil

}

func resourceInterfaceWireguardUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	interface_wireguard := new(roscl.InterfaceWireguard)
	interface_wireguard.Name = d.Get("name").(string)
	interface_wireguard.Ranges = d.Get("ranges").(string)

	res, err := c.UpdateInterfaceWireguard(d.Id(), interface_wireguard)

	if err != nil {
		return fmt.Errorf("error updating ip address: %s", err.Error())
	}

	d.SetId(res.ID)

	return nil
}

func resourceInterfaceWireguardDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	interface_wireguard, _ := c.ReadInterfaceWireguard(d.Id())
	err := c.DeleteInterfaceWireguard(interface_wireguard)
	if err != nil {
		return fmt.Errorf("error deleting ip address: %s", err.Error())
	}
	d.SetId("")
	return nil
}
