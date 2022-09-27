package routeros

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
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
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: testAccSystemIdentityConfig("TestRouter_" + name),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckSystemIdentityExists(testSystemIdentityTask),
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

func testAccCheckSystemIdentityExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("not found: %s", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no id is set")
		}

		return nil
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
