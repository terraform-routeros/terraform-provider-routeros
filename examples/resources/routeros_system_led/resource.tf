resource "routeros_system_led" "sfp1" {
  interface = "sfp1"
  leds      = ["sfp-led"]
  type      = "interface-activity"
}
