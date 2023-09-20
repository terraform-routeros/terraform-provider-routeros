
resource "routeros_system_logging" "log_snmp_disk" {
  action = "disk"
  topics = ["snmp"]
}