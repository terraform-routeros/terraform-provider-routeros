package routeros

import (
	"fmt"
	"strconv"

	roscl "github.com/gnewbury1/terraform-provider-routeros/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceIPRoutes() *schema.Resource {
	return &schema.Resource{
		Read: datasourceIPRoutesRead,
		Schema: map[string]*schema.Schema{
			"routes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
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
				},
			},
		},
	}
}

func datasourceIPRoutesRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	res, err := c.ReadIPRoutes()

	if err != nil {
		return fmt.Errorf("error fetching ip routees: %s", err.Error())
	}

	routes := make([]map[string]interface{}, len(res))
	route := make(map[string]interface{})
	for k, v := range res {
		route["id"] = v.ID
		route["active"], _ = strconv.ParseBool(v.Active)
		route["dhcp"], _ = strconv.ParseBool(v.Dhcp)
		route["distance"], _ = strconv.Atoi(v.Distance)
		route["dst_address"] = v.DstAddress
		route["dynamic"], _ = strconv.ParseBool(v.Dynamic)
		route["ecmp"], _ = strconv.ParseBool(v.Ecmp)
		route["gateway"] = v.Gateway
		route["hw_offloaded"], _ = strconv.ParseBool(v.HwOffloaded)
		route["immediate_gw"] = v.ImmediateGw
		route["inactive"], _ = strconv.ParseBool(v.Inactive)
		route["pref_src"] = v.PrefSrc
		route["routing_table"] = v.RoutingTable
		route["scope"], _ = strconv.Atoi(v.Scope)
		route["suppress_hw_offload"], _ = strconv.ParseBool(v.SuppressHwOffload)
		route["target_scope"], _ = strconv.Atoi(v.TargetScope)
		route["vrf_interface"] = v.VrfInterface
		routes[k] = route
	}

	d.SetId(resource.UniqueId())
	if err := d.Set("routes", routes); err != nil {
		return err
	}

	return nil

}
