resource "routeros_capsman_channel" "test_channel" {
  name                  = "test_channel"
  comment               = "test_channel"
  band                  = "2ghz-b/g/n"
  control_channel_width = "10mhz"
  extension_channel     = "eCee"
  frequency             = [2412]
  reselect_interval     = "1h"
  save_selected         = true
  secondary_frequency   = ["disabled"]
  skip_dfs_channels     = true
  tx_power              = 20
}
