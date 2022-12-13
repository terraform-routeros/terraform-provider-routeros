# Terraform Provider RouterOS

![module testing workflow](https://github.com/vaerh/terraform-provider-routeros/actions/workflows/module_testing.yml/badge.svg?branch=main)

## Purpose

This provider allows you to configure Mikrotik routers using [old API](https://help.mikrotik.com/docs/display/ROS/API) 
or [REST API](https://help.mikrotik.com/docs/display/ROS/REST+API), using or not using TLS.

Relative to the [original repository](https://github.com/GNewbury1/terraform-provider-routeros), 
the project structure has been redesigned and support for the old API has been added.

## Using the provider

To get started with the provider, include the following in your Terraform manifests:

```terraform
terraform {
  required_providers {
    routeros = {
      source = "vaerh/routeros"
    }
  }
}

provider "routeros" {
  hosturl  = "(https|api|apis)://my.router.local[:port]"
  username = "my_username"
  password = "my_super_secret_password"
}

```

For more in-depth documentation about each of the resources and datasources, please read the 
[documentation on Hashicorp's Provider registry](https://registry.terraform.io/providers/vaerh/routeros/latest/docs)

### Tested RouterOS versions

- go 1.19 and ROS 7.5 & 7.6 (stable)

## Changelog

For a detailed changelog, please see the [changelog](CHANGELOG.md).

## Contributing
This version of the module greatly simplifies the process of adding new resources.
You are welcome!
