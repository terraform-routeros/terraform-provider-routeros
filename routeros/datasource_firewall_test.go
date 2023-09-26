package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testDatasourceFirewall = "data.routeros_firewall.fw"

func TestAccDatasourceFirewallTest_basic(t *testing.T) {
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
						Config: testAccDatasourceFirewallConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testDatasourceFirewall),
						),
					},
				},
			})

		})
	}
}

func testAccDatasourceFirewallConfig() string {
	return providerConfig + `
data "routeros_firewall" "fw" {
  address_list {}
  mangle {}
  nat {}
  rules {}
}
`
}
