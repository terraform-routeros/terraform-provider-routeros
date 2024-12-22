package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpIpsecPeer = "routeros_ip_ipsec_peer.test"

func TestAccIpIpsecPeerTest_basic(t *testing.T) {
	// t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/ipsec/peer", "routeros_ip_ipsec_peer"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpIpsecPeerConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpIpsecPeer),
							resource.TestCheckResourceAttr(testIpIpsecPeer, "address", "lv20.nordvpn.com"),
							resource.TestCheckResourceAttr(testIpIpsecPeer, "exchange_mode", "ike2"),
							resource.TestCheckResourceAttr(testIpIpsecPeer, "name", "NordVPN"),
							resource.TestCheckResourceAttr(testIpIpsecPeer, "profile", "default"),
						),
					},
				},
			})

		})
	}
}

func testAccIpIpsecPeerConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_ip_ipsec_peer" "test" {
  address       = "lv20.nordvpn.com"
  exchange_mode = "ike2"
  name          = "NordVPN"
}

`, providerConfig)
}
