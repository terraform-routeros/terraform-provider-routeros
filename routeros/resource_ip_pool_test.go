package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpPoolAddress = "routeros_ip_pool.test_pool"

func TestAccIpPoolTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/pool", "routeros_ip_pool"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpPoolConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpPoolAddress),
							resource.TestCheckResourceAttr(testIpPoolAddress, "name", "test_pool"),
						),
					},
				},
			})

		})
	}
}

func testAccIpPoolConfig() string {
	return providerConfig + `

resource "routeros_ip_pool" "test_pool" {
	name   = "test_pool"
	ranges = ["10.0.0.100-10.0.0.200"]
  }

`
}
