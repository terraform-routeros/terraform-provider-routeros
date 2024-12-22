package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfaceWirelessAccessList = "routeros_interface_wireless_access_list.test"

func TestAccInterfaceWirelessAccessListTest_basic(t *testing.T) {
	if !testCheckMaxVersion(t, testCapsManChannelMaxVersion) {
		t.Logf("Test skipped, the maximum supported version is %v", testCapsManChannelMaxVersion)
		return
	}

	// t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: testCheckResourceDestroy("/interface/wireless/access-list",
					"routeros_interface_wireless_access_list"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceWirelessAccessListConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceWirelessAccessList),
							resource.TestCheckResourceAttr(testInterfaceWirelessAccessList, "signal_range", "-100..100"),
							resource.TestCheckResourceAttr(testInterfaceWirelessAccessList, "time", "3h3m-5h,mon,tue,wed,thu,fri"),
							resource.TestCheckResourceAttr(testInterfaceWirelessAccessList, "mac_address", "00:AA:BB:CC:DD:EE"),
						),
					},
				},
			})

		})
	}
}

func testAccInterfaceWirelessAccessListConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_interface_wireless_access_list" "test" {
  signal_range = "-100..100"
  time         = "3h3m-5h,mon,tue,wed,thu,fri"
  mac_address  = "00:AA:BB:CC:DD:EE"
}
`, providerConfig)
}
