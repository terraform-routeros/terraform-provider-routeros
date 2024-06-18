package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testDatasourceIpFirewall = "data.routeros_ip_firewall.fw"

func TestAccDatasourceIpFirewallTest_basic(t *testing.T) {
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
						Config: testAccDatasourceIpFirewallConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testDatasourceIpFirewall),
						),
					},
				},
			})

		})
	}
}

func testAccDatasourceIpFirewallConfig() string {
	return providerConfig + `
data "routeros_ip_firewall" "fw" {
  address_list {}
  mangle {}
  nat {}
  rules {}
}
`
}
