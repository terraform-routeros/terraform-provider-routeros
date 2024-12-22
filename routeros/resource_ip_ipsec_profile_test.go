package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testIpIpsecProfile = "routeros_ip_ipsec_profile.test"

func TestAccIpIpsecProfileTest_basic(t *testing.T) {
	// t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/ipsec/profile", "routeros_ip_ipsec_profile"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpIpsecProfileConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpIpsecProfile),
							resource.TestCheckResourceAttr(testIpIpsecProfile, "name", "test-profile"),
							resource.TestCheckResourceAttr(testIpIpsecProfile, "hash_algorithm", "sha256"),
							resource.TestCheckResourceAttr(testIpIpsecProfile, "enc_algorithm.#", "2"),
							resource.TestCheckResourceAttr(testIpIpsecProfile, "enc_algorithm.0", "aes-192"),
							resource.TestCheckResourceAttr(testIpIpsecProfile, "enc_algorithm.1", "aes-256"),
							resource.TestCheckResourceAttr(testIpIpsecProfile, "nat_traversal", "false"),
						),
					},
					{
						Config:        testAccIpIpsecProfileConfig(),
						ResourceName:  testIpIpsecProfile,
						ImportStateId: `name=test-profile`,
						ImportState:   true,
						ImportStateCheck: func(states []*terraform.InstanceState) error {
							if len(states) != 1 {
								return fmt.Errorf("more than 1 states received, only one expected")
							}
							return nil
						},
					},
				},
			})

		})
	}
}

func testAccIpIpsecProfileConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_ip_ipsec_profile" "test" {
  name           = "test-profile"
  hash_algorithm = "sha256"
  enc_algorithm  = ["aes-192", "aes-256"]
  dh_group       = ["ecp384", "ecp521"]
  nat_traversal  = false
}
`, providerConfig)
}
