package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testDatasourceIpAddressesAddress = "data.routeros_ip_addresses.addresses"

func TestAccDatasourceIpAddressesTest_basic(t *testing.T) {
	for _, name := range testNames {
		testSetTransportEnv(t, name)
		t.Run(name, func(t *testing.T) {

			resource.Test(t, resource.TestCase{
				PreCheck:  func() { testAccPreCheck(t) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: testAccDatasourceIpAddressesConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckDatasourceIpAddressesExists(testDatasourceIpAddressesAddress),
						),
					},
				},
			})

		})
	}
}

func testAccCheckDatasourceIpAddressesExists(name string) resource.TestCheckFunc {
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

func testAccDatasourceIpAddressesConfig() string {
	return `

provider "routeros" {
	insecure = true
}

data "routeros_ip_addresses" "addresses" {}
`
}
