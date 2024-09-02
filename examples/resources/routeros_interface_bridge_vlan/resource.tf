resource "routeros_interface_bridge_vlan" "bridge_vlan" {
  bridge   = "bridge"
  vlan_ids = ["50"]
  tagged = [
    "bridge",
    "ether1"
  ]
  untagged = [
    "ether5"
  ]
}

resource "routeros_interface_bridge_vlan" "bridge_vlan" {
  bridge   = "bridge"
  vlan_ids = ["4", "10", "20", "50", "100", "101", "102", "103", "112", "201", "202", "220", "254"]
}

resource "routeros_interface_bridge_vlan" "bridge_vlan" {
  bridge   = "bridge"
  vlan_ids = ["4", "10", "20", "50", "100-103", "112", "201", "202", "220", "254"]
}

resource "routeros_interface_bridge_vlan" "bridge_vlan" {
  bridge   = "bridge"
  vlan_ids = ["100-115", "120", "122", "128-130"]
}
