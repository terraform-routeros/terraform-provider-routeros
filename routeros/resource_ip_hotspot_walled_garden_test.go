package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpHotspotWalledGarden = "routeros_ip_hotspot_walled_garden.test"

func TestAccIpHotspotWalledGardenTest_basic(t *testing.T) {
	// t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/hotspot/walled-garden", "routeros_ip_hotspot_walled_garden"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpHotspotWalledGardenConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpHotspotWalledGarden),
							resource.TestCheckResourceAttr(testIpHotspotWalledGarden, "action", "deny"),
							resource.TestCheckResourceAttr(testIpHotspotWalledGarden, "dst_host", "1.2.3.4"),
							resource.TestCheckResourceAttr(testIpHotspotWalledGarden, "dst_port", "!443"),
						),
					},
				},
			})

		})
	}
}

func testAccIpHotspotWalledGardenConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_ip_hotspot_walled_garden" "test" {
  action   = "deny"
  dst_host = "1.2.3.4"
  dst_port = "!443"
}
`, providerConfig)
}
