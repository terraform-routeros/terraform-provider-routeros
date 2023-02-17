package routeros

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type DataValidateFunc func(d *schema.ResourceData) diag.Diagnostics

var errorNoLongerExists = errors.New("resource no longer exists")

// Dynamic resource ID lookup to save us from situations where we are trying to delete a resource
// that has been destroyed outside of Terraform. Always returns only the internal Mikrotik id!
func dynamicIdLookup(idType IdType, path string, c Client, d *schema.ResourceData) (string, error) {
	// Dynamic lookup id.
	res, err := ReadItems(&ItemId{idType, d.Id()}, path, c)
	if err != nil {
		// API/REST client error.
		return "", err
	}

	// Resource not found.
	if len(*res) == 0 {
		d.SetId("")
		return "", errorNoLongerExists
	}

	return (*res)[0].GetID(Id), nil
}

// ResourceCreate Creation of a resource in accordance with the TF Schema.
func ResourceCreate(ctx context.Context, s map[string]*schema.Schema, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	item, metadata := TerraformResourceDataToMikrotik(s, d)

	res, err := CreateItem(item, metadata.Path, m.(Client))
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

	// At this time, we have a successfully created resource,
	// regardless of the success of its reading.
	switch metadata.IdType {
	case Id:
		// Response ID.
		d.SetId(res.GetID(Id))
	case Name:
		// Resource ID.
		d.SetId(item.GetID(Name))
	}

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
	}

	//spew.Dump(res)
	return MikrotikResourceDataToTerraform(res, s, d)
}

// ResourceRead Reading some information about one specific resource.
func ResourceRead(ctx context.Context, s map[string]*schema.Schema, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	metadata := GetMetadata(s)

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

	d.SetId((*res)[0].GetID(metadata.IdType))

	return MikrotikResourceDataToTerraform((*res)[0], s, d)
}

// ResourceUpdate Updating the resource in accordance with the TF Schema.
func ResourceUpdate(ctx context.Context, s map[string]*schema.Schema, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	item, metadata := TerraformResourceDataToMikrotik(s, d)

	// d.Id() can be the name of a resource or its identifier.
	// Mikrotik only operates on resource ID!
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

	return MikrotikResourceDataToTerraform(res, s, d)
}

// ResourceDelete Deleting the resource.
func ResourceDelete(ctx context.Context, s map[string]*schema.Schema, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	metadata := GetMetadata(s)

	id, err := dynamicIdLookup(metadata.IdType, metadata.Path, m.(Client), d)
	if err != nil {
		if err != errorNoLongerExists {
			ColorizedDebug(ctx, fmt.Sprintf(ErrorMsgDelete, err))
			return diag.FromErr(err)
		}

		// We inform the user that the resource no longer exists.
		d.SetId("")
		return diag.Diagnostics{
			diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  errorNoLongerExists.Error(),
			},
		}
	}

	if err := DeleteItem(&ItemId{Id, id}, metadata.Path, m.(Client)); err != nil {
		ColorizedDebug(ctx, fmt.Sprintf(ErrorMsgDelete, err))
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}

func DefaultCreate(s map[string]*schema.Schema) schema.CreateContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

func DefaultDelete(s map[string]*schema.Schema) schema.DeleteContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		return ResourceDelete(ctx, s, d, m)
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
