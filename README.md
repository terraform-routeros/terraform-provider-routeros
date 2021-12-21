# Terraform Provider RouterOS

![client testing workflow](https://github.com/gnewbury1/terraform-provider-routeros/actions/workflows/client_tests.yml/badge.svg?branch=main)

![provider testing workflow](https://github.com/gnewbury1/terraform-provider-routeros/actions/workflows/provider_tests.yml/badge.svg?branch=main)

## Purpose

This repository implements two things:
- Go client to use with RouterOS v7 REST API
- Terraform provider which makes use of said go client

This is to allow configuration of Mikrotik's RouterOS at scale using Terraform. Please note that this provider is still in development, so functionality is missing in some areas. Have a look at [the documentation]([documentation on Hashicorp's Provider registry](https://registry.terraform.io/providers/GNewbury1/routeros/latest/docs)) for functionality that has been implemented. If there is a particular function you'd like to see, please raise an issue to it can be added to the [roadmap](https://github.com/users/GNewbury1/projects/2).

## Using the provider

To get started with the provider, include the following in your Terraform manifests:

```terraform
terraform {
  required_providers {
    routeros = {
      source = "GNewbury1/routeros"
    }
  }
}

provider "routeros" {
  hosturl  = "https://my.router.local"
  username = "my_username"
  password = "my_super_secret_password"
}

```

For more in-depth documentation about each of the resources and datasources, please read the [documentation on Hashicorp's Provider registry](https://registry.terraform.io/providers/GNewbury1/routeros/latest/docs)

## Changelog and Roadmap

For a detailed changelog, please see the [changelog.md](changelog.md).
For a roadmap, please see [the github project](https://github.com/users/GNewbury1/projects/2) I created to cover this. To get things added to the roadmap (such as bugs or features), please raise an issue and describe what needs to be fixed/implemented. As you can imagine, I can only test in the environment I have, which may not work in the environment you have.

## Contributing

Please do raise a pull request if you have a contribution to make. Any and all contributions are welcome!