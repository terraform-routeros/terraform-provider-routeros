# Terraform Provider RouterOS

![module testing workflow](https://github.com/GNewbury1/terraform-provider-routeros/actions/workflows/release.yml/badge.svg?branch=main)

**Note**: Release `1.35.1` does not contain artifacts. Please do not use `v1.35.1`.

## Purpose

This provider allows you to configure Mikrotik routers using [old API](https://help.mikrotik.com/docs/display/ROS/API) or [REST API](https://help.mikrotik.com/docs/display/ROS/REST+API), using or not using TLS.
Compatibility testing is only performed within ROS version 7.x.

From version 1.0.0, the provider has been rewritten by [vaerh](https://github.com/vaerh), and their [fork](https://github.com/vaerh/terraform-provider-routeros) has now been merged. This version drastically improves adding new endpoints to the provider, enabling significantly easier development. [vaerh](https://github.com/vaerh) has been added as a maintainer to this project.

_We are not affiliated in any way with Mikrotik or the development of RouterOS_
## Using the provider

To get started with the provider, you first need to enable the REST API on your router. [You can follow the Mikrotik documentation on this](https://help.mikrotik.com/docs/display/ROS/REST+API), but the gist is to create an SSL cert (in `/system/certificates`) and enable the `web-ssl` service (in `/ip/services`) which uses that certificate. After that, include the following in your Terraform manifests:

```terraform
terraform {
  required_providers {
    routeros = {
      source = "terraform-routeros/routeros"
    }
  }
}

provider "routeros" {
  hosturl  = "(https|api|apis)://my.router.local[:port]"
  username = "my_username"
  password = "my_super_secret_password"
}

```

For more in-depth documentation about each of the resources and datasources, please read the [documentation on Hashicorp's Provider registry](https://registry.terraform.io/providers/terraform-routeros/routeros/latest/docs)

### Versions tested

- go 1.21 and ROS 7.7, 7.8, 7.9 (stable)

## Changelog

For a detailed changelog, please see the [changelog.md](CHANGELOG.md).

## Contributing
This version of the module greatly simplifies the process of adding new resources.
You are welcome!
