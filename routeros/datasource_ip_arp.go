package routeros

// Script generated from sampled device MikroTik 7.11.2 (stable) on CHR AMD-x86_64

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceIpArp() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceIpArpRead,
		Schema: map[string]*schema.Schema{
			MetaResourcePath: PropResourcePath("/ip/arp"),
			MetaId:           PropId(Id),

			KeyFilter: PropFilterRw,
			"data": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": { // Sample = .id: "*1"
							Type:     schema.TypeString,
							Computed: true,
						},
						"dhcp": { // Sample = DHCP: "false"
							Type:     schema.TypeBool,
							Computed: true,
						},
						"address": { // Sample = address: "192.168.9.10"
							Type:     schema.TypeString,
							Computed: true,
						},
						"complete": { // Sample = complete: "true"
							Type:     schema.TypeBool,
							Computed: true,
						},
						"disabled": { // Sample = disabled: "false"
							Type:     schema.TypeBool,
							Computed: true,
						},
						"dynamic": { // Sample = dynamic: "true"
							Type:     schema.TypeBool,
							Computed: true,
						},
						"interface": { // Sample = interface: "ether1"
							Type:     schema.TypeString,
							Computed: true,
						},
						"invalid": { // Sample = invalid: "false"
							Type:     schema.TypeBool,
							Computed: true,
						},
						"mac_address": { // Sample = mac-address: "70:85:C2:37:5A:21"
							Type:     schema.TypeString,
							Computed: true,
						},
						"published": { // Sample = published: "false"
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func datasourceIpArpRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	s := DatasourceIpArp().Schema
	path := s[MetaResourcePath].Default.(string)

	res, err := ReadItemsFiltered(buildReadFilter(d.Get(KeyFilter).(map[string]interface{})), path, m.(Client))
	if err != nil {
		return diag.FromErr(err)
	}

	return MikrotikResourceDataToTerraformDatasource(res, "data", s, d)
}
