package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfaceL2tpServerSettings = "routeros_interface_l2tp_server.settings"

func TestAccInterfaceL2tpServerTest_basic(t *testing.T) {
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
						Config: testAccInterfaceL2tpServerConfig("false"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceL2tpServerSettings),
							resource.TestCheckResourceAttr(testInterfaceL2tpServerSettings, "enabled", "false"),
						),
					},
					{
						Config: testAccInterfaceL2tpServerConfig("true"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceL2tpServerSettings),
							resource.TestCheckResourceAttr(testInterfaceL2tpServerSettings, "enabled", "true"),
						),
					},
				},
			})

		})
	}
}

func testAccInterfaceL2tpServerConfig(param string) string {
	return fmt.Sprintf(`%v

resource "routeros_interface_l2tp_server" "settings" {
  enabled = %v
}
`, providerConfig, param)
}
