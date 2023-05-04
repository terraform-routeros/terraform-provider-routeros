package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testIpServiceAddress = "routeros_ip_service.telnet"

func TestAccIpServiceTest_basic(t *testing.T) {
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
						Config: testAccIpServiceConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckIpServiceExists(testIpServiceAddress),
							resource.TestCheckResourceAttr(testIpServiceAddress, "name", "telnet"),
						),
					},
				},
			})
		})
	}
}

func testAccCheckIpServiceExists(name string) resource.TestCheckFunc {
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

func testAccIpServiceConfig() string {
	return providerConfig + `

resource "routeros_ip_service" "telnet" {
	numbers  = "telnet"
	disabled = true
	port     = 23
}
`
}
