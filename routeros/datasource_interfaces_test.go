package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
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
							testResourcePrimaryInstanceId(testDatasourceInterfaces),
						),
					},
				},
			})

		})
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
