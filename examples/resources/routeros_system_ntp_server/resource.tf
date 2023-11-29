resource "routeros_system_ntp_server" "test" {
  enabled             = true
  broadcast           = true
  multicast           = true
  manycast            = true
  use_local_clock     = true
  local_clock_stratum = 3
}