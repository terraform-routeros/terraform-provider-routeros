# Terraform Provider RouterOS

## Purpose

This provider is intended to be used with Router OS 7 and above. It makes use of the REST API introduced with Router OS v7.

### Todo for 0.1.0 (initial pre-release)

#### Resources:
- [ ] Add resources for following objects normally found under `/ip`
    - [x] IP addresses
    - [ ] DHCP client
    - [ ] DHCP server
    - [ ] Pools
    - [ ] Routes
- [x] Add resources for the following objects normally found under `/interface`
    - [x] VLANs
    - [ ] Bridge
        - [ ] Ports
        - [ ] VLAN

#### Data:
- [ ] Add interface list as data object
- [ ] Add route list as data object
- [ ] Add IP address list as data object

## Usage

Please refer to the [documentation](docs/)

## Changelog

See [changelog.md](changelog.md)