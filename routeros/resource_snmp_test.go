package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testSNMPMinVersion = "7.8"
const testSNMPAddress = "routeros_snmp.test"

func TestAccSNMPTest_basic(t *testing.T) {
	if !testCheckMinVersion(t, testSNMPMinVersion) {
		t.Logf("Test skipped, the minimum required version is %v", testSNMPMinVersion)
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
				Steps: []resource.TestStep{
					{
						Config: testAccSNMPConfig(0),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckSNMPExists(testSNMPAddress),
							resource.TestCheckResourceAttr(testSNMPAddress, "enabled", "true"),
						),
					},
					{
						Config: testAccSNMPConfig(1),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckSNMPExists(testSNMPAddress),
							resource.TestCheckResourceAttr(testSNMPAddress, "enabled", "false"),
						),
					},
				},
			})
		})
	}
}

func testAccCheckSNMPExists(name string) resource.TestCheckFunc {
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

func testAccSNMPConfig(n int) string {
	var conf = []string{
		`
resource "routeros_snmp_community" "test" {
	name = "private"
}

resource "routeros_snmp" "test" {
	contact          = "John D."
	enabled          = true
	engine_id_suffix = "8a3c"
	location         = "Backyard"
	trap_community   = "private"
	trap_generators  = "start-trap"
	trap_version     = 3
	depends_on = [routeros_snmp_community.test]
}`,
		`
resource "routeros_snmp" "test" {
	contact          = ""
	enabled          = false
	engine_id_suffix = ""
	location         = ""
	trap_community   = "public"
	trap_generators  = "temp-exception"
	trap_version     = 1
}`,
	}
	return providerConfig + conf[n]
}
