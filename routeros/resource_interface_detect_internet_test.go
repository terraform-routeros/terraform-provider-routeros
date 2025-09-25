package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfaceDetectInternet = "routeros_interface_detect_internet.test"

func TestAccInterfaceDetectInternetTest_basic(t *testing.T) {
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
						Config: testAccInterfaceDetectInternetConfig("all"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceDetectInternet),
							resource.TestCheckResourceAttr(testInterfaceDetectInternet, "internet_interface_list", "all"),
							resource.TestCheckResourceAttr(testInterfaceDetectInternet, "wan_interface_list", "all"),
							resource.TestCheckResourceAttr(testInterfaceDetectInternet, "lan_interface_list", "all"),
							resource.TestCheckResourceAttr(testInterfaceDetectInternet, "detect_interface_list", "all"),
						),
					},
					{
						Config: testAccInterfaceDetectInternetConfig("none"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceDetectInternet),
							resource.TestCheckResourceAttr(testInterfaceDetectInternet, "internet_interface_list", "none"),
							resource.TestCheckResourceAttr(testInterfaceDetectInternet, "wan_interface_list", "none"),
							resource.TestCheckResourceAttr(testInterfaceDetectInternet, "lan_interface_list", "none"),
							resource.TestCheckResourceAttr(testInterfaceDetectInternet, "detect_interface_list", "none"),
						),
					},
				},
			})

		})
	}
}

func testAccInterfaceDetectInternetConfig(param string) string {
	return fmt.Sprintf(`%v

resource "routeros_interface_detect_internet" "test" {
  internet_interface_list = "%v"
  wan_interface_list      = "%v"
  lan_interface_list      = "%v"
  detect_interface_list   = "%v"
}
`, providerConfig, param, param, param, param)
}
