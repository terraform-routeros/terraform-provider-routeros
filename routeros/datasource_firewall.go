package routeros

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var firewallSections = []string{"address_list", "nat", "mangle", "rules"}

func DatasourceFirewall() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceFirewallFilterRead,
		Description: `This datasource contains all supported firewall resources:
- address_list
- nat
- mangle
- rules (aka filter)
`,
		Schema: map[string]*schema.Schema{
			"address_list": getFirewallAddrListSchema(),
			"mangle":       getFirewallMangleSchema(),
			"nat":          getFirewallNatSchema(),
			"rules":        getFirewallFilterSchema(),
		},
	}
}

func datasourceFirewallFilterRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	basePath := "/ip/firewall/"

	s := DatasourceFirewall().Schema

	var isEmpty = true
	for _, section := range firewallSections {
		isEmpty = isEmpty && len(d.Get(section).([]interface{})) == 0
	}

	if isEmpty {
		return diag.Diagnostics{
			diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "You must specify at least one return section of the resource.",
				Detail: "Please specify one or more sections of the firewall, information from which will be " +
					"returned as a result of the data source query: rules{}, nat { filter = {...}}, etc.",
			},
		}
	}

	for _, section := range firewallSections {
		if len(d.Get(section).([]interface{})) == 0 {
			continue
		}

		path := basePath
		// The filtering section is named 'rules' to avoid confusion: filter { filter = { ... }}.
		if section == "rules" {
			path += "filter"
		} else {
			// Kebab case!
			path += SnakeToKebab(section)
		}

		// Snake case!
		var res []MikrotikItem

		for _, sectionResourceData := range d.Get(section).([]interface{}) {
			filter := sectionResourceData.(map[string]interface{})[KeyFilter].(map[string]interface{})

			r, err := ReadItemsFiltered(buildReadFilter(filter), path, m.(Client))
			if err != nil {
				return diag.FromErr(err)
			}

			res = append(res, *r...)
		}
		diags = append(diags, MikrotikResourceDataToTerraformDatasource(&res, section, s, d)...)
	}
	return diags
}
