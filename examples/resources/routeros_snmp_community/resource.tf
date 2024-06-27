resource "routeros_snmp_community" "test" {
  authentication_password = "authpasswd"
  authentication_protocol = "MD5"
  comment                 = "Comment"
  disabled                = true
  encryption_password     = "encpassword"
  encryption_protocol     = "DES"
  name                    = "private"
  read_access             = true
  security                = "private"
  write_access            = true
}

resource "routeros_snmp_community" "mything" {
  addresses = ["10.0.1.12", "10.10.0.129"]
  name      = "mything"
}
