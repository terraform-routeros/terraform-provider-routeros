resource "routeros_wireguard_keys" "wgk" {
  number = 3
}

output "wg_keys" {
  value     = routeros_wireguard_keys.wgk.keys[*]
  sensitive = true
}

output "wg_key" {
  value = nonsensitive(routeros_wireguard_keys.wgk.keys[1].public)
}