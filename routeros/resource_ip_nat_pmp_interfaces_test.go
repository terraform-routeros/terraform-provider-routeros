package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testNatPmpInterfaces = "routeros_ip_nat_pmp_interfaces.test"

func TestAccNatPmpInterfacesTest_basic(t *testing.T) {
	if !testCheckMinVersion(t, testNatPmpMinVersion) {
		t.Logf("Test skipped, the minimum required version is %v", testNatPmpMinVersion)
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
						Config: testAccNatPmpInterfacesConfig("ether1", "external", `forced_ip = "0.0.0.0"`),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testNatPmpInterfaces),
							resource.TestCheckResourceAttr(testNatPmpInterfaces, "interface", "ether1"),
							resource.TestCheckResourceAttr(testNatPmpInterfaces, "type", "external"),
							resource.TestCheckResourceAttr(testNatPmpInterfaces, "forced_ip", "0.0.0.0"),
						),
					},
					{
						Taint:  []string{testNatPmpInterfaces},
						Config: testAccNatPmpInterfacesConfig("ether1", "internal", ""),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(testNatPmpInterfaces, "interface", "ether1"),
							resource.TestCheckResourceAttr(testNatPmpInterfaces, "type", "internal"),
							resource.TestCheckNoResourceAttr(testNatPmpInterfaces, "forced_ip"),
						),
					},
				},
			})
		})
	}
}

func testAccNatPmpInterfacesConfig(s1, s2, s3 string) string {
	return fmt.Sprintf(`%v

resource "routeros_ip_nat_pmp_interfaces" "test" {
	interface           = "%v"
	type                = "%v"
	%v
}
`, providerConfig, s1, s2, s3)
}
