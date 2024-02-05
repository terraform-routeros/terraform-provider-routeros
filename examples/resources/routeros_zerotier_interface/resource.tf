resource "zerotier_identity" "identity" {}

resource "zerotier_network" "network" {
  name = "test"
}

resource "zerotier_member" "member" {
  authorized              = true
  member_id               = zerotier_identity.identity.id
  name                    = "test"
  network_id              = zerotier_network.network.id
  hidden                  = false
  allow_ethernet_bridging = true
  no_auto_assign_ips      = true
}

resource "routeros_zerotier" "zt1" {
  comment    = "ZeroTier Central"
  identity   = zerotier_identity.identity.private_key
  interfaces = ["all"]
  name       = "zt1"
}

resource "routeros_zerotier_interface" "zerotier1" {
  allow_default = false
  allow_global  = false
  allow_managed = false
  instance      = routeros_zerotier.zt1.name
  name          = "zerotier1"
  network       = zerotier_network.network.id
}
