package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
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
							testResourcePrimaryInstanceId(testIpRouteAddress),
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

func testAccIpRouteConfig() string {
	return providerConfig + `

resource "routeros_ip_route" "test_route" {
	distance      = 1
	dst_address   = "10.0.0.0/24"
	gateway		  = "192.168.103.1"
	check_gateway = "bfd-multihop"
}

`
}
