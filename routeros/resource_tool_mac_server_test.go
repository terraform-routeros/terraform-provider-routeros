package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testToolsMacServer = "routeros_tool_mac_server.test"

func TestAccToolsMacServerTest_basic(t *testing.T) {
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
						Config: testAccToolsMacServerConfig("none"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testToolsMacServer),
							resource.TestCheckResourceAttr(testToolsMacServer, "allowed_interface_list", "none"),
						),
					},
					{
						Config: testAccToolsMacServerConfig("all"),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(testToolsMacServer, "allowed_interface_list", "all"),
						),
					},
				},
			})
		})
	}
}

const testToolsMacServerWinBox = "routeros_tool_mac_server_winbox.test"

func TestAccToolsMacServerWinBoxTest_basic(t *testing.T) {
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
						Config: testAccToolsMacServerWinBoxConfig("none"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testToolsMacServerWinBox),
							resource.TestCheckResourceAttr(testToolsMacServerWinBox, "allowed_interface_list", "none"),
						),
					},
					{
						Config: testAccToolsMacServerWinBoxConfig("LAN"),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(testToolsMacServerWinBox, "allowed_interface_list", "LAN"),
						),
					},
				},
			})
		})
	}
}

func testAccToolsMacServerConfig(acl string) string {
	return fmt.Sprintf(`%v

resource "routeros_tool_mac_server" "test" {
	allowed_interface_list  = "%v"
}
`, providerConfig, acl)
}

func testAccToolsMacServerWinBoxConfig(acl string) string {
	return fmt.Sprintf(`%v

resource "routeros_tool_mac_server_winbox" "test" {
	allowed_interface_list  = "%v"
}
`, providerConfig, acl)
}
