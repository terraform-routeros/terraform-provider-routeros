package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpv6DhcpServerOptionSets = "routeros_ipv6_dhcp_server_option_sets.test"

func TestAccIpv6DhcpServerOptionSetsTest_basic(t *testing.T) {
	// t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ipv6/dhcp-server/option/sets", "routeros_ipv6_dhcp_server_option_sets"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpv6DhcpServerOptionSetsConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpv6DhcpServerOptionSets),
							resource.TestCheckResourceAttr(testIpv6DhcpServerOptionSets, "name", "test-set"),
							resource.TestCheckResourceAttr(testIpv6DhcpServerOptionSets, "options.#", "1"),
							resource.TestCheckResourceAttr(testIpv6DhcpServerOptionSets, "options.0", "domain-search-o24"),
						),
					},
				},
			})

		})
	}
}

func testAccIpv6DhcpServerOptionSetsConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_ipv6_dhcp_server_option" "domain-search" {
  name  = "domain-search-o24"
  code  = 24
  value = "0x07'example'0x05'local'0x00"
}

resource "routeros_ipv6_dhcp_server_option_sets" "test" {
  name = "test-set"
  options = [routeros_ipv6_dhcp_server_option.domain-search.name]
}
`, providerConfig)
}
