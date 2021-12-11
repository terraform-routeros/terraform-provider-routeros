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
				Type:     schema.TypeString,
				Required: true,
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
	vlan_obj.VlanID = d.Get("vlan_id").(string)

	res, err := c.CreateVLAN(vlan_obj)
	if err != nil {
		return fmt.Errorf("error creating vlan: %s", err.Error())
	}

	d.SetId(res.ID)
	return nil
}

func resourceInterfaceVlanRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	vlan, err := c.ReadVLAN(d.Id())

	if err != nil {
		return fmt.Errorf("error fetching vlan: %s", err.Error())
	}

	d.SetId(vlan.ID)

	return nil

}

func resourceInterfaceVlanUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	vlan_obj := new(roscl.VLAN)
	vlan_obj.Name = d.Get("name").(string)
	vlan_obj.Disabled = strconv.FormatBool(d.Get("disabled").(bool))
	vlan_obj.Interface = d.Get("interface").(string)
	vlan_obj.VlanID = d.Get("vlan_id").(string)

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
