resource "routeros_capsman_datapath" "test_rule" {
  comment                   = "Catch-all"
  interface                 = "cap1"
  signal_range              = "-120..-85"
  allow_signal_out_of_range = "20s"
  action                    = "reject"
}
