package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingSettings = "routeros_routing_settings.test"

func TestAccRoutingSettingsTest_basic(t *testing.T) {
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
						Config: testAccRoutingSettingsConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingSettings),
							resource.TestCheckResourceAttr(testRoutingSettings, "check_gateway_ping_count", "3"),
						),
					},
				},
			})
		})
	}
}

func testAccRoutingSettingsConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_routing_settings" "test" {
  check_gateway_ping_count = 3
}
`, providerConfig)
}