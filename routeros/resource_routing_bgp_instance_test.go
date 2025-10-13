package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingBgpInstance = "routeros_routing_bgp_instance.test"

func TestAccRoutingBgpInstanceTest_basic(t *testing.T) {
	t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/routing/bgp/instance", "routeros_routing_bgp_instance"),
				Steps: []resource.TestStep{
					{
						Config: testAccRoutingBgpInstanceConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingBgpInstance),
							resource.TestCheckResourceAttr(testRoutingBgpInstance, "as", "65000"),
							resource.TestCheckResourceAttr(testRoutingBgpInstance, "name", "bgp-instance-1"),
						),
					},
				},
			})

		})
	}
}

func testAccRoutingBgpInstanceConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_routing_bgp_instance" "test" {
  as   = "65000"
  name = "bgp-instance-1"
}
`, providerConfig)
}
