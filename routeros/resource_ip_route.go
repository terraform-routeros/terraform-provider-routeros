package routeros

import (
	"fmt"
	"strconv"

	roscl "github.com/gnewbury1/terraform-provider-routeros/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIPRoute() *schema.Resource {
	return &schema.Resource{
		Create: resourceIPRouteCreate,
		Read:   resourceIPRouteRead,
		Update: resourceIPRouteUpdate,
		Delete: resourceIPRouteDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"active": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"dhcp": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"distance": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"dst_address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"dynamic": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ecmp": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"gateway": {
				Type:     schema.TypeString,
				Required: true,
			},
			"hw_offloaded": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"immediate_gw": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"inactive": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"pref_src": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"routing_table": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "main",
			},
			"scope": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"suppress_hw_offload": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"target_scope": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vrf_interface": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIPRouteCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)

	ip_route := new(roscl.IPRoute)
	ip_route.Distance = strconv.Itoa(d.Get("distance").(int))
	ip_route.DstAddress = d.Get("dst_address").(string)
	ip_route.Gateway = d.Get("gateway").(string)
	ip_route.PrefSrc = d.Get("pref_src").(string)
	ip_route.RoutingTable = d.Get("routing_table").(string)

	res, err := c.CreateIPRoute(ip_route)
	if err != nil {
		return fmt.Errorf("error creating ip pool: %s", err.Error())
	}

	d.SetId(res.ID)
	return nil
}

func resourceIPRouteRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	ip_route, err := c.ReadIPRoute(d.Id())

	if err != nil {
		return fmt.Errorf("error fetching ip pool: %s", err.Error())
	}

	active, _ := strconv.ParseBool(ip_route.Active)
	dhcp, _ := strconv.ParseBool(ip_route.Dhcp)
	distance, _ := strconv.Atoi(ip_route.Distance)
	dynamic, _ := strconv.ParseBool(ip_route.Dynamic)
	ecmp, _ := strconv.ParseBool(ip_route.Ecmp)
	hw_offloaded, _ := strconv.ParseBool(ip_route.HwOffloaded)
	inactive, _ := strconv.ParseBool(ip_route.Inactive)
	scope, _ := strconv.Atoi(ip_route.Scope)
	suppress_hw_offload, _ := strconv.ParseBool(ip_route.SuppressHwOffload)
	targetscope, _ := strconv.Atoi(ip_route.TargetScope)

	d.SetId(ip_route.ID)
	d.Set("active", active)
	d.Set("dhcp", dhcp)
	d.Set("distance", distance)
	d.Set("dst_address", ip_route.DstAddress)
	d.Set("dynamic", dynamic)
	d.Set("ecmp", ecmp)
	d.Set("gateway", ip_route.Gateway)
	d.Set("hw_offloaded", hw_offloaded)
	d.Set("immediate_gw", ip_route.ImmediateGw)
	d.Set("inactive", inactive)
	d.Set("pref_src", ip_route.PrefSrc)
	d.Set("routing_table", ip_route.RoutingTable)
	d.Set("scope", scope)
	d.Set("suppress_hw_offload", suppress_hw_offload)
	d.Set("target_scope", targetscope)
	d.Set("vrf_interface", ip_route.VrfInterface)

	return nil

}

func resourceIPRouteUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	ip_route := new(roscl.IPRoute)
	ip_route.Distance = strconv.Itoa(d.Get("distance").(int))
	ip_route.DstAddress = d.Get("dst_address").(string)
	ip_route.Gateway = d.Get("gateway").(string)
	ip_route.Inactive = strconv.FormatBool(d.Get("inactive").(bool))
	ip_route.PrefSrc = d.Get("pref_src").(string)
	ip_route.RoutingTable = d.Get("routing_table").(string)

	res, err := c.UpdateIPRoute(d.Id(), ip_route)

	if err != nil {
		return fmt.Errorf("error updating ip address: %s", err.Error())
	}

	d.SetId(res.ID)

	return nil
}

func resourceIPRouteDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	ip_route, _ := c.ReadIPRoute(d.Id())
	err := c.DeleteIPRoute(ip_route)
	if err != nil {
		return fmt.Errorf("error deleting ip address: %s", err.Error())
	}
	d.SetId("")
	return nil
}
