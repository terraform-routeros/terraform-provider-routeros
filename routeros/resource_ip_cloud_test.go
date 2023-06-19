package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testIpCloudAddress = "routeros_ip_cloud.test"

func TestAccIpCloudTest_basic(t *testing.T) {
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
						Config: testAccIpCloudConfig(0),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckIpCloudExists(testIpCloudAddress),
							resource.TestCheckResourceAttr(testIpCloudAddress, "ddns_update_interval", "none"),
						),
					},
					{
						Config: testAccIpCloudConfig(1),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckIpCloudExists(testIpCloudAddress),
							resource.TestCheckResourceAttr(testIpCloudAddress, "ddns_update_interval", "15m"),
						),
					},
					{
						Config: testAccIpCloudConfig(2),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckIpCloudExists(testIpCloudAddress),
							resource.TestCheckResourceAttr(testIpCloudAddress, "ddns_update_interval", "10m"),
						),
					},
					{
						Config: testAccIpCloudConfig(3),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckIpCloudExists(testIpCloudAddress),
							resource.TestCheckResourceAttr(testIpCloudAddress, "ddns_update_interval", "none"),
						),
					},
				},
			})
		})
	}
}

func testAccCheckIpCloudExists(name string) resource.TestCheckFunc {
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

func testAccIpCloudConfig(n int) string {
	var conf = []string{
		`
resource "routeros_ip_cloud" "test" {
  ddns_enabled         = true
  ddns_update_interval = "none"
}`,
		`
resource "routeros_ip_cloud" "test" {
  ddns_enabled         = true
  ddns_update_interval = "15m"
}`,
		`
resource "routeros_ip_cloud" "test" {
  ddns_enabled         = true
  ddns_update_interval = "600"
}`,
		`
resource "routeros_ip_cloud" "test" {
  ddns_enabled = true
}`,
	}
	return providerConfig + conf[n]
}
