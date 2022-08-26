package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testInterfaceWireguardAddress = "routeros_wireguard.test_wg_interface"

func TestAccInterfaceWireguardTest_basic(t *testing.T) {
	for _, name := range testNames {
		testSetTransportEnv(t, name)
		t.Run(name, func(t *testing.T) {

			resource.Test(t, resource.TestCase{
				PreCheck:     func() { testAccPreCheck(t) },
				Providers:    testAccProviders,
				CheckDestroy: testCheckResourceDestroy("/interface/wireguard", "routeros_wireguard"),
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

resource "routeros_wireguard" "test_wg_interface" {
	name   		= "test_wg_interface"
	listen_port = "13231"
  }

`
}
