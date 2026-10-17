package routeros

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DatasourceApiCommand is a generic, read-only data source that runs an
// arbitrary RouterOS API command (typically a '.../print'-style command)
// against any menu path and returns the raw response.
//
// It exists as an escape hatch for menus/queries that don't have a
// dedicated typed resource or data source yet (e.g. counting IPv6 firewall
// address-list entries, which routeros_ipv6_firewall does not expose), and
// it works transparently with both the REST ('http(s)://') and the classic
// API ('api(s)://') provider transports, using the same Client abstraction
// (ReadItemsRaw -> Client.SendRequest) as every other data source.
func DatasourceApiCommand() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceApiCommandRead,
		Description: "Runs an arbitrary, read-only RouterOS API command (a '.../print'-style command) against any " +
			"menu path and returns its result. Use this to query things that don't have a dedicated typed data " +
			"source yet, e.g. counting entries in an IPv6 firewall address-list.",
		Schema: map[string]*schema.Schema{
			"path": {
				Type:     schema.TypeString,
				Required: true,
				Description: "RouterOS API menu path to run the command against, e.g. `/ipv6/firewall/address-list`. " +
					"A trailing `/print` is accepted and stripped automatically, since the print command is always " +
					"the one that gets executed.",
			},
			"words": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Description: "Raw RouterOS API sentence words/arguments for the command, e.g. " +
					"`[\"=count-only=\", \"?list=blocklist-ipv64-dshield1\"]`. Words starting with `?` are treated " +
					"as print filters (RouterOS 'where' clauses, e.g. `?list=blocklist-ipv64-dshield1`); all other " +
					"words are passed as literal command arguments (a leading `=` is optional and is stripped if " +
					"present, e.g. `=count-only=` and `count-only` are equivalent).",
			},
			"result": {
				Type:     schema.TypeList,
				Computed: true,
				Description: "The rows returned by the command, each as a map of RouterOS property name to string " +
					"value. Commands that return a single value instead of rows (e.g. with `count-only`) yield a " +
					"single map, typically containing a `ret` key with the value.",
				Elem: &schema.Schema{
					Type: schema.TypeMap,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
		},
	}
}

func datasourceApiCommandRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	path := strings.TrimSuffix(d.Get("path").(string), "/print")

	var filter []string
	item := MikrotikItem{}

	for _, w := range d.Get("words").([]interface{}) {
		word := w.(string)

		if strings.HasPrefix(word, "?") {
			filter = append(filter, strings.TrimPrefix(word, "?"))
			continue
		}

		key, value, _ := strings.Cut(strings.TrimPrefix(word, "="), "=")
		item[key] = value
	}

	// Keep parity with the plain-read code path when no arguments are given.
	if len(item) == 0 {
		item = nil
	}

	res, err := ReadItemsRaw(filter, item, path, m.(Client))
	if err != nil {
		return diag.FromErr(err)
	}

	result := make([]interface{}, 0, len(*res))
	for _, mikrotikItem := range *res {
		row := make(map[string]interface{}, len(mikrotikItem))
		for k, v := range mikrotikItem {
			row[k] = v
		}
		result = append(result, row)
	}

	if err := d.Set("result", result); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(UniqueId())

	return nil
}
