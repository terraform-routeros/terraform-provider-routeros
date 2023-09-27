package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testDatasourceIpv6AddressesAddress = "data.routeros_ipv6_addresses.addresses"

func TestAccDatasourceIpv6AddressesTest_basic(t *testing.T) {
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
						Config: testAccDatasourceIpv6AddressesConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testDatasourceIpv6AddressesAddress),
						),
					},
				},
			})
		})
	}
}

func testAccDatasourceIpv6AddressesConfig() string {
	return providerConfig + `

data "routeros_ipv6_addresses" "addresses" {}
`
}
