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
