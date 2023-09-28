package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
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
							testResourcePrimaryInstanceId(testInterfaceIPv6AddressAddress),
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

func testAccInterfaceIPv6AddressConfig() string {
	return providerConfig + `

resource "routeros_ipv6_address" "test_v6_address" {
	interface	= "bridge"
	address		= "fc00:3::/64"
	advertise   = true
	eui_64		= true
	no_dad		= true
	// address		= "::1/64"
	// from_pool 	= "pool1"
}

resource "routeros_ip_route" "d2_rmark_isp22" {
	distance      = 2
	gateway       = "10.10.10.1"
	routing_table = "main"
	target_scope  = 12
}
`
}
