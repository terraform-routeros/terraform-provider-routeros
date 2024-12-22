package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpHotspotWalledGardenIp = "routeros_ip_hotspot_walled_garden_ip.test"

func TestAccIpHotspotWalledGardenIpTest_basic(t *testing.T) {
	// t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/hotspot/walled-garden/ip", "routeros_ip_hotspot_walled_garden_ip"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpHotspotWalledGardenIpConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpHotspotWalledGardenIp),
							resource.TestCheckResourceAttr(testIpHotspotWalledGardenIp, "action", "reject"),
							resource.TestCheckResourceAttr(testIpHotspotWalledGardenIp, "dst_address", "!0.0.0.0"),
							resource.TestCheckResourceAttr(testIpHotspotWalledGardenIp, "dst_address_list", "dlist"),
							resource.TestCheckResourceAttr(testIpHotspotWalledGardenIp, "dst_port", "0-65535"),
							resource.TestCheckResourceAttr(testIpHotspotWalledGardenIp, "protocol", "tcp"),
							resource.TestCheckResourceAttr(testIpHotspotWalledGardenIp, "src_address", "0.0.0.0"),
							resource.TestCheckResourceAttr(testIpHotspotWalledGardenIp, "src_address_list", "slist"),
						),
					},
				},
			})

		})
	}
}

func testAccIpHotspotWalledGardenIpConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_ip_hotspot_walled_garden_ip" "test" {
  action = "reject"
  dst_address = "!0.0.0.0"
  dst_address_list = "dlist"
  dst_port = "0-65535"
  protocol = "tcp"
  src_address = "0.0.0.0"
  src_address_list = "slist"
}
`, providerConfig)
}
