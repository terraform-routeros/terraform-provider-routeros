package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testUPNPInterfaces = "routeros_ip_upnp_interfaces.test"

func TestAccUPNPInterfacesTest_basic(t *testing.T) {
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
						Config: testAccUPNPInterfacesConfig("ether1", "external", `forced_ip = "0.0.0.0"`),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testUPNPInterfaces),
							resource.TestCheckResourceAttr(testUPNPInterfaces, "interface", "ether1"),
							resource.TestCheckResourceAttr(testUPNPInterfaces, "type", "external"),
							resource.TestCheckResourceAttr(testUPNPInterfaces, "forced_ip", "0.0.0.0"),
						),
					},
					{
						Taint:  []string{testUPNPInterfaces},
						Config: testAccUPNPInterfacesConfig("ether1", "internal", ""),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(testUPNPInterfaces, "interface", "ether1"),
							resource.TestCheckResourceAttr(testUPNPInterfaces, "type", "internal"),
							resource.TestCheckNoResourceAttr(testUPNPInterfaces, "forced_ip"),
						),
					},
				},
			})
		})
	}
}

func testAccUPNPInterfacesConfig(s1, s2, s3 string) string {
	return fmt.Sprintf(`%v

resource "routeros_ip_upnp_interfaces" "test" {
	interface           = "%v"
	type                = "%v"
	%v
}
`, providerConfig, s1, s2, s3)
}
