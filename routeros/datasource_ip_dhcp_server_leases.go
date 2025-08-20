package routeros

// Script generated from sampled device MikroTik 7.11.2 (stable) on CHR AMD-x86_64

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceIpDhcpServerLeases() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceIpDhcpServerLeasesRead,
		Schema: map[string]*schema.Schema{
			MetaResourcePath: PropResourcePath("/ip/dhcp-server/lease"),
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
						"active_client_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"active_server": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"active_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"active_mac_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"address": { // Sample = address: "192.168.0.1"
							Type:     schema.TypeString,
							Computed: true,
						},
						"address_lists": { // Sample = address-lists: ""
							Type:     schema.TypeString,
							Computed: true,
						},
						"age": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"allow_dual_stack_queue": {
    						Type:     schema.TypeBool,
    						Computed: true,
						},
						"blocked": { // Sample = blocked: "false"
							Type:     schema.TypeBool,
							Computed: true,
						},
						"comment": { // Sample = comment: "server1 "
							Type:     schema.TypeString,
							Computed: true,
						},
						"class_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"client_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"dhcp_option": { // Sample = dhcp-option: ""
							Type:     schema.TypeString,
							Computed: true,
						},
						"disabled": { // Sample = disabled: "true"
							Type:     schema.TypeBool,
							Computed: true,
						},
						"dynamic": { // Sample = dynamic: "false"
							Type:     schema.TypeBool,
							Computed: true,
						},
						"expires_after": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"host_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_seen": { // Sample = last-seen: "never"
							Type:     schema.TypeString,
							Computed: true,
						},
						"mac_address": { // Sample = mac-address: "00:0C:29:00:01:A0"
							Type:     schema.TypeString,
							Computed: true,
						},
						"radius": { // Sample = radius: "false"
							Type:     schema.TypeBool,
							Computed: true,
						},
						"server": { // Sample = server: "bridge_dhcp_lan"
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": { // Sample = status: "waiting"
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func datasourceIpDhcpServerLeasesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	s := DatasourceIpDhcpServerLeases().Schema
	path := s[MetaResourcePath].Default.(string)

	res, err := ReadItemsFiltered(buildReadFilter(d.Get(KeyFilter).(map[string]interface{})), path, m.(Client))
	if err != nil {
		return diag.FromErr(err)
	}

	return MikrotikResourceDataToTerraformDatasource(res, "data", s, d)
}
