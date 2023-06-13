package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
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
							testAccCheckRoutingOspfInstanceExists(testRoutingOspfInstance),
							resource.TestCheckResourceAttr(testRoutingOspfInstance, "name", "test_routing_ospf_instance"),
						),
					},
				},
			})

		})
	}
}

func testAccCheckRoutingOspfInstanceExists(name string) resource.TestCheckFunc {
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

func testAccCheckRoutingOspfInstanceConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_routing_ospf_instance" "test_routing_ospf_instance" {
	name   		= "test_routing_ospf_instance"
	disabled	= false
  }

`
}
