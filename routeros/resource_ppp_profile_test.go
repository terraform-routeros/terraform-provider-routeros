package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testPPPProfile = "routeros_ppp_profile.test"

func TestAccPPPProfileTest_basic(t *testing.T) {
	t.Parallel()
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
							testAccCheckPPPProfileExists(testPPPProfile),
							resource.TestCheckResourceAttr(testPPPProfile, "name", "profile1"),
						),
					},
				},
			})

		})
	}
}

func testAccCheckPPPProfileExists(name string) resource.TestCheckFunc {
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
