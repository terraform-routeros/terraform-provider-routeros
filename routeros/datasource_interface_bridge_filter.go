package routeros

// Script generated from sampled device MikroTik 7.11.2 (stable) on CHR AMD-x86_64

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceInterfaceBridgeFilter() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceInterfaceBridgeFiltersRead,
		Schema: map[string]*schema.Schema{
			MetaResourcePath: PropResourcePath("/interface/bridge/filter"),
			MetaId:           PropId(Id),

			KeyFilter: PropFilterRw,
			"filters": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": { // Sample = .id: "*1"
							Type:     schema.TypeString,
							Computed: true,
						},
						"action": { // Sample = action: "drop"
							Type:     schema.TypeString,
							Computed: true,
						},
						"bytes": { // Sample = bytes: "0"
							Type:     schema.TypeInt,
							Computed: true,
						},
						"in_interface": { // Sample = chain: "ether1"
							Type:     schema.TypeString,
							Computed: true,
						},
						"chain": { // Sample = chain: "forward"
							Type:     schema.TypeString,
							Computed: true,
						},
						"comment": { // Sample = comment: "Drop data between cast ports"
							Type:     schema.TypeString,
							Computed: true,
						},
						"dynamic": { // Sample = dynamic: "false"
							Type:     schema.TypeBool,
							Computed: true,
						},
						"invalid": { // Sample = invalid: "false"
							Type:     schema.TypeBool,
							Computed: true,
						},
						"packets": { // Sample = packets: "0"
							Type:     schema.TypeInt,
							Computed: true,
						},
						"mac_protocol": { // Sample = mac_protocol: "0x890D"
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func datasourceInterfaceBridgeFiltersRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	s := DatasourceInterfaceBridgeFilter().Schema
	path := s[MetaResourcePath].Default.(string)

	res, err := ReadItemsFiltered(buildReadFilter(d.Get(KeyFilter).(map[string]interface{})), path, m.(Client))
	if err != nil {
		return diag.FromErr(err)
	}

	return MikrotikResourceDataToTerraformDatasource(res, "filters", s, d)
}
