resource "routeros_system_note" "test" {
  note              = "For authorized use only."
  show_at_login     = true
  show_at_cli_login = true
}
