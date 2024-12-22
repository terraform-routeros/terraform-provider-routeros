package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpHotspotUser = "routeros_ip_hotspot_user.test"

func TestAccIpHotspotUserTest_basic(t *testing.T) {
	// t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/hotspot/user", "routeros_ip_hotspot_user"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpHotspotUserConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpHotspotUser),
							resource.TestCheckResourceAttr(testIpHotspotUser, "name", "user-1"),
							resource.TestCheckResourceAttr(testIpHotspotUser, "profile", "default"),
						),
					},
				},
			})

		})
	}
}

func testAccIpHotspotUserConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_ip_hotspot_user" "test" {
  name = "user-1"
}
`, providerConfig)
}
