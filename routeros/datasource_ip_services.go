package routeros

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceIPServices() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceIPServicesRead,
		Schema: map[string]*schema.Schema{
			MetaResourcePath: PropResourcePath("/ip/service"),
			MetaId:           PropId(Id),

			KeyFilter: PropFilterRw,
			"services": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"certificate": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"connection": {
							Type:     schema.TypeBool,
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
						"invalid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"port": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"tls_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vrf": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func datasourceIPServicesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	s := DatasourceIPServices().Schema
	path := s[MetaResourcePath].Default.(string)

	var filter = d.Get(KeyFilter).(map[string]interface{})
	filter["dynamic"] = "false"

	res, err := ReadItemsFiltered(buildReadFilter(filter), path, m.(Client))
	if err != nil {
		return diag.FromErr(err)
	}

	return MikrotikResourceDataToTerraformDatasource(res, "services", s, d)
}
