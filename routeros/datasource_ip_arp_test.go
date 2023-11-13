package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testDatasourceIpArp = "data.routeros_ip_arp.data"

func TestAccDatasourceIpArpTest_basic(t *testing.T) {
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
						Config: testAccDatasourceIpArpConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testDatasourceIpArp),
						),
					},
				},
			})

		})
	}
}

func testAccDatasourceIpArpConfig() string {
	return providerConfig + `

data "routeros_ip_arp" "data" {}
`
}
