package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testDatasourceIpAddressesAddress = "data.routeros_ip_addresses.addresses"

func TestAccDatasourceIpAddressesTest_basic(t *testing.T) {
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
						Config: testAccDatasourceIpAddressesConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testDatasourceIpAddressesAddress),
						),
					},
				},
			})

		})
	}
}

func testAccDatasourceIpAddressesConfig() string {
	return `

provider "routeros" {
	insecure = true
}

data "routeros_ip_addresses" "addresses" {}
`
}
