
resource "routeros_system_logging_action" "disk" {
  name               = "disk"
  target             = "remote"
  remote             = "192.168.1.1"
  bsd_syslog         = true
  syslog_facility    = "user"
  syslog_severity    = "notice"
  syslog_time_format = "iso8601"
}
