resource "routeros_snmp" "test" {
  contact          = "John D."
  enabled          = true
  engine_id_suffix = "8a3c"
  location         = "Backyard"
  trap_community   = "private"
  trap_generators  = "start-trap"
  trap_version     = 3
}
