package routeros

import (
	"fmt"
	"strconv"
	"strings"

	roscl "github.com/gnewbury1/terraform-provider-routeros/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceBridgePort() *schema.Resource {
	return &schema.Resource{
		Create: resourceInterfaceBridgePortCreate,
		Read:   resourceInterfaceBridgePortRead,
		Update: resourceInterfaceBridgePortUpdate,
		Delete: resourceInterfaceBridgePortDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"nextid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"auto_isolate": {
				Type:     schema.TypeBool,
				Computed: true,
				Optional: true,
			},
			"bpdu_guard": {
				Type:     schema.TypeBool,
				Computed: true,
				Optional: true,
			},
			"bridge": {
				Type:     schema.TypeString,
				Required: true,
			},
			"broadcast_flood": {
				Type:     schema.TypeBool,
				Computed: true,
				Optional: true,
			},
			"comment": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  "",
			},
			"debug_info": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"edge": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"edge_port": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"edge_port_discovery": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"external_fdb_status": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"fast_leave": {
				Type:     schema.TypeBool,
				Computed: true,
				Optional: true,
			},
			"forwarding": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"frame_types": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"horizon": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"hw": {
				Type:     schema.TypeBool,
				Computed: true,
				Optional: true,
			},
			"hw_offload": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"hw_offload_group": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"inactive": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ingress_filtering": {
				Type:     schema.TypeBool,
				Computed: true,
				Optional: true,
			},
			"interface": {
				Type:     schema.TypeString,
				Required: true,
			},
			"internal_path_cost": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceInterfaceBridgePortCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)

	untagged := strings.Join(ConvSInterfaceToSString(d.Get("untagged").([]interface{})), ",")
	tagged := strings.Join(ConvSInterfaceToSString(d.Get("tagged").([]interface{})), ",")

	bridge_port := new(roscl.InterfaceBridgePort)
	bridge_port.Bridge = d.Get("bridge").(string)
	bridge_port.Disabled = strconv.FormatBool(d.Get("disabled").(bool))
	bridge_port.Tagged = tagged
	bridge_port.Untagged = untagged
	bridge_port.VlanIds = strconv.Itoa(d.Get("vlan_id").(int))

	res, err := c.CreateInterfaceBridgePort(bridge_port)
	if err != nil {
		return fmt.Errorf("error creating ip pool: %s", err.Error())
	}

	current_tagged := ConvSStringToSInterface(strings.Split(res.CurrentTagged, ","))
	current_untagged := ConvSStringToSInterface(strings.Split(res.CurrentUntagged, ","))
	dynamic, _ := strconv.ParseBool(bridge_port.Dynamic)

	d.SetId(res.ID)
	d.Set("current_tagged", current_tagged)
	d.Set("current_untagged", current_untagged)
	d.Set("dynamic", dynamic)
	return nil
}

func resourceInterfaceBridgePortRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	bridge_port, err := c.ReadInterfaceBridgePort(d.Id())

	if err != nil {
		return fmt.Errorf("error fetching ip pool: %s", err.Error())
	}

	current_tagged := ConvSStringToSInterface(strings.Split(bridge_port.CurrentTagged, ","))
	current_untagged := ConvSStringToSInterface(strings.Split(bridge_port.CurrentUntagged, ","))
	tagged := ConvSStringToSInterface(strings.Split(bridge_port.Tagged, ","))
	untagged := ConvSStringToSInterface(strings.Split(bridge_port.Untagged, ","))
	vlan_id, _ := strconv.Atoi(bridge_port.VlanIds)
	disabled, _ := strconv.ParseBool(bridge_port.Disabled)
	dynamic, _ := strconv.ParseBool(bridge_port.Dynamic)

	d.SetId(bridge_port.ID)
	d.Set("bridge", bridge_port.Bridge)
	d.Set("current_tagged", current_tagged)
	d.Set("current_untagged", current_untagged)
	d.Set("disabled", disabled)
	d.Set("dynamic", dynamic)
	d.Set("tagged", tagged)
	d.Set("untagged", untagged)
	d.Set("vlan_id", vlan_id)

	return nil

}

func resourceInterfaceBridgePortUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)

	bridge_port := new(roscl.InterfaceBridgePort)
	bridge_port.Bridge = d.Get("bridge").(string)
	bridge_port.Disabled = strconv.FormatBool(d.Get("disabled").(bool))
	bridge_port.Tagged = strings.Join(d.Get("tagged").([]string), ",")
	bridge_port.Untagged = strings.Join(d.Get("untagged").([]string), ",")
	bridge_port.VlanIds = strconv.Itoa(d.Get("vlan_id").(int))

	res, err := c.UpdateInterfaceBridgePort(d.Id(), bridge_port)

	if err != nil {
		return fmt.Errorf("error updating ip address: %s", err.Error())
	}

	d.SetId(res.ID)

	return nil
}

func resourceInterfaceBridgePortDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	bridge_port, _ := c.ReadInterfaceBridgePort(d.Id())
	err := c.DeleteInterfaceBridgePort(bridge_port)
	if err != nil {
		return fmt.Errorf("error deleting ip address: %s", err.Error())
	}
	d.SetId("")
	return nil
}
