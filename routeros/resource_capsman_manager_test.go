package routeros

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testCapsManAaaAddress = "routeros_capsman_aaa.test_3a"
const testCapsManManagerAddress = "routeros_capsman_manager.test_manager"
const testCapsManManagerInterfaceAddress = "routeros_capsman_manager_interface.test_manager_interface"

func TestAccCapsManManagerTest_basic(t *testing.T) {
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
						Config: testAccCapsManManagerUnitConfig(name, "routeros_capsman_manager"),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckCapsManResourceExists(testCapsManManagerAddress),
							resource.TestCheckResourceAttr(testCapsManManagerAddress, "id", "caps-man.manager"),
						),
					},
					{
						Config: testAccCapsManManagerUnitConfig(name, "routeros_capsman_aaa"),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckCapsManResourceExists(testCapsManAaaAddress),
							resource.TestCheckResourceAttr(testCapsManAaaAddress, "id", "caps-man.aaa"),
						),
					},
					{
						Config: testAccCapsManManagerUnitConfig(name, "routeros_capsman_manager_interface"),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckCapsManResourceExists(testCapsManManagerInterfaceAddress),
							resource.TestCheckResourceAttr(testCapsManManagerInterfaceAddress, "interface", "ether1"),
						),
					},
				},
			})
		})
	}
}

func testAccCheckCapsManResourceExists(address string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[address]
		if !ok {
			return fmt.Errorf("not found: %s", address)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no id is set")
		}

		return nil
	}
}

func testAccCapsManManagerUnitConfig(testName, resourceName string) string {
	conf := `
	provider "routeros" {
		insecure = true
	}
`

	switch resourceName {
	// AAA
	case "routeros_capsman_aaa":
		if strings.Contains(testName, "API") {
			// API
			conf += `
			resource "routeros_capsman_aaa" "test_3a" {
				called_format  = "ssid"
				mac_mode       = "as-username-and-password"
			}
			`
		} else {
			// REST
			conf += `
			resource "routeros_capsman_aaa" "test_3a" {
				called_format  = "mac:ssid"
				mac_mode       = "as-username"
			}
			`
		}
	// Manager
	case "routeros_capsman_manager":
		if strings.Contains(testName, "API") {
			// API
			conf += `
			resource "routeros_capsman_manager" "test_manager" {
				enabled        = true
				upgrade_policy = "require-same-version"
			}		
			`
		} else {
			// REST
			conf += `
			resource "routeros_capsman_manager" "test_manager" {
				enabled        = false
				upgrade_policy = "none"
			}		
			`
		}
	// CAPsMAN interfaces
	case "routeros_capsman_manager_interface":
		conf += `
		resource "routeros_capsman_manager_interface" "test_manager_interface" {
			interface = "ether1"
			forbid    = true
		}
		`
	}

	return conf
}
