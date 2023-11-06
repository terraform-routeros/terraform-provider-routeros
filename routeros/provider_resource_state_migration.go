package routeros

import (
	"context"
	"fmt"

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
