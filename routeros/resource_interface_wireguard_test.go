package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testInterfaceWireguardAddress = "routeros_interface_wireguard.test_wg_interface"

func TestAccInterfaceWireguardTest_basic(t *testing.T) {
	t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/wireguard", "routeros_interface_wireguard"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceWireguardConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckInterfaceWireguardExists(testInterfaceWireguardAddress),
							resource.TestCheckResourceAttr(testInterfaceWireguardAddress, "listen_port", "13231"),
						),
					},
				},
			})

		})
	}
}

func testAccCheckInterfaceWireguardExists(name string) resource.TestCheckFunc {
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

func testAccInterfaceWireguardConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_interface_wireguard" "test_wg_interface" {
	name   		= "test_wg_interface"
	listen_port = "13231"
  }

`
}
