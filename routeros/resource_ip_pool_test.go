package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testIpPoolAddress = "routeros_ip_pool.test_pool"

func TestAccIpPoolTest_basic(t *testing.T) {
	for _, name := range testNames {
		testSetTransportEnv(t, name)
		t.Run(name, func(t *testing.T) {

			resource.Test(t, resource.TestCase{
				PreCheck:     func() { testAccPreCheck(t) },
				Providers:    testAccProviders,
				CheckDestroy: testCheckResourceDestroy("/ip/pool", "routeros_ip_pool"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpPoolConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckIpPoolExists(testIpPoolAddress),
							resource.TestCheckResourceAttr(testIpPoolAddress, "name", "test_pool"),
						),
					},
				},
			})

		})
	}
}

func testAccCheckIpPoolExists(name string) resource.TestCheckFunc {
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

func testAccIpPoolConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_ip_pool" "test_pool" {
	name   = "test_pool"
	ranges = ["10.0.0.100-10.0.0.200"]
  }

`
}
