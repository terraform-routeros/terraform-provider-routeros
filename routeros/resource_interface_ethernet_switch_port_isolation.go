package routeros

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  ".id": "*1",
  "forwarding-override": "ether1",
  "invalid": "false",
  "name": "ether1",
  "switch": "switch1"
}
*/

// https://help.mikrotik.com/docs/display/ROS/Switch+Chip+Features#SwitchChipFeatures-Portisolation
func ResourceInterfaceEthernetSwitchPortIsolation() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath:   PropResourcePath("/interface/ethernet/switch/port-isolation"),
		MetaId:             PropId(Id),
		MetaSkipFields:     PropSkipFields("name"),
		MetaSetUnsetFields: PropSetUnsetFields("forwarding_override"),

		KeyInvalid: PropInvalidRo,
		KeyName:    PropName("Port name."),
		"switch": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Name of the switch.",
		},
		"forwarding_override": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Forces ingress traffic to be forwarded to a specific interface. Multiple interfaces can be specified by separating them with a comma.",
		},
	}

	resCreateUpdate := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		item, metadata := TerraformResourceDataToMikrotik(resSchema, d)

		res, err := ReadItems(&ItemId{Name, d.Get("name").(string)}, metadata.Path, m.(Client))
		if err != nil {
			// API/REST client error.
			ColorizedDebug(ctx, fmt.Sprintf(ErrorMsgPatch, err))
			return diag.FromErr(err)
		}

		// Resource not found.
		if len(*res) == 0 {
			d.SetId("")
			ColorizedDebug(ctx, fmt.Sprintf(ErrorMsgPatch, err))
			return diag.FromErr(errorNoLongerExists)
		}

		d.SetId((*res)[0].GetID(Id))
		item[".id"] = d.Id()

		var resUrl string
		if m.(Client).GetTransport() == TransportREST {
			resUrl = "/set"
		}

		err = m.(Client).SendRequest(crudPost, &URL{Path: metadata.Path + resUrl}, item, nil)
		if err != nil {
			return diag.FromErr(err)
		}

		return ResourceRead(ctx, resSchema, d, m)
	}

	return &schema.Resource{
		CreateContext: resCreateUpdate,
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: resCreateUpdate,
		DeleteContext: DefaultSystemDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: ImportStateCustomContext(resSchema),
		},

		Schema: resSchema,
	}
}
