variable "rule" {
  type = list(object({
    chain              = string
    action             = string
    connection_state   = optional(string)
    in_interface_list  = optional(string, "all")
    out_interface_list = optional(string)
    src_address        = optional(string, "0.0.0.0/0")
    dst_address        = optional(string)
    src_port           = optional(string)
    dst_port           = optional(string)
    protocol           = optional(string)
    comment            = optional(string, "(terraform-defined)")
    log                = optional(bool, false)
    disabled           = optional(bool, true)
  }))

  default = [
    { chain = "input", action = "accept", comment = "00" },
    { chain = "input", action = "accept", comment = "01" },
    { chain = "input", action = "accept", comment = "02" },
    { chain = "input", action = "accept", comment = "03" },
    { chain = "input", action = "accept", comment = "04" },
    { chain = "input", action = "accept", comment = "05" },
    { chain = "input", action = "accept", comment = "06" },
    { chain = "input", action = "accept", comment = "07" },
    { chain = "input", action = "accept", comment = "08" },
    { chain = "input", action = "accept", comment = "09" },
    { chain = "input", action = "accept", comment = "10" },
    { chain = "input", action = "accept", comment = "11" },
    { chain = "input", action = "accept", comment = "12" },
    { chain = "input", action = "accept", comment = "13" },
    { chain = "input", action = "accept", comment = "14" },
    { chain = "input", action = "accept", comment = "15" },
    { chain = "input", action = "accept", comment = "16" },
    { chain = "input", action = "accept", comment = "17" },
    { chain = "input", action = "accept", comment = "18" },
    { chain = "input", action = "accept", comment = "19" },
    { chain = "input", action = "accept", comment = "20" },
    { chain = "input", action = "accept", comment = "21" },
    { chain = "input", action = "accept", comment = "22" },
    { chain = "input", action = "accept", comment = "23" },
    { chain = "input", action = "accept", comment = "24" },
    { chain = "input", action = "accept", comment = "25" },
    { chain = "input", action = "accept", comment = "26" },
    { chain = "input", action = "accept", comment = "27" },
    { chain = "input", action = "accept", comment = "28" },
    { chain = "input", action = "accept", comment = "29" },
    { chain = "input", action = "accept", comment = "30" },
    { chain = "input", action = "accept", comment = "31" },
  ]
}

locals {
  # https://discuss.hashicorp.com/t/does-map-sort-keys/12056/2
  # Map keys are always iterated in lexicographical order!
  rule_map = { for idx, rule in var.rule : format("%03d", idx) => rule }
}

resource "routeros_ip_firewall_filter" "rules" {
  for_each          = local.rule_map
  chain             = each.value.chain
  action            = each.value.action
  comment           = each.value.comment
  log               = each.value.log
  disabled          = each.value.disabled
  connection_state  = each.value.connection_state
  in_interface_list = each.value.in_interface_list
  src_address       = each.value.src_address
  dst_port          = each.value.dst_port
  protocol          = each.value.protocol
}

resource "routeros_move_items" "fw_rules" {
  #  resource_name = "routeros_ip_firewall_filter"
  resource_path = "/ip/firewall/filter"
  sequence      = [for i, _ in local.rule_map : routeros_ip_firewall_filter.rules[i].id]
  depends_on    = [routeros_ip_firewall_filter.rules]
}
