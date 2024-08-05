resource "routeros_tool_netwatch" "test" {
  name      = "watch-google-pdns"
  host      = "8.8.8.8"
  interval  = "30s"
  up_script = ":log info \"Ping to 8.8.8.8 successful\""
}
