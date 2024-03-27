resource "zerotier_identity" "identity" {}

resource "routeros_zerotier" "zt1" {
  comment    = "ZeroTier Central"
  identity   = zerotier_identity.identity.private_key
  interfaces = ["all"]
  name       = "zt1"
}
