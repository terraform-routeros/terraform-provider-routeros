package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testSystemCertificatesAddress = "routeros_system_certificate.root_ca"

func TestAccSystemCertificatesTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/certificate", "routeros_system_certificate"),
				Steps: []resource.TestStep{
					{
						Config: testAccSystemCertificatesConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testSystemCertificatesAddress),
							resource.TestCheckResourceAttr(testSystemCertificatesAddress, "name", "Test-Root-CA"),
						),
					},
				},
			})

		})
	}
}

func testAccSystemCertificatesConfig() string {
	return providerConfig + `

resource "routeros_system_certificate" "root_ca" {
	name        = "Test-Root-CA"
	common_name = "RootCA"
	key_usage   = ["key-cert-sign", "crl-sign"]
	trusted     = true
	sign {
	}
}	  

resource "routeros_system_certificate" "server_crt" {
	name        = "Server-Certificate"
	common_name = "server.crt"
	// KUs: igitalSignature, keyEncipherment or keyAgreement
	key_usage   = ["digital-signature", "key-encipherment", "tls-server"]
	sign {
		ca = routeros_system_certificate.root_ca.name
	}
}

resource "routeros_system_certificate" "client_crt" {
	name        = "Client-Certificate"
	common_name = "client.crt"
	key_size    = "prime256v1"
	// KUs: digitalSignature and/or keyAgreement
	key_usage   = ["digital-signature", "key-agreement", "tls-client"]
	sign {
		ca = routeros_system_certificate.root_ca.name
	}
}

resource "routeros_system_certificate" "unsigned_crt" {
	name             = "Unsigned-Certificate"
	common_name      = "unsigned.crt"
	key_size         = "1024"
	subject_alt_name = "DNS:router.lan,DNS:myrouter.lan,IP:192.168.88.1"
}
`
}
