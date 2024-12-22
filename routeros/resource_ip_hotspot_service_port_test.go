package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpHotspotServicePort = "routeros_ip_hotspot_service_port.test"

func TestAccIpHotspotServicePortTest_basic(t *testing.T) {
	// t.Parallel()
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
						Config: testAccIpHotspotServicePortConfig("true"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpHotspotServicePort),
							resource.TestCheckResourceAttr(testIpHotspotServicePort, "disabled", "true"),
							resource.TestCheckResourceAttr(testIpHotspotServicePort, "name", "ftp"),
						),
					},
					{
						Config: testAccIpHotspotServicePortConfig("false"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpHotspotServicePort),
							resource.TestCheckResourceAttr(testIpHotspotServicePort, "disabled", "false"),
							resource.TestCheckResourceAttr(testIpHotspotServicePort, "name", "ftp"),
						),
					},
				},
			})

		})
	}
}

func testAccIpHotspotServicePortConfig(param string) string {
	return fmt.Sprintf(`%v

resource "routeros_ip_hotspot_service_port" "test" {
  name    = "ftp"
  disabled = %v
}
`, providerConfig, param)
}
