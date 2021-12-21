# Changelog

## Pre-release

### 0.3.0 (Upcoming)
#### Features:
- Added ability to manage Wireguard interfaces
- Added ability to manage Wireguard peers

#### Misc:
- Updated documentation and readme
- Added Github Actions to test code in pull requests to `main` and to branches matching `release/*`
    - CI tests check ROS versions `7.1`, `7.1.1`, and `7.2rc1`
- Added Github Actions to publish a release

### 0.2.1 (2021-12-21)
#### Bug Fix
- Fixed a bug where the Arp of a `resource.routeros_interface_bridge` was read as a `bool` rather than a `string` in the Get function. This caused Terraform to throw an error whenever it needed to read a `routeros_interface_bridge` resource.

### 0.2.0 (2021-12-18)
- Added ability to manage /ip/firewall/filter [(Issue #16)](https://github.com/GNewbury1/terraform-provider-routeros/issues/16)

### 0.1.1 (2021-12-13)
#### Bug fixes:
- Fixed IP Address attributes not setting correctly

### 0.1.0 (2021-12-13)
#### Features:
- Initial Release
