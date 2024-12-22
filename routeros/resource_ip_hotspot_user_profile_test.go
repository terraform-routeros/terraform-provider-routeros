package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpHotspotUserProfile = "routeros_ip_hotspot_user_profile.test"

func TestAccIpHotspotUserProfileTest_basic(t *testing.T) {
	// t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/hotspot/user/profile", "routeros_ip_hotspot_user_profile"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpHotspotUserProfileConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpHotspotUserProfile),
							resource.TestCheckResourceAttr(testIpHotspotUserProfile, "advertise", "true"),
						),
					},
				},
			})

		})
	}
}

func testAccIpHotspotUserProfileConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_ip_hotspot_user_profile" "test" {
  add_mac_cookie     = true
  address_list       = "list-1"
  idle_timeout       = "none"
  keepalive_timeout  = "2m"
  mac_cookie_timeout = "3d"
  name               = "new-profile"
  shared_users       = 3
  status_autorefresh = "2m"
  transparent_proxy  = true
  advertise          = true
}
`, providerConfig)
}
