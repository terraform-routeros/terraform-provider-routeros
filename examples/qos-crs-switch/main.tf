# End-to-end DSCP-based QoS on a MikroTik CRS3xx/CRS5xx switch chip.
#
# Reproduces the switch-chip QoS stack used on a CRS309-1G-8S+ ToR switch:
#   * 7 traffic classes mapped from DSCP
#   * L3 (IP/DSCP) classification map
#   * per-port trust configuration
#   * egress scheduler (tx-manager) weights
#
# Prerequisite: qos-hw-offloading must be enabled on the switch chip.

terraform {
  required_providers {
    routeros = {
      source = "terraform-routeros/routeros"
    }
  }
}

# --- Switch chip: enable L3 + QoS hardware offloading -----------------------
resource "routeros_interface_ethernet_switch" "switch1" {
  switch_id         = "*0"
  l3_hw_offloading  = true
  qos_hw_offloading = true
}

# --- QoS profiles: DSCP -> internal traffic class ---------------------------
locals {
  profiles = {
    "1-Real-Time (EF/VoIP/Video)"           = { dscp = 46, tc = 7 }
    "2-Critical (CS6/Sync DB)"              = { dscp = 48, tc = 6 }
    "3-LatencyCritical (AF41/etcd/DB-Async)" = { dscp = 34, tc = 5 }
    "4-Interactive (AF31/Web/SSH)"          = { dscp = 26, tc = 4 }
    "6-Bulk Data (CS1/Downloads)"           = { dscp = 10, tc = 0 }
    "7-Backups (CS1/Backup)"                = { dscp = 8, tc = 0 }
  }
}

resource "routeros_interface_ethernet_switch_qos_profile" "p" {
  for_each      = local.profiles
  name          = each.key
  dscp          = each.value.dscp
  traffic_class = each.value.tc

  depends_on = [routeros_interface_ethernet_switch.switch1]
}

# The best-effort default profile (no DSCP, default traffic class).
resource "routeros_interface_ethernet_switch_qos_profile" "standard" {
  name       = "5-Standard (BE/Default)"
  depends_on = [routeros_interface_ethernet_switch.switch1]
}

# --- IP (L3/DSCP) classification map ----------------------------------------
resource "routeros_interface_ethernet_switch_qos_map_ip" "m" {
  for_each = local.profiles
  dscp     = each.value.dscp
  profile  = each.key

  depends_on = [routeros_interface_ethernet_switch_qos_profile.p]
}

resource "routeros_interface_ethernet_switch_qos_map_ip" "be" {
  dscp    = 0
  profile = routeros_interface_ethernet_switch_qos_profile.standard.name
}

# --- Per-port trust ---------------------------------------------------------
resource "routeros_interface_ethernet_switch_qos_port" "trunk" {
  for_each = toset(["sfp-sfpplus1", "sfp-sfpplus3", "sfp-sfpplus4"])
  name     = each.key
  profile  = routeros_interface_ethernet_switch_qos_profile.standard.name
  trust_l2 = "trust"
  trust_l3 = "trust"
}

# --- Egress scheduler (tx-manager) weights ----------------------------------
resource "routeros_interface_ethernet_switch_qos_tx_manager_queue" "tc5" {
  traffic_class = 5
  schedule      = "high-priority-group"
  weight        = 8
}

resource "routeros_interface_ethernet_switch_qos_tx_manager_queue" "tc6" {
  traffic_class = 6
  schedule      = "high-priority-group"
  weight        = 16
}

# --- Software queue assignment on an interface ------------------------------
resource "routeros_queue_interface" "ether1" {
  interface = "ether1"
  queue     = "ethernet-default"
}
