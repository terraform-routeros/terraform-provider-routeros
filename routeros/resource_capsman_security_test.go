package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testCapsManSecurityMaxVersion = "7.12.2"
const testCapsManSecurityAddress = "routeros_capsman_security.test_security"

func TestAccCapsManSecurityTest_basic(t *testing.T) {
	if !testCheckMaxVersion(t, testCapsManSecurityMaxVersion) {
		t.Logf("Test skipped, the maximum supported version is %v", testCapsManSecurityMaxVersion)
		return
	}

	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/caps-man/security", "routeros_capsman_security"),
				Steps: []resource.TestStep{
					{
						Config: testAccCapsManSecurityConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testCapsManSecurityAddress),
							resource.TestCheckResourceAttr(testCapsManSecurityAddress, "name", "test_security"),
						),
					},
				},
			})

		})
	}
}

func testAccCapsManSecurityConfig() string {
	return providerConfig + `

resource "routeros_capsman_security" "test_security" {
	name                  = "test_security"
	comment               = "test_security"
	authentication_types  = ["wpa-psk", "wpa-eap", "wpa2-psk"] // Unordered items!
	disable_pmkid         = true
	eap_methods           = "eap-tls,passthrough"
	eap_radius_accounting = true
	encryption            = ["tkip", "aes-ccm"]  // Unordered items!
	group_encryption      = "aes-ccm"
	group_key_update      = "1h"
	passphrase            = "0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF0123456789ABCDE" // Max length check
	tls_certificate       = "none"
	tls_mode              = "verify-certificate"
}
`
}
