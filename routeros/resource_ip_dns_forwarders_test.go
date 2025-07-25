package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpDnsForwarders = "routeros_ip_dns_forwarders.test"

func TestAccIpDnsForwardersTest_basic(t *testing.T) {
	t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/dns/forwarders", "routeros_ip_dns_forwarders"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpDnsForwardersConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpDnsForwarders),
							resource.TestCheckResourceAttr(testIpDnsForwarders, "disabled", "true"),
							resource.TestCheckResourceAttr(testIpDnsForwarders, "dns_servers.0", "1.1.1.1"),
							resource.TestCheckResourceAttr(testIpDnsForwarders, "doh_servers.0", "2.2.2.2"),
							resource.TestCheckResourceAttr(testIpDnsForwarders, "name", "test"),
							resource.TestCheckResourceAttr(testIpDnsForwarders, "verify_doh_cert", "false"),
						),
					},
				},
			})

		})
	}
}

func testAccIpDnsForwardersConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_ip_dns_forwarders" "test" {
  disabled        = true
  dns_servers     = ["1.1.1.1"]
  doh_servers     = ["2.2.2.2"]
  name            = "test"
  verify_doh_cert = "false"
}
`, providerConfig)
}
