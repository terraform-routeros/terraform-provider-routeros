package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpv6DhcpServerOption = "routeros_ipv6_dhcp_server_option.test"

func TestAccIpv6DhcpServerOptionTest_basic(t *testing.T) {
	// t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ipv6/dhcp-server/option", "routeros_ipv6_dhcp_server_option"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpv6DhcpServerOptionConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpv6DhcpServerOption),
							resource.TestCheckResourceAttr(testIpv6DhcpServerOption, "name", "domain-search"),
							resource.TestCheckResourceAttr(testIpv6DhcpServerOption, "code", "24"),
							resource.TestCheckResourceAttr(testIpv6DhcpServerOption, "value", "0x07'example'0x05'local'0x00"),
						),
					},
				},
			})

		})
	}
}

func testAccIpv6DhcpServerOptionConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_ipv6_dhcp_server_option" "test" {
  name  = "domain-search"
  code  = 24
  value = "0x07'example'0x05'local'0x00"
}
`, providerConfig)
}
