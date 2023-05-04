package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testInterfaceListMemberAddress = "routeros_interface_list_member.test_list_member"

func TestAccInterfaceListMemberTest_basic(t *testing.T) {
	t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/list/member", "routeros_interface_list_member"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceListMemberConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckInterfaceListMemberExists(testInterfaceListMemberAddress),
							resource.TestCheckResourceAttr(testInterfaceListMemberAddress, "interface", "ether1"),
						),
					},
				},
			})

		})
	}
}

func testAccCheckInterfaceListMemberExists(name string) resource.TestCheckFunc {
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

func testAccInterfaceListMemberConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_interface_list_member" "test_list_member" {
	interface      = "ether1"
	list           = "list"
}
`
}
