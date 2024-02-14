package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIPv6DhcpClientOption = "routeros_ipv6_dhcp_client_option.test_dhcp"

func TestAccIPv6DhcpClientOptionTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ipv6/dhcp-client/option", "routeros_ipv6_dhcp_client_option"),
				Steps: []resource.TestStep{
					{
						Config: testAccIPv6DhcpClientOptionConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIPv6DhcpClientOption),
							resource.TestCheckResourceAttr(testIPv6DhcpClientOption, "name", "my-dhcp-option"),
							resource.TestCheckResourceAttr(testIPv6DhcpClientOption, "code", "60"),
						),
					},
				},
			})

		})
	}
}

func testAccIPv6DhcpClientOptionConfig() string {
	return providerConfig + `

resource "routeros_ipv6_dhcp_client_option" "test_dhcp" {
	name = "my-dhcp-option"
	code = 60
  }

`
}
