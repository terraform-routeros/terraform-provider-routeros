package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testIpIpsecPolicy = "routeros_ip_ipsec_policy.policy"

func TestAccIpIpsecPolicyTest_basic(t *testing.T) {
	// t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/ipsec/policy", "routeros_ip_ipsec_policy"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpIpsecPolicyConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpIpsecPolicy),
							resource.TestCheckResourceAttr(testIpIpsecPolicy, "dst_address", "0.0.0.0/0"),
							resource.TestCheckResourceAttr(testIpIpsecPolicy, "group", "test-group-p"),
							resource.TestCheckResourceAttr(testIpIpsecPolicy, "proposal", "default"),
							resource.TestCheckResourceAttr(testIpIpsecPolicy, "src_address", "0.0.0.0/0"),
							resource.TestCheckResourceAttr(testIpIpsecPolicy, "template", "true"),
						),
					},
					{
						Config:        testAccIpIpsecPolicyConfig(),
						ResourceName:  testIpIpsecPolicy,
						ImportStateId: `group=test-group-p`,
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

func testAccIpIpsecPolicyConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_ip_ipsec_policy_group" "group-for-policy" {
  name = "test-group-p"
}

resource "routeros_ip_ipsec_policy" "policy" {
  dst_address = "0.0.0.0/0"
  group       = routeros_ip_ipsec_policy_group.group-for-policy.name
  proposal    = "default"
  src_address = "0.0.0.0/0"
  template    = true
}
`, providerConfig)
}
