package routeros

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceInterfaces() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceInterfacesRead,
		Schema: map[string]*schema.Schema{
			MetaResourcePath: PropResourcePath("/interface"),
			MetaId:           PropId(Id),
			MetaSkipFields: PropSkipFields(
				"fp_rx_byte", "fp_rx_packet", "fp_tx_byte", "fp_tx_packet", "link_downs",
				"rx_byte", "rx_drop", "rx_error", "rx_packet", "tx_byte",
				"tx_drop", "tx_error", "tx_packet", "tx_queue_drop",
			),
			KeyFilter: PropFilterRw,
			"interfaces": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"actual_mtu": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"default_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"dynamic": {
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
						"l2mtu": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"last_link_down_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_link_up_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mac_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"max_l2mtu": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						// Can be - 'auto'
						"mtu": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"running": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"slave": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func datasourceInterfacesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	s := DatasourceInterfaces().Schema
	path := s[MetaResourcePath].Default.(string)

	res, err := ReadItemsFiltered(buildReadFilter(d.Get(KeyFilter).(map[string]interface{})), path, m.(Client))
	if err != nil {
		return diag.FromErr(err)
	}

	return MikrotikResourceDataToTerraformDatasource(res, "interfaces", s, d)
}
