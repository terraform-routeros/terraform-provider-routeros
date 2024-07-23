package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testOpenVPNServerMinVersion = "7.8"
const testOpenVPNServer = "routeros_ovpn_server.server"
const testInterfaceOpenVPNServer = "routeros_interface_ovpn_server.user1"

func TestAccOpenVPNServerTest_basic(t *testing.T) {
	if !testCheckMinVersion(t, testOpenVPNServerMinVersion) {
		t.Logf("Test skipped, the minimum required version is %v", testOpenVPNServerMinVersion)
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
				Steps: []resource.TestStep{
					{
						Config: testAccOpenVPNServerConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testOpenVPNServer),
							testResourcePrimaryInstanceId(testInterfaceOpenVPNServer),
							resource.TestCheckResourceAttr(testOpenVPNServer, "id", "interface.ovpn-server.server"),
							resource.TestCheckResourceAttr(testInterfaceOpenVPNServer, "name", "ovpn-in1"),
							resource.TestCheckResourceAttr(testInterfaceOpenVPNServer, "user", "user1"),
						),
					},
				},
			})

		})
	}
}

// Complex test for OpenVPN server resources.
func testAccOpenVPNServerConfig() string {
	return providerConfig + `
	resource "routeros_system_certificate" "ovpn_ca" {
		name        = "OpenVPN-Root-CA"
		common_name = "OpenVPN Root CA"
		key_size    = "prime256v1"
		key_usage   = ["key-cert-sign", "crl-sign"]
		trusted     = true
		sign {
		}
	  }
	  
	  resource "routeros_system_certificate" "ovpn_server_crt" {
		name        = "OpenVPN-Server-Certificate"
		common_name = "Mikrotik OpenVPN"
		key_size    = "prime256v1"
		key_usage   = ["digital-signature", "key-encipherment", "tls-server"]
		sign {
		  ca = routeros_system_certificate.ovpn_ca.name
		}
	  }
	  
	  resource "routeros_ovpn_server" "server" {
		enabled          = false
		certificate      = routeros_system_certificate.ovpn_server_crt.name
		auth             = ["sha256", "sha512"]
		redirect_gateway = ["def1", "ipv6"]
		tls_version      = "only-1.2"
	  }
	  
	  # The resource should be created only after the OpenVPN server is enabled!
      resource "routeros_interface_ovpn_server" "user1" {
        name       = "ovpn-in1"
        user       = "user1"
        depends_on = [routeros_ovpn_server.server]
      }
`
}
