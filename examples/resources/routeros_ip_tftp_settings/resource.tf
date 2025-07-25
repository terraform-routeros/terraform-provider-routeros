resource "routeros_ip_tftp" "file" {
  ip_addresses  = ["10.0.0.0/24"]
  req_filename  = "file.txt"
  real_filename = routeros_file.file.name
  read_only     = true
}