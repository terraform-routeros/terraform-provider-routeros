package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingOspfAreaRange = "routeros_routing_ospf_area_range.test"

func TestAccRoutingOspfAreaRangeTest_basic(t *testing.T) {
	t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/routing/ospf/area/range", "routeros_routing_ospf_area_range"),
				Steps: []resource.TestStep{
					{
						Config: testAccRoutingOspfAreaRangeConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingOspfAreaRange),
							resource.TestCheckResourceAttr(testRoutingOspfAreaRange, "area", "ospf-area-oar"),
							resource.TestCheckResourceAttr(testRoutingOspfAreaRange, "advertise", "true"),
							resource.TestCheckResourceAttr(testRoutingOspfAreaRange, "prefix", "::/64"),
							resource.TestCheckResourceAttr(testRoutingOspfAreaRange, "disabled", "true"),
						),
					},
				},
			})

		})
	}
}

func testAccRoutingOspfAreaRangeConfig() string {
	return fmt.Sprintf(`%v
resource "routeros_routing_ospf_instance" "ospf-instance-oar" {
  name   		= "ospf-instance-oar"
  disabled	= false
}

resource "routeros_routing_ospf_area" "ospf-area-oar" {
  name   		= "ospf-area-oar"
  disabled	= true
  instance 	= routeros_routing_ospf_instance.ospf-instance-oar.name
}

resource "routeros_routing_ospf_area_range" "test" {
  area      = routeros_routing_ospf_area.ospf-area-oar.name
  advertise = true
  prefix    = "::/64"
  disabled  = true
}
`, providerConfig)
}
