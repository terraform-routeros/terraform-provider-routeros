package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testUserAddress = "routeros_system_user.test"

func TestAccUserTest_basic(t *testing.T) {
	// t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/user", "routeros_system_user"),
				Steps: []resource.TestStep{
					{
						Config: testAccUserConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testUserAddress),
							resource.TestCheckResourceAttr(testUserAddress, "name", "test-user-1"),
						),
					},
				},
			})

		})
	}
}

func testAccUserConfig() string {
	return providerConfig + `

resource "routeros_system_user" "test" {
	name        = "test-user-1"
	address     = "0.0.0.0/0"
	group       = "read"
	password    = "secret"
	comment     = "Test User"
}
`
}
