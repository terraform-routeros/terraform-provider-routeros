# THIS WORKS.
# terraform {
#   required_providers {
#     routeros = {
#       source  = "terraform-routeros/routeros"
#       version = "1.99.1"
#     }
#   }
# }



terraform {
  required_providers {
    routeros = {
      source  = "terraform.local/local/routeros"
      version = "1.0.0"
    }
  }
}

provider "routeros" {
  hosturl        = "http://localhost:8080"
  username       = "admin"
  password       = "password"
  insecure       = true
}

resource "routeros_interface_bridge" "bridge1" {
  name              = "bridge1"
  ageing_time       = "300s"
  ingress_filtering = true
  protocol_mode     = "rstp"
  priority          = "0x8000"
  igmp_snooping     = false
  vlan_filtering    = false
}


resource "routeros_ip_arp" "arp" {
  address = "192.168.88.128"
  interface = "bridge1"
  mac_address = "00:CA:FE:BA:BE:00"
  published = true
}