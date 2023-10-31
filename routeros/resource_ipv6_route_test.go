package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
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
							testResourcePrimaryInstanceId(testIPv6RouteAddress),
							resource.TestCheckResourceAttr(testIPv6RouteAddress, "distance", "1"),
							resource.TestCheckResourceAttr(testIPv6RouteAddress, "dst_address", "::/0"),
							resource.TestCheckResourceAttr(testIPv6RouteAddress, "gateway", "2001:db8:1000::1"),
						),
					},
				},
			})

		})
	}
}

func testAccIPv6RouteConfig() string {
	return providerConfig + `

resource "routeros_ipv6_route" "test_route" {
	distance    = 1
	dst_address = "::/0"
	gateway		= "2001:db8:1000::1"
  }

`
}
