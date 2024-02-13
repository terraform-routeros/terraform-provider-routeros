resource "routeros_container_mounts" "caddyfile" {
  name = "Caddyfile"
  src  = "/usb1-part1/containers/caddy/Caddyfile"
  dst  = "/etc/caddy/Caddyfile"
}
