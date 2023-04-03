package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testCapsManDatapathAddress = "routeros_capsman_datapath.test_datapath"

func TestAccCapsManDatapathTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/caps-man/datapath", "routeros_capsman_datapath"),
				Steps: []resource.TestStep{
					{
						Config: testAccCapsManDatapathConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckCapsManDatapathExists(testCapsManDatapathAddress),
							resource.TestCheckResourceAttr(testInterfaceBridgeAddress, "name", "test_datapath"),
						),
					},
				},
			})
		})
	}
}

func testAccCheckCapsManDatapathExists(name string) resource.TestCheckFunc {
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

func testAccCapsManDatapathConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_capsman_datapath" "test_datapath" {
	name   = "test_datapath"
  }

`
}
