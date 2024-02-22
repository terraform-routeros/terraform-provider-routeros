package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIPv6NeighborDiscoveryAddress = "routeros_ipv6_neighbor_discovery.test"

func TestAccIPv6FNeighborDiscoveryTest_full(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ipv6/nd", "routeros_ipv6_neighbor_discovery"),
				Steps: []resource.TestStep{
					{
						Config: testAccFullIPv6NeighborDiscoveryConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIPv6NeighborDiscoveryAddress),
							resource.TestCheckResourceAttr(testIPv6NeighborDiscoveryAddress, "interface", "ether1"),
							resource.TestCheckResourceAttr(testIPv6NeighborDiscoveryAddress, "hop_limit", "33"),
							resource.TestCheckResourceAttr(testIPv6NeighborDiscoveryAddress, "advertise_dns", "false"),
							resource.TestCheckResourceAttr(testIPv6NeighborDiscoveryAddress, "advertise_mac_address", "true"),
							resource.TestCheckResourceAttr(testIPv6NeighborDiscoveryAddress, "disabled", "false"),
							resource.TestCheckResourceAttr(testIPv6NeighborDiscoveryAddress, "managed_address_configuration", "true"),
							resource.TestCheckResourceAttr(testIPv6NeighborDiscoveryAddress, "mtu", "9000"),
							resource.TestCheckResourceAttr(testIPv6NeighborDiscoveryAddress, "other_configuration", "true"),
							resource.TestCheckResourceAttr(testIPv6NeighborDiscoveryAddress, "ra_delay", "3s"),
							resource.TestCheckResourceAttr(testIPv6NeighborDiscoveryAddress, "ra_interval", "3m20s-10m"),
							resource.TestCheckResourceAttr(testIPv6NeighborDiscoveryAddress, "ra_lifetime", "30m"),
							resource.TestCheckResourceAttr(testIPv6NeighborDiscoveryAddress, "ra_preference", "high"),
							resource.TestCheckResourceAttr(testIPv6NeighborDiscoveryAddress, "reachable_time", "10m"),
							resource.TestCheckResourceAttr(testIPv6NeighborDiscoveryAddress, "retransmit_interval", "12m"),
						),
					},
				},
			})
		})
	}
}

func TestAccIPv6FNeighborDiscoveryTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ipv6/nd", "routeros_ipv6_neighbor_discovery"),
				Steps: []resource.TestStep{
					{
						Config: testAccSimpleIPv6NeighborDiscoveryConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIPv6NeighborDiscoveryAddress),
							resource.TestCheckResourceAttr(testIPv6NeighborDiscoveryAddress, "interface", "ether1"),
						),
					},
				},
			})
		})
	}
}

func testAccFullIPv6NeighborDiscoveryConfig() string {
	return providerConfig + `

resource "routeros_ipv6_neighbor_discovery" "test" {
 		interface 						= "ether1"
		hop_limit 						= 33
        advertise_dns 					= false
        advertise_mac_address 		    = true
        disabled 						= false		
        managed_address_configuration	= true
        mtu 							= 9000
        other_configuration				= true
        pref64_prefixes					= []
        ra_delay						= "3s"
        ra_interval						= "3m20s-10m"
        ra_lifetime						= "30m"
        ra_preference					= "high"
        reachable_time					= "10m"
        retransmit_interval				= "12m"
  }
`
}

func testAccSimpleIPv6NeighborDiscoveryConfig() string {
	return providerConfig + `

resource "routeros_ipv6_neighbor_discovery" "test" {
 		interface 						= "ether1"
  }
`
}
