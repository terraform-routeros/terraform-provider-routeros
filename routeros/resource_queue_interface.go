package routeros

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  ".id": "*A",
  "active-queue": "ethernet-default",
  "default-queue": "only-hardware-queue",
  "interface": "ether1",
  "queue": "ethernet-default"
}
*/

// https://help.mikrotik.com/docs/display/ROS/Queues
// ResourceQueueInterface assigns an interface queue type to an interface. The
// rows are auto-created by RouterOS (one per interface), so this resource adopts
// an existing row keyed by `interface` and can only modify it, never add or
// delete it.
func ResourceQueueInterface() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/queue/interface"),
		MetaId:           PropId(Id),

		"interface": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "Name of the interface whose queue is being configured.",
		},
		"queue": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the interface queue type to assign (e.g. `ethernet-default`, `only-hardware-queue`, `no-queue`).",
		},
		"active_queue": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Queue type currently active on the interface.",
		},
		"default_queue": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Default queue type for the interface.",
		},
	}

	resCreateUpdate := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		item, metadata := TerraformResourceDataToMikrotik(resSchema, d)

		res, err := ReadItemsFiltered([]string{"interface=" + d.Get("interface").(string)}, metadata.Path, m.(Client))
		if err != nil {
			ColorizedDebug(ctx, fmt.Sprintf(ErrorMsgPatch, err))
			return diag.FromErr(err)
		}

		if len(*res) == 0 {
			d.SetId("")
			return diag.FromErr(errorNoLongerExists)
		}

		d.SetId((*res)[0].GetID(Id))
		item[".id"] = d.Id()

		// interface identifies the (auto-created) row and is read-only on it.
		delete(item, "interface")

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
