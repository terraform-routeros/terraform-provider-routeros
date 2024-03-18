resource "routeros_container" "busybox" {
  remote_image  = "library/busybox:1.35.0"
  cmd           = "/bin/httpd -f -p 8080"
  interface     = routeros_interface_veth.busybox.name
  logging       = true
  root_dir      = "/usb1-part1/containers/busybox/root"
  start_on_boot = true
}
