resource "routeros_ip_arp" "arp" {
  address = "192.168.88.128"
  interface = "bridge1"
  mac_address = "00:CA:FE:BA:BE:00"
  published = true
}