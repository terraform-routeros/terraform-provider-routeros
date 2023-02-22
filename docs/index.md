---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "RouterOS Provider"
subcategory: ""
description: |-
    A provider to integrate with the REST API introduced in RouterOS v7
  
---

# RouterOS Provider

To get started with the provider, you first need to enable the REST API on your router. [You can follow the Mikrotik documentation on this](https://help.mikrotik.com/docs/display/ROS/REST+API), but the gist is to create an SSL cert (in `/system/certificates`) and enable the `web-ssl` service (in `/ip/services`) which uses that certificate.


## Example Usage

```terraform
terraform {
  required_providers {
    routeros = {
      source = "GNewbury1/routeros"
    }
  }
}

provider "routeros" {
  hosturl        = "https://router.local"        # env ROS_HOSTURL or MIKROTIK_HOST
  username       = "admin"                       # env ROS_USERNAME or MIKROTIK_USER
  password       = ""                            # env ROS_PASSWORD or MIKROTIK_PASSWORD
  ca_certificate = "/path/to/ca/certificate.pem" # env ROS_CA_CERTIFICATE or MIKROTIK_CA_CERTIFICATE
  insecure       = true                          # env ROS_INSECURE or MIKROTIK_INSECURE
}

resource "routeros_interface_gre" "gre_hq" {
  name           = "gre-hq-1"
  remote_address = "10.77.3.26"
  disabled       = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `hosturl` (String) URL of the ROS router. Include the scheme (http/https)

### Optional

- `insecure` (Boolean) Whether to verify the SSL certificate or not
- `password` (String, Sensitive) Password for the ROS user
- `username` (String) Username for the ROS user