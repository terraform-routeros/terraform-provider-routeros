package routeros

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceIPRoutes() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceIPRoutesRead,
		Schema: map[string]*schema.Schema{
			MetaResourcePath: PropResourcePath("/ip/route"),
			MetaId:           PropId(Id),

			KeyFilter: PropFilterRw,
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
						"blackhole": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						KeyComment: {
							Type:     schema.TypeString,
							Computed: true,
						},
						"connect": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"dhcp": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"disabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"dst_address": {
							Type:     schema.TypeString,
							Computed: true,
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
							Computed: true,
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
						"local_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"pref_src": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"routing_table": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"static": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"suppress_hw_offload": {
							Type:     schema.TypeBool,
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

func datasourceIPRoutesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	s := DatasourceIPRoutes().Schema
	path := s[MetaResourcePath].Default.(string)

	res, err := ReadItemsFiltered(buildReadFilter(d.Get(KeyFilter).(map[string]interface{})), path, m.(Client))
	if err != nil {
		return diag.FromErr(err)
	}

	return MikrotikResourceDataToTerraformDatasource(res, "routes", s, d)
}
