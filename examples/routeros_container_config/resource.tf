resource "routeros_container_config" "config" {
  registry_url = "https://registry-1.docker.io"
  ram_high     = "0"
  tmpdir       = "/usb1-part1/containers/tmp"
  layer_dir    = "/usb1-part1/containers/layers"
}
