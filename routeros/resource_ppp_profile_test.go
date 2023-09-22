package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testPPPProfile = "routeros_ppp_profile.test"

func TestAccPPPProfileTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ppp/profile", "routeros_ppp_profile"),
				Steps: []resource.TestStep{
					{
						Config: testAccPPPProfileConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testPPPProfile),
							resource.TestCheckResourceAttr(testPPPProfile, "name", "profile1"),
						),
					},
				},
			})

		})
	}
}

func testAccPPPProfileConfig() string {
	return providerConfig + `
	resource "routeros_ppp_profile" "test" {
		name        = "profile1"
		rate_limit  = "10M/200k"
		use_upnp    = "no"
		dns_server  = ["8.8.8.8", "1.1.1.1"]
	}
`
}
