package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testInterfaceIPAddressAddress = "routeros_ip_address.test_ip_address"

func TestAccInterfaceIPAddressTest_basic(t *testing.T) {
	for _, name := range testNames {
		testSetTransportEnv(t, name)
		t.Run(name, func(t *testing.T) {

			resource.Test(t, resource.TestCase{
				PreCheck:     func() { testAccPreCheck(t) },
				Providers:    testAccProviders,
				CheckDestroy: testCheckResourceDestroy("/ip/address", "routeros_ip_address"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceIPAddressConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckInterfaceIPAddressExists(testInterfaceIPAddressAddress),
							resource.TestCheckResourceAttr(testInterfaceIPAddressAddress, "address", "172.16.255.254/32"),
						),
					},
				},
			})

		})
	}
}

func testAccCheckInterfaceIPAddressExists(name string) resource.TestCheckFunc {
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

func testAccInterfaceIPAddressConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_ip_address" "test_ip_address" {
	interface	= "bridge"
	address		= "172.16.255.254/32"
  }

`
}
