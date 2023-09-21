package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
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
							testResourcePrimaryInstanceId(testIpCloudAddress),
							resource.TestCheckResourceAttr(testIpCloudAddress, "ddns_update_interval", "none"),
						),
					},
					{
						Config: testAccIpCloudConfig(1),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpCloudAddress),
							resource.TestCheckResourceAttr(testIpCloudAddress, "ddns_update_interval", "15m"),
						),
					},
					{
						Config: testAccIpCloudConfig(2),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpCloudAddress),
							resource.TestCheckResourceAttr(testIpCloudAddress, "ddns_update_interval", "10m"),
						),
					},
					{
						Config: testAccIpCloudConfig(3),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpCloudAddress),
							resource.TestCheckResourceAttr(testIpCloudAddress, "ddns_update_interval", "none"),
						),
					},
				},
			})
		})
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
