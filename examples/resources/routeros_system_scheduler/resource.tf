resource "routeros_system_scheduler" "schedule1" {
  name     = "schedule1"
  on-event = "script name"
}