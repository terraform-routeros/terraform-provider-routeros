package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testSystemKeymatProvider = "routeros_system_keymat_provider.test"

func TestAccSystemKeymatProviderTest_basic(t *testing.T) {
	t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/system/keymat-provider", "routeros_system_keymat_provider"),
				Steps: []resource.TestStep{
					{
						Config: testAccSystemKeymatProviderConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testSystemKeymatProvider),
							resource.TestCheckResourceAttr(testSystemKeymatProvider, "name", "keymat-test"),
							resource.TestCheckResourceAttr(testSystemKeymatProvider, "type", "hkdf"),
							resource.TestCheckResourceAttr(testSystemKeymatProvider, "hkdf_hash", "sha256"),
						),
					},
				},
			})
		})
	}
}

func testAccSystemKeymatProviderConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_system_keymat_provider" "test" {
  name        = "keymat-test"
  type        = "hkdf"
  hkdf_hash   = "sha256"
  hkdf_secret = "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f"
}
`, providerConfig)
}