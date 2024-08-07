
resource "routeros_ip_dhcp_server_option" "jumbo_frame_opt" {
  code  = 77
  name  = "jumbo-mtu-opt"
  value = "0x2336"
}

resource "routeros_ip_dhcp_server_option" "tftp_option" {
  code  = 66
  name  = "tftpserver-66"
  value = "s'10.10.10.22'"
}

resource "routeros_ip_dhcp_server_option_set" "lan_option_set" {
  name    = "lan-option-set"
  options = join(",", [routeros_ip_dhcp_server_option.jumbo_frame_opt.name, routeros_ip_dhcp_server_option.tftp_option.name])
}