package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfaceListAddress = "routeros_interface_list.test_list"

func TestAccInterfaceListTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/list", "routeros_interface_list"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceListConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceListAddress),
							resource.TestCheckResourceAttr(testInterfaceListAddress, "name", "test_list"),
						),
					},
				},
			})

		})
	}
}

func testAccInterfaceListConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_interface_list" "test_list" {
	name      = "test_list"
}
`
}
