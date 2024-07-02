resource "routeros_capsman_channel" "channel1" {
  name      = "1"
  band      = "2ghz-g/n"
  frequency = [2412]
}

resource "routeros_capsman_interface" "cap1" {
  name = "cap1"

  channel = {
    config = routeros_capsman_channel.channel1.name
  }
}
