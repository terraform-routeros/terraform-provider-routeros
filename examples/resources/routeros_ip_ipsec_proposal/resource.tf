resource "routeros_ip_ipsec_proposal" "test" {
  name      = "NordVPN"
  pfs_group = "none"
  lifetime  = "45m10s"
}
