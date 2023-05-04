package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testDatasourceFirewall = "data.routeros_firewall.fw"

func TestAccDatasourceFirewallTest_basic(t *testing.T) {
	t.Parallel()
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
							testAccCheckDatasourceFirewallExists(testDatasourceFirewall),
						),
					},
				},
			})

		})
	}
}

func testAccCheckDatasourceFirewallExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("not found: %s", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no id is set")
		}

		return nil
	}
}

func testAccDatasourceFirewallConfig() string {
	return `

provider "routeros" {
	insecure = true
}

data "routeros_firewall" "fw" {
  address_list {}
  mangle {}
  nat {}
  rules {}
}
`
}
