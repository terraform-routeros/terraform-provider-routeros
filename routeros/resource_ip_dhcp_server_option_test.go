package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testIpDhcpServerOption = "routeros_ip_dhcp_server_option.test_option"

func TestAccIpDhcpServerNetworkOption_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/dhcp-server/network", "routeros_ip_dhcp_server_option"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpDhcpServerOptionConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckIpDhcpServerOptionExists(testIpDhcpServerOption),
							resource.TestCheckResourceAttr(testIpDhcpServerOption, "code", "77"),
							resource.TestCheckResourceAttr(testIpDhcpServerOption, "name", "test-opt"),
							resource.TestCheckResourceAttr(testIpDhcpServerOption, "value", "s'10.10.10.22'"),

							resource.TestCheckResourceAttr(testIpDhcpServerOption, "raw_value", "31302e31302e31302e3232"),
						),
					},
				},
			})

		})
	}
}

func testAccCheckIpDhcpServerOptionExists(name string) resource.TestCheckFunc {
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

func testAccIpDhcpServerOptionConfig() string {
	return providerConfig + `
resource "routeros_ip_dhcp_server_option" "test_option" {
	code    = 77
	name    = "test-opt"
    value   = "s'10.10.10.22'"
  }
`
}
