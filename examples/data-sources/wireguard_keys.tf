data "routeros_wireguard_keys" "wgk" {
  number = 3
}

output "wg_keys" {
  value = data.routeros_wireguard_keys.wgk.keys[*]
}

output "wg_key" {
  value = data.routeros_wireguard_keys.wgk.keys[2].private
}