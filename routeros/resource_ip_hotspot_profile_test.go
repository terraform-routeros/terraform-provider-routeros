package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpHotspotProfile = "routeros_ip_hotspot_profile.test"

func TestAccIpHotspotProfileTest_basic(t *testing.T) {
	// t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/hotspot/profile", "routeros_ip_hotspot_profile"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpHotspotProfileConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpHotspotProfile),
							resource.TestCheckResourceAttr(testIpHotspotProfile, "name", "hsprof-1"),
							resource.TestCheckResourceAttr(testIpHotspotProfile, "login_by.0", "https"),
							resource.TestCheckResourceAttr(testIpHotspotProfile, "login_by.1", "mac"),
							resource.TestCheckResourceAttr(testIpHotspotProfile, "login_by.2", "trial"),
							resource.TestCheckResourceAttr(testIpHotspotProfile, "use_radius", "true"),
						),
					},
				},
			})

		})
	}
}

func testAccIpHotspotProfileConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_ip_hotspot_profile" "test" {
  name       = "hsprof-1"
  login_by   = ["mac", "https", "trial"]
  use_radius = true
}
`, providerConfig)
}
