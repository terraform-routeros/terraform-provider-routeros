resource "routeros_system_scheduler" "schedule1" {
  name     = "schedule1"
  on_event = "script name"
}