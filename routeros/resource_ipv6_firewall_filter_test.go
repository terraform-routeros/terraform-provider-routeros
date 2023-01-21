package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testIPv6FirewallFilterAddress = "routeros_ipv6_firewall_filter.rule"

func TestAccIPv6FirewallFilterTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ipv6/firewall/filter", "routeros_ipv6_firewall_filter"),
				Steps: []resource.TestStep{
					{
						Config: testAccIPv6FirewallFilterConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckIPv6FirewallFilterExists(testIPv6FirewallFilterAddress),
							resource.TestCheckResourceAttr(testIPv6FirewallFilterAddress, "action", "drop"),
						),
					},
				},
			})

		})
	}
}

func testAccCheckIPv6FirewallFilterExists(name string) resource.TestCheckFunc {
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

func testAccIPv6FirewallFilterConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_ipv6_firewall_filter" "rule" {
	// add action=drop chain=forward comment="" hop-limit= protocol=icmpv6
	action 		= "drop"
	chain   	= "forward"
	comment 	= "Check drop hop-limit=1 + ipv6 multicast"
	src_address = "ff00::/8"
	hop_limit	= "equal:1"
	protocol 	= "icmpv6"
  }

`
}
