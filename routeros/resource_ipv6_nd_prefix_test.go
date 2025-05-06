package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpv6NdPrefix = "routeros_ipv6_nd_prefix.test"

func TestAccIpv6NdPrefixTest_basic(t *testing.T) {
	t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ipv6/nd/prefix", "routeros_ipv6_nd_prefix"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpv6NdPrefixConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpv6NdPrefix),
							resource.TestCheckResourceAttr(testIpv6NdPrefix, "interface", "ether1"),
							resource.TestCheckResourceAttr(testIpv6NdPrefix, "prefix", "fd55::/64"),
							resource.TestCheckResourceAttr(testIpv6NdPrefix, "preferred_lifetime", "1w"),
						),
					},
				},
			})

		})
	}
}

func testAccIpv6NdPrefixConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_ipv6_nd_prefix" "test" {
  prefix             = "fd55::/64"
  interface          = "ether1"
  preferred_lifetime = "6d24h"
}
`, providerConfig)
}
