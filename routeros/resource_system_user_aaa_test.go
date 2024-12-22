package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testUserAAA = "routeros_system_user_aaa.settings"

func TestAccUserAAATest_basic(t *testing.T) {
	// t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				Steps: []resource.TestStep{
					{
						Config: testAccUserAAAConfig("true"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testUserAAA),
							resource.TestCheckResourceAttr(testUserAAA, "use_radius", "true"),
						),
					},
					{
						Config: testAccUserAAAConfig("false"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testUserAAA),
							resource.TestCheckResourceAttr(testUserAAA, "use_radius", "false"),
						),
					},
				},
			})

		})
	}
}

func testAccUserAAAConfig(param string) string {
	return fmt.Sprintf(`%v

resource "routeros_system_user_aaa" "settings" {
	use_radius = %v
}
`, providerConfig, param)
}
