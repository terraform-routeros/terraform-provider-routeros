package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testDatasourceSystemResource = "data.routeros_system_resource.data"

func TestAccDatasourceSystemResourceTest_basic(t *testing.T) {
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
						Config: testAccDatasourceSystemResourceConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testDatasourceSystemResource),
						),
					},
				},
			})

		})
	}
}

func testAccDatasourceSystemResourceConfig() string {
	return providerConfig + `

data "routeros_system_resource" "data" {}
`
}
