resource "routeros_system_script" "mode_button" {
  name   = "mode-button"
  source = <<EOF
    :log info message=("mode button pressed")
  EOF
}

resource "routeros_system_routerboard_button_mode" "settings" {
  enabled   = true
  hold_time = "0s..1m"
  on_event  = routeros_system_script.mode_button.name
}
