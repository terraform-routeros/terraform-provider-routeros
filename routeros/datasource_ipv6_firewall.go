package routeros

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var ipv6firewallSections = []string{"rules"}

func DatasourceIPv6Firewall() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceIPv6FirewallFilterRead,
		Description: `This datasource contains all supported firewall resources:
- rules (aka filter)
`,
		Schema: map[string]*schema.Schema{
			"rules": getIPv6FirewallFilterSchema(),
		},
	}
}

func datasourceIPv6FirewallFilterRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	basePath := "/ipv6/firewall/"

	s := DatasourceIPv6Firewall().Schema

	var isEmpty = true
	for _, section := range ipv6firewallSections {
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

	for _, section := range ipv6firewallSections {
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
