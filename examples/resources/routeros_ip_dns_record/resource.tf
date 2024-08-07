resource "routeros_ip_dns_record" "name_record" {
  name    = "router.lan"
  address = "192.168.88.1"
  type    = "A"
}

resource "routeros_ip_dns_record" "regexp_record" {
  regexp  = ".*pool.ntp.org"
  address = "192.168.88.1"
  type    = "A"
}

resource "routeros_dns_record" "aaaa_record" {
  name    = "ipv6.lan"
  address = "ff00::1"
  type    = "AAAA"
}

resource "routeros_dns_record" "cname_record" {
  name  = "cname.lan"
  cname = "ipv4.lan"
  type  = "CNAME"
}

resource "routeros_dns_record" "fwd_record" {
  name       = "fwd.lan"
  forward_to = "127.0.0.1"
  type       = "FWD"
}

resource "routeros_dns_record" "mx_record" {
  name          = "mx.lan"
  mx_exchange   = "127.0.0.1"
  mx_preference = 10
  type          = "MX"
}

resource "routeros_dns_record" "ns_record" {
  name = "ns.lan"
  ns   = "127.0.0.1"
  type = "NS"
}

resource "routeros_dns_record" "nxdomain_record" {
  name = "nxdomain.lan"
  type = "NXDOMAIN"
}

resource "routeros_dns_record" "srv_record" {
  name         = "srv.lan"
  srv_port     = 8080
  srv_priority = 10
  srv_target   = "127.0.0.1"
  srv_weight   = 100
  type         = "SRV"
}

resource "routeros_dns_record" "txt_record" {
  name = "_acme-challenge.yourwebsite.com"
  text = "dW6MrI3nBy3eJgYWH3QAg1Cwk_TvjFESOuKo+mp6nm1"
  type = "TXT"
}
