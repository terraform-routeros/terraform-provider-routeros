# Certificate rotation

Original [issue](https://github.com/terraform-routeros/terraform-provider-routeros/issues/584)

## Example
```terraform
resource "tls_private_key" "ca_key" {
  algorithm = "RSA"
}

resource "tls_self_signed_cert" "ca_cert" {
  subject {
    common_name  = "testCA"
    organization = "test"
  }

  private_key_pem       = tls_private_key.ca_key.private_key_pem
  allowed_uses          = ["digital_signature", "cert_signing", "crl_signing"]
  validity_period_hours = 24 * 365 * 5
  is_ca_certificate     = true
}

resource "tls_private_key" "server_key" {
  algorithm = "RSA"
}

resource "tls_cert_request" "server_csr" {
  private_key_pem = tls_private_key.server_key.private_key_pem
  subject {
    common_name  = "mikrotik.example.com"
    organization = "test"
  }
}

resource "tls_locally_signed_cert" "server_cert" {
  cert_request_pem   = tls_cert_request.server_csr.cert_request_pem
  ca_private_key_pem = tls_private_key.ca_key.private_key_pem
  ca_cert_pem        = tls_self_signed_cert.ca_cert.cert_pem

  validity_period_hours = 12

  allowed_uses = [
    "key_encipherment",
    "digital_signature",
    "server_auth",
  ]
}

output "cert_serial_number_expected" {
  value = format("%x", tls_locally_signed_cert.server_cert.id)
}

resource "routeros_file" "server_key" {
  name     = "server.key"
  contents = tls_private_key.server_key.private_key_pem
}

resource "routeros_file" "server_cert" {
  name     = "server.crt"
  contents = tls_locally_signed_cert.server_cert.cert_pem
}

resource “routeros_system_certificate” “server_cert” {
  name = “server”
  common_name = tls_cert_request.server_csr.subject[0].common_name
  import {
    cert_file_name = routeros_file.server_cert.name
    key_file_name = routeros_file.server_key.name
  }
  depends_on = [routeros_file.server_cert, routeros_file.server_key]
  lifecycle {
    replace_triggered_by = [
      tls_locally_signed_cert.server_cert.cert_pem
    ]
  }
}

output "cert_serial_nubmer_on_device" {
  value = routeros_system_certificate.server_cert.serial_number
}
```