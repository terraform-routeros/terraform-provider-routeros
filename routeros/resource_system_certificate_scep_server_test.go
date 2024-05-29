package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceCertificateScepServer_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: resource.ComposeAggregateTestCheckFunc(
					testCheckResourceDestroy("/certificate", "routeros_system_certificate"),
					testCheckResourceDestroy("/certificate/scep-server", "routeros_certificate_scep_server"),
				),
				Steps: []resource.TestStep{
					{
						Config: testAccResourceCertificateScepServerConfig,
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId("routeros_certificate_scep_server.test_acc"),
							resource.TestCheckResourceAttr("routeros_certificate_scep_server.test_acc", "ca_cert", "test_acc_root_ca"),
							resource.TestCheckResourceAttr("routeros_certificate_scep_server.test_acc", "path", "/scep/test_acc_scep_server"),
						),
					},
					{
						Config:            testAccResourceCertificateScepServerConfigUpdated,
						ImportState:       true,
						ResourceName:      "routeros_certificate_scep_server.test_acc",
						ImportStateVerify: true,
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId("routeros_certificate_scep_server.test_acc"),
							resource.TestCheckResourceAttr("routeros_certificate_scep_server.test_acc", "ca_cert", "test_acc_root_ca"),
							resource.TestCheckResourceAttr("routeros_certificate_scep_server.test_acc", "next_ca_cert", "test_acc_root_ca"),
							resource.TestCheckResourceAttr("routeros_certificate_scep_server.test_acc", "path", "/scep/test_acc_scep_server_updated"),
							resource.TestCheckResourceAttr("routeros_certificate_scep_server.test_acc", "days_valid", "30"),
							resource.TestCheckResourceAttr("routeros_certificate_scep_server.test_acc", "request_lifetime", "2h"),
						),
					},
				},
			})
		})
	}
}

var testAccResourceCertificateScepServerConfigDeps = providerConfig + `
resource "routeros_system_certificate" "test_acc" {
  name        = "test_acc_root_ca"
  common_name = "Test Acc Root CA"
  key_usage   = ["key-cert-sign", "crl-sign"]
  trusted     = true
  sign {
  }
}
`

var testAccResourceCertificateScepServerConfig = testAccResourceCertificateScepServerConfigDeps + `
resource "routeros_certificate_scep_server" "test_acc" {
  ca_cert = routeros_system_certificate.test_acc.name
  path    = "/scep/test_acc_scep_server"
}
`

var testAccResourceCertificateScepServerConfigUpdated = testAccResourceCertificateScepServerConfigDeps + `
resource "routeros_certificate_scep_server" "test_acc" {
  ca_cert          = routeros_system_certificate.test_acc.name
  next_ca_cert     = routeros_system_certificate.test_acc.name
  path             = "/scep/test_acc_scep_server_updated"
  days_valid       = 30
  request_lifetime = "2h"
}
`
