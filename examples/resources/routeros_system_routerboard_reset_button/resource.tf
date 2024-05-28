resource "routeros_system_script" "reset_button" {
  name   = "reset-button"
  source = <<EOF
    :log info message=("reset button pressed")
  EOF
}

resource "routeros_system_routerboard_button_reset" "settings" {
  enabled   = true
  hold_time = "0s..1m"
  on_event  = routeros_system_script.reset_button.name
}
