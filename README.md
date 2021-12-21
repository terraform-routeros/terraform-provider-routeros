# Terraform Provider RouterOS

![client testing workflow](https://github.com/gnewbury1/terraform-provider-routeros/actions/workflows/client_tests.yml/badge.svg?branch=main)

![provider testing workflow](https://github.com/gnewbury1/terraform-provider-routeros/actions/workflows/provider_tests.yml/badge.svg?branch=main)

## Purpose

This provider is intended to be used with Router OS 7 and above. It makes use of the REST API introduced with Router OS v7.

## Currently Implemented

- Interface
    - Bridge
        - Bridge VLAN
        - Bridge Port
    - VLAN
- IP
    - DHCP Client
    - DHCP Server
    - Firewall
        - Filter
    - Pool
    - Route

## Usage

Please refer to the [documentation](docs/)

## Changelog

See [changelog.md](changelog.md)