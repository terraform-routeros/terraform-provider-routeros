package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingIgmpProxyInterface = "routeros_routing_igmp_proxy_interface.test"

func TestAccRoutingIgmpProxyInterfaceTest_basic(t *testing.T) {
	t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/routing/igmp-proxy/interface", "routeros_routing_igmp_proxy_interface"),
				Steps: []resource.TestStep{
					{
						Config: testAccRoutingIgmpProxyInterfaceConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingIgmpProxyInterface),
							resource.TestCheckResourceAttr(testRoutingIgmpProxyInterface, "alternative_subnets.0", "0.0.0.1/32"),
							resource.TestCheckResourceAttr(testRoutingIgmpProxyInterface, "alternative_subnets.1", "0.0.0.2/32"),
							resource.TestCheckResourceAttr(testRoutingIgmpProxyInterface, "disabled", "true"),
							resource.TestCheckResourceAttr(testRoutingIgmpProxyInterface, "interface", "all"),
							resource.TestCheckResourceAttr(testRoutingIgmpProxyInterface, "threshold", "5"),
						),
					},
				},
			})

		})
	}
}

func testAccRoutingIgmpProxyInterfaceConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_routing_igmp_proxy_interface" "test" {
  alternative_subnets = ["0.0.0.1/32", "0.0.0.2/32"]
  disabled            = true
  interface           = "all"
  threshold           = 5
}
`, providerConfig)
}
