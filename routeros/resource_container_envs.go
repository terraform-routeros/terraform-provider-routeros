package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// https://help.mikrotik.com/docs/display/ROS/Container#Container-Addenvironmentvariablesandmounts(optional)
func ResourceContainerEnvs() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/container/envs"),
		MetaId:           PropId(Id),

		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the environment variables list.",
		},
		"key": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Key of the environment variable.",
		},
		"value": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Value of the environment variable.",
		},
	}

	return &schema.Resource{
		CreateContext: DefaultCreate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultUpdate(resSchema),
		DeleteContext: DefaultDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}

// TODO: cleaner would be if we model it like envlist directly
/*
resource "routeros_container_envs" "test" {
  name = "test"

  env {
    key   = "foo"
    value = "bar"
  }

  env {
    key   = "hello"
    value = "world"
  }
}

resource "routeros_container_envs" "test_foo" {
	name  = "test"
	key   = "foo"
	value = "bar"
  }

  resource "routeros_container_envs" "test_hello" {
	name  = "test"
	key   = "hello"
	value = "wold"
  }
*/
