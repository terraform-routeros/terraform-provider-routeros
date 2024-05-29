package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIPFirewallRawAddress = "routeros_ip_firewall_raw.rule"

func TestAccIPFirewallRawTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/firewall/raw", "routeros_ip_firewall_raw"),
				Steps: []resource.TestStep{
					{
						Config: testAccIPFirewallRawConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIPFirewallRawAddress),
							resource.TestCheckResourceAttr(testIPFirewallRawAddress, "action", "accept"),
						),
					},
				},
			})

		})
	}
}

func testAccIPFirewallRawConfig() string {
	return providerConfig + `
resource "routeros_ip_firewall_raw" "rule" {
	action 		= "accept"
	chain   	= "prerouting"
	src_address = "10.0.0.1"
	dst_address = "10.0.1.1"
	dst_port 	= "443"
	protocol 	= "tcp"
}
`
}
