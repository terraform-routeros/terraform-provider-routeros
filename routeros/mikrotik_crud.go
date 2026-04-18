package routeros

import (
	"context"
	"fmt"
)

// resource path is '/interface/vlan' etc.
// resource query is '/id' or '?.id=*39'.
var (
	errEmptyId   = fmt.Errorf("the resource id not defined")
	errEmptyItem = fmt.Errorf("the item is null")
	errEmptyPath = fmt.Errorf("the resource path not defined")
)

// https://help.mikrotik.com/docs/display/ROS/REST+API

func CreateItem(ctx context.Context, item MikrotikItem, resourcePath string, c Client) (MikrotikItem, error) {
	if item == nil {
		return nil, errEmptyItem
	}
	if resourcePath == "" {
		return nil, errEmptyPath
	}

	var crud = crudCreate
	invalidationPath := resourcePath

	if cm := ctxGetCrudMethod(ctx); cm != crudUnknown {
		crud = cm
		if c.GetTransport() == TransportREST {
			// apiMethodName[crud] is CLI path
			resourcePath += apiMethodName[crud]
		}
	}

	res := MikrotikItem{}
	err := c.SendRequest(crud, &URL{Path: resourcePath}, item, &res)
	if err == nil {
		if cache := c.GetBulkCache(); cache != nil {
			cache.Invalidate(invalidationPath)
		}
	}

	return res, err
}

func ReadItems(id *ItemId, resourcePath string, c Client) (*[]MikrotikItem, error) {
	// id can be empty.

	if resourcePath == "" {
		return nil, errEmptyPath
	}

	// Bulk-read cache fast path. Only kicks in for .id lookups on a managed
	// resource when the provider has "bulk_read_refresh" enabled. Datasource
	// reads (id==nil) and name-based lookups fall through to the original
	// filtered GET so their semantics are unchanged.
	if id != nil && id.Type == Id {
		if cache := c.GetBulkCache(); cache != nil {
			if item, itemFound, pathCached := cache.Lookup(resourcePath, id.Value); pathCached {
				if itemFound {
					return &[]MikrotikItem{item}, nil
				}
				return &[]MikrotikItem{}, nil
			}
			all, err := cache.DoBulkFetch(resourcePath, func() ([]MikrotikItem, error) {
				var items []MikrotikItem
				if err := c.SendRequest(crudRead, &URL{Path: resourcePath}, nil, &items); err != nil {
					return nil, err
				}
				return items, nil
			})
			if err != nil {
				return nil, err
			}
			for _, it := range all {
				if it.GetID(Id) == id.Value {
					return &[]MikrotikItem{it}, nil
				}
			}
			return &[]MikrotikItem{}, nil
		}
	}

	url := &URL{Path: resourcePath}

	// If the 'id' is nil, then this is a Datasource reading (resource Path only).
	if id != nil {
		// REST: prevent 404 'Not Found' error by direct resource request (/interface/vlan/*39).
		// Error occurs when a resource has been deleted outside terraform control.
		// But in the case below we have an empty [] or non-empty array [{...}].

		// /interface/vlan?.id=*39
		url.Query = []string{"?" + id.Type.String() + "=" + id.Value}
	}

	var res []MikrotikItem
	err := c.SendRequest(crudRead, url, nil, &res)

	return &res, err
}

func ReadItemsFiltered(filter []string, resourcePath string, c Client) (*[]MikrotikItem, error) {
	if resourcePath == "" {
		return nil, errEmptyPath
	}

	// Filter format: name=value
	// REST query: name=value; name=value
	// API  query: ?=name=value; ?=name=value
	if c.GetTransport() == TransportAPI {
		for i, s := range filter {
			filter[i] = "?=" + s
		}
	}
	url := &URL{Path: resourcePath, Query: filter}

	var res []MikrotikItem
	err := c.SendRequest(crudRead, url, nil, &res)

	return &res, err
}

func UpdateItem(id *ItemId, resourcePath string, item MikrotikItem, c Client) (MikrotikItem, error) {
	if id.Value == "" {
		return nil, errEmptyId
	}
	if resourcePath == "" {
		return nil, errEmptyPath
	}

	invalidationPath := resourcePath

	if c.GetTransport() == TransportREST {
		// /interface/vlan/*39
		resourcePath += "/" + id.Value
	} else {
		item[".id"] = id.Value
	}

	res := MikrotikItem{}
	err := c.SendRequest(crudUpdate, &URL{Path: resourcePath}, item, &res)
	if err == nil {
		if cache := c.GetBulkCache(); cache != nil {
			cache.Invalidate(invalidationPath)
		}
	}

	return res, err
}

func DeleteItem(id *ItemId, resourcePath string, c Client) error {
	if id.Value == "" {
		return errEmptyId
	}
	if resourcePath == "" {
		return errEmptyPath
	}

	invalidationPath := resourcePath
	url := &URL{Path: resourcePath}

	if c.GetTransport() == TransportREST {
		// This method is used to delete the record with a specified ID from the menu encoded in the URL.
		// If the deletion has been succeeded, the server responds with an empty response.
		// For example, call to delete the record twice, on second call router will return 404 error.

		// /interface/vlan/*39
		url.Path += "/" + id.Value
	} else {
		url.Query = []string{"=.id=" + id.Value}
	}

	err := c.SendRequest(crudDelete, url, nil, &MikrotikItem{})
	if err == nil {
		if cache := c.GetBulkCache(); cache != nil {
			cache.Invalidate(invalidationPath)
		}
	}
	return err
}
