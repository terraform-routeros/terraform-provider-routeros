package routeros

import (
	"log"
	"strconv"

	roscl "github.com/gnewbury1/terraform-provider-routeros/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCapsManDatapath() *schema.Resource {
	return &schema.Resource{
		Create: resourceCapsManDatapathCreate,
		Read:   resourceCapsManDatapathRead,
		Update: resourceCapsManDatapathUpdate,
		Delete: resourceCapsManDatapathDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"bridge": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bridge_cost": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bridge_horizon": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface_list": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"l2mtu": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_forwarding": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_to_client_forwarding": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mtu": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"openflow_switch": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceCapsManDatapathCreate(d *schema.ResourceData, m interface{}) error {

	c := m.(*roscl.Client)
	datapath_obj := new(roscl.CapsManDatapath)

	datapath_obj.Name = d.Get("name").(string)
	datapath_obj.Bridge = d.Get("bridge").(string)
	datapath_obj.BridgeCost = d.Get("bridge_cost").(string)
	datapath_obj.Comment = d.Get("comment").(string)
	datapath_obj.BridgeHorizon = d.Get("bridge_horizon").(string)
	datapath_obj.InterfaceList = d.Get("interface_list").(string)
	datapath_obj.L2MTU = d.Get("l2mtu").(string)
	datapath_obj.MTU = d.Get("mtu").(string)
	datapath_obj.LocalForwarding = d.Get("local_forwarding").(string)
	datapath_obj.ClientToClientForwarding = d.Get("client_to_client_forwarding").(string)
	datapath_obj.OpenFlowSwitch = d.Get("openflow_switch").(string)
	datapath_obj.VlanMode = d.Get("vlan_mode").(string)
	vlan_id, is_set := d.GetOk("vlan_id")
	if is_set {
		datapath_obj.VlanID = strconv.Itoa(vlan_id.(int))
	}

	res, err := c.CreateCapsManDatapath(datapath_obj)
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PUT request to the API")
		log.Fatal(err.Error())
		return err
	}

	d.SetId(res.ID)
	return resourceCapsManDatapathRead(d, m)
}

func resourceCapsManDatapathRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	datapath, err := c.ReadCapsManDatapath(d.Id())

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a GET request to the API")
		log.Fatal(err.Error())
		return err
	}

	vlan_id, _ := strconv.Atoi(datapath.VlanID)

	d.SetId(datapath.ID)
	d.Set("name", datapath.Name)
	d.Set("bridge_cost", datapath.BridgeCost)
	d.Set("bridge", datapath.Bridge)
	d.Set("comment", datapath.Comment)
	d.Set("bridge_horizon", datapath.BridgeHorizon)
	d.Set("interface_list", datapath.BridgeHorizon)
	d.Set("l2mtu", datapath.BridgeHorizon)
	d.Set("mtu", datapath.BridgeHorizon)
	d.Set("client_to_client_forwarding", datapath.ClientToClientForwarding)
	d.Set("openflow_switch", datapath.OpenFlowSwitch)
	d.Set("vlan_mode", datapath.VlanMode)
	d.Set("vlan_id", vlan_id)

	return nil

}

func resourceCapsManDatapathUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	datapath_obj := new(roscl.CapsManDatapath)

	datapath_obj.Name = d.Get("name").(string)
	datapath_obj.Bridge = d.Get("bridge").(string)
	datapath_obj.BridgeCost = d.Get("bridge_cost").(string)
	datapath_obj.Comment = d.Get("comment").(string)
	datapath_obj.BridgeHorizon = d.Get("bridge_horizon").(string)
	datapath_obj.InterfaceList = d.Get("interface_list").(string)
	datapath_obj.L2MTU = d.Get("l2mtu").(string)
	datapath_obj.MTU = d.Get("mtu").(string)
	datapath_obj.LocalForwarding = d.Get("local_forwarding").(string)
	datapath_obj.ClientToClientForwarding = d.Get("client_to_client_forwarding").(string)
	datapath_obj.OpenFlowSwitch = d.Get("openflow_switch").(string)
	datapath_obj.VlanMode = d.Get("vlan_mode").(string)
	vlan_id, is_set := d.GetOk("vlan_id")
	if is_set {
		datapath_obj.VlanID = strconv.Itoa(vlan_id.(int))
	}

	res, err := c.UpdateCapsManDatapath(d.Id(), datapath_obj)

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PATCH request to the API")
		log.Fatal(err.Error())
		return err
	}

	d.SetId(res.ID)

	return nil
}

func resourceCapsManDatapathDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	err := c.DeleteCapsManDatapath(d.Id())
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a DELETE request to the API")
		log.Fatal(err.Error())
		return err
	}
	d.SetId("")
	return nil
}
