variable "bridge_filter_rule" {
  type = list(object({
    chain              = string
    action             = string
    connection_state   = optional(string)
    in_interface_list  = optional(string, "all")
    out_interface_list = optional(string)
    src_address        = optional(string)
    dst_address        = optional(string)
    src_port           = optional(string)
    dst_port           = optional(string)
    jump_target        = optional(string)
    protocol           = optional(string)
    comment            = optional(string, "(terraform-defined)")
    log                = optional(bool, false)
    log_prefix         = optional(string, "")
    disabled           = optional(bool, false)
  }))

  default = [
    { "action" = "drop", "chain" = "forward", "comment" = "Drop data between bridge ports" },
    { "action" = "drop", "chain" = "forward", "comment" = "Block VLAN encap", "log_prefix" = "Block VLAN encap", "mac_protocol" = "vlan" },
    { "action" = "accept", "chain" = "forward", "comment" = "", "disabled" = "true", "dst_address" = "224.0.0.251/32", "ip_protocol" = "udp", "log_prefix" = "Allow bonjour", "mac_protocol" = "ip" },
  ]
}

locals {
  rule_map = { for idx, rule in var.bridge_filter_rule : format("%03d", idx) => rule }
}

resource "routeros_interface_bridge_filter" "rules" {
  for_each          = local.rule_map
  chain             = each.value.chain
  action            = each.value.action
  comment           = each.value.comment
  log               = each.value.log
  log_prefix        = each.value.log_prefix
  disabled          = each.value.disabled
  connection_state  = each.value.connection_state
  in_interface_list = each.value.in_interface_list
  dst_port          = each.value.dst_port
  protocol          = each.value.protocol
  src_address       = each.value.src_address
  jump_target       = each.value.jump_target
}

resource "routeros_move_items" "bridge_filter_rules" {
  #  resource_name = "routeros_interface_bridge_filter"
  resource_path = "/interface/bridge/filter"
  sequence      = [for i, _ in local.rule_map : routeros_interface_bridge_filter.rules[i].id]
  depends_on    = [routeros_interface_bridge_filter.rules]
}
