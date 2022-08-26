package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testDatasourceIpRoutes = "data.routeros_ip_routes.routes"

func TestAccDatasourceIpRoutesTest_basic(t *testing.T) {
	for _, name := range testNames {
		testSetTransportEnv(t, name)
		t.Run(name, func(t *testing.T) {

			resource.Test(t, resource.TestCase{
				PreCheck:  func() { testAccPreCheck(t) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: testAccDatasourceIpRoutesConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckDatasourceIpRoutesExists(testDatasourceIpRoutes),
						),
					},
				},
			})

		})
	}
}

func testAccCheckDatasourceIpRoutesExists(name string) resource.TestCheckFunc {
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

func testAccDatasourceIpRoutesConfig() string {
	return `

provider "routeros" {
	insecure = true
}

data "routeros_ip_routes" "routes" {}
`
}
