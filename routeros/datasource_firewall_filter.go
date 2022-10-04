package routeros

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func DatasourceFirewallFilter() *schema.Resource {
	firewallFilterSchema := &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				KeyFilter: PropFilterRw,

				"id": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"action": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"address_list_timeout": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"bytes": {
					Type:     schema.TypeInt,
					Computed: true,
				},
				"chain": {
					Type:     schema.TypeString,
					Computed: true,
				},
				KeyComment: {
					Type:     schema.TypeString,
					Computed: true,
				},
				"connection_bytes": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"connection_limit": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"connection_mark": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"connection_nat_state": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"connection_rate": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"connection_state": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"connection_type": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"content": {
					Type:     schema.TypeString,
					Computed: true,
				},
				KeyDisabled: {
					Type:     schema.TypeBool,
					Computed: true,
				},
				"dscp": {
					Type:     schema.TypeInt,
					Computed: true,
				},
				"dst_address": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"dst_address_list": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"dst_address_type": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"dst_limit": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"dst_port": {
					Type:     schema.TypeString,
					Computed: true,
				},
				KeyDynamic: {
					Type:     schema.TypeBool,
					Computed: true,
				},
				"fragment": {
					Type:     schema.TypeBool,
					Computed: true,
				},
				"hotspot": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"icmp_options": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"in_bridge_port": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"in_bridge_port_list": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"in_interface": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"in_interface_list": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"ingress_priority": {
					Type:     schema.TypeInt,
					Computed: true,
				},
				"invalid": {
					Type:     schema.TypeBool,
					Computed: true,
				},
				"ipsec_policy": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"ipv4_options": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"jump_target": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"hw_offload": {
					Type:     schema.TypeBool,
					Computed: true,
				},
				"layer7_protocol": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"limit": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"log": {
					Type:     schema.TypeBool,
					Computed: true,
				},
				"log_prefix": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"nth": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"out_bridge_port": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"out_bridge_port_list": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"out_interface": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"out_interface_list": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"packets": {
					Type:     schema.TypeInt,
					Computed: true,
				},
				"packet_mark": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"packet_size": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"per_connection_classifier": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"port": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"priority": {
					Type:     schema.TypeInt,
					Computed: true,
				},
				"protocol": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"psd": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"random": {
					Type:     schema.TypeInt,
					Computed: true,
				},
				"reject_with": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"routing_table": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"routing_mark": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"src_address": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"src_address_list": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"src_address_type": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"src_port": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"src_mac_address": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"tcp_flags": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"tcp_mss": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"time": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"tls_host": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"ttl": {
					Type:     schema.TypeString,
					Computed: true,
				},
			},
		},
	}

	return &schema.Resource{
		ReadContext: datasourceFirewallFilterRead,
		Schema: map[string]*schema.Schema{
			//MetaId:           PropId(Id),
			//MetaResourcePath: PropResourcePath("/ip/firewall/"), // + section name

			// section = ["rules", "nat"]
			"section": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
					// Kebab case!
					ValidateFunc: validation.StringInSlice([]string{
						"address-list", "mangle", "nat", "rules",
					}, false),
				},
				Description: "Resource return sections: address-list, nat, mangle, rules (filter).",
			},

			"rules": firewallFilterSchema,
		},
	}
}

func datasourceFirewallFilterRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	s := DatasourceFirewallFilter().Schema
	basePath := "/ip/firewall/"

	sectionSet := d.Get("section").(*schema.Set)
	if sectionSet.Len() == 0 {
		return diag.Diagnostics{
			diag.Diagnostic{
				Severity: diag.Error,
				Summary:  `You must specify at least one return type of the resource: type = ["rules", "mangle", ...]`,
			},
		}
	}

	for _, v := range sectionSet.List() {
		path := basePath
		section := v.(string)
		// The filtering section is named 'rules' to avoid confusion: filter { filter = { ... }}.
		if section == "rules" {
			path += "filter"
		} else {
			path += section
		}

		ColorizedDebug(ctx, path)

		// Snake case!
		var res []MikrotikItem
		for _, sectionResourceData := range d.Get(KebabToSnake(section)).([]interface{}) {
			filter := sectionResourceData.(map[string]interface{})[KeyFilter].(map[string]interface{})
			r, err := ReadItemsFiltered(buildReadFilter(filter), path, m.(Client))
			if err != nil {
				return diag.FromErr(err)
			}

			res = append(res, *r...)
		}
		println(section)
		diags = append(diags, MikrotikResourceDataToTerraformDatasource(&res, section, s, d)...)

		//return MikrotikResourceDataToTerraformDatasource(res, "rules", s, d)
	}
	return diags
}
