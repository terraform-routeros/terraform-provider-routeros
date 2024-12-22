package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testIpIpsecIdentity = "routeros_ip_ipsec_identity.identity"

func TestAccIpIpsecIdentityTest_basic(t *testing.T) {
	// t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/ipsec/identity", "routeros_ip_ipsec_identity"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpIpsecIdentityConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpIpsecIdentity),
							resource.TestCheckResourceAttr(testIpIpsecIdentity, "auth_method", "eap"),
							resource.TestCheckResourceAttr(testIpIpsecIdentity, "certificate", ""),
							resource.TestCheckResourceAttr(testIpIpsecIdentity, "eap_methods", "eap-mschapv2"),
							resource.TestCheckResourceAttr(testIpIpsecIdentity, "generate_policy", "port-strict"),
							resource.TestCheckResourceAttr(testIpIpsecIdentity, "mode_config", "NordVPN-i"),
							resource.TestCheckResourceAttr(testIpIpsecIdentity, "peer", "NordVPN-i"),
							resource.TestCheckResourceAttr(testIpIpsecIdentity, "username", "support@mikrotik.com"),
							resource.TestCheckResourceAttr(testIpIpsecIdentity, "password", "secret"),
						),
					},
					{
						Config:        testAccIpIpsecIdentityConfig(),
						ResourceName:  testIpIpsecIdentity,
						ImportStateId: `peer=NordVPN-i`,
						ImportState:   true,
						ImportStateCheck: func(states []*terraform.InstanceState) error {
							if len(states) != 1 {
								return fmt.Errorf("more than 1 states received, only one expected")
							}
							return nil
						},
					},
				},
			})

		})
	}
}

func testAccIpIpsecIdentityConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_ip_ipsec_mode_config" "mode-for-identity" {
  name      = "NordVPN-i"
  responder = false
}

resource "routeros_ip_ipsec_peer" "peer-for-identity" {
  address       = "lv30.nordvpn.com"
  exchange_mode = "ike2"
  name          = "NordVPN-i"
}

resource "routeros_ip_ipsec_identity" "identity" {
  auth_method     = "eap"
  certificate     = ""
  eap_methods     = "eap-mschapv2"
  generate_policy = "port-strict"
  mode_config     = routeros_ip_ipsec_mode_config.mode-for-identity.name
  peer            = routeros_ip_ipsec_peer.peer-for-identity.name
  username        = "support@mikrotik.com"
  password        = "secret"
}
`, providerConfig)
}
