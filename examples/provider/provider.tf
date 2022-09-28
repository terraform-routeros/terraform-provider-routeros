terraform {
  required_providers {
    routeros = {
      source = "github.com/vaerh/routeros"
    }
  }
}

provider "routeros" {
  hosturl        = "https://router.local"         # env ROS_HOSTURL or MIKROTIK_HOST
  username       = "admin"                        # env ROS_USERNAME or MIKROTIK_USER
  password       = ""                             # env ROS_PASSWORD or MIKROTIK_PASSWORD
  ca_certificate = "/path/to/ca/certificate.pem"  # env ROS_CA_CERTIFICATE or MIKROTIK_CA_CERTIFICATE
  insecure       = true                           # env ROS_INSECURE or MIKROTIK_INSECURE
}

resource "routeros_gre" "gre_hq" {
  name      = "gre-hq-1"
  remote_address = "10.77.3.26"
  disabled  = true
}