package routeros

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceIPv6Addresses() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceIPv6AddressesRead,
		Schema: map[string]*schema.Schema{
			MetaResourcePath: PropResourcePath("/ipv6/address"),
			MetaId:           PropId(Id),

			KeyFilter: PropFilterRw,
			"addresses": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"actual_interface": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"advertise": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"deprecated": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"comment": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"disabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"dynamic": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"eui_64": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"from_pool": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"interface": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"invalid": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"link_local": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"no_dad": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"slave": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func datasourceIPv6AddressesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	s := DatasourceIPv6Addresses().Schema
	path := s[MetaResourcePath].Default.(string)

	res, err := ReadItemsFiltered(buildReadFilter(d.Get(KeyFilter).(map[string]interface{})), path, m.(Client))
	if err != nil {
		return diag.FromErr(err)
	}

	return MikrotikResourceDataToTerraformDatasource(res, "addresses", s, d)
}
