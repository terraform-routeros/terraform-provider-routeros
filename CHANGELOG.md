## [1.46.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.46.0...v1.46.1) (2024-04-21)


### Bug Fixes

* **nat:** [Backward compatibility] resource "routeros_ip_firewall_nat" ([ee8f992](https://github.com/terraform-routeros/terraform-provider-routeros/commit/ee8f992098efb9bfdc636b81a4912b302046c839)), closes [#431](https://github.com/terraform-routeros/terraform-provider-routeros/issues/431)
* **wg_peer:** Invalid syntax for 'client_keepalive' field ([df24de2](https://github.com/terraform-routeros/terraform-provider-routeros/commit/df24de28b406a1e431056266f6d5b429447d9f62)), closes [#432](https://github.com/terraform-routeros/terraform-provider-routeros/issues/432)

## [1.46.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.45.0...v1.46.0) (2024-04-16)


### Features

* add routeros_certificate_scep_server resource ([#420](https://github.com/terraform-routeros/terraform-provider-routeros/issues/420)) ([b80b52d](https://github.com/terraform-routeros/terraform-provider-routeros/commit/b80b52d9ef5f8ce8d4b1396092ba1dc816765489))


### Bug Fixes

* **no-release:** [Backward compatibility] routeros_interface_bridge NOT WORKING as expected ([dd46e95](https://github.com/terraform-routeros/terraform-provider-routeros/commit/dd46e95ac82b18d3ff2f7ef5a2c1cba30dc6f2f1)), closes [#417](https://github.com/terraform-routeros/terraform-provider-routeros/issues/417)
* **no-release:** [Backward compatibility] routeros_interface_bridge_port ([dca44bc](https://github.com/terraform-routeros/terraform-provider-routeros/commit/dca44bc1c23246fc63d2310f8f2a7d9f15ec5519)), closes [#419](https://github.com/terraform-routeros/terraform-provider-routeros/issues/419)
* **no-release:** [Backward compatibility] routeros_ip_firewall_filter NOT WORKING AS EXPECTED ([6cc2072](https://github.com/terraform-routeros/terraform-provider-routeros/commit/6cc2072110ca5d67501a65ce5cd69cbb25e4fe12)), closes [#418](https://github.com/terraform-routeros/terraform-provider-routeros/issues/418)
* **no-release:** Field 'tx_carrier_sense_error' not found in the schema ([c3d2eb2](https://github.com/terraform-routeros/terraform-provider-routeros/commit/c3d2eb2343e6593f673d57e2bbe80bd32fc863a9)), closes [#416](https://github.com/terraform-routeros/terraform-provider-routeros/issues/416)
* **no-release:** Fix comparison of time & hex not provided by user ([4f85ccf](https://github.com/terraform-routeros/terraform-provider-routeros/commit/4f85ccf3303e6b91e6c0aee6dc0b4602d04604a3))

## [1.45.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.44.3...v1.45.0) (2024-04-09)


### Features

* **dhcp-relay:** Add DHCP Relay support ([6eb5901](https://github.com/terraform-routeros/terraform-provider-routeros/commit/6eb590183e2c67dc5f797bbfbe545d502f39a04c)), closes [#413](https://github.com/terraform-routeros/terraform-provider-routeros/issues/413)
* **upnp:** Add UPNP & UPNP Interfaces ([8fd3da6](https://github.com/terraform-routeros/terraform-provider-routeros/commit/8fd3da66a671e1e3e6e62a38749fa0b58f3c1595)), closes [#412](https://github.com/terraform-routeros/terraform-provider-routeros/issues/412)


### Bug Fixes

* Export/import certificates ([205a221](https://github.com/terraform-routeros/terraform-provider-routeros/commit/205a22110a712a123a6225ae0f71d9f763043c02)), closes [#404](https://github.com/terraform-routeros/terraform-provider-routeros/issues/404)
* **upnp:** Fix "unknown parameter forced-external-ip" ([74333f2](https://github.com/terraform-routeros/terraform-provider-routeros/commit/74333f2ec7eeba08ff018084cf75003d102fac32))

## [1.44.3](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.44.2...v1.44.3) (2024-04-09)


### Bug Fixes

* **Importer:** Add extended resource import function ([713d9ee](https://github.com/terraform-routeros/terraform-provider-routeros/commit/713d9ee6318b0f71dea33207687b5062af605859)), closes [#403](https://github.com/terraform-routeros/terraform-provider-routeros/issues/403)

## [1.44.2](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.44.1...v1.44.2) (2024-04-04)


### Bug Fixes

* Fix for release ([3c54e07](https://github.com/terraform-routeros/terraform-provider-routeros/commit/3c54e076d86ef551223c1d83b5fa72a0c5742db8))

## [1.44.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.44.0...v1.44.1) (2024-04-04)


### Bug Fixes

* routeros_ip_service  not working ([c896837](https://github.com/terraform-routeros/terraform-provider-routeros/commit/c8968373e42e60bc1d0d53d23031ac8231086d7f)), closes [#407](https://github.com/terraform-routeros/terraform-provider-routeros/issues/407)

## [1.44.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.43.0...v1.44.0) (2024-04-04)


### Features

* Add `routeros_capsman_interface` resource to manage CAPsMAN interfaces ([65b9717](https://github.com/terraform-routeros/terraform-provider-routeros/commit/65b9717b8afa7d218012f613af88c53fe0a21063))


### Bug Fixes

* Fix inline configuration properties not to update when untouched in `routeros_capsman_configuration` resources ([e2a5b55](https://github.com/terraform-routeros/terraform-provider-routeros/commit/e2a5b5527a77a99a84f35e8381a59add6ab63d11))

## [1.43.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.42.0...v1.43.0) (2024-04-01)


### Features

*  Add NTP client resource ([0a678cf](https://github.com/terraform-routeros/terraform-provider-routeros/commit/0a678cf09c3ff350857c89c7d86cfc8a6de7cf00)), closes [#386](https://github.com/terraform-routeros/terraform-provider-routeros/issues/386)
* Add IP neighbor discovery resource ([706bccb](https://github.com/terraform-routeros/terraform-provider-routeros/commit/706bccbec25c25f5be162f941664954769fe4cfd)), closes [#388](https://github.com/terraform-routeros/terraform-provider-routeros/issues/388)
* Add SSH Server resource ([99883e4](https://github.com/terraform-routeros/terraform-provider-routeros/commit/99883e49dcd05ad9eca46c2042c2996c1225c938)), closes [#389](https://github.com/terraform-routeros/terraform-provider-routeros/issues/389)
* Add tool/mac-server support ([13b565c](https://github.com/terraform-routeros/terraform-provider-routeros/commit/13b565c40f3178e57b475bdd85880d52c7f5eb3b)), closes [#387](https://github.com/terraform-routeros/terraform-provider-routeros/issues/387)


### Bug Fixes

* Fix `isEmpty` function for Lists with ObjectType ([dbbb710](https://github.com/terraform-routeros/terraform-provider-routeros/commit/dbbb710f1ef97e3fed2ef7477e73a78e91104901))
* Set "AlwaysPresentNotUserProvided" for PropVrfRw ([fd58ea1](https://github.com/terraform-routeros/terraform-provider-routeros/commit/fd58ea19ae18b6fd2fbd1d1945b45c0936f1f0b9))
* **v7.12.2:** Change `routeros_interface_wireguard_peer` schema ([3f937c1](https://github.com/terraform-routeros/terraform-provider-routeros/commit/3f937c1b44d4661212b8aad93ebef326d52dbb55))
* **v7.12.2:** Change `routeros_ipv6_neighbor_discovery` schema ([598319a](https://github.com/terraform-routeros/terraform-provider-routeros/commit/598319adc29a02a34f6e4c52a87a979cc36a4a0b))
* **v7.12.2:** Change `routeros_routing_bgp_connection` schema ([205e43d](https://github.com/terraform-routeros/terraform-provider-routeros/commit/205e43dac4600fdf2e632380f9cd687c55cc06b4))

## [1.42.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.41.0...v1.42.0) (2024-03-29)


### Features

* Add `http` scheme support for the REST API ([c326052](https://github.com/terraform-routeros/terraform-provider-routeros/commit/c3260525afad9dad0e92b37ed6bb409a13e32d2e))

## [1.41.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.40.0...v1.41.0) (2024-03-27)


### Features

* Add function to create skip metadata based on slices ([ac54292](https://github.com/terraform-routeros/terraform-provider-routeros/commit/ac54292f53b6384efa5b7e4f2e1ce0761ce47b3e))


### Bug Fixes

* Fix fields in new commits ([205919f](https://github.com/terraform-routeros/terraform-provider-routeros/commit/205919f49fdb522357db562def3518194c3b81e4))
* Simplify the procedure for generating field conversion lists. ([a67b772](https://github.com/terraform-routeros/terraform-provider-routeros/commit/a67b772da55e9590390b92eb0e16b83041977058))

## [1.40.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.39.0...v1.40.0) (2024-03-26)


### Features

* Add `routeros_zerotier_controller` resource to manage ZeroTier controller ([68cb99a](https://github.com/terraform-routeros/terraform-provider-routeros/commit/68cb99a350261e5350f5bba987b5a42e8b6a637c))
* Add `routeros_zerotier_interface` resource to manage ZeroTier interface ([481709c](https://github.com/terraform-routeros/terraform-provider-routeros/commit/481709cee5b3a4fe3122173909e3ae19e43ccc18))
* Add `routeros_zerotier` resource to manage ZeroTier instances ([f182f0c](https://github.com/terraform-routeros/terraform-provider-routeros/commit/f182f0c93ab0ada397db2305eacd868e4a9b837a))

## [1.39.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.38.0...v1.39.0) (2024-03-26)


### Features

* Add `routeros_interface_wireless_cap` resource to manage CAPsMAN client ([b1558f9](https://github.com/terraform-routeros/terraform-provider-routeros/commit/b1558f92722ab69ccfe9cba86dc509f5ca591819))
* Add `routeros_ip_cloud_advanced` resource to manage advanced cloud settings ([405827d](https://github.com/terraform-routeros/terraform-provider-routeros/commit/405827d0d843431abe71fb303a65105ae29e86b0))
* Add `routeros_ppp_aaa` resource to manage authentication and accounting ([87f06c0](https://github.com/terraform-routeros/terraform-provider-routeros/commit/87f06c01c2a5b1b33bada1e4d6c69bcc559a38d3))

## [1.38.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.37.2...v1.38.0) (2024-03-23)


### Features

* Add support for interface macvlan ([24d940b](https://github.com/terraform-routeros/terraform-provider-routeros/commit/24d940b53e9b87487a19f7e7c44f7fe3a27e3c3f))


### Bug Fixes

* **no-release:** Resolve warnings due to some missing properties ([#381](https://github.com/terraform-routeros/terraform-provider-routeros/issues/381)) ([864ed27](https://github.com/terraform-routeros/terraform-provider-routeros/commit/864ed278e304e6789cb4401d4ef8ea33ca2e7bd2))

## [1.37.2](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.37.1...v1.37.2) (2024-03-18)


### Bug Fixes

* **dhcp-server:** Add a missing DHCP server option attribute ([c375754](https://github.com/terraform-routeros/terraform-provider-routeros/commit/c375754f0651196cbdb3ec986149d60f16fa6847)), closes [#376](https://github.com/terraform-routeros/terraform-provider-routeros/issues/376)

## [1.37.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.37.0...v1.37.1) (2024-03-18)


### Bug Fixes

* Enable importing ethernet interfaces ([#379](https://github.com/terraform-routeros/terraform-provider-routeros/issues/379)) ([3676f3f](https://github.com/terraform-routeros/terraform-provider-routeros/commit/3676f3f210acaa51e6f0632369d8ba5f35f56b5e))

## [1.37.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.36.1...v1.37.0) (2024-03-18)


### Features

* Add `routeros_wifi` resource to manage WiFi interfaces ([5f234c4](https://github.com/terraform-routeros/terraform-provider-routeros/commit/5f234c456ca8c790e214d5b3e2dd8a8b1ab69b88))


### Bug Fixes

* Add reusable L2MTU property ([93a7495](https://github.com/terraform-routeros/terraform-provider-routeros/commit/93a749559aaf7df3235c8d9512cc7c5a1e417cea))
* Fix the `routeros_wifi_configuration` resource to suppress pristine inline parameters ([77d807c](https://github.com/terraform-routeros/terraform-provider-routeros/commit/77d807c2fd1e8bfe8d67de011c2b3b4851477d41))
* Refactor `AlwaysPresentNotUserProvided` helper to self-contain empty value check ([669fd68](https://github.com/terraform-routeros/terraform-provider-routeros/commit/669fd689008aba4673205fcfca1252c5b1a7f795))
* Refactor `AlwaysPresentNotUserProvided` helper to support map type ([038fe7c](https://github.com/terraform-routeros/terraform-provider-routeros/commit/038fe7c192ba87c46f1ed13334183f33d1cc6717))

## [1.36.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.36.0...v1.36.1) (2024-03-18)


### Bug Fixes

* stop using Int32MaxValue as upper bound for validation ([fcae1bd](https://github.com/terraform-routeros/terraform-provider-routeros/commit/fcae1bdfb8ef0b3622133eab4b04822b122ed405))

## [1.36.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.35.1...v1.36.0) (2024-03-17)


### Features

* Add routeros_system_script ([bf72b32](https://github.com/terraform-routeros/terraform-provider-routeros/commit/bf72b327d25605216ba8a5a44cb2aed43775dd6f)), closes [#373](https://github.com/terraform-routeros/terraform-provider-routeros/issues/373)

## [1.35.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.35.0...v1.35.1) (2024-02-28)


### Bug Fixes

* **no-release:** Fix max value for MTU - Fix Main ([#365](https://github.com/terraform-routeros/terraform-provider-routeros/issues/365)) ([b2c57de](https://github.com/terraform-routeros/terraform-provider-routeros/commit/b2c57de6967d92dfc1dd445566f711544913483e))
* Undo commit "Change l2mtu property " ([#367](https://github.com/terraform-routeros/terraform-provider-routeros/issues/367)) ([8657054](https://github.com/terraform-routeros/terraform-provider-routeros/commit/86570540ecf5a3cdcd1a7391074e6626bbef8bc0)), closes [#326](https://github.com/terraform-routeros/terraform-provider-routeros/issues/326)

## [1.35.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.34.0...v1.35.0) (2024-02-22)


### Features

* Add resource routeros_ipv6_neighbor_discovery ([#362](https://github.com/terraform-routeros/terraform-provider-routeros/issues/362)) ([13fb7b7](https://github.com/terraform-routeros/terraform-provider-routeros/commit/13fb7b75ba398b97551b58ff7cd8f99e26a15e12))


### Bug Fixes

* **no-release:** Add gateway6 field to /interface/veth ([#358](https://github.com/terraform-routeros/terraform-provider-routeros/issues/358)) ([b0385f6](https://github.com/terraform-routeros/terraform-provider-routeros/commit/b0385f66815ac1205606d64c3df80fcd9b315d08))

## [1.34.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.33.1...v1.34.0) (2024-02-15)


### Features

* **ipv6:** Add support for /ipv6/dhcp-client/option ([df43330](https://github.com/terraform-routeros/terraform-provider-routeros/commit/df43330ca7bed7e15d8a7f18d76adcacd3289fc0))

## [1.33.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.33.0...v1.33.1) (2024-02-14)


### Features

* **no-release:** Add support for /container ([#356](https://github.com/terraform-routeros/terraform-provider-routeros/issues/356)) ([afacc8f](https://github.com/terraform-routeros/terraform-provider-routeros/commit/afacc8fa8d382cf1db23336d448f141f8cabea56))


### Bug Fixes

* **ipv6:** Add fields for /ipv6/dhcp-client ([#357](https://github.com/terraform-routeros/terraform-provider-routeros/issues/357)) ([17a26d9](https://github.com/terraform-routeros/terraform-provider-routeros/commit/17a26d9a5c2f6a92415da688a64074f7bf97e9ce))
* **no-release:** Add missing ethernet interface fields to SkipFields ([#359](https://github.com/terraform-routeros/terraform-provider-routeros/issues/359)) ([55409c4](https://github.com/terraform-routeros/terraform-provider-routeros/commit/55409c4f762d2664da7b5cef7169e36013e2f203))

## [1.33.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.32.6...v1.33.0) (2024-02-13)


### Features

* **file:** Add support for /file ([#355](https://github.com/terraform-routeros/terraform-provider-routeros/issues/355)) ([2a1d9fd](https://github.com/terraform-routeros/terraform-provider-routeros/commit/2a1d9fd7f866b427fcbff4589b6b19d42c9267a1))


### Bug Fixes

* **no-release:** Change the Name property for the ipip resource ([4899cf1](https://github.com/terraform-routeros/terraform-provider-routeros/commit/4899cf18fae8dd85a3df64ac7b75d0d7752f4c52))

## [1.32.6](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.32.5...v1.32.6) (2024-02-06)


### Features

* **no-release:** add interface/ipip ([#346](https://github.com/terraform-routeros/terraform-provider-routeros/issues/346)) ([e7bd8dd](https://github.com/terraform-routeros/terraform-provider-routeros/commit/e7bd8ddaaf66ad101bd901d1c9045d52f69ab45b))


### Bug Fixes

* dhcp client add script string parameter ([#348](https://github.com/terraform-routeros/terraform-provider-routeros/issues/348)) ([3df2b69](https://github.com/terraform-routeros/terraform-provider-routeros/commit/3df2b69b155368383af121ac51a86f0d7b9a2896))
* **no-release:** Mandatory use of ID as ipip resource identifier ([55f8a62](https://github.com/terraform-routeros/terraform-provider-routeros/commit/55f8a629297512a8d6397de6786b9b6fb3320a6c)), closes [#346](https://github.com/terraform-routeros/terraform-provider-routeros/issues/346)

## [1.32.5](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.32.4...v1.32.5) (2024-01-25)


### Bug Fixes

* Change l2mtu property ([9260354](https://github.com/terraform-routeros/terraform-provider-routeros/commit/92603541fd6b3f11dcbfd1ff2c1904f6687062a9)), closes [#326](https://github.com/terraform-routeros/terraform-provider-routeros/issues/326)
* **cloud:** Add missing field ([235e3b0](https://github.com/terraform-routeros/terraform-provider-routeros/commit/235e3b07de33ca0b1edb82f728c0e3d241256a26)), closes [#331](https://github.com/terraform-routeros/terraform-provider-routeros/issues/331)
* **no-release:** Fix Delete. Not an system object ([#343](https://github.com/terraform-routeros/terraform-provider-routeros/issues/343)) ([89ee0ea](https://github.com/terraform-routeros/terraform-provider-routeros/commit/89ee0ea2934f923dbf99744bd7d455bbf95910d2))
* **no-release:** Fix Delete. Not an Systemresource ([#342](https://github.com/terraform-routeros/terraform-provider-routeros/issues/342)) ([e560128](https://github.com/terraform-routeros/terraform-provider-routeros/commit/e560128554dd157b8b9a3399724baa2e5def521e))
* **wg peer:** Add missing fields ([48018c4](https://github.com/terraform-routeros/terraform-provider-routeros/commit/48018c4d525319b48e77923d57b4816ceb53cd30)), closes [#332](https://github.com/terraform-routeros/terraform-provider-routeros/issues/332)

## [1.32.4](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.32.3...v1.32.4) (2024-01-24)


### Bug Fixes

* Removed all default values ([546764e](https://github.com/terraform-routeros/terraform-provider-routeros/commit/546764e1a14795936cd6890aad6530942b104c3e)), closes [#326](https://github.com/terraform-routeros/terraform-provider-routeros/issues/326)

## [1.32.3](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.32.2...v1.32.3) (2024-01-19)


### Bug Fixes

* **switch:** Correct schema errors ([c4b3421](https://github.com/terraform-routeros/terraform-provider-routeros/commit/c4b3421fc39b2b1984aeab00b36c475844c4edf4)), closes [#325](https://github.com/terraform-routeros/terraform-provider-routeros/issues/325)

## [1.32.2](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.32.1...v1.32.2) (2024-01-16)


### Bug Fixes

* **switch:** Incorrect procedure for deleting the resource ([1e2327c](https://github.com/terraform-routeros/terraform-provider-routeros/commit/1e2327c43d934da59fc2a6a36a485b4caf18db4d)), closes [#325](https://github.com/terraform-routeros/terraform-provider-routeros/issues/325)

## [1.32.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.32.0...v1.32.1) (2024-01-16)


### Bug Fixes

* Add missing `default` property in `routeros_routing_bgp_template` ([0a2863f](https://github.com/terraform-routeros/terraform-provider-routeros/commit/0a2863ff01cfc26a904dab246716b4c9a7610f81))
* Add missing generated certificate properties in `routeros_capsman_manager` ([e4378af](https://github.com/terraform-routeros/terraform-provider-routeros/commit/e4378af30e5da4b1f4e6cb18cad5cb4b3ce61c2a))
* Add new `port_cost_mode` property to `routeros_interface_bridge` ([cab4fb6](https://github.com/terraform-routeros/terraform-provider-routeros/commit/cab4fb67e768b2e0b59c3ceac0edba07a1ba1138))
* Fix validator of the `address` property in `routeros_user_manager_router` ([1c7a46b](https://github.com/terraform-routeros/terraform-provider-routeros/commit/1c7a46b1bea703cafc5e1dd4782fce299d3383c6))
* Fix validator of the `advertise` property in `routeros_interface_ethernet` ([6c98bb7](https://github.com/terraform-routeros/terraform-provider-routeros/commit/6c98bb7360cbcaabbe3ef6d4688a0683577ef211))

## [1.32.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.31.0...v1.32.0) (2024-01-15)


### Features

* **switch:** Add support for /interface/ethernet/switch/host ([c267e5d](https://github.com/terraform-routeros/terraform-provider-routeros/commit/c267e5dbeab24578f3a7dbfcd9f5150b790391ba)), closes [#325](https://github.com/terraform-routeros/terraform-provider-routeros/issues/325)
* **switch:** Add support for /interface/ethernet/switch/port ([a3e7921](https://github.com/terraform-routeros/terraform-provider-routeros/commit/a3e7921a1c59a6e4d79965ec3ec8bc1cf3f0ce09)), closes [#325](https://github.com/terraform-routeros/terraform-provider-routeros/issues/325)
* **switch:** Add support for /interface/ethernet/switch/vlan ([2690373](https://github.com/terraform-routeros/terraform-provider-routeros/commit/2690373800ca29672b3c279fda4e99128e9f0860)), closes [#325](https://github.com/terraform-routeros/terraform-provider-routeros/issues/325)
* **switch:** Add support for interface/ethernet/switch/rule ([163ebbe](https://github.com/terraform-routeros/terraform-provider-routeros/commit/163ebbe4d54e5a2d7444d015306ab0cdd4565055)), closes [#325](https://github.com/terraform-routeros/terraform-provider-routeros/issues/325)

## [1.31.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.30.0...v1.31.0) (2024-01-15)


### Features

* Add WiFi access list resource ([e5a213a](https://github.com/terraform-routeros/terraform-provider-routeros/commit/e5a213abc3e10586471ddce5c0806a010997bcaa))
* Add Wifi accounting and authentication resource ([edd8299](https://github.com/terraform-routeros/terraform-provider-routeros/commit/edd8299edbad5b6c866d25659b262966d71fe09d))
* Add WiFi CAP resource ([d68fc6c](https://github.com/terraform-routeros/terraform-provider-routeros/commit/d68fc6c5c3d4207464446ee87451df00e4a885b1))
* Add WiFi CAPsMAN resource ([00197cf](https://github.com/terraform-routeros/terraform-provider-routeros/commit/00197cfc130d8316a43c351a66c1c757aaeb3c2b))
* Add WiFi channel resource ([3fd1f34](https://github.com/terraform-routeros/terraform-provider-routeros/commit/3fd1f34b040ee1e1eacce246a221a092315b8aea))
* Add WiFi configuration resource ([1505bff](https://github.com/terraform-routeros/terraform-provider-routeros/commit/1505bffb0d19c27553f7614b352dfc1a087c9806))
* Add WiFi datapath resource ([4193151](https://github.com/terraform-routeros/terraform-provider-routeros/commit/4193151e790ebe4b4b0074b3dd8e0cf5a4878ad8))
* Add WiFi interworking resource ([f748368](https://github.com/terraform-routeros/terraform-provider-routeros/commit/f748368c3ec2262eea531a51ec34706a30842f24))
* Add WiFi provisioning resource ([7a80781](https://github.com/terraform-routeros/terraform-provider-routeros/commit/7a80781f6bfe6bd1818c26a1f0f094317581fddc))
* Add WiFi security resource ([0c7e3af](https://github.com/terraform-routeros/terraform-provider-routeros/commit/0c7e3af28d23b27eff730331f0589484eef88b6d))
* Add WiFi steering resource ([3794b0f](https://github.com/terraform-routeros/terraform-provider-routeros/commit/3794b0fbf8bbd847deea81e5fc0554ffb010355f))

## [1.30.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.29.1...v1.30.0) (2024-01-12)


### Features

* **routing:** Manage Route Filters ([0c29e53](https://github.com/terraform-routeros/terraform-provider-routeros/commit/0c29e531be76e8c322062babd4c082d7e685ff7f)), closes [#330](https://github.com/terraform-routeros/terraform-provider-routeros/issues/330)

## [1.29.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.29.0...v1.29.1) (2024-01-11)


### Bug Fixes

* **helpers:** Fix the style of writing filter parameters ([adf342f](https://github.com/terraform-routeros/terraform-provider-routeros/commit/adf342ff668cde440265fb34461d0d105602fc9f))

## [1.29.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.28.1...v1.29.0) (2024-01-05)


### Features

* Add system user authentication and accounting settings resource ([1899983](https://github.com/terraform-routeros/terraform-provider-routeros/commit/18999832de2ff864008b1f84cc0801fdfc9697c5))
* Add system user group resource ([c7cc658](https://github.com/terraform-routeros/terraform-provider-routeros/commit/c7cc658b69c5e75033c4eddfebc42f928bccd4e2))
* Add system user settings resource ([f889b7b](https://github.com/terraform-routeros/terraform-provider-routeros/commit/f889b7bd95bf4f559888edcfb32640b8d4fd731a))


### Bug Fixes

* Change the validator ([0b1881f](https://github.com/terraform-routeros/terraform-provider-routeros/commit/0b1881ff69cd3bc56a30d9fd2e239e08ee480b0d))
* Fix the growth of the 'valid' slice ([206cfc2](https://github.com/terraform-routeros/terraform-provider-routeros/commit/206cfc277a1ce5caaa9842d5b33503ee861745b1))

## [1.28.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.28.0...v1.28.1) (2023-12-25)


### Bug Fixes

* **firewall:** Fix the error of deleting field-lists ([3f2f7b1](https://github.com/terraform-routeros/terraform-provider-routeros/commit/3f2f7b1f6d335f99e3ff73e82f147a678325d4cf))

## [1.28.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.27.2...v1.28.0) (2023-12-15)


### Features

* Data source to list DHCP leases ([71f9571](https://github.com/terraform-routeros/terraform-provider-routeros/commit/71f95712255dcb1c1420c243c5a4426ba5328afd)), closes [#316](https://github.com/terraform-routeros/terraform-provider-routeros/issues/316)


### Bug Fixes

* Comparison of MAC addresses in different character cases ([9acc3cc](https://github.com/terraform-routeros/terraform-provider-routeros/commit/9acc3cc705dd13a6e75d1487c1088eb7fea04609))

## [1.27.2](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.27.1...v1.27.2) (2023-12-06)


### Bug Fixes

* Firewall filter rules order ([#314](https://github.com/terraform-routeros/terraform-provider-routeros/issues/314)) ([3d32136](https://github.com/terraform-routeros/terraform-provider-routeros/commit/3d321362445fcfd401317f9829b1a206326dbae9)), closes [#293](https://github.com/terraform-routeros/terraform-provider-routeros/issues/293)

## [1.27.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.27.0...v1.27.1) (2023-12-06)


### Bug Fixes

* Add 2.5Gbps to validation set ([#313](https://github.com/terraform-routeros/terraform-provider-routeros/issues/313)) ([47d6781](https://github.com/terraform-routeros/terraform-provider-routeros/commit/47d67811bc70776b9aea58f4f8ec0dc0166eed42)), closes [#311](https://github.com/terraform-routeros/terraform-provider-routeros/issues/311)

## [1.27.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.26.0...v1.27.0) (2023-12-02)


### Features

* Add float type support ([4cc485b](https://github.com/terraform-routeros/terraform-provider-routeros/commit/4cc485b316747c59dee7e7efc3ddda32ee1c83b9))
* Add user manager advanced settings resource ([287db08](https://github.com/terraform-routeros/terraform-provider-routeros/commit/287db0886c31dfb1a91a7193ea944ffccc30747a))
* Add user manager attribute resource ([fa7dd30](https://github.com/terraform-routeros/terraform-provider-routeros/commit/fa7dd302356bafcb912d95eb41ec3ea339c4b78c))
* Add user manager database resource ([ca4e490](https://github.com/terraform-routeros/terraform-provider-routeros/commit/ca4e490fe75c85f41023d16f5316adef06c5cc37))
* Add user manager limitation resource ([a97bc45](https://github.com/terraform-routeros/terraform-provider-routeros/commit/a97bc45c75c99a525875046a01c2940069b624ba))
* Add user manager profile limitation resource ([733932c](https://github.com/terraform-routeros/terraform-provider-routeros/commit/733932c5bbca367fb5f461f3508b485ed428fb14))
* Add user manager profile resource ([d95bed1](https://github.com/terraform-routeros/terraform-provider-routeros/commit/d95bed153373122ef958d96a606b9199fa3f01f5))
* Add user manager router resource ([162e01e](https://github.com/terraform-routeros/terraform-provider-routeros/commit/162e01e556d9d92e0833ce470ff6486f645a0d34))
* Add user manager settings resource ([23761cf](https://github.com/terraform-routeros/terraform-provider-routeros/commit/23761cf5d2b13b28f93c6d2f2bc064f7156d58ca))
* Add user manager user group resource ([fff1568](https://github.com/terraform-routeros/terraform-provider-routeros/commit/fff156866c8f2a47edb52e407a4d129ad54a2185))
* Add user manager user profile resource ([9500635](https://github.com/terraform-routeros/terraform-provider-routeros/commit/95006359d071dc3c4ef2c3c19bf5847ab6f3dbbe))
* Add user manager user resource ([727cd9b](https://github.com/terraform-routeros/terraform-provider-routeros/commit/727cd9b196680d3649cccd506a7f88530aa735a4))


### Bug Fixes

* Fix the `AlwaysPresentNotUserProvided` helper to handle lists and sets correctly ([7c68003](https://github.com/terraform-routeros/terraform-provider-routeros/commit/7c68003df94e4fb3755101996e4418a41bb306b4))

## [1.26.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.25.1...v1.26.0) (2023-11-29)


### Features

* **/system/ntp/server:** Add NTP Server resource ([f7851ca](https://github.com/terraform-routeros/terraform-provider-routeros/commit/f7851ca3af4d26c3ebff5677153213b44b2d4657)), closes [#306](https://github.com/terraform-routeros/terraform-provider-routeros/issues/306)

## [1.25.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.25.0...v1.25.1) (2023-11-27)


### Bug Fixes

* Add missing schema fields ([#303](https://github.com/terraform-routeros/terraform-provider-routeros/issues/303)) ([bd245b9](https://github.com/terraform-routeros/terraform-provider-routeros/commit/bd245b91820bc3920b38cdc984c7332c36ce03d0))
* **no-release:** Fix incorrect addition of skip fields ([9ee6d70](https://github.com/terraform-routeros/terraform-provider-routeros/commit/9ee6d70170d4c290609742ae27bac48228edaf7c))

## [1.25.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.24.0...v1.25.0) (2023-11-15)


### Features

* Add RADIUS incoming resource ([6fb0a23](https://github.com/terraform-routeros/terraform-provider-routeros/commit/6fb0a23a98f6089493a35031d9722b978110bb41))
* Add RADIUS resource ([eb0b5c3](https://github.com/terraform-routeros/terraform-provider-routeros/commit/eb0b5c38c0758efc866441fe630fc529f84af195))


### Bug Fixes

* Add compatibility layer for the VRF property in RADIUS incoming resource ([96a5354](https://github.com/terraform-routeros/terraform-provider-routeros/commit/96a5354b95f4f94d5198b55f0d017b8477c8f0b0))
* **no-release:** Skip computed stat fields ([#299](https://github.com/terraform-routeros/terraform-provider-routeros/issues/299)) ([579f0a0](https://github.com/terraform-routeros/terraform-provider-routeros/commit/579f0a0ea9cc73ab4c038b8bf9a2721ce6a9f99f))

## [1.24.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.23.0...v1.24.0) (2023-11-13)


### Features

* **ds:** Add /ip/arp datasource ([6ecd622](https://github.com/terraform-routeros/terraform-provider-routeros/commit/6ecd622973facf7296fd0009304c0249f1e3e369))
* **ds:** Add /system/resource datasource ([79a599e](https://github.com/terraform-routeros/terraform-provider-routeros/commit/79a599eb980bd73b677098d0cc54505b66b6069e))


### Bug Fixes

* **REST:** Return possible error on JSON parsing ([7134eac](https://github.com/terraform-routeros/terraform-provider-routeros/commit/7134eac723886d14214ef60de0faaa532387483c))

## [1.23.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.22.1...v1.23.0) (2023-11-10)


### Features

* Add 802.1X client resource ([db76369](https://github.com/terraform-routeros/terraform-provider-routeros/commit/db763696e4003399b76cb474ea32614a4e8028db))
* Add 802.1X server resource ([05894af](https://github.com/terraform-routeros/terraform-provider-routeros/commit/05894afc3145a8af232b3eb517c95ac224f2cfe4))


### Bug Fixes

* Bump Go version from 1.19 to 1.21 ([#297](https://github.com/terraform-routeros/terraform-provider-routeros/issues/297)) ([1a56503](https://github.com/terraform-routeros/terraform-provider-routeros/commit/1a565038228d408d55d405423288a54f006e967c))

## [1.23.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.22.1...v1.23.0) (2023-11-10)


### Features

* Add 802.1X client resource ([db76369](https://github.com/terraform-routeros/terraform-provider-routeros/commit/db763696e4003399b76cb474ea32614a4e8028db))
* Add 802.1X server resource ([05894af](https://github.com/terraform-routeros/terraform-provider-routeros/commit/05894afc3145a8af232b3eb517c95ac224f2cfe4))

## [1.22.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.22.0...v1.22.1) (2023-11-09)


### Bug Fixes

* Unexpected value: aes256 for RouterOS v7.7 ([#294](https://github.com/terraform-routeros/terraform-provider-routeros/issues/294)) ([6fc8149](https://github.com/terraform-routeros/terraform-provider-routeros/commit/6fc8149e77d45a53312e7d9c9b50b451d4791460)), closes [#291](https://github.com/terraform-routeros/terraform-provider-routeros/issues/291)

## [1.22.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.21.0...v1.22.0) (2023-11-07)


### Features

* Add DHCP server config resource ([#288](https://github.com/terraform-routeros/terraform-provider-routeros/issues/288)) ([0e9fbbf](https://github.com/terraform-routeros/terraform-provider-routeros/commit/0e9fbbf52484962789bd28b2caaab9be238bff86))

## [1.21.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.20.1...v1.21.0) (2023-11-06)


### Features

* Add ethernet switch settings ([162c1da](https://github.com/terraform-routeros/terraform-provider-routeros/commit/162c1da0233e2e909d98dd02f011e89513233c9a)), closes [#285](https://github.com/terraform-routeros/terraform-provider-routeros/issues/285) [#282](https://github.com/terraform-routeros/terraform-provider-routeros/issues/282)
* Add MLAG settings ([6b8cfd2](https://github.com/terraform-routeros/terraform-provider-routeros/commit/6b8cfd246d893abefc88ee3933e039f7cf1de508)), closes [#268](https://github.com/terraform-routeros/terraform-provider-routeros/issues/268)


### Bug Fixes

* **bridge:** Add Name-Id migration ([84a7f3c](https://github.com/terraform-routeros/terraform-provider-routeros/commit/84a7f3c7cafeb700fe7fef8f367208b6d4ba2dc5))
* **CAPsMAN:** Add Name-Id migration ([5d0effa](https://github.com/terraform-routeros/terraform-provider-routeros/commit/5d0effa4d24365671dcbf268da66460b4483c25f))
* **dhcp_server:** Add Name-Id migration ([c8c9ff8](https://github.com/terraform-routeros/terraform-provider-routeros/commit/c8c9ff89d6b81253c7e8e8f020aba7cc47e03159))
* **eoip:** Add Name-Id migration ([fdbd68f](https://github.com/terraform-routeros/terraform-provider-routeros/commit/fdbd68f21f2ef5284b0e8529095a4cb5a3a76067))
* **eoip:** Fix the resource ID type ([7916c30](https://github.com/terraform-routeros/terraform-provider-routeros/commit/7916c30fcf36899631dfb25a362c378dc3ebbac4))
* **gre:** Add Name-Id migration ([6e811c3](https://github.com/terraform-routeros/terraform-provider-routeros/commit/6e811c3c3aa054f0c91d554cc9814ed3ff032b62))
* **interface_list:** Add Name-Id migration ([6b326c0](https://github.com/terraform-routeros/terraform-provider-routeros/commit/6b326c0dd62241e7b70307ac6400421f44ff94b1))
* **ip_pool:** Add Name-Id migration ([98c17c2](https://github.com/terraform-routeros/terraform-provider-routeros/commit/98c17c2130e798092a86e2e724b99c7a90980f15))
* **scheduler:** Add Name-Id migration ([c34b994](https://github.com/terraform-routeros/terraform-provider-routeros/commit/c34b99483f96eb9d1b49552326883fe9190c3c93))
* **vlan:** Add Name-Id migration ([7592f2f](https://github.com/terraform-routeros/terraform-provider-routeros/commit/7592f2f2fba7ea7d3a7875906b71a75126fd0f90))
* **vrrp:** Add Name-Id migration ([30c8f37](https://github.com/terraform-routeros/terraform-provider-routeros/commit/30c8f37a004c0b3b67dce85965535a0836b57af7))
* **wg:** Add Name-Id migration ([d676b8b](https://github.com/terraform-routeros/terraform-provider-routeros/commit/d676b8b26c64e114caeb108191ed911319c0f4ab))

## [1.20.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.20.0...v1.20.1) (2023-11-02)


### Bug Fixes

* Fix empty value check to handle default numeric values correctly ([#286](https://github.com/terraform-routeros/terraform-provider-routeros/issues/286)) ([661e49c](https://github.com/terraform-routeros/terraform-provider-routeros/commit/661e49ccbd0ca87eec024187e0e9ad6c2cb9890b))
* **no-release:** Some boolean params can't be reset and the provider does not understand the value ([#269](https://github.com/terraform-routeros/terraform-provider-routeros/issues/269)) ([678c9a4](https://github.com/terraform-routeros/terraform-provider-routeros/commit/678c9a4c67d2553e3bc72dc6a29d14d415520fa6)), closes [#253](https://github.com/terraform-routeros/terraform-provider-routeros/issues/253)

## [1.20.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.19.0...v1.20.0) (2023-10-31)


### Features

* Add EoIP tunnel support ([#283](https://github.com/terraform-routeros/terraform-provider-routeros/issues/283)) ([bcab0fb](https://github.com/terraform-routeros/terraform-provider-routeros/commit/bcab0fb38634a992250ff92271543c5f0dd309cc))

## [1.19.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.18.4...v1.19.0) (2023-10-29)


### Features

* Add CAPsMAN access-list resource ([#281](https://github.com/terraform-routeros/terraform-provider-routeros/issues/281)) ([a0379c9](https://github.com/terraform-routeros/terraform-provider-routeros/commit/a0379c9b22edff87dc5dbedb1f74d8b30d010f09))

## [1.18.4](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.18.3...v1.18.4) (2023-10-26)


### Bug Fixes

* Fix enumerated values in CAPsMAN resources ([978576a](https://github.com/terraform-routeros/terraform-provider-routeros/commit/978576a59768e79369e1ba0fdff00044113260cd))

## [1.18.3](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.18.2...v1.18.3) (2023-10-09)


### Bug Fixes

* Fix double slash at the end of a hostname ([d33aa79](https://github.com/terraform-routeros/terraform-provider-routeros/commit/d33aa79a78d4932d2c8f89c854d0aa8940e6642c)), closes [#275](https://github.com/terraform-routeros/terraform-provider-routeros/issues/275)

## [1.18.2](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.18.1...v1.18.2) (2023-10-01)


### Bug Fixes

* Improvements on the resource routeros_interface_ethernet ([#266](https://github.com/terraform-routeros/terraform-provider-routeros/issues/266)) ([099185b](https://github.com/terraform-routeros/terraform-provider-routeros/commit/099185beff487b514379a8472c0c208bdb6a6215))

## [1.18.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.18.0...v1.18.1) (2023-09-27)


### Bug Fixes

* Move WG  keys from datasource to resource ([#265](https://github.com/terraform-routeros/terraform-provider-routeros/issues/265)) ([a4eaf8c](https://github.com/terraform-routeros/terraform-provider-routeros/commit/a4eaf8c5fd00e56f7d69ddac5bfa575e8487ff60))

## [1.18.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.17.0...v1.18.0) (2023-09-24)


### Features

* Creating key sets for WireGuard tunnels ([e2d28a3](https://github.com/terraform-routeros/terraform-provider-routeros/commit/e2d28a3d8d1184ab2fb4118cb7b44147cb8fbbc3))

## [1.17.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.16.0...v1.17.0) (2023-09-22)


### Features

* Ip firewall connection tracking ([#260](https://github.com/terraform-routeros/terraform-provider-routeros/issues/260)) ([9d39bf8](https://github.com/terraform-routeros/terraform-provider-routeros/commit/9d39bf82ebbff621888bb6535fe57148488f0215))

## [1.16.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.15.0...v1.16.0) (2023-09-21)


### Features

* Implement routeros_system_logging resource ([#261](https://github.com/terraform-routeros/terraform-provider-routeros/issues/261)) ([f8c89aa](https://github.com/terraform-routeros/terraform-provider-routeros/commit/f8c89aa7a197ae765fbbaf3138dda06b1a8787e4))

## [1.15.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.14.0...v1.15.0) (2023-09-20)


### Features

* Add routeros_ip_dhcp_server_option and routeros_ip_dhcp_server_option_set ([#259](https://github.com/terraform-routeros/terraform-provider-routeros/issues/259)) ([3722afb](https://github.com/terraform-routeros/terraform-provider-routeros/commit/3722afb8574250a7a7e3211f2e40f3b4acfdc56f))

## [1.14.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.13.3...v1.14.0) (2023-09-19)


### Features

* Implementation of routeos_interface_ethernet ([#256](https://github.com/terraform-routeros/terraform-provider-routeros/issues/256))([#255](https://github.com/terraform-routeros/terraform-provider-routeros/issues/255)) ([0d848bf](https://github.com/terraform-routeros/terraform-provider-routeros/commit/0d848bf4b3d12b438bc9cbb137c91dca616b9d6a))

## [1.13.3](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.13.2...v1.13.3) (2023-09-18)


### Bug Fixes

* ip_service drift + failure ([#257](https://github.com/terraform-routeros/terraform-provider-routeros/issues/257)) ([b53b31b](https://github.com/terraform-routeros/terraform-provider-routeros/commit/b53b31bc4ada26245c272a3e597c3ff4e4ed5d6c)), closes [#254](https://github.com/terraform-routeros/terraform-provider-routeros/issues/254)

## [1.13.2](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.13.1...v1.13.2) (2023-07-20)


### Bug Fixes

* Add SNMP Settings ([#242](https://github.com/terraform-routeros/terraform-provider-routeros/issues/242)) ([e3a0d36](https://github.com/terraform-routeros/terraform-provider-routeros/commit/e3a0d36acb60c98088a298a21220746db0259dc4)), closes [#232](https://github.com/terraform-routeros/terraform-provider-routeros/issues/232)

## [1.13.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.13.0...v1.13.1) (2023-07-19)


### Bug Fixes

* no updates when modifying the cod ([23d175b](https://github.com/terraform-routeros/terraform-provider-routeros/commit/23d175bf6ea2361330900628696f9c31641aebdb)), closes [#240](https://github.com/terraform-routeros/terraform-provider-routeros/issues/240)

## [1.13.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.12.1...v1.13.0) (2023-07-13)


### Features

* Add an SNMP resource ([43c1ec9](https://github.com/terraform-routeros/terraform-provider-routeros/commit/43c1ec9f53a40b96f9bc47d7446e86c6c724f3e2)), closes [#232](https://github.com/terraform-routeros/terraform-provider-routeros/issues/232)
* Add SNMP community resource ([eeea040](https://github.com/terraform-routeros/terraform-provider-routeros/commit/eeea04099b6431f508eb68cdd4d924653a7d17ff))

## [1.12.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.12.0...v1.12.1) (2023-07-12)


### Bug Fixes

* Fix the ParseDuration function ([1995d9e](https://github.com/terraform-routeros/terraform-provider-routeros/commit/1995d9ef8b8eac2808a34f3e433a24bd978e3bbc))
* Resource firewall filter ([fa04b82](https://github.com/terraform-routeros/terraform-provider-routeros/commit/fa04b820b7754d2e7cfdbab0d1064075415bf31d)), closes [#237](https://github.com/terraform-routeros/terraform-provider-routeros/issues/237)

## [1.12.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.11.0...v1.12.0) (2023-06-19)


### Features

* Add IP Cloud ([#234](https://github.com/terraform-routeros/terraform-provider-routeros/issues/234)) ([675e9f3](https://github.com/terraform-routeros/terraform-provider-routeros/commit/675e9f3b57735ca035af34b43a0568fe2ee71c28)), closes [#231](https://github.com/terraform-routeros/terraform-provider-routeros/issues/231)

## [1.11.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.10.4...v1.11.0) (2023-06-19)


### Features

* New OSPF resource ([4d473ea](https://github.com/terraform-routeros/terraform-provider-routeros/commit/4d473ea2e958a52f6544742cf47f0a1190e36508))


### Bug Fixes

* Add a helper for the attribute 'inactive' ([adca988](https://github.com/terraform-routeros/terraform-provider-routeros/commit/adca988ca08d36d6cab9a839fbad891827e72a81))
* Fix for error "no-summaries only valid for stubby areas" ([f222f71](https://github.com/terraform-routeros/terraform-provider-routeros/commit/f222f71691fd6e67b9967d66ef541e9d88376cea))

## [1.10.4](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.10.3...v1.10.4) (2023-05-30)


### Bug Fixes

* Patching firewall rules with place_before ([#224](https://github.com/terraform-routeros/terraform-provider-routeros/issues/224)) ([5ef738e](https://github.com/terraform-routeros/terraform-provider-routeros/commit/5ef738e1d530bb0c2aeb0abdeb8fd71f535150b7)), closes [#223](https://github.com/terraform-routeros/terraform-provider-routeros/issues/223)

## [1.10.3](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.10.2...v1.10.3) (2023-05-27)


### Bug Fixes

* ci fix expired token ([#220](https://github.com/terraform-routeros/terraform-provider-routeros/issues/220)) ([e6a8585](https://github.com/terraform-routeros/terraform-provider-routeros/commit/e6a85853e60c4fafd49d59a7e702a9bae53f4678))
* **docs:** Update release.yml ([#221](https://github.com/terraform-routeros/terraform-provider-routeros/issues/221)) ([44ba77d](https://github.com/terraform-routeros/terraform-provider-routeros/commit/44ba77d1ffe8b9863a3e01a16379dea23406c8ee)), closes [#219](https://github.com/terraform-routeros/terraform-provider-routeros/issues/219)
* Wrong field names in example files ([#219](https://github.com/terraform-routeros/terraform-provider-routeros/issues/219)) ([b0105ef](https://github.com/terraform-routeros/terraform-provider-routeros/commit/b0105ef1b295f7468debab14735e557e12eb01f3)), closes [#218](https://github.com/terraform-routeros/terraform-provider-routeros/issues/218)

## [1.10.2](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.10.1...v1.10.2) (2023-05-23)


### Bug Fixes

* Remove extra space after passthrough in validation ([#217](https://github.com/terraform-routeros/terraform-provider-routeros/issues/217)) ([3061910](https://github.com/terraform-routeros/terraform-provider-routeros/commit/306191072d3ceb57acc4e0533ed878e1f6a18646))

## [1.10.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.10.0...v1.10.1) (2023-05-19)


### Bug Fixes

* ipv6 addr-list addr w/o netmask adds /128 netmask [#216](https://github.com/terraform-routeros/terraform-provider-routeros/issues/216)  ([d6f7fad](https://github.com/terraform-routeros/terraform-provider-routeros/commit/d6f7fadfed9bac3e9bcc60640d935c08499053d2))

## [1.10.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.9.1...v1.10.0) (2023-05-17)


### Features

* Add support for /interface/pppoe-client ([#215](https://github.com/terraform-routeros/terraform-provider-routeros/issues/215)) ([a8cbe7d](https://github.com/terraform-routeros/terraform-provider-routeros/commit/a8cbe7d78ec868752f11184131c09b662e561a4c)), closes [#202](https://github.com/terraform-routeros/terraform-provider-routeros/issues/202)

## [1.9.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.9.0...v1.9.1) (2023-05-17)


### Bug Fixes

* Field 'comment' not found in the schema ([#214](https://github.com/terraform-routeros/terraform-provider-routeros/issues/214)) ([01a7f10](https://github.com/terraform-routeros/terraform-provider-routeros/commit/01a7f101ade024981f8c59a56775aa1f4bdae442)), closes [#213](https://github.com/terraform-routeros/terraform-provider-routeros/issues/213)

## [1.9.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.8.0...v1.9.0) (2023-05-17)


### Features

* Support ipv6 firewall address lists ([878fbf7](https://github.com/terraform-routeros/terraform-provider-routeros/commit/878fbf70d8e78c8993105da01f41a8ce8b9df4cb)), closes [#212](https://github.com/terraform-routeros/terraform-provider-routeros/issues/212)

## [1.8.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.7.2...v1.8.0) (2023-05-16)


### Features

* Support bridge settings ([0bea447](https://github.com/terraform-routeros/terraform-provider-routeros/commit/0bea447ceed39579b565cd9041fdb98769e21f46)), closes [#209](https://github.com/terraform-routeros/terraform-provider-routeros/issues/209)

## [1.7.2](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.7.1...v1.7.2) (2023-05-15)


### Bug Fixes

* nil pointer on bgp ([93cf45e](https://github.com/terraform-routeros/terraform-provider-routeros/commit/93cf45e45352b09cf24e82273d6905e10c0b1f13)), closes [#207](https://github.com/terraform-routeros/terraform-provider-routeros/issues/207)

## [1.7.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.7.0...v1.7.1) (2023-05-15)


### Bug Fixes

* Fix resource names [#183](https://github.com/terraform-routeros/terraform-provider-routeros/issues/183) ([a4314d0](https://github.com/terraform-routeros/terraform-provider-routeros/commit/a4314d0c2765f7d6eea1cc672bf8c5dc633e9941))
* Fix the gateway field (veth) ([97b933b](https://github.com/terraform-routeros/terraform-provider-routeros/commit/97b933b2221005433cf249403086bce6d970c202))

## [1.7.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.6.0...v1.7.0) (2023-05-14)


### Features

* BGP connection ([3874d90](https://github.com/terraform-routeros/terraform-provider-routeros/commit/3874d909ff245e5493368a4e3d472e45cdcad65c)), closes [#183](https://github.com/terraform-routeros/terraform-provider-routeros/issues/183)
* BGP templates ([7984574](https://github.com/terraform-routeros/terraform-provider-routeros/commit/7984574c9282019894e02b3f4b3fab04461c80a5)), closes [#183](https://github.com/terraform-routeros/terraform-provider-routeros/issues/183)
* Processing nested fields in a list ([23928a0](https://github.com/terraform-routeros/terraform-provider-routeros/commit/23928a02c724d40b533ef16ec07deb2551497fb2))
* Support for /interface/bonding [#203](https://github.com/terraform-routeros/terraform-provider-routeros/issues/203) ([a7de21f](https://github.com/terraform-routeros/terraform-provider-routeros/commit/a7de21fd630450590a08be5b94449f82c33e6bbf))
* Support for veth interfaces [#206](https://github.com/terraform-routeros/terraform-provider-routeros/issues/206) ([a6fdcf8](https://github.com/terraform-routeros/terraform-provider-routeros/commit/a6fdcf80a2af0d2b14b66683fb578160e7555a99))


### Bug Fixes

* Changing the signature isEmpty + fixing the result for boolean values ([aedc90e](https://github.com/terraform-routeros/terraform-provider-routeros/commit/aedc90e5efbc6cc98adb558893cc17f727adeda9))
* Correct the logic of isEmpty ([18f4bf1](https://github.com/terraform-routeros/terraform-provider-routeros/commit/18f4bf19c58b4a5a8139e86946c254d7310ba013))
* Use helpers to process data for TypeMap ([280c994](https://github.com/terraform-routeros/terraform-provider-routeros/commit/280c994d060af676620cd592d26bcd988cc90405))

## [1.6.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.5.0...v1.6.0) (2023-05-05)


### Features

* Support creating users [#200](https://github.com/terraform-routeros/terraform-provider-routeros/issues/200) ([#201](https://github.com/terraform-routeros/terraform-provider-routeros/issues/201)) ([78191e2](https://github.com/terraform-routeros/terraform-provider-routeros/commit/78191e2038607af5081d06dfaabc208010f6d667))

## [1.5.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.4.0...v1.5.0) (2023-05-04)


### Features

* Add OpenVPN Server support ([6477fcd](https://github.com/terraform-routeros/terraform-provider-routeros/commit/6477fcdc61d5769a685fb32e551b41609f6f6aa6))


### Bug Fixes

* Rename the PropNameRw property and add a new one without forced re-creation ([a37f926](https://github.com/terraform-routeros/terraform-provider-routeros/commit/a37f926d2a1d8ffbddfbf5e75c5a28591f33e44c))

## [1.4.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.3.9...v1.4.0) (2023-05-01)


### Features

* Support for ip/services ([#195](https://github.com/terraform-routeros/terraform-provider-routeros/issues/195)) ([591096d](https://github.com/terraform-routeros/terraform-provider-routeros/commit/591096d4a249acb6f4484bc16a5aec691577453c)), closes [#182](https://github.com/terraform-routeros/terraform-provider-routeros/issues/182)

## [1.3.9](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.3.8...v1.3.9) (2023-05-01)


### Bug Fixes

* Fix the creation of resources when renaming them ([c229d27](https://github.com/terraform-routeros/terraform-provider-routeros/commit/c229d27e27802507901381218585f27131db6c2a)), closes [#192](https://github.com/terraform-routeros/terraform-provider-routeros/issues/192)

## [1.3.8](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.3.7...v1.3.8) (2023-04-26)


### Bug Fixes

* Warnings for primary_ntp and secondary_ntp when using routeros_ip_dhcp_client ([#190](https://github.com/terraform-routeros/terraform-provider-routeros/issues/190)) ([a7fc49f](https://github.com/terraform-routeros/terraform-provider-routeros/commit/a7fc49f1899f441fabbc63148a1b5c075cbaf27c)), closes [#189](https://github.com/terraform-routeros/terraform-provider-routeros/issues/189)

## [1.3.7](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.3.6...v1.3.7) (2023-04-23)


### Bug Fixes

* Add `check_gateway` field to `routeros_ip_route` ([#187](https://github.com/terraform-routeros/terraform-provider-routeros/issues/187)) ([20b84ae](https://github.com/terraform-routeros/terraform-provider-routeros/commit/20b84aea4b4ce3af725ebd0f5165cca010f5692a)), closes [#186](https://github.com/terraform-routeros/terraform-provider-routeros/issues/186)

## [1.3.6](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.3.5...v1.3.6) (2023-04-22)


### Bug Fixes

* ip/route dst_address should not be mandatory ([#185](https://github.com/terraform-routeros/terraform-provider-routeros/issues/185)) ([9cf42c7](https://github.com/terraform-routeros/terraform-provider-routeros/commit/9cf42c7ced6b813a9c0cf8465d1814ab1a5bce98)), closes [#184](https://github.com/terraform-routeros/terraform-provider-routeros/issues/184)

## [1.3.5](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.3.4...v1.3.5) (2023-04-20)


### Bug Fixes

* Ability to set clamp-tcp-mss on mangle rule ([3226a91](https://github.com/terraform-routeros/terraform-provider-routeros/commit/3226a91963be950ba4111fd133b5e4f104a4fbbe)), closes [#178](https://github.com/terraform-routeros/terraform-provider-routeros/issues/178)
* disabled mangle not seen ([ffb53d6](https://github.com/terraform-routeros/terraform-provider-routeros/commit/ffb53d66cb4edf8b7952f890c1d8e14f6f11b60b)), closes [#175](https://github.com/terraform-routeros/terraform-provider-routeros/issues/175)

## [1.3.4](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.3.3...v1.3.4) (2023-04-20)


### Bug Fixes

* dns servers cannot be removed ([#179](https://github.com/terraform-routeros/terraform-provider-routeros/issues/179)) ([3db9080](https://github.com/terraform-routeros/terraform-provider-routeros/commit/3db9080137431e5031de9f27d53e8f022154f40a)), closes [#174](https://github.com/terraform-routeros/terraform-provider-routeros/issues/174)

## [1.3.3](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.3.2...v1.3.3) (2023-04-19)


### Bug Fixes

* dns servers cannot be removed ([#177](https://github.com/terraform-routeros/terraform-provider-routeros/issues/177)) ([34a73be](https://github.com/terraform-routeros/terraform-provider-routeros/commit/34a73bed579e84568bfc24a3ead4bf8c1c62bbe9)), closes [#174](https://github.com/terraform-routeros/terraform-provider-routeros/issues/174)

## [1.3.2](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.3.1...v1.3.2) (2023-04-18)


### Bug Fixes

* dns servers cannot be removed ([#176](https://github.com/terraform-routeros/terraform-provider-routeros/issues/176)) ([1ebc4d9](https://github.com/terraform-routeros/terraform-provider-routeros/commit/1ebc4d98072c86499bb972081ba1649e2af52ef0)), closes [#174](https://github.com/terraform-routeros/terraform-provider-routeros/issues/174)

## [1.3.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.3.0...v1.3.1) (2023-04-12)


### Bug Fixes

* Remove default for VRRP interface group ([0cd9b5d](https://github.com/terraform-routeros/terraform-provider-routeros/commit/0cd9b5d41e10932ba631ae4d00b05c0ef948bbf0))

## [1.3.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.2.0...v1.3.0) (2023-04-10)


### Features

* Add support for a "certificate" resource ([898d2ad](https://github.com/terraform-routeros/terraform-provider-routeros/commit/898d2adf540ddcc04d4e535a36aee91fa3558fcd))

## [1.2.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.1.11...v1.2.0) (2023-04-03)


### Features

* Add support for CAPsMAN resources ([514b51f](https://github.com/terraform-routeros/terraform-provider-routeros/commit/514b51fd39250569e4c4112f7d349f98f885d743))
* Add support for composite types (TypeMap), TypeList.Int, TypeSet.Int, TypeSet.String ([8698a18](https://github.com/terraform-routeros/terraform-provider-routeros/commit/8698a18aa3179b281156c4fc311ba0ed5f5692a8))
* Add support for transforming the composite fields of Mikrotik. ([47d9ad3](https://github.com/terraform-routeros/terraform-provider-routeros/commit/47d9ad388753ff22eaee2bc83158920c37b60fd7))
* Add the default actions for system resources ([b3fb513](https://github.com/terraform-routeros/terraform-provider-routeros/commit/b3fb5138177b42e8dfb35a6d489f807fa8032be8))


### Bug Fixes

* Fix the import path ([a195c45](https://github.com/terraform-routeros/terraform-provider-routeros/commit/a195c45c1599b2e3f920c43cbc4e6dbabf8c895d))
* The 'disabled' property must be Computed (read-only) ([c4b85f6](https://github.com/terraform-routeros/terraform-provider-routeros/commit/c4b85f6ae3580f557acc4eaa14cacc371638ebda))

## [1.1.11](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.1.10...v1.1.11) (2023-03-23)


### Bug Fixes

* routeros_ip_dns - new fields in 7.8 ([#170](https://github.com/terraform-routeros/terraform-provider-routeros/issues/170)) ([c3d3eb3](https://github.com/terraform-routeros/terraform-provider-routeros/commit/c3d3eb3bdb9ac21bded109717bdba5075a1720ee)), closes [#169](https://github.com/terraform-routeros/terraform-provider-routeros/issues/169)

## [1.1.10](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.1.9...v1.1.10) (2023-03-21)


### Bug Fixes

* Fix the order of document generation in CI ([3793cde](https://github.com/terraform-routeros/terraform-provider-routeros/commit/3793cde6785344aa9c5a092ed9142263a340949e))

## [1.1.9](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.1.8...v1.1.9) (2023-03-21)


### Bug Fixes

* Fix [#165](https://github.com/terraform-routeros/terraform-provider-routeros/issues/165) for REST responses containing escape sequences  ([#167](https://github.com/terraform-routeros/terraform-provider-routeros/issues/167)) ([646ba4f](https://github.com/terraform-routeros/terraform-provider-routeros/commit/646ba4f6843904ec0122eb89285044104b051aa2))

## [1.1.8](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.1.7...v1.1.8) (2023-03-12)


### Bug Fixes

* routeros_dns_record does not change resource type and data correctly after [#158](https://github.com/terraform-routeros/terraform-provider-routeros/issues/158) ([4d95e80](https://github.com/terraform-routeros/terraform-provider-routeros/commit/4d95e80e73f8f494be7d3ea5fca382cc4e3f2fc5)), closes [#159](https://github.com/terraform-routeros/terraform-provider-routeros/issues/159)

## [1.1.7](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.1.6...v1.1.7) (2023-03-11)


### Bug Fixes

* /ip/dns/static errors when trying to change the resource type ([0a935cd](https://github.com/terraform-routeros/terraform-provider-routeros/commit/0a935cd4affd9effdd8c0e8190415d055e4aafa9)), closes [#156](https://github.com/terraform-routeros/terraform-provider-routeros/issues/156)

## [1.1.6](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.1.5...v1.1.6) (2023-03-10)


### Bug Fixes

* /ip/route - field disabled missing ([0baf464](https://github.com/terraform-routeros/terraform-provider-routeros/commit/0baf464ff8064fa38cf4458e4c56d3a0733f9865)), closes [#149](https://github.com/terraform-routeros/terraform-provider-routeros/issues/149)

## [1.1.5](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.1.4...v1.1.5) (2023-03-10)


### Bug Fixes

* /ip/dns/record - field type is missing ([2072d14](https://github.com/terraform-routeros/terraform-provider-routeros/commit/2072d1486c2c1ded8259a628dc2e447a519a2a92)), closes [#150](https://github.com/terraform-routeros/terraform-provider-routeros/issues/150)

## [1.1.4](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.1.3...v1.1.4) (2023-03-04)


### Bug Fixes

* Fix "post-test destroy" error ([38e79a4](https://github.com/terraform-routeros/terraform-provider-routeros/commit/38e79a497b1ff364ffc5c3b3a6e0d11c958d3616))
* Fix /ip/dhcp-server/network required field ([e9c69be](https://github.com/terraform-routeros/terraform-provider-routeros/commit/e9c69be9eca4b6791157c2a00bae0fdc436fec74))

## [1.1.3](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.1.2...v1.1.3) (2023-02-24)

## [1.1.2](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.1.1...v1.1.2) (2023-02-24)

## [1.1.1](https://github.com/GNewbury1/terraform-provider-routeros/compare/v1.1.0...v1.1.1) (2023-02-23)


### Bug Fixes

* Set key correctly ([a299647](https://github.com/GNewbury1/terraform-provider-routeros/commit/a299647b7a23891bf6574bfe503bfa3b6d397cbd))

# [1.1.0](https://github.com/GNewbury1/terraform-provider-routeros/compare/v1.0.13...v1.1.0) (2023-02-23)


### Features

* Add new signing key for new org ([7c0364a](https://github.com/GNewbury1/terraform-provider-routeros/commit/7c0364aa3bdfe3905cc7f588f9a114e98cbc76c8))

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
