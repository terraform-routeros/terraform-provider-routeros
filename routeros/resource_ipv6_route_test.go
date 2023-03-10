package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testIPv6RouteAddress = "routeros_ipv6_route.test_route"

func TestAccIPv6RouteTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ipv6/route", "routeros_ipv6_route"),
				Steps: []resource.TestStep{
					{
						Config: testAccIPv6RouteConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckIPv6RouteExists(testIpRouteAddress),
							resource.TestCheckResourceAttr(testIPv6RouteAddress, "distance", "1"),
							resource.TestCheckResourceAttr(testIPv6RouteAddress, "dst_address", "::/0"),
							resource.TestCheckResourceAttr(testIPv6RouteAddress, "gateway", "2001:DB8:1000::1"),
						),
					},
				},
			})

		})
	}
}

func testAccCheckIPv6RouteExists(name string) resource.TestCheckFunc {
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

func testAccIPv6RouteConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_ip_route" "test_route" {
	distance    = 1
	dst_address = "::/0"
	gateway		= "2001:DB8:1000::1"
  }

`
}
