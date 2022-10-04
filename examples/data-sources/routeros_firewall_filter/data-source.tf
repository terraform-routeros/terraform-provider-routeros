data "routeros_firewall_filter" "fw" {
  filter = {
    chain = "input"
    comment = "rule_2"
  }
}

output "rules" {
  value = [for value in data.routeros_firewall_filter.fw.rules: [value.id, value.src_address]]
}

resource "routeros_firewall_filter" "rule_3" {
  action = "accept"
  chain  = "input"
  comment = "rule_3"
  src_address = "192.168.0.5"
  place_before = "${data.routeros_firewall_filter.fw.rules[0].id}"
}
