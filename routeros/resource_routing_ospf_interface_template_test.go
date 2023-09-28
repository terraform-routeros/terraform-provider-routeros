package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingOspfInterfaceTemplate = "routeros_routing_ospf_interface_template.test_routing_ospf_interface_template"

func TestAccRoutingOspfInterfaceTemplateTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/routing/ospf/interface-template", "routeros_routing_ospf_interface_template"),
				Steps: []resource.TestStep{
					{
						Config: testAccCheckRoutingOspfInterfaceTemplateConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingOspfInterfaceTemplate),
							resource.TestCheckResourceAttr(testRoutingOspfInterfaceTemplate, "area", "test_routing_ospf_area"),
						),
					},
				},
			})

		})
	}
}

func testAccCheckRoutingOspfInterfaceTemplateConfig() string {
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
	disabled	= false
	instance 	= routeros_routing_ospf_instance.test_routing_ospf_instance.name
}

resource "routeros_routing_ospf_interface_template" "test_routing_ospf_interface_template" {
	area       = routeros_routing_ospf_area.test_routing_ospf_area.name
	interfaces = ["ether3"]
	passive    = true
}

`
}
