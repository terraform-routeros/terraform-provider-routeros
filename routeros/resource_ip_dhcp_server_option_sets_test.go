package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpDhcpServerOptionSet = "routeros_ip_dhcp_server_option_set.test_option_set"

func TestAccIpDhcpServerNetworkOptionSet_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/dhcp-server/option/sets", "routeros_ip_dhcp_server_option_set"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpDhcpServerOptionSetConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpDhcpServerOptionSet),
							resource.TestCheckResourceAttr(testIpDhcpServerOptionSet, "name", "test-opt-set"),
							resource.TestCheckResourceAttr(testIpDhcpServerOptionSet, "options", "test-opt1,test-opt2"),
						),
					},
				},
			})

		})
	}
}

func testAccIpDhcpServerOptionSetConfig() string {
	return providerConfig + `
resource "routeros_ip_dhcp_server_option" "test_option_1" {
	code    = 77
	name    = "test-opt1"
    value   = "s'10.10.10.22'"
  }

resource "routeros_ip_dhcp_server_option" "test_option_2" {
	code    = 90
	name    = "test-opt2"
    value   = "s'10.10.10.22'"
  }

resource "routeros_ip_dhcp_server_option_set" "test_option_set" {
	name      = "test-opt-set"
    options   = join(",", [routeros_ip_dhcp_server_option.test_option_1.name, routeros_ip_dhcp_server_option.test_option_2.name])
  }
`
}
