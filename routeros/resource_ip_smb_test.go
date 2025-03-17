package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpSMBTask = "routeros_ip_smb.test"

func TestAccIpSMBTest_basic(t *testing.T) {
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
						Config: testAccIpSMBConfig("auto", "MSHOME", "MikrotikSMB", "all"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpSMBTask),
							resource.TestCheckResourceAttr(testIpSMBTask, "enabled", "auto"),
							resource.TestCheckResourceAttr(testIpSMBTask, "domain", "MSHOME"),
							resource.TestCheckResourceAttr(testIpSMBTask, "comment", "MikrotikSMB"),
							resource.TestCheckResourceAttr(testIpSMBTask, "interfaces", "all"),
						),
					},
				},
			})

		})
	}
}

func testAccIpSMBConfig(enabled, domain, comment, _interface string) string {
	return fmt.Sprintf(`%v
resource "routeros_ip_smb" "test" {
  enabled    = "%v"
  domain     = "%v"
  comment    = "%v"
  interfaces = ["%v"]
}
`, providerConfig, enabled, domain, comment, _interface)
}
