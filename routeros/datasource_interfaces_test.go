package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testDatasourceInterfaces = "data.routeros_interfaces.interfaces"

func TestAccDatasourceInterfacesTest_basic(t *testing.T) {
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
						Config: testAccDatasourceInterfacesConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckDatasourceInterfacesExists(testDatasourceInterfaces),
						),
					},
				},
			})

		})
	}
}

func testAccCheckDatasourceInterfacesExists(name string) resource.TestCheckFunc {
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

func testAccDatasourceInterfacesConfig() string {
	return `

provider "routeros" {
	insecure = true
}

data "routeros_interfaces" "interfaces" {}
`
}
