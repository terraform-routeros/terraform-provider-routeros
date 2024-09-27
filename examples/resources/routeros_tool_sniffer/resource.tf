resource "routeros_tool_sniffer" "test" {
  enabled = true

  streaming_enabled = true
  streaming_server  = "192.168.88.5:37008"
  filter_stream     = true

  filter_interface                = ["ether2"]
  filter_direction                = "rx"
  filter_operator_between_entries = "and"
}
