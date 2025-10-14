package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingBfdMinVersion = "7.20"
const testRoutingBfdConfiguration = "routeros_routing_bfd_configuration.test"

func TestAccRoutingBfdConfigurationTest_basic(t *testing.T) {
	if !testCheckMinVersion(t, testRoutingBfdMinVersion) {
		t.Logf("Test skipped, the minimum required version is %v", testRoutingBfdMinVersion)
		return
	}

	t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/routing/bfd/configuration", "routeros_routing_bfd_configuration"),
				Steps: []resource.TestStep{
					{
						Config: testAccRoutingBfdConfigurationConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingBfdConfiguration),
							resource.TestCheckResourceAttr(testRoutingBfdConfiguration, "vrf", "main"),
							resource.TestCheckResourceAttr(testRoutingBfdConfiguration, "forbid_bfd", "true"),
						),
					},
					{
						Config: testAccRoutingBfdConfigurationConfig2(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingBfdConfiguration),
							resource.TestCheckResourceAttr(testRoutingBfdConfiguration, "vrf", "main"),
							resource.TestCheckResourceAttr(testRoutingBfdConfiguration, "interfaces.#", "0"),
						),
					},
				},
			})

		})
	}
}

func testAccRoutingBfdConfigurationConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_routing_bfd_configuration" "test" {
  interfaces = ["lo", "ether2"]
  vrf        = "main"
  forbid_bfd = true
}
`, providerConfig)
}

func testAccRoutingBfdConfigurationConfig2() string {
	return fmt.Sprintf(`%v

resource "routeros_routing_bfd_configuration" "test" {
  vrf        = "main"
}
`, providerConfig)
}
