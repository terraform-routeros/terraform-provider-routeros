package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testInterfaceIPv6AddressAddress = "routeros_ipv6_address.test_v6_address"

func TestAccInterfaceIPv6AddressTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ipv6/address", "routeros_ipv6_address"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceIPv6AddressConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckInterfaceIPv6AddressExists(testInterfaceIPv6AddressAddress),
							resource.TestCheckResourceAttrWith(testInterfaceIPv6AddressAddress, "address",
								func(value string) error {
									if value[:7] != "fc00:3:" {
										return fmt.Errorf(`Attribute 'address' expected "fc00:3:", got "%s"`, value)
									}
									return nil
								}),
						),
					},
				},
			})

		})
	}
}

func testAccCheckInterfaceIPv6AddressExists(name string) resource.TestCheckFunc {
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

func testAccInterfaceIPv6AddressConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_ipv6_address" "test_v6_address" {
	interface	= "bridge"
	address		= "fc00:3::/64"
	advertise   = true
	eui_64		= true
	no_dad		= true
	// address		= "::1/64"
	// from_pool 	= "pool1"
  }

`
}
