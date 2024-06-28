package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testDatasourceIPv6Firewall = "data.routeros_ipv6_firewall.fw"

func TestAccDatasourceIPv6FirewallTest_basic(t *testing.T) {
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
						Config: testAccDatasourceIPv6FirewallConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testDatasourceIPv6Firewall),
						),
					},
				},
			})

		})
	}
}

func testAccDatasourceIPv6FirewallConfig() string {
	return providerConfig + `
data "routeros_ipv6_firewall" "fw" {
  rules {}
}
`
}
