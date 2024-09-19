resource "routeros_dns_adlist" "test" {
  url        = "https://raw.githubusercontent.com/StevenBlack/hosts/master/hosts"
  ssl_verify = false
}
