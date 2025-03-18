resource "routeros_ip_smb" "test" {
  enabled    = "auto"
  domain     = "MSHOME"
  comment    = "MikrotikSMB"
  interfaces = ["all"]
}
