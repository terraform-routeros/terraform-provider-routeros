package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfaceOpenVPNClient = "routeros_interface_ovpn_client.ovpn-in1"

func TestAccOpenVPNClientTest_basic(t *testing.T) {
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
						Config: testAccOpenVPNClientConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceOpenVPNClient),
							resource.TestCheckResourceAttr(testInterfaceOpenVPNClient, "name", "ovpn-in1"),
							resource.TestCheckResourceAttr(testInterfaceOpenVPNClient, "user", "user1"),
							resource.TestCheckResourceAttr(testInterfaceOpenVPNClient, "connect_to", "192.168.1.1"),
						),
					},
				},
			})

		})
	}
}

// Complex test for OpenVPN client resources.
func testAccOpenVPNClientConfig() string {
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
	  
	  resource "routeros_system_certificate" "ovpn_client_crt" {
		name        = "OpenVPN-Client-Certificate"
		common_name = "Mikrotik OpenVPN Client"
		key_size    = "prime256v1"
		key_usage   = ["digital-signature", "key-encipherment", "tls-client"]
		sign {
		  ca = routeros_system_certificate.ovpn_ca.name
		}
	  }
	  
      resource "routeros_interface_ovpn_client" "ovpn-in1" {
        name        = "ovpn-in1"
        user        = "user1"
		connect_to  = "192.168.1.1"
		certificate = routeros_system_certificate.ovpn_client_crt.name
      }
`
}
