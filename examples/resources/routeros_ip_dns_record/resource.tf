resource "routeros_ip_dns_record" "name_record" {
  name    = "router.lan"
  address = "192.168.88.1"
}

resource "routeros_ip_dns_record" "regexp_record" {
  regexp  = ".*pool.ntp.org"
  address = "192.168.88.1"
}