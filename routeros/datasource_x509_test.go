package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testDatasourceX509 = "data.routeros_x509.cert"

func TestAccDatasourceX509Test_basic(t *testing.T) {
	t.Run("X509", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProviderFactories: testAccProviderFactories,
			Steps: []resource.TestStep{
				{
					Config: testAccDatasourceX509Config(),
					Check: resource.ComposeTestCheckFunc(
						testResourcePrimaryInstanceId(testDatasourceX509),
					),
				},
			},
		})

	})
}

func testAccDatasourceX509Config() string {
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
`
}
