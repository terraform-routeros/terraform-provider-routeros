package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testDatasourceInterfaceBridgeFilter = "data.routeros_interface_bridge_filter.rules"

func TestAccDatasourceInterfaceBridgeFilterTest_basic(t *testing.T) {
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
						Config: testAccDatasourceInterfaceBridgeFilterConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testDatasourceInterfaceBridgeFilter),
						),
					},
				},
			})

		})
	}
}

func testAccDatasourceInterfaceBridgeFilterConfig() string {
	return providerConfig + `

data "routeros_interface_bridge_filter" "rules" {}
`
}
