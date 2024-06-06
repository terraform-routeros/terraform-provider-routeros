package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testSystemCertificatesImportAddress = "routeros_system_certificate.external"

func TestAccSystemCertificatesTest_import(t *testing.T) {
	// ROS 7.12.x does not return IDs for created files.
	if !testCheckMinVersion(t, testFileMinVersion) {
		t.Logf("Test skipped, the minimum required version is %v", testFileMinVersion)
		return
	}

	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			var externalCrt MikrotikItem

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/certificate", "routeros_system_certificate"),
				Steps: []resource.TestStep{
					{
						Config: testAccSystemCertificatesImportConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testSystemCertificatesImportAddress),
							// external_crt
							testCheckResourceExists("routeros_system_certificate.external", "/certificate", &externalCrt),
							testCheckMikrotikItemAttr("routeros_system_certificate.external", &externalCrt, "name", "external.crt"),
							resource.TestCheckResourceAttr("routeros_system_certificate.external", "name", "external.crt"),
							resource.TestCheckResourceAttr("routeros_system_certificate.external", "common_name", "External Certificate"),
							resource.TestCheckResourceAttr("routeros_system_certificate.external", "private_key", "true"),
						),
					},
				},
			})

		})
	}
}

func testAccSystemCertificatesImportConfig() string {
	return providerConfig + `
data "routeros_x509" "cert" {
	data = <<EOT
	-----BEGIN CERTIFICATE-----
	MIIBlTCCATugAwIBAgIINLsws71B5zIwCgYIKoZIzj0EAwIwHzEdMBsGA1UEAwwU
	RXh0ZXJuYWwgQ2VydGlmaWNhdGUwHhcNMjQwNTE3MjEyOTUzWhcNMjUwNTE3MjEy
	OTUzWjAfMR0wGwYDVQQDDBRFeHRlcm5hbCBDZXJ0aWZpY2F0ZTBZMBMGByqGSM49
	AgEGCCqGSM49AwEHA0IABKE1g0Qj4ujIold9tklu2z4BUu/K7xDFF5YmedtOfJyM
	1/80APNboqn71y4m4XNE1JNtQuR2bSZPHVrzODkR16ujYTBfMA8GA1UdEwEB/wQF
	MAMBAf8wDgYDVR0PAQH/BAQDAgG2MB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEF
	BQcDAjAdBgNVHQ4EFgQUNXd5bvluIV9YAhGc5yMHc6OzXpMwCgYIKoZIzj0EAwID
	SAAwRQIhAODte/qS6CE30cvnQpxP/ObWBPIPZnHtkFHIIC1AOSXwAiBGCGQE+aJY
	W72Rw0Y1ckvlt6sU0urkzGuj5wxVF/gSYA==
	-----END CERTIFICATE-----
EOT
}

resource "routeros_file" "key" {
	name     = "external.key"
	contents = <<EOT
-----BEGIN ENCRYPTED PRIVATE KEY-----
MIHeMEkGCSqGSIb3DQEFDTA8MBsGCSqGSIb3DQEFDDAOBAiy/wEW6/MglgICCAAw
HQYJYIZIAWUDBAEqBBD6v8dLA2FjPn62Xz57pcu9BIGQhclivPw1eC2b14ea58Tw
nzDdbYN6/yUiMqapW2xZaT7ZFnbEai4n9/utgtEDnfKHlZvZj2kRhvYoWrvTkt/W
1mkd5d/runsn+B5GO+CMHFHh4t41WMpZysmg+iP8FiiehOQEsWyEZFaedxfYYtSL
Sk+abxJ+NMQoh+S5d73niu1CO8uqQjOd8BoSOurURsOh
-----END ENCRYPTED PRIVATE KEY-----
EOT
}  

resource "routeros_file" "cert" {
	name     = "external.crt"
	contents = data.routeros_x509.cert.pem
}  

resource "routeros_system_certificate" "external" {
	name        = "external.crt"
	common_name = data.routeros_x509.cert.common_name
	import {
		cert_file_name  = routeros_file.cert.name
		key_file_name   = routeros_file.key.name
		passphrase      = "11111111"
	}
	depends_on = [routeros_file.key, routeros_file.cert]
}
`
}
