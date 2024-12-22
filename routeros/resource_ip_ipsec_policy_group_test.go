package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testIpIpsecPolicyGroup = "routeros_ip_ipsec_policy_group.test"

func TestAccIpIpsecPolicyGroupTest_basic(t *testing.T) {
	// t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/ipsec/policy/group", "routeros_ip_ipsec_policy_group"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpIpsecPolicyGroupConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpIpsecPolicyGroup),
							resource.TestCheckResourceAttr(testIpIpsecPolicyGroup, "name", "test-group"),
						),
					},
					{
						Config:        testAccIpIpsecPolicyGroupConfig(),
						ResourceName:  testIpIpsecPolicyGroup,
						ImportStateId: `name=test-group`,
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

func testAccIpIpsecPolicyGroupConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_ip_ipsec_policy_group" "test" {
  name    = "test-group"
}
`, providerConfig)
}
