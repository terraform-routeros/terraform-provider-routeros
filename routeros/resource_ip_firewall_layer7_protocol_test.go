package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpFirewallLayer7Protocol = "routeros_ip_firewall_layer7_protocol.test"

func TestAccIpFirewallLayer7ProtocolTest_basic(t *testing.T) {
	t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/firewall/layer7-protocol", "routeros_ip_firewall_layer7_protocol"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpFirewallLayer7ProtocolConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpFirewallLayer7Protocol),
							resource.TestCheckResourceAttr(testIpFirewallLayer7Protocol, "name", "rdp"),
							resource.TestCheckResourceAttr(testIpFirewallLayer7Protocol, "regexp", "rdpdr.*cliprdr.*rdpsnd"),
						),
					},
				},
			})

		})
	}
}

func testAccIpFirewallLayer7ProtocolConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_ip_firewall_layer7_protocol" "test" {
  name   = "rdp"
  regexp = "rdpdr.*cliprdr.*rdpsnd"
}
`, providerConfig)
}
