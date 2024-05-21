resource "routeros_system_certificate" "root_ca" {
  name        = "Test-Root-CA"
  common_name = "RootCA"
  key_usage   = ["key-cert-sign", "crl-sign"]
  trusted     = true
  # Sign Root CA.
  sign {
  }
}

# digitalSignature: Used for entity and data origin authentication with integrity.
# keyEncipherment:  Used to encrypt symmetric key, which is then transferred to target.
# keyAgreement:     Enables use of key agreement to establish symmetric key with target. 

resource "routeros_system_certificate" "server_crt" {
  name        = "Server-Certificate"
  common_name = "server.crt"
  #  KUs: igitalSignature, keyEncipherment or keyAgreement
  key_usage = ["digital-signature", "key-encipherment", "tls-server"]
  sign {
    ca = routeros_system_certificate.root_ca.name
  }
}

resource "routeros_system_certificate" "client_crt" {
  name        = "Client-Certificate"
  common_name = "client.crt"
  key_size    = "prime256v1"
  #  KUs: digitalSignature and/or keyAgreement
  key_usage = ["digital-signature", "key-agreement", "tls-client"]
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
    scep_url = "http://scep.server/scep/test"
  }
}

#  Import certificate
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
  name = "external.key"
  # The lines of the certificate must not contain indentation.
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
  name = "external.crt"
  # Normalized certificate
  contents = data.routeros_x509.cert.pem
}

resource "routeros_system_certificate" "external" {
  name        = "external.crt"
  common_name = data.routeros_x509.cert.common_name
  import {
    cert_file_name = routeros_file.cert.name
    key_file_name  = routeros_file.key.name
    passphrase     = "11111111"
  }
  depends_on = [routeros_file.key, routeros_file.cert]
}
