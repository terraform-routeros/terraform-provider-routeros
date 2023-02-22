## [1.0.13](https://github.com/GNewbury1/terraform-provider-routeros/compare/v1.0.12...v1.0.13) (2023-02-22)


### Bug Fixes

* **#122:** Add missing fields to interface list member ([3debe19](https://github.com/GNewbury1/terraform-provider-routeros/commit/3debe192f123463e85a00f0318f92f7996d06906)), closes [#122](https://github.com/GNewbury1/terraform-provider-routeros/issues/122)

## [1.0.12](https://github.com/GNewbury1/terraform-provider-routeros/compare/v1.0.11...v1.0.12) (2023-02-22)


### Bug Fixes

* **#106:** "root_path_cost" not found. ([4d568f5](https://github.com/GNewbury1/terraform-provider-routeros/commit/4d568f54db78297f3b71f4b1403f65214585c4ac)), closes [#106](https://github.com/GNewbury1/terraform-provider-routeros/issues/106)

## [1.0.11](https://github.com/GNewbury1/terraform-provider-routeros/compare/v1.0.10...v1.0.11) (2023-02-20)


### Bug Fixes

* **#106:** Fix internal validation (for release). ([fa3bb93](https://github.com/GNewbury1/terraform-provider-routeros/commit/fa3bb93f22b22dbb6c50296cf6b12f030920f8d6))

## [1.0.9](https://github.com/GNewbury1/terraform-provider-routeros/compare/v1.0.8...v1.0.9) (2023-02-20)


### Bug Fixes

* **#106:** Added "multicast_router" field. ([128efe1](https://github.com/GNewbury1/terraform-provider-routeros/commit/128efe12b91fd2bf6a16536bd618047d4d6200b8)), closes [#106](https://github.com/GNewbury1/terraform-provider-routeros/issues/106)
* **#110:** "host_name" set to Computed. ([d8c80dc](https://github.com/GNewbury1/terraform-provider-routeros/commit/d8c80dcc50dcee15f88aaa4a54574c4f5889a856)), closes [#110](https://github.com/GNewbury1/terraform-provider-routeros/issues/110)

## [1.0.8](https://github.com/GNewbury1/terraform-provider-routeros/compare/v1.0.7...v1.0.8) (2023-02-19)


### Bug Fixes

* **#110:** Typo in hostname field for dhcp lease ([1113a36](https://github.com/GNewbury1/terraform-provider-routeros/commit/1113a3641e44fed551245415f387f1db34690d52)), closes [#110](https://github.com/GNewbury1/terraform-provider-routeros/issues/110)

## [1.0.7](https://github.com/GNewbury1/terraform-provider-routeros/compare/v1.0.6...v1.0.7) (2023-02-17)


### Bug Fixes

* [#106](https://github.com/GNewbury1/terraform-provider-routeros/issues/106) /interface/bridge resource schema (IGMP snooping). ([dd9baaa](https://github.com/GNewbury1/terraform-provider-routeros/commit/dd9baaa9f81571c596c73f649919b6f475f6a327))
* [#109](https://github.com/GNewbury1/terraform-provider-routeros/issues/109) /interface/bridge/port resource schema (STP). ([182b067](https://github.com/GNewbury1/terraform-provider-routeros/commit/182b0679e0f9ec9996aff2cce170906c7cf5bf51))
* /ip/dhcp-server/lease resource schema. ([68a67d4](https://github.com/GNewbury1/terraform-provider-routeros/commit/68a67d48763e7bc8f7f5e9bd141044c5782637f4))

## [1.0.7](https://github.com/GNewbury1/terraform-provider-routeros/compare/v1.0.6...v1.0.7) (2023-02-17)


### Bug Fixes

* [#106](https://github.com/GNewbury1/terraform-provider-routeros/issues/106) /interface/bridge resource schema (IGMP snooping). ([dd9baaa](https://github.com/GNewbury1/terraform-provider-routeros/commit/dd9baaa9f81571c596c73f649919b6f475f6a327))
* [#109](https://github.com/GNewbury1/terraform-provider-routeros/issues/109) /interface/bridge/port resource schema (STP). ([182b067](https://github.com/GNewbury1/terraform-provider-routeros/commit/182b0679e0f9ec9996aff2cce170906c7cf5bf51))
* /ip/dhcp-server/lease resource schema. ([68a67d4](https://github.com/GNewbury1/terraform-provider-routeros/commit/68a67d48763e7bc8f7f5e9bd141044c5782637f4))

## [1.0.6](https://github.com/GNewbury1/terraform-provider-routeros/compare/v1.0.5...v1.0.6) (2023-02-17)


### Bug Fixes

* **#110:** Add missing fields to DhcpServerLease ([100af8f](https://github.com/GNewbury1/terraform-provider-routeros/commit/100af8f7da96ff38a879536f1894118fd9bc858d)), closes [#110](https://github.com/GNewbury1/terraform-provider-routeros/issues/110)

## [1.0.6](https://github.com/GNewbury1/terraform-provider-routeros/compare/v1.0.5...v1.0.6) (2023-02-17)


### Bug Fixes

* **#110:** Add missing fields to DhcpServerLease ([100af8f](https://github.com/GNewbury1/terraform-provider-routeros/commit/100af8f7da96ff38a879536f1894118fd9bc858d)), closes [#110](https://github.com/GNewbury1/terraform-provider-routeros/issues/110)

## [1.0.5](https://github.com/GNewbury1/terraform-provider-routeros/compare/v1.0.4...v1.0.5) (2023-02-17)


### Bug Fixes

* Spaces in resource names ([#102](https://github.com/GNewbury1/terraform-provider-routeros/issues/102) - [#104](https://github.com/GNewbury1/terraform-provider-routeros/issues/104)). ([6dafa4b](https://github.com/GNewbury1/terraform-provider-routeros/commit/6dafa4bd26ea406c5f4e481f201da1f16dd9b747))

## [1.0.4](https://github.com/GNewbury1/terraform-provider-routeros/compare/v1.0.3...v1.0.4) (2023-02-16)


### Bug Fixes

* Add gpg fingerprint to CI ([b315e13](https://github.com/GNewbury1/terraform-provider-routeros/commit/b315e130d338de82a1347c8f91cd4ba442d8d7c3))

## [1.0.2](https://github.com/GNewbury1/terraform-provider-routeros/compare/v1.0.1...v1.0.2) (2023-02-15)


### Bug Fixes

* Create multiple names for the same resource to aid compatibility ([5ed67a7](https://github.com/GNewbury1/terraform-provider-routeros/commit/5ed67a7e78cbf167320e2092c7c276e7410041bd))
* Interface child items had incorrect reference ([be14cb6](https://github.com/GNewbury1/terraform-provider-routeros/commit/be14cb6a2feec52cf5c34cc51924c1df09c90023))

## [1.0.1](https://github.com/GNewbury1/terraform-provider-routeros/compare/v1.0.0...v1.0.1) (2023-02-14)


### Bug Fixes

* IP validation fix ([12c1a23](https://github.com/GNewbury1/terraform-provider-routeros/commit/12c1a230aac5636b80636bd060c4024167d67f64))
