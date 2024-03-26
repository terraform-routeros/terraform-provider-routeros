package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testPppAaaAddress = "routeros_ppp_aaa.test"

func TestAccPppAaaTest_basic(t *testing.T) {
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
						Config: testAccPppAaaConfig(0),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testPppAaaAddress),
							resource.TestCheckResourceAttr(testPppAaaAddress, "accounting", "true"),
							resource.TestCheckResourceAttr(testPppAaaAddress, "use_radius", "true"),
						),
					},
					{
						Config: testAccPppAaaConfig(1),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testPppAaaAddress),
							resource.TestCheckResourceAttr(testPppAaaAddress, "accounting", "false"),
							resource.TestCheckResourceAttr(testPppAaaAddress, "use_radius", "false"),
						),
					},
				},
			})

		})
	}
}

func testAccPppAaaConfig(n int) string {
	var conf = []string{
		`
resource "routeros_ppp_aaa" "test" {
  accounting = true
  use_radius = true
}`,
		`
resource "routeros_ppp_aaa" "test" {
  accounting = false
  use_radius = false
}`,
	}
	return providerConfig + conf[n]
}
