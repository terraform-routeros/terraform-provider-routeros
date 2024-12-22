package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testUserGroupAddress = "routeros_system_user_group.test"

func TestAccUserGroupTest_basic(t *testing.T) {
	// t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/user/group", "routeros_system_user_group"),
				Steps: []resource.TestStep{
					{
						Config: testAccUserGroupConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testUserGroupAddress),
							resource.TestCheckResourceAttr(testUserGroupAddress, "name", "terraform"),
						),
					},
				},
			})

		})
	}
}

func testAccUserGroupConfig() string {
	return providerConfig + `

resource "routeros_system_user_group" "test" {
	name        = "terraform"
	policy      = ["api", "!ftp", "!local", "password", "policy", "read", "!reboot", "!rest-api", "!romon", "sensitive", "!sniff", "!ssh", "!telnet", "!test", "!web", "!winbox", "write"]
}	
`
}
