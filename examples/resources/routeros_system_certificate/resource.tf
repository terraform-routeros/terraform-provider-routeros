resource "routeros_system_certificate" "root_ca" {
  name        = "Test-Root-CA"
  common_name = "RootCA"
  key_usage   = ["key-cert-sign", "crl-sign"]
  trusted     = true
  // Sign Root CA.
  sign {
  }
}

// digitalSignature: Used for entity and data origin authentication with integrity.
// keyEncipherment:  Used to encrypt symmetric key, which is then transferred to target.
// keyAgreement:     Enables use of key agreement to establish symmetric key with target. 

resource "routeros_system_certificate" "server_crt" {
  name        = "Server-Certificate"
  common_name = "server.crt"
  // KUs: igitalSignature, keyEncipherment or keyAgreement
  key_usage = ["digital-signature", "key-encipherment", "tls-server"]
  sign {
    ca = routeros_system_certificate.root_ca.name
  }
}

resource "routeros_system_certificate" "client_crt" {
  name        = "Client-Certificate"
  common_name = "client.crt"
  key_size    = "prime256v1"
  // KUs: digitalSignature and/or keyAgreement
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