package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingOspfArea = "routeros_routing_ospf_area.test_routing_ospf_area"

func TestAccRoutingOspfInstanceArea_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/routing/ospf/area", "routeros_routing_ospf_area"),
				Steps: []resource.TestStep{
					{
						Config: testAccCheckRoutingOspfAreaConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingOspfArea),
							resource.TestCheckResourceAttr(testRoutingOspfArea, "name", "test_routing_ospf_area"),
						),
					},
				},
			})

		})
	}
}

func testAccCheckRoutingOspfAreaConfig() string {
	return providerConfig + `
resource "routeros_routing_ospf_instance" "test_routing_ospf_instance" {
	name   		= "test_routing_ospf_instance"
	disabled	= false
  }

resource "routeros_routing_ospf_area" "test_routing_ospf_area" {
	name   		= "test_routing_ospf_area"
	disabled	= true
	instance 	= routeros_routing_ospf_instance.test_routing_ospf_instance.name
}

`
}
