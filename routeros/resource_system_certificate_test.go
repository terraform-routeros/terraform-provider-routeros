package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testSystemCertificatesAddress = "routeros_system_certificate.root_ca"

func TestAccSystemCertificatesTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			var rootCa, caCrt, serverCrt, clientCrt, scepClient, unsignedCrt MikrotikItem

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
							// root_ca
							testCheckResourceExists(testSystemCertificatesAddress, "/certificate", &rootCa),
							testCheckMikrotikItemAttr(testSystemCertificatesAddress, &rootCa, "name", "Test-Root-CA"),
							testCheckMikrotikItemAttrSet(testSystemCertificatesAddress, &rootCa, "fingerprint"),
							resource.TestCheckResourceAttr(testSystemCertificatesAddress, "name", "Test-Root-CA"),
							resource.TestCheckResourceAttrSet(testSystemCertificatesAddress, "fingerprint"),
							// ca_crt
							testCheckResourceExists("routeros_system_certificate.ca_crt", "/certificate", &caCrt),
							testCheckMikrotikItemAttr("routeros_system_certificate.ca_crt", &caCrt, "name", "CA-Certificate"),
							resource.TestCheckResourceAttr("routeros_system_certificate.ca_crt", "name", "CA-Certificate"),
							// server_crt
							testCheckResourceExists("routeros_system_certificate.server_crt", "/certificate", &serverCrt),
							testCheckMikrotikItemAttr("routeros_system_certificate.server_crt", &serverCrt, "name", "Server-Certificate"),
							testCheckMikrotikItemAttr("routeros_system_certificate.server_crt", &serverCrt, "ca", "Test-Root-CA"),
							testCheckMikrotikItemAttrSet("routeros_system_certificate.server_crt", &serverCrt, "fingerprint"),
							resource.TestCheckResourceAttr("routeros_system_certificate.server_crt", "name", "Server-Certificate"),
							resource.TestCheckResourceAttr("routeros_system_certificate.server_crt", "ca", "Test-Root-CA"),
							resource.TestCheckResourceAttrSet("routeros_system_certificate.server_crt", "fingerprint"),
							// client_crt
							testCheckResourceExists("routeros_system_certificate.client_crt", "/certificate", &clientCrt),
							testCheckMikrotikItemAttr("routeros_system_certificate.client_crt", &clientCrt, "name", "Client-Certificate"),
							resource.TestCheckResourceAttr("routeros_system_certificate.client_crt", "name", "Client-Certificate"),
							// unsigned_crt
							testCheckResourceExists("routeros_system_certificate.unsigned_crt", "/certificate", &unsignedCrt),
							testCheckMikrotikItemAttr("routeros_system_certificate.unsigned_crt", &unsignedCrt, "name", "Unsigned-Certificate"),
							testCheckNoMikrotikItemAttr("routeros_system_certificate.unsigned_crt", &unsignedCrt, "ca"),
							testCheckNoMikrotikItemAttr("routeros_system_certificate.unsigned_crt", &unsignedCrt, "fingerprint"),
							resource.TestCheckResourceAttr("routeros_system_certificate.unsigned_crt", "name", "Unsigned-Certificate"),
							resource.TestCheckNoResourceAttr("routeros_system_certificate.unsigned_crt", "ca"),
							resource.TestCheckNoResourceAttr("routeros_system_certificate.unsigned_crt", "fingerprint"),
							// scep_client
							testCheckResourceExists("routeros_system_certificate.scep_client", "/certificate", &scepClient),
							testCheckMikrotikItemAttr("routeros_system_certificate.scep_client", &scepClient, "name", "SCEP-Client"),
							testCheckMikrotikItemAttr("routeros_system_certificate.scep_client", &scepClient, "common-name", "scep-client.crt"),
							testCheckMikrotikItemAttr("routeros_system_certificate.scep_client", &scepClient, "key-usage", "digital-signature,key-agreement,tls-client"),
							testCheckMikrotikItemAttr("routeros_system_certificate.scep_client", &scepClient, "scep-url", "http://scep.server/scep/test"),
							testCheckMikrotikItemAttr("routeros_system_certificate.scep_client", &scepClient, "challenge-password", "12345"),
							testCheckMikrotikItemAttrSet("routeros_system_certificate.scep_client", &scepClient, "status"),
							resource.TestCheckResourceAttr("routeros_system_certificate.scep_client", "name", "SCEP-Client"),
							resource.TestCheckResourceAttr("routeros_system_certificate.scep_client", "common_name", "scep-client.crt"),
							resource.TestCheckResourceAttr("routeros_system_certificate.scep_client", "key_usage.#", "3"),
							resource.TestCheckResourceAttr("routeros_system_certificate.scep_client", "key_usage.0", "digital-signature"),
							resource.TestCheckResourceAttr("routeros_system_certificate.scep_client", "key_usage.1", "key-agreement"),
							resource.TestCheckResourceAttr("routeros_system_certificate.scep_client", "key_usage.2", "tls-client"),
							resource.TestCheckResourceAttr("routeros_system_certificate.scep_client", "scep_url", "http://scep.server/scep/test"),
							resource.TestCheckResourceAttr("routeros_system_certificate.scep_client", "challenge_password", "12345"),
							resource.TestCheckResourceAttrSet("routeros_system_certificate.scep_client", "status"),
							resource.TestCheckResourceAttrPair("routeros_system_certificate.scep_client", "scep_url", "routeros_system_certificate.scep_client", "sign_via_scep.0.scep_url"),
							resource.TestCheckResourceAttrPair("routeros_system_certificate.scep_client", "challenge_password", "routeros_system_certificate.scep_client", "sign_via_scep.0.challenge_password"),
						),
					},
					{
						Config: testAccSystemCertificatesConfigUpdated(),
						Check: resource.ComposeAggregateTestCheckFunc(
							// root_ca
							testCheckResourceExists(testSystemCertificatesAddress, "/certificate", &rootCa),
							testCheckMikrotikItemAttr(testSystemCertificatesAddress, &rootCa, "name", "Test-Root-CA"),
							// ca_crt (added ca_crl_host)
							testCheckResourceExists("routeros_system_certificate.ca_crt", "/certificate", &caCrt),
							testCheckMikrotikItemAttr("routeros_system_certificate.ca_crt", &caCrt, "name", "CA-Certificate"),
							testCheckMikrotikItemAttr("routeros_system_certificate.ca_crt", &caCrt, "ca-crl-host", "10.0.0.1"),
							resource.TestCheckResourceAttr("routeros_system_certificate.ca_crt", "name", "CA-Certificate"),
							resource.TestCheckResourceAttr("routeros_system_certificate.ca_crt", "ca_crl_host", "10.0.0.1"),
							// server_crt (changed name)
							testCheckResourceExists("routeros_system_certificate.server_crt", "/certificate", &serverCrt),
							testCheckMikrotikItemAttr("routeros_system_certificate.server_crt", &serverCrt, "name", "Server-Certificate-New"),
							testCheckMikrotikItemAttr("routeros_system_certificate.server_crt", &serverCrt, "ca", "Test-Root-CA"),
							resource.TestCheckResourceAttr("routeros_system_certificate.server_crt", "name", "Server-Certificate-New"),
							resource.TestCheckResourceAttr("routeros_system_certificate.server_crt", "ca", "Test-Root-CA"),
							// client_crt (added country, renamed to avoid 'name exists' error)
							testCheckResourceExists("routeros_system_certificate.client_crt", "/certificate", &clientCrt),
							testCheckMikrotikItemAttr("routeros_system_certificate.client_crt", &clientCrt, "name", "Client-Certificate-New"),
							testCheckMikrotikItemAttr("routeros_system_certificate.client_crt", &clientCrt, "country", "US"),
							resource.TestCheckResourceAttr("routeros_system_certificate.client_crt", "name", "Client-Certificate-New"),
							resource.TestCheckResourceAttr("routeros_system_certificate.client_crt", "country", "US"),
							// unsigned_crt (certificate requested to be signed)
							testCheckResourceExists("routeros_system_certificate.unsigned_crt", "/certificate", &unsignedCrt),
							testCheckMikrotikItemAttr("routeros_system_certificate.unsigned_crt", &unsignedCrt, "name", "Unsigned-Certificate"),
							testCheckMikrotikItemAttr("routeros_system_certificate.unsigned_crt", &unsignedCrt, "ca", "Test-Root-CA"),
							testCheckMikrotikItemAttrSet("routeros_system_certificate.unsigned_crt", &unsignedCrt, "fingerprint"),
							resource.TestCheckResourceAttr("routeros_system_certificate.unsigned_crt", "name", "Unsigned-Certificate"),
							resource.TestCheckResourceAttr("routeros_system_certificate.unsigned_crt", "ca", "Test-Root-CA"),
							resource.TestCheckResourceAttrSet("routeros_system_certificate.unsigned_crt", "fingerprint"),
							// scep_client (changed scep_url)
							testCheckResourceExists("routeros_system_certificate.scep_client", "/certificate", &scepClient),
							testCheckMikrotikItemAttr("routeros_system_certificate.scep_client", &scepClient, "name", "SCEP-Client"),
							testCheckMikrotikItemAttr("routeros_system_certificate.scep_client", &scepClient, "common-name", "scep-client.crt"),
							testCheckMikrotikItemAttr("routeros_system_certificate.scep_client", &scepClient, "scep-url", "http://scep.server/scep/test_updated"),
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

resource "routeros_system_certificate" "ca_crt" {
	name        = "CA-Certificate"
	common_name = "ca.crt"
	key_usage   = ["key-cert-sign", "crl-sign"]
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

resource "routeros_system_certificate" "scep_client" {
  name        = "SCEP-Client"
  common_name = "scep-client.crt"
  key_usage   = ["digital-signature", "key-agreement", "tls-client"]

  sign_via_scep {
    scep_url           = "http://scep.server/scep/test"
    challenge_password = "12345"
  }
}
`
}

func testAccSystemCertificatesConfigUpdated() string {
	return providerConfig + `

resource "routeros_system_certificate" "root_ca" {
  name        = "Test-Root-CA"
  common_name = "RootCA"
  key_usage   = ["key-cert-sign", "crl-sign"]
  trusted     = true
  sign {
  }
}

# Added ca_crl_host
resource "routeros_system_certificate" "ca_crt" {
  name        = "CA-Certificate"
  common_name = "ca.crt"
  key_usage   = ["key-cert-sign", "crl-sign"]
  sign {
    ca_crl_host = "10.0.0.1"
  }
}

# Changed name
resource "routeros_system_certificate" "server_crt" {
  name        = "Server-Certificate-New"
  common_name = "server.crt"
  // KUs: igitalSignature, keyEncipherment or keyAgreement
  key_usage   = ["digital-signature", "key-encipherment", "tls-server"]
  sign {
    ca = routeros_system_certificate.root_ca.name
  }
}

# Added country, have to change a name
resource "routeros_system_certificate" "client_crt" {
  name        = "Client-Certificate-New"
  common_name = "client.crt"
  country     = "US"
  key_size    = "prime256v1"
  // KUs: digitalSignature and/or keyAgreement
  key_usage   = ["digital-signature", "key-agreement", "tls-client"]
  sign {
    ca = routeros_system_certificate.root_ca.name
  }
}

# Sign unsigned certificate
resource "routeros_system_certificate" "unsigned_crt" {
  name             = "Unsigned-Certificate"
  common_name      = "unsigned.crt"
  key_size         = "1024"
  subject_alt_name = "DNS:router.lan,DNS:myrouter.lan,IP:192.168.88.1"
  sign {
    ca = routeros_system_certificate.root_ca.name
  }
}

# Changed scep_url
resource "routeros_system_certificate" "scep_client" {
  name        = "SCEP-Client"
  common_name = "scep-client.crt"
  key_usage   = ["digital-signature", "key-agreement", "tls-client"]

  sign_via_scep {
    scep_url           = "http://scep.server/scep/test_updated"
    challenge_password = "12345"
  }
}
`
}
