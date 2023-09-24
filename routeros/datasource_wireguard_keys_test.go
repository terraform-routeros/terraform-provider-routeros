package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testDatasourceWireGuardKeys = "data.routeros_wireguard_keys.keys"

func TestAccDatasourceWireGuardKeys_basic(t *testing.T) {
	t.Run("WG keys", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProviderFactories: testAccProviderFactories,
			Steps: []resource.TestStep{
				{
					Config: testAccDatasourceWireGuardKeysConfig(),
					Check: resource.ComposeTestCheckFunc(
						testResourcePrimaryInstanceId(testDatasourceWireGuardKeys),
					),
				},
			},
		})

	})
}

func testAccDatasourceWireGuardKeysConfig() string {
	return `

provider "routeros" {
	insecure = true
}

data "routeros_wireguard_keys" "keys" {
	number = 3
}
`
}
