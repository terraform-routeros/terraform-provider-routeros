package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpCloudAdvancedAddress = "routeros_ip_cloud_advanced.test"

func TestAccIpCloudAdvancedTest_basic(t *testing.T) {
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
						Config: testAccIpCloudAdvancedConfig(0),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpCloudAdvancedAddress),
							resource.TestCheckResourceAttr(testIpCloudAdvancedAddress, "use_local_address", "true"),
						),
					},
					{
						Config: testAccIpCloudAdvancedConfig(1),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpCloudAdvancedAddress),
							resource.TestCheckResourceAttr(testIpCloudAdvancedAddress, "use_local_address", "false"),
						),
					},
				},
			})
		})
	}
}

func testAccIpCloudAdvancedConfig(n int) string {
	var conf = []string{
		`
resource "routeros_ip_cloud_advanced" "test" {
  use_local_address = true
}`,
		`
resource "routeros_ip_cloud_advanced" "test" {
  use_local_address = false
}`,
	}
	return providerConfig + conf[n]
}
