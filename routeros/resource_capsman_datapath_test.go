package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testCapsManDatapathAddress = "routeros_capsman_datapath.test_datapath"

func TestAccCapsManDatapathTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/caps-man/datapath", "routeros_capsman_datapath"),
				Steps: []resource.TestStep{
					{
						Config: testAccCapsManDatapathConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckCapsManDatapathExists(testCapsManDatapathAddress),
							resource.TestCheckResourceAttr(testCapsManDatapathAddress, "name", "test_datapath"),
						),
					},
				},
			})
		})
	}
}

func testAccCheckCapsManDatapathExists(name string) resource.TestCheckFunc {
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

func testAccCapsManDatapathConfig() string {
	return providerConfig + `

resource "routeros_capsman_datapath" "test_datapath" {
	name                        = "test_datapath"
	comment                     = "test_datapath"
	arp                         = "local-proxy-arp"
	bridge                      = "bridge"
	bridge_cost                 = 100
	bridge_horizon              = 200
	client_to_client_forwarding = true
	interface_list              = "static"
	l2mtu                       = 1450
	local_forwarding            = true
	mtu                         = 1500
	vlan_id                     = 101
	vlan_mode                   = "no-tag"
//  openflow_switch             = "aaa"
}

`
}
