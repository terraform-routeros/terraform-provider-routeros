package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfaceIPAddressAddress = "routeros_ip_address.test_ip_address"

func TestAccInterfaceIPAddressTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/address", "routeros_ip_address"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceIPAddressConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceIPAddressAddress),
							resource.TestCheckResourceAttr(testInterfaceIPAddressAddress, "address", "172.16.255.254/32"),
						),
					},
				},
			})

		})
	}
}

func testAccInterfaceIPAddressConfig() string {
	return providerConfig + `

resource "routeros_ip_address" "test_ip_address" {
	interface	= "bridge"
	address		= "172.16.255.254/32"
  }
`
}

func TestIPAddressUpdateOmitsUnchangedVRF(t *testing.T) {
	item := MikrotikItem{"address": "172.16.255.254/32", "interface": "vlan10", "comment": "updated", KeyVrf: "main"}
	filterIPAddressUpdateItem(item, false)
	if _, ok := item[KeyVrf]; ok {
		t.Fatal("unchanged vrf must not be included in an IP address PATCH")
	}
}

func TestIPAddressUpdatePreservesExplicitVRFChange(t *testing.T) {
	item := MikrotikItem{KeyVrf: "blue"}
	filterIPAddressUpdateItem(item, true)
	if got := item[KeyVrf]; got != "blue" {
		t.Fatalf("explicit vrf change serialized as %q, want %q", got, "blue")
	}
}
