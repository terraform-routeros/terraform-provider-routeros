package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpServiceAddress = "routeros_ip_service.telnet"

func TestAccIpServiceTest_basic(t *testing.T) {
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
							testResourcePrimaryInstanceId(testIpServiceAddress),
							resource.TestCheckResourceAttr(testIpServiceAddress, "name", "telnet"),
						),
					},
				},
			})
		})
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
