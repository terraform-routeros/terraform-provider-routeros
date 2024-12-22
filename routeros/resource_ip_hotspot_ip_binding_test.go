package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpHotspotIpBinding = "routeros_ip_hotspot_ip_binding.test"

func TestAccIpHotspotIpBindingTest_basic(t *testing.T) {
	// t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/hotspot/ip-binding", "routeros_ip_hotspot_ip_binding"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpHotspotIpBindingConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpHotspotIpBinding),
							resource.TestCheckResourceAttr(testIpHotspotIpBinding, "address", "0.0.0.1"),
							resource.TestCheckResourceAttr(testIpHotspotIpBinding, "comment", "comment"),
							resource.TestCheckResourceAttr(testIpHotspotIpBinding, "mac_address", "00:00:00:00:01:10"),
							resource.TestCheckResourceAttr(testIpHotspotIpBinding, "to_address", "0.0.0.2"),
						),
					},
				},
			})

		})
	}
}

func testAccIpHotspotIpBindingConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_ip_hotspot_ip_binding" "test" {
   address     = "0.0.0.1"
   comment     = "comment"
   mac_address = "00:00:00:00:01:10"
   to_address  = "0.0.0.2" 
}
`, providerConfig)
}
