package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpv6Pool = "routeros_ipv6_pool.test"

func TestAccIpv6PoolTest_basic(t *testing.T) {
	// t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ipv6/pool", "routeros_ipv6_pool"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpv6PoolConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpv6Pool),
							resource.TestCheckResourceAttr(testIpv6Pool, "name", "test-pool"),
							resource.TestCheckResourceAttr(testIpv6Pool, "prefix", "2001:db8:12::/48"),
							resource.TestCheckResourceAttr(testIpv6Pool, "prefix_length", "64"),
						),
					},
				},
			})

		})
	}
}

func testAccIpv6PoolConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_ipv6_pool" "test" {
  name          = "test-pool"
  prefix        = "2001:db8:12::/48"
  prefix_length = 64
}
`, providerConfig)
}
