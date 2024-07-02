package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testCapsManChannelMaxVersion = "7.12.2"
const testCapsManChannelAddress = "routeros_capsman_channel.test_channel"

func TestAccCapsManChannelTest_basic(t *testing.T) {
	if !testCheckMaxVersion(t, testCapsManChannelMaxVersion) {
		t.Logf("Test skipped, the maximum supported version is %v", testCapsManChannelMaxVersion)
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
				CheckDestroy:      testCheckResourceDestroy("/caps-man/channel", "routeros_capsman_channel"),
				Steps: []resource.TestStep{
					{
						Config: testAccCapsManChannelConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testCapsManChannelAddress),
							resource.TestCheckResourceAttr(testCapsManChannelAddress, "name", "test_channel"),
						),
					},
				},
			})
		})
	}
}

func testAccCapsManChannelConfig() string {
	return providerConfig + `

resource "routeros_capsman_channel" "test_channel" {
	name                  = "test_channel"
	comment               = "test_channel"
	band                  = "2ghz-b/g/n"
	control_channel_width = "10mhz"
	extension_channel     = "eCee"
	frequency             = [2412]
	reselect_interval     = "1h"
	save_selected         = true
	secondary_frequency   = ["disabled"]
	skip_dfs_channels     = true
	tx_power              = 20
}
`
}
