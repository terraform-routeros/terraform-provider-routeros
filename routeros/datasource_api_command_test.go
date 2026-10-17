package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testDatasourceApiCommand = "data.routeros_api_command.interfaces"
const testDatasourceApiCommandCount = "data.routeros_api_command.interface_count"

func TestAccDatasourceApiCommandTest_basic(t *testing.T) {
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
						Config: testAccDatasourceApiCommandConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testDatasourceApiCommand),
							resource.TestCheckResourceAttrSet(testDatasourceApiCommand, "result.0.name"),
							// 'count-only' answers on the trailing sentence instead of a
							// row, on both transports.
							resource.TestCheckResourceAttrSet(testDatasourceApiCommandCount, "result.0.ret"),
						),
					},
				},
			})

		})
	}
}

func testAccDatasourceApiCommandConfig() string {
	return providerConfig + `

data "routeros_api_command" "interfaces" {
  path = "/interface"
}

data "routeros_api_command" "interface_count" {
  path  = "/interface"
  words = ["count-only"]
}
`
}
