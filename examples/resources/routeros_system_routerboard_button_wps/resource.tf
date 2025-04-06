resource "routeros_system_script" "wps_button" {
  name   = "wps-button"
  source = <<EOF
    :foreach iface in=[/interface/wifi find where configuration.mode="ap"] do={
      /interface/wifi wps-push-button $iface
    }
  EOF
}

resource "routeros_system_routerboard_button_wps" "settings" {
  enabled   = true
  hold_time = "0s..1m"
  on_event  = routeros_system_script.wps_button.name
}
