package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testToolGraphingInterfaceTask = "routeros_tool_graphing_interface.test"

func TestAccToolGraphingInterfaceTest_basic(t *testing.T) {
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
						Config: testAccToolGraphingInterfaceConfig("all", "0.0.0.0/0", true),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testToolGraphingInterfaceTask),
							resource.TestCheckResourceAttr(testToolGraphingInterfaceTask, "interface", "all"),
							resource.TestCheckResourceAttr(testToolGraphingInterfaceTask, "allow_address", "0.0.0.0/0"),
							resource.TestCheckResourceAttr(testToolGraphingInterfaceTask, "store_on_disk", "true"),
						),
					},
				},
			})

		})
	}
}

const testToolGraphingQueueTask = "routeros_tool_graphing_queue.test"

func TestAccToolGraphingQueueTest_basic(t *testing.T) {
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
						Config: testAccToolGraphingQueueConfig("all", "0.0.0.0/0", true, true),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testToolGraphingQueueTask),
							resource.TestCheckResourceAttr(testToolGraphingQueueTask, "simple_queue", "all"),
							resource.TestCheckResourceAttr(testToolGraphingQueueTask, "allow_address", "0.0.0.0/0"),
							resource.TestCheckResourceAttr(testToolGraphingQueueTask, "allow_target", "true"),
							resource.TestCheckResourceAttr(testToolGraphingQueueTask, "store_on_disk", "true"),
						),
					},
				},
			})

		})
	}
}

const testToolGraphingResourceTask = "routeros_tool_graphing_resource.test"

func TestAccToolGraphingResourceTest_basic(t *testing.T) {
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
						Config: testAccToolGraphingResourceConfig("0.0.0.0/0", true),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testToolGraphingResourceTask),
							resource.TestCheckResourceAttr(testToolGraphingResourceTask, "allow_address", "0.0.0.0/0"),
							resource.TestCheckResourceAttr(testToolGraphingResourceTask, "store_on_disk", "true"),
						),
					},
				},
			})

		})
	}
}

func testAccToolGraphingInterfaceConfig(_interface string, allowAddress string, storeOnDisk bool) string {
	return fmt.Sprintf(`%v

resource "routeros_tool_graphing_interface" "test" {
  interface     = "%v"
  allow_address = "%v"
  store_on_disk = %v
}
`, providerConfig, _interface, allowAddress, storeOnDisk)
}

func testAccToolGraphingQueueConfig(simpleQueue string, allowAddress string, allowTarget bool, storeOnDisk bool) string {
	return fmt.Sprintf(`%v

resource "routeros_tool_graphing_queue" "test" {
  simple_queue  = "%v"
  allow_address = "%v"
  allow_target  = %v
  store_on_disk = %v
}
`, providerConfig, simpleQueue, allowAddress, allowTarget, storeOnDisk)
}

func testAccToolGraphingResourceConfig(allowAddress string, storeOnDisk bool) string {
	return fmt.Sprintf(`%v

resource "routeros_tool_graphing_resource" "test" {
  allow_address = "%v"
  store_on_disk = %v
}
`, providerConfig, allowAddress, storeOnDisk)
}
