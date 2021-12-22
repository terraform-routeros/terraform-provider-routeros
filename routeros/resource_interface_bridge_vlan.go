package routeros

import (
	"log"
	"strconv"
	"strings"

	roscl "github.com/gnewbury1/terraform-provider-routeros/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceBridgeVlan() *schema.Resource {
	return &schema.Resource{
		Create: resourceInterfaceBridgeVlanCreate,
		Read:   resourceInterfaceBridgeVlanRead,
		Update: resourceInterfaceBridgeVlanUpdate,
		Delete: resourceInterfaceBridgeVlanDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"bridge": {
				Type:     schema.TypeString,
				Required: true,
			},
			"current_tagged": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"current_untagged": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"dynamic": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"tagged": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"untagged": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"vlan_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func resourceInterfaceBridgeVlanCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)

	untagged := strings.Join(ConvSInterfaceToSString(d.Get("untagged").([]interface{})), ",")
	tagged := strings.Join(ConvSInterfaceToSString(d.Get("tagged").([]interface{})), ",")

	bridge_vlan := new(roscl.InterfaceBridgeVlan)
	bridge_vlan.Bridge = d.Get("bridge").(string)
	bridge_vlan.Disabled = strconv.FormatBool(d.Get("disabled").(bool))
	bridge_vlan.Tagged = tagged
	bridge_vlan.Untagged = untagged
	bridge_vlan.VlanIds = strconv.Itoa(d.Get("vlan_id").(int))

	res, err := c.CreateInterfaceBridgeVlan(bridge_vlan)
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PUT request to the API")
		log.Fatal(err.Error())
		return err
	}

	current_tagged := ConvSStringToSInterface(strings.Split(res.CurrentTagged, ","))
	current_untagged := ConvSStringToSInterface(strings.Split(res.CurrentUntagged, ","))
	dynamic, _ := strconv.ParseBool(bridge_vlan.Dynamic)

	d.SetId(res.ID)
	d.Set("current_tagged", current_tagged)
	d.Set("current_untagged", current_untagged)
	d.Set("dynamic", dynamic)
	return nil
}

func resourceInterfaceBridgeVlanRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	bridge_vlan, err := c.ReadInterfaceBridgeVlan(d.Id())

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a GET request to the API")
		log.Fatal(err.Error())
		return err
	}

	current_tagged := ConvSStringToSInterface(strings.Split(bridge_vlan.CurrentTagged, ","))
	current_untagged := ConvSStringToSInterface(strings.Split(bridge_vlan.CurrentUntagged, ","))
	tagged := ConvSStringToSInterface(strings.Split(bridge_vlan.Tagged, ","))
	untagged := ConvSStringToSInterface(strings.Split(bridge_vlan.Untagged, ","))
	vlan_id, _ := strconv.Atoi(bridge_vlan.VlanIds)
	disabled, _ := strconv.ParseBool(bridge_vlan.Disabled)
	dynamic, _ := strconv.ParseBool(bridge_vlan.Dynamic)

	d.SetId(bridge_vlan.ID)
	d.Set("bridge", bridge_vlan.Bridge)
	d.Set("current_tagged", current_tagged)
	d.Set("current_untagged", current_untagged)
	d.Set("disabled", disabled)
	d.Set("dynamic", dynamic)
	d.Set("tagged", tagged)
	d.Set("untagged", untagged)
	d.Set("vlan_id", vlan_id)

	return nil

}

func resourceInterfaceBridgeVlanUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)

	bridge_vlan := new(roscl.InterfaceBridgeVlan)
	bridge_vlan.Bridge = d.Get("bridge").(string)
	bridge_vlan.Disabled = strconv.FormatBool(d.Get("disabled").(bool))
	bridge_vlan.Tagged = strings.Join(d.Get("tagged").([]string), ",")
	bridge_vlan.Untagged = strings.Join(d.Get("untagged").([]string), ",")
	bridge_vlan.VlanIds = strconv.Itoa(d.Get("vlan_id").(int))

	res, err := c.UpdateInterfaceBridgeVlan(d.Id(), bridge_vlan)

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PATCH request to the API")
		log.Fatal(err.Error())
		return err
	}

	d.SetId(res.ID)

	return nil
}

func resourceInterfaceBridgeVlanDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	bridge_vlan, _ := c.ReadInterfaceBridgeVlan(d.Id())
	err := c.DeleteInterfaceBridgeVlan(bridge_vlan)
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a DELETE request to the API")
		log.Fatal(err.Error())
		return err
	}
	d.SetId("")
	return nil
}
