package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpArpAddress = "routeros_ip_arp.test"

func TestAccIpArpTest(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/dhcp-server", "routeros_ip_dhcp_server"),
				Steps: []resource.TestStep{
					{
						Config: ipArpConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpArpAddress),
							resource.TestCheckResourceAttr(testIpArpAddress, "interface", "bridge"),
						),
					},
				},
			})

		})
	}
}

func ipArpConfig() string {
	return providerConfig + `
resource "routeros_ip_arp" "test_arp" {
	name	    = "192.168.88.128"
	interface   = "bridge1"
	mac_address = "00:CA:FE:BA:BE:00"
    published   = true
}
`
}
