package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testCapsManChannelAddress = "routeros_capsman_channel.test_channel"

func TestAccCapsManChannelTest_basic(t *testing.T) {
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
							testAccCheckCapsManChannelExists(testCapsManChannelAddress),
							resource.TestCheckResourceAttr(testCapsManChannelAddress, "name", "test_channel"),
						),
					},
				},
			})
		})
	}
}

func testAccCheckCapsManChannelExists(name string) resource.TestCheckFunc {
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

func testAccCapsManChannelConfig() string {
	return providerConfig + `

resource "routeros_capsman_channel" "test_channel" {
	name                  = "test_channel"
	comment               = "test_channel"
	band                  = "2ghz-b/g/n"
	control_channel_width = "10mhz"
	extension_channel     = "eCee"
	frequency             = 2412
	reselect_interval     = "1h"
	save_selected         = true
	secondary_frequency   = "disabled"
	skip_dfs_channels     = true
	tx_power              = 20
}
`
}
