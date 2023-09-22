package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfaceListMemberAddress = "routeros_interface_list_member.test_list_member"

func TestAccInterfaceListMemberTest_basic(t *testing.T) {
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
							testResourcePrimaryInstanceId(testInterfaceListMemberAddress),
							resource.TestCheckResourceAttr(testInterfaceListMemberAddress, "interface", "ether1"),
						),
					},
				},
			})

		})
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
