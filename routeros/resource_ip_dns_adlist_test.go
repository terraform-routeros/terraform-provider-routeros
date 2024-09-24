package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testDnsAdlistMinVersion = "7.15"
const testResourceDnsAdlist = "routeros_ip_dns_adlist.test"

func TestAccResourceDnsAdlistTest_basic(t *testing.T) {
	if !testCheckMinVersion(t, testDnsAdlistMinVersion) {
		t.Logf("Test skipped, the minimum required version is %v", testDnsAdlistMinVersion)
		return
	}

	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/dns/adlist", "routeros_ip_dns_adlist"),
				Steps: []resource.TestStep{
					{
						Config: testAccResourceDnsAdlistConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testResourceDnsAdlist),
							resource.TestCheckResourceAttr(testResourceDnsAdlist, "url", "https://raw.githubusercontent.com/StevenBlack/hosts/master/hosts"),
							resource.TestCheckResourceAttr(testResourceDnsAdlist, "ssl_verify", "false"),
						),
					},
				},
			})

		})
	}
}

func testAccResourceDnsAdlistConfig() string {
	return providerConfig + `
resource "routeros_ip_dns_adlist" "test" {
	url        = "https://raw.githubusercontent.com/StevenBlack/hosts/master/hosts"
	ssl_verify = false
}`
}
