# Terraform Provider RouterOS

![module testing workflow](https://github.com/GNewbury1/terraform-provider-routeros/actions/workflows/module_testing.yml/badge.svg?branch=main)

## Purpose

This provider allows you to configure Mikrotik routers using [old API](https://help.mikrotik.com/docs/display/ROS/API) 
or [REST API](https://help.mikrotik.com/docs/display/ROS/REST+API), using or not using TLS.

Relative to the [original repository](https://github.com/GNewbury1/terraform-provider-routeros), 
the project structure has been redesigned and support for the old API has been added.

## Using the provider

To get started with the provider, you first need to enable the REST API on your router. [You can follow the Mikrotik documentation on this](https://help.mikrotik.com/docs/display/ROS/REST+API), but the gist is to create an SSL cert (in `/system/certificates`) andenable the `web-ssl` service (in `/ip/services`) which uses that certificate. After that, include the following in your Terraform manifests:

```terraform
terraform {
  required_providers {
    routeros = {
      source = "GNewbury1/routeros"
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
[documentation on Hashicorp's Provider registry](https://registry.terraform.io/providers/GNewbury1/routeros/latest/docs)

### Versions tested

- go 1.19 and ROS 7.5, 7.6, 7.7 (stable)

## Changelog

For a detailed changelog, please see the [changelog.md](CHANGELOG.md).
For a roadmap, please see [the github project](https://github.com/users/GNewbury1/projects/2) I created to cover this. To get things added to the roadmap (such as bugs or features), please raise an issue and describe what needs to be fixed/implemented. As you can imagine, I can only test in the environment I have, which may not work in the environment you have.

## Contributing
This version of the module greatly simplifies the process of adding new resources.
You are welcome!
