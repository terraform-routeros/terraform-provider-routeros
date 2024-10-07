package routeros

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DefaultCreate(s map[string]*schema.Schema) schema.CreateContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		return ResourceCreate(ctx, s, d, m)
	}
}

func DefaultCreateWithTimeout(s map[string]*schema.Schema, t time.Duration) schema.CreateContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		return ResourceCreate(ctx, s, d, m)
	}
}

func DefaultValidateCreate(s map[string]*schema.Schema, f DataValidateFunc) schema.CreateContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		if f != nil {
			if diags := f(d); diags.HasError() {
				return diags
			}
		}
		return ResourceCreate(ctx, s, d, m)
	}
}

func DefaultRead(s map[string]*schema.Schema) schema.ReadContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		return ResourceRead(ctx, s, d, m)
	}
}

func DefaultUpdate(s map[string]*schema.Schema) schema.UpdateContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		return ResourceUpdate(ctx, s, d, m)
	}
}

func DefaultValidateUpdate(s map[string]*schema.Schema, f DataValidateFunc) schema.UpdateContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		if f != nil {
			if diags := f(d); diags.HasError() {
				return diags
			}
		}
		return ResourceUpdate(ctx, s, d, m)
	}
}

func DefaultDelete(s map[string]*schema.Schema) schema.DeleteContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		return ResourceDelete(ctx, s, d, m)
	}
}

// Function to update resources that are present in the system by default out of the box.
// The distinctive feature of such resources is that they cannot be deleted, but they can be modified.
// For example, enabling/disabling the resource.
//
// FIXME Replace fucntions in resources: ResourceInterfaceEthernetSwitchPortIsolation, ResourceInterfaceEthernetSwitchPort
// ResourceInterfaceEthernetSwitch, ResourceInterfaceLte, ResourceIpService
func DefaultCreateUpdate(s map[string]*schema.Schema) func(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics {
	return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		item, metadata := TerraformResourceDataToMikrotik(s, d)

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

		return ResourceRead(ctx, s, d, m)
	}
}
