package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpDhcpClientOptionAddress = "routeros_ip_dhcp_client_option.test_dhcp"

func TestAccIpDhcpClientOptionTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/dhcp-client/option", "routeros_ip_dhcp_client_option"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpDhcpClientOptionConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpDhcpClientOptionAddress),
							resource.TestCheckResourceAttr(testIpDhcpClientOptionAddress, "name", "my-dhcp-option"),
							resource.TestCheckResourceAttr(testIpDhcpClientOptionAddress, "code", "60"),

						),
					},
				},
			})

		})
	}
}

func testAccIpDhcpClientOptionConfig() string {
	return providerConfig + `

resource "routeros_ip_dhcp_client_option" "test_dhcp" {
	name = "my-dhcp-option"
	code = 60
  }

`
}
