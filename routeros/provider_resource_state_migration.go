package routeros

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func stateMigrationNameToId(resourcePath string) schema.StateUpgradeFunc {
	return func(ctx context.Context, rawState map[string]interface{}, m interface{}) (map[string]interface{}, error) {
		ColorizedMessage(ctx, INFO, fmt.Sprintf("ID attribute before migration: %#v", rawState["id"]))

		if rawState["id"] != nil {
			res, err := ReadItems(&ItemId{Name, rawState["id"].(string)}, resourcePath, m.(Client))
			if err != nil {
				return nil, err
			}

			// Resource not found.
			if len(*res) == 0 {
				rawState["id"] = ""
				ColorizedMessage(ctx, WARN, "No resource found, but the scheme has been updated.",
					map[string]interface{}{"path": resourcePath, "id": rawState["id"]})
				return rawState, nil
			}

			rawState["id"] = (*res)[0].GetID(Id)
		}

		ColorizedMessage(ctx, INFO, fmt.Sprintf("ID attribute after migration: %#v", rawState["id"]))

		return rawState, nil
	}
}

// stateMigrationClearInjectedDefault resets an attribute to an empty string if the state contains the given
// value. It is used to drop an obsolete schema-injected default from previously recorded states so that the
// attribute is no longer serialized unless it is explicitly configured.
func stateMigrationClearInjectedDefault(key, value string) schema.StateUpgradeFunc {
	return func(ctx context.Context, rawState map[string]interface{}, m interface{}) (map[string]interface{}, error) {
		if rawState[key] == value {
			rawState[key] = ""
		}

		return rawState, nil
	}
}

func stateMigrationScalarToList(keys ...string) schema.StateUpgradeFunc {
	return func(ctx context.Context, rawState map[string]interface{}, m interface{}) (map[string]interface{}, error) {
		for _, key := range keys {
			if rawState[key] == nil {
				continue
			}

			value := reflect.ValueOf(rawState[key])
			if value.IsZero() {
				rawState[key] = []interface{}{}
			}

			if reflect.ValueOf(value).Kind() == reflect.String {
				rawState[key] = strings.Split(rawState[key].(string), ",")
			}

			slice := reflect.MakeSlice(reflect.SliceOf(value.Type()), 0, 1)
			reflect.Append(slice, value)
			rawState[key] = slice.Interface()
		}

		return rawState, nil
	}
}
