resource "routeros_interface_l2tp_client" "test" {
  name       = "l2tp-test-out"
  connect_to = "127.0.0.1"
  user       = "aaa"
  password   = "bbb"
}
