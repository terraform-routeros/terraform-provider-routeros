package routeros

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceMoveItems() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/path"),
		MetaId:           PropId(Id),

		"resource_name": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Resource name in the notation ```routeros_ip_firewall_filter```.",
			ValidateFunc: validation.StringMatch(regexp.MustCompile(`^routeros(_\w+)+$`), ""),
			AtLeastOneOf: []string{"resource_name", "resource_path"},
		},
		"resource_path": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "URL path of the resource in the notation ```/ip/firewall/filter```.",
			ValidateFunc: validation.StringMatch(regexp.MustCompile(`^(/\w+)+$`), ""),
			AtLeastOneOf: []string{"resource_name", "resource_path"},
		},
		"sequence": {
			Type:     schema.TypeList,
			Required: true,
			Description: "List identifiers in the required sequence. To locate the ```sequence``` before an " +
				"existing rule, add its ```id``` to the last element of the ```sequence```.",
			Elem: &schema.Schema{
				Type:     schema.TypeString,
				MinItems: 2,
			},
		},
		// "anchor_rule": {
		// 	Type:     schema.TypeString,
		// 	Optional: true,
		// 	Description: "The rule before which the ```sequence``` of rules will be placed. If this field is not specified, " +
		// 		"the rules will be placed before the last element of the ```sequence```.",
		// },
	}
	resRead := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		path, ok := d.GetOk("resource_path")
		if !ok {
			path = d.Get("resource_name")
			path = strings.TrimPrefix(path.(string), "routeros_")
			path = strings.ReplaceAll(path.(string), "_", "/")
		}

		res, err := ReadItems(nil, path.(string), m.(Client))
		if err != nil {
			ColorizedDebug(ctx, fmt.Sprintf(ErrorMsgGet, err))
			return diag.FromErr(err)
		}

		// Resource not found.
		if len(*res) == 0 {
			d.SetId("")
			return nil
		}

		var conf = make(map[string]struct{})
		for _, v := range d.Get("sequence").([]any) {
			conf[v.(string)] = struct{}{}
		}

		// TODO It is necessary to sometime transfer the logic for monotonous reading of the sequence, so that it is possible
		// to control the correctness of the position of the rules and possibly introduce an anchor rule.
		// This check will not work in all cases!!!
		var list []string
		for _, r := range *res {
			if id, ok := r[".id"]; ok {
				if _, ok := conf[id]; ok {
					list = append(list, id)
				}
			}
		}

		d.Set("sequence", list)
		return nil
	}

	resCreateUpdate := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		var list []string
		for _, v := range d.Get("sequence").([]any) {
			list = append(list, v.(string))
		}

		// if anchor, ok := d.GetOk("anchor_rule"); ok {
		// 	item = MikrotikItem{
		// 		"numbers":     strings.Join(list, ","),
		// 		"destination": anchor.(string),
		// 	}
		// } else {
		var item MikrotikItem = MikrotikItem{
			"numbers":     strings.Join(list[:len(list)-1], ","),
			"destination": list[len(list)-1],
		}
		// }

		path, ok := d.GetOk("resource_path")
		if !ok {
			path = d.Get("resource_name")
			path = strings.TrimPrefix(path.(string), "routeros_")
			path = strings.ReplaceAll(path.(string), "_", "/")
		}

		if m.(Client).GetTransport() == TransportREST {
			path = path.(string) + "/move"
		}
		err := m.(Client).SendRequest(crudMove, &URL{Path: path.(string)}, item, nil)
		if err != nil {
			ColorizedDebug(ctx, fmt.Sprintf(ErrorMsgPut, err))
			return diag.FromErr(err)
		}

		d.SetId(strings.ReplaceAll(strings.TrimLeft(path.(string), "/"), "/", "."))

		return resRead(ctx, d, m)
	}

	return &schema.Resource{
		CreateContext: resCreateUpdate,
		ReadContext:   resRead,
		UpdateContext: resCreateUpdate,
		DeleteContext: DefaultSystemDelete(resSchema),

		Schema: resSchema,
	}
}
