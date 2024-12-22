package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpIpsecKey = "routeros_ip_ipsec_key.test"

func TestAccIpIpsecKeyTest_basic(t *testing.T) {
	// t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/ipsec/key", "routeros_ip_ipsec_key"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpIpsecKeyConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpIpsecKey),
							resource.TestCheckResourceAttr(testIpIpsecKey, "name", "test-key"),
							resource.TestCheckResourceAttr(testIpIpsecKey, "key_size", "2048"),
						),
					},
				},
			})
		})
	}
}

func testAccIpIpsecKeyConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_ip_ipsec_key" "test" {
  name     = "test-key"
  key_size = 2048
}
`, providerConfig)
}
