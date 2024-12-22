package routeros

import (
	"testing"
)

const testDatasourceSystemRouterboard = "data.routeros_system_routerboard.data"

func TestAccDatasourceSystemRouterboardTest_basic(t *testing.T) {
	t.Log("The test is skipped, the resource is only available on real hardware.")
	/*
		// t.Parallel()
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
							Config: testAccDatasourceSystemRouterboardConfig(),
							Check: resource.ComposeTestCheckFunc(
								testResourcePrimaryInstanceId(testDatasourceSystemRouterboard),
							),
						},
					},
				})

			})
		}
	*/
}

/*
func testAccDatasourceSystemRouterboardConfig() string {
	return providerConfig + `

data "routeros_system_routerboard" "data" {}
`
}
*/
