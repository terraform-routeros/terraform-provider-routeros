# Terraform Provider RouterOS

## Purpose

This provider is intended to be used with Router OS 7 and above. It makes use of the REST API introduced with Router OS v7.

## Usage

```terraform
provider "routeros" {
    hosturl  = "https://my.router.lan"
    username = "user"
    password = "password
}
```

## Changelog

See [changelog.md](changelog.md)