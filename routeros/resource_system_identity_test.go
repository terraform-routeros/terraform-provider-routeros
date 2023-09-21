package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testSystemIdentityTask = "routeros_identity.test"

func TestAccSystemIdentityTest_basic(t *testing.T) {
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
						Config: testAccSystemIdentityConfig("TestRouter_" + name),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testSystemIdentityTask),
							resource.TestCheckResourceAttr(testSystemIdentityTask, "name", "TestRouter_"+name),
						),
					},
					{
						Config: testAccSystemIdentityConfig("MikroTik"),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(testSystemIdentityTask, "name", "MikroTik"),
						),
					},
				},
			})

		})
	}
}

func testAccSystemIdentityConfig(name string) string {
	return fmt.Sprintf(`

provider "routeros" {
	insecure = true
}

resource "routeros_identity" "test" {
	name = "%v"
}
`, name)
}
