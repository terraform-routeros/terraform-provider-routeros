package routeros

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

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

// Passing the called CRUD method on creation through an existing context.
type ctxCrudMethod string

const ctxCrudMethodKey = "crudMethod"

// Specifies a CRUD method as part of the resource schema description.
func ctxSetCrudMethod(ctx context.Context, m crudMethod) context.Context {
	return context.WithValue(ctx, ctxCrudMethod(ctxCrudMethodKey), m)
}

// Retrieve a CRUD method as part of the processing  of a resource creation request.
func ctxGetCrudMethod(ctx context.Context) crudMethod {
	if v := ctx.Value(ctxCrudMethod(ctxCrudMethodKey)); v != nil {
		return v.(crudMethod)
	}
	return crudUnknown
}

// ResourceCreate - Creation of a resource in accordance with the TF Schema.
// It is possible to transparently pass the request type (CRUD Method) within an existing context.
//
//	CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
//		return ResourceCreate(ctxSetCrudMethod(ctx, crudGenerateKey), resSchema, d, m)
//	},
func ResourceCreate(ctx context.Context, s map[string]*schema.Schema, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	item, metadata := TerraformResourceDataToMikrotik(s, d)

	res, err := CreateItem(ctx, item, metadata.Path, m.(Client))
	if err != nil {
		ColorizedDebug(ctx, fmt.Sprintf(ErrorMsgPut, err))
		return diag.FromErr(err)
	}

	// Some resources may return an empty array as a response when executing commands other than 'create'.
	// For these cases, we will try to find the created element by name (if available).
	if res.GetID(Id) == "" && item[KeyName] != "" {
		items, err := ReadItems(&ItemId{Name, item[KeyName]}, metadata.Path, m.(Client))
		if err != nil {
			return diag.FromErr(err)
		}

		if items != nil && len(*items) == 1 {
			res = (*items)[0]
		}
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

func ResourceCreateAndWait(ctx context.Context, s map[string]*schema.Schema, d *schema.ResourceData, m interface{}, timeout time.Duration) diag.Diagnostics {
	item, metadata := TerraformResourceDataToMikrotik(s, d)
	if item[KeyName] == "" {
		panic("Asynchronous resource creation should be applied to objects that have the 'name' attribute.")
	}
	ColorizedDebug(ctx, fmt.Sprintf("Wait timeout is %s", timeout))

	// The lifetime of a REST session is 60 seconds.
	_, err := CreateItem(ctx, item, metadata.Path, m.(Client))
	if err != nil {
		// context deadline exceeded (Client.Timeout exceeded while awaiting headers)
		// {"detail":"Session closed","error":400,"message":"Bad Request"}
		// from RouterOS device: action timed out - try again, if error continues contact MikroTik support and send a supout file (13)
		if !strings.Contains(err.Error(), context.DeadlineExceeded.Error()) && !strings.Contains(err.Error(), "Session closed") &&
			!strings.Contains(err.Error(), "action timed out - try again") {
			ColorizedDebug(ctx, fmt.Sprintf(ErrorMsgPut, err))
			return diag.FromErr(err)
		}
		ColorizedDebug(ctx, "Timeout, the Create context is canceled, waiting for the resource to be created. "+
			"Session termination by MikroTik is ignored.")
	}

	// context deadline exceeded
	var res MikrotikItem

	// During RSA key generation, we can get 100% CPU utilization of the MT.
	// During this time MT may stop accepting external requests!
	localCtx, cancel := context.WithTimeout(context.Background(), timeout)
	attempt := 0
	for {
		defer cancel()

		// We will try to find the created element by name (if available).
		items, err := ReadItems(&ItemId{Name, item[KeyName]}, metadata.Path, m.(Client))
		if err != nil {
			ColorizedMessage(ctx, TRACE, fmt.Sprintf("Timeout, the Read context is canceled, waiting for the resource to be created. "+
				"Session termination by MikroTik is ignored. Attempt #%v", attempt))
			attempt++
		}

		if items != nil && len(*items) == 1 {
			res = (*items)[0]
			break
		}

		select {
		case <-localCtx.Done():
			// The context deadline has been exceeded.
			return diag.Diagnostics{
				diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "The resource ID was not found in the response",
				},
			}
		default:
			time.Sleep(15 * time.Second)
		}

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

// SystemResourceRead The difference from the normal reading is in the method of generation of Id.
func SystemResourceRead(ctx context.Context, s map[string]*schema.Schema, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	metadata := GetMetadata(s)

	res := MikrotikItem{}
	err := m.(Client).SendRequest(crudRead, &URL{Path: metadata.Path}, nil, &res)
	if err != nil {
		return diag.FromErr(err)
	}

	// We make a unique Id, it does not affect the work with the Mikrotik.
	// Id: /caps-man/manager -> caps-man.manager
	d.SetId(strings.ReplaceAll(strings.TrimLeft(metadata.Path, "/"), "/", "."))

	return MikrotikResourceDataToTerraform(res, s, d)
}

// SystemResourceCreateUpdate A resource cannot be created, it can only be changed.
func SystemResourceCreateUpdate(ctx context.Context, s map[string]*schema.Schema, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	item, metadata := TerraformResourceDataToMikrotik(s, d)

	var resUrl string
	if m.(Client).GetTransport() == TransportREST {
		// https://router/rest/system/identity/set
		// https://router/rest/caps-man/manager/set
		resUrl = "/set"
	}

	err := m.(Client).SendRequest(crudPost, &URL{Path: metadata.Path + resUrl}, item, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	return SystemResourceRead(ctx, s, d, m)
}

// SystemResourceDelete Delete function will remove the object from the Terraform state
// No delete functionality provided by API for System Resources.
func SystemResourceDelete(ctx context.Context, s map[string]*schema.Schema, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("")
	if m.(Client).GetExtraParams().SuppressSysODelWarn {
		return nil
	}
	return DeleteSystemObject
}

// ImportStateCustomContext is an implementation of StateContextFunc that can be used to
// import resources with the ability to explicitly or implicitly specify a key field.
// `terraform [global options] import [options] ADDR ID`.
// During import the content of the `ID` is checked and depending on the specified string it is possible to automatically search for the internal Mikrotik identifier.
// Logic of `ID` processing
// - The first character of the string contains an asterisk (standard Mikrotik identifier `*3E`): import without additional search.
// - String containing no "=" character (`wifi-01`): the "name" field is used for searching.
// - String containing only one "=" character (`"comment=hAP-ac3"`): the "word left" and "word right" pair is used for searching.
func ImportStateCustomContext(s map[string]*schema.Schema) schema.StateContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
		id := d.Id()
		fieldName := "name"

		if len(id) == 0 || id[0] == '*' {
			return []*schema.ResourceData{d}, nil
		} else {
			// By default, we filter by the "name" field
			if s := strings.Split(id, "="); len(s) == 2 {
				// field=value
				fieldName = s[0]
				id = s[1]
			}
		}

		path := s[MetaResourcePath].Default.(string)

		res, err := ReadItemsFiltered([]string{SnakeToKebab(fieldName) + "=" + id}, path, m.(Client))
		if err != nil {
			return nil, err
		}

		switch len(*res) {
		case 0:
			return nil, fmt.Errorf("resource not found: %v=%v", fieldName, id)
		case 1:
			retId, ok := (*res)[0][Id.String()]
			if !ok {
				return nil, fmt.Errorf("attribute %v not found in the response", Id.String())
			}
			d.SetId(retId)
		default:
			return nil, fmt.Errorf("more than one resource found: %v=%v", fieldName, id)
		}

		return []*schema.ResourceData{d}, nil
	}
}
