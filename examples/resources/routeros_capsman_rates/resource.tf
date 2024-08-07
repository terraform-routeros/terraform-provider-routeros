resource "routeros_capsman_rates" "test_rates" {
  name              = "test_rates"
  comment           = "test_rates"
  basic             = ["1Mbps", "5.5Mbps", "6Mbps", "18Mbps", "36Mbps", "54Mbps"]
  ht_basic_mcs      = ["mcs-0", "mcs-7", "mcs-11", "mcs-14", "mcs-16", "mcs-21"]
  ht_supported_mcs  = ["mcs-3", "mcs-8", "mcs-10", "mcs-13", "mcs-17", "mcs-18"]
  supported         = ["2Mbps", "11Mbps", "9Mbps", "12Mbps", "24Mbps", "48Mbps"]
  vht_basic_mcs     = "none"
  vht_supported_mcs = "mcs0-9,mcs0-7"
}