package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testDatasourceIpv6AddressesAddress = "data.routeros_ipv6_addresses.addresses"

func TestAccDatasourceIpv6AddressesTest_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDatasourceIpv6AddressesConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDatasourceIpv6AddressesExists(testDatasourceIpv6AddressesAddress),
				),
			},
		},
	})
}

func testAccCheckDatasourceIpv6AddressesExists(name string) resource.TestCheckFunc {
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

func testAccDatasourceIpv6AddressesConfig() string {
	return `

provider "routeros" {
	insecure = true
}

data "routeros_ipv6_addresses" "addresses" {}
`
}
