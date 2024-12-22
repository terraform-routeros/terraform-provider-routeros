package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpHotspot = "routeros_ip_hotspot.test"

func TestAccIpHotspotTest_basic(t *testing.T) {
	// t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/hotspot", "routeros_ip_hotspot"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpHotspotConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpHotspot),
							resource.TestCheckResourceAttr(testIpHotspot, "name", "server-1"),
							resource.TestCheckResourceAttr(testIpHotspot, "interface", "ether3"),
						),
					},
				},
			})

		})
	}
}

func testAccIpHotspotConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_ip_hotspot" "test" {
  name      = "server-1"
  interface = "ether3"
}
`, providerConfig)
}
