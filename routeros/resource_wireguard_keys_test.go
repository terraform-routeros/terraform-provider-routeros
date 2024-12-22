package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testResourceWireGuardKeys = "routeros_wireguard_keys.keys"

func TestAccResourceWireGuardKeys_basic(t *testing.T) {
	// t.Parallel()
	t.Run("WG keys", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProviderFactories: testAccProviderFactories,
			Steps: []resource.TestStep{
				{
					Config: testAccResourceWireGuardKeysConfig(),
					Check: resource.ComposeTestCheckFunc(
						testResourcePrimaryInstanceId(testResourceWireGuardKeys),
						resource.TestCheckResourceAttr(testResourceWireGuardKeys, "number", "3"),
					),
				},
			},
		})

	})
}

func testAccResourceWireGuardKeysConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_wireguard_keys" "keys" {
	number = 3
}
`
}
