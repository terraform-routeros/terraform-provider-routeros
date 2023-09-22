package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingOspfInstance = "routeros_routing_ospf_instance.test_routing_ospf_instance"

func TestAccRoutingOspfInstanceTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/routing/ospf/instance", "routeros_routing_ospf_instance"),
				Steps: []resource.TestStep{
					{
						Config: testAccCheckRoutingOspfInstanceConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingOspfInstance),
							resource.TestCheckResourceAttr(testRoutingOspfInstance, "name", "test_routing_ospf_instance"),
						),
					},
				},
			})

		})
	}
}

func testAccCheckRoutingOspfInstanceConfig() string {
	return providerConfig + `
resource "routeros_routing_ospf_instance" "test_routing_ospf_instance" {
	name   		= "test_routing_ospf_instance"
	disabled	= false
  }
`
}
