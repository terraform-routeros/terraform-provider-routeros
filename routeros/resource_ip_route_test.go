package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testIpRouteAddress = "routeros_ip_route.test_route"

func TestAccIpRouteTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/route", "routeros_ip_route"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpRouteConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckIpRouteExists(testIpRouteAddress),
							resource.TestCheckResourceAttr(testIpRouteAddress, "distance", "1"),
							resource.TestCheckResourceAttr(testIpRouteAddress, "dst_address", "10.0.0.0/24"),
							resource.TestCheckResourceAttr(testIpRouteAddress, "gateway", "192.168.103.1"),
						),
					},
				},
			})

		})
	}
}

func testAccCheckIpRouteExists(name string) resource.TestCheckFunc {
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

func testAccIpRouteConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_ip_route" "test_route" {
	distance      = 1
	dst_address   = "10.0.0.0/24"
	gateway		  = "192.168.103.1"
	check_gateway = "bfd-multihop"
}

`
}
