package routeros

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  {
    ".id": "*0",
    "dynamic": "true",
    "fib": "",
    "invalid": "false",
    "name": "main"
  },
*/

// https://help.mikrotik.com/docs/display/ROS/Policy+Routing
func ResourceRoutingTable() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/routing/table"),
		MetaId:           PropId(Id),

		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		KeyDynamic:  PropDynamicRo,
		"fib": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			ForceNew:    true,
			Description: "fib parameter should be specified if the routing table is intended to push routes to the FIB.",
		},
		KeyInvalid: PropInvalidRo,
		"name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Routing table name.",
		},
	}

	return &schema.Resource{
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			metadata := GetMetadata(resSchema)

			res, err := ReadItems(&ItemId{metadata.IdType, d.Id()}, metadata.Path, m.(Client))
			if err != nil {
				ColorizedDebug(ctx, fmt.Sprintf(ErrorMsgGet, err))
				return diag.FromErr(err)
			}

			// Resource not found.
			if len(*res) == 0 {
				d.SetId("")
				return nil
			}

			r := (*res)[0]

			d.SetId(r.GetID(metadata.IdType))

			// These are crutches, but I don't know an easier way to implement this strange logic.
			if _, ok := r["fib"]; ok {
				r["fib"] = "yes"
			} else {
				r["fib"] = "no"
			}

			return MikrotikResourceDataToTerraform(r, resSchema, d)
		},

		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			item, metadata := TerraformResourceDataToMikrotik(resSchema, d)

			if b, ok := item["fib"]; ok {
				if b == "no" {
					delete(item, "fib")
				} else {
					item["fib"] = ""
				}
			}

			res, err := CreateItem(ctx, item, metadata.Path, m.(Client))
			if err != nil {
				ColorizedDebug(ctx, fmt.Sprintf(ErrorMsgPut, err))
				return diag.FromErr(err)
			}

			// ... If no ID is set, Terraform assumes the resource was not created successfully;
			// as a result, no state will be saved for that resource.
			if res.GetID(Id) == "" {
				return diag.Diagnostics{
					diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "The resource ID was not found in the response",
					},
				}
			}

			// Response ID.
			d.SetId(res.GetID(Id))

			// We ask for information again in the case of API.
			if m.(Client).GetTransport() == TransportAPI {
				r, err := ReadItems(&ItemId{Id, res.GetID(Id)}, metadata.Path, m.(Client))
				if err != nil {
					ColorizedDebug(ctx, fmt.Sprintf(ErrorMsgPut, err))
					return diag.FromErr(err)
				}

				if len(*r) == 0 {
					return diag.Diagnostics{
						diag.Diagnostic{
							Severity: diag.Error,
							Summary: fmt.Sprintf("Mikrotik resource path='%v' id='%v' not found",
								metadata.Path, res.GetID(Id)),
						},
					}
				}

				res = (*r)[0]
				if _, ok := res["fib"]; ok {
					res["fib"] = "yes"
				} else {
					res["fib"] = "no"
				}
			}

			return MikrotikResourceDataToTerraform(res, resSchema, d)
		},

		UpdateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			item, metadata := TerraformResourceDataToMikrotik(resSchema, d)

			if b, ok := item["fib"]; ok {
				if b == "no" {
					delete(item, "fib")
				} else {
					item["fib"] = ""
				}
			}

			id, err := dynamicIdLookup(metadata.IdType, metadata.Path, m.(Client), d)
			if err != nil {
				// There is nothing to update, because resource id not found
				// or some other error.
				ColorizedDebug(ctx, fmt.Sprintf(ErrorMsgPatch, err))
				return diag.FromErr(err)
			}

			res, err := UpdateItem(&ItemId{Id, id}, metadata.Path, item, m.(Client))
			if err != nil {
				ColorizedDebug(ctx, fmt.Sprintf(ErrorMsgPatch, err))
				return diag.FromErr(err)
			}

			if _, ok := res["fib"]; ok {
				res["fib"] = "yes"
			} else {
				res["fib"] = "no"
			}

			return MikrotikResourceDataToTerraform(res, resSchema, d)
		},

		DeleteContext: DefaultDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
