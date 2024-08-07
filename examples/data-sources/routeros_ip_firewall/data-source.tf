data "routeros_ip_firewall" "fw" {
  rules {
    filter = {
      chain   = "input"
      comment = "rule_2"
    }
  }

  rules {
    filter = {
      chain = "forward"
    }
  }

  nat {}
}

output "rules" {
  value = [for value in data.routeros_ip_firewall.fw.rules : [value.id, value.src_address]]
}

output "nat" {
  value = [for value in data.routeros_ip_firewall.fw.nat : [value.id, value.comment]]
}

resource "routeros_ip_firewall" "rule_3" {
  action       = "accept"
  chain        = "input"
  comment      = "rule_3"
  src_address  = "192.168.0.5"
  place_before = data.routeros_ip_firewall_filter.fw.rules[0].id
}
