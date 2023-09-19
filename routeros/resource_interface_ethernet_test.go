package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testInterfaceEthernetAddress = "routeros_interface_ethernet.test"

func TestAccInterfaceEthernetTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceEthernetConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckInterfaceEthernetExists(testInterfaceEthernetAddress),
							resource.TestCheckResourceAttr(testInterfaceEthernetAddress, "name", "terraform"),
							resource.TestCheckResourceAttr(testInterfaceEthernetAddress, "mtu", "9000"),
							resource.TestCheckResourceAttr(testInterfaceEthernetAddress, "advertise", "10000M-full"),
							resource.TestCheckResourceAttr(testInterfaceEthernetAddress, "arp", "disabled"),
							resource.TestCheckResourceAttr(testInterfaceEthernetAddress, "auto_negotiation", "false"),
							resource.TestCheckResourceAttr(testInterfaceEthernetAddress, "speed", "100Mbps"),
						),
					},
				},
			})
		})
	}
}

func testAccCheckInterfaceEthernetExists(name string) resource.TestCheckFunc {
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

func testAccInterfaceEthernetConfig() string {
	return providerConfig + `

resource "routeros_interface_ethernet" "test" {
  factory_name     = "ether2"
  name             = "terraform"
  mtu              = "9000"
  advertise        = "10000M-full"
  arp              = "disabled"
  auto_negotiation = false
  speed            = "100Mbps"
}
`
}
