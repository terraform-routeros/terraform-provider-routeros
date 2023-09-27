package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testDatasourceIpRoutes = "data.routeros_ip_routes.routes"

func TestAccDatasourceIpRoutesTest_basic(t *testing.T) {
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
						Config: testAccDatasourceIpRoutesConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testDatasourceIpRoutes),
						),
					},
				},
			})

		})
	}
}

func testAccDatasourceIpRoutesConfig() string {
	return providerConfig + `

data "routeros_ip_routes" "routes" {}
`
}
