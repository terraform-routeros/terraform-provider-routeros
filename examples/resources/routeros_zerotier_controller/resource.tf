resource "zerotier_identity" "identity" {}

resource "routeros_zerotier" "zt1" {
  comment    = "ZeroTier Central"
  identity   = zerotier_identity.identity.private_key
  interfaces = ["all"]
  name       = "zt1"
}

resource "routeros_zerotier_controller" "test" {
  instance = routeros_zerotier.zt1.name
  name     = "test"
  network  = "1234567812345678"
  private  = true
}
