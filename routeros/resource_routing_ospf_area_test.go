package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
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
							testAccCheckRoutingOspfAreaExists(testRoutingOspfArea),
							resource.TestCheckResourceAttr(testRoutingOspfArea, "name", "test_routing_ospf_area"),
						),
					},
				},
			})

		})
	}
}

func testAccCheckRoutingOspfAreaExists(name string) resource.TestCheckFunc {
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

func testAccCheckRoutingOspfAreaConfig() string {
	return `

provider "routeros" {
	insecure = true
}

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
