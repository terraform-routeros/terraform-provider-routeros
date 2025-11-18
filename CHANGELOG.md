## [1.92.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.92.0...v1.92.1) (2025-11-18)

### Bug Fixes

* routeros_interface_macvlan add missing attributes ([#865](https://github.com/terraform-routeros/terraform-provider-routeros/issues/865)) ([5b21944](https://github.com/terraform-routeros/terraform-provider-routeros/commit/5b21944a2ff183e8029b63ef468493bbcf1fa0ad))

## [1.92.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.91.0...v1.92.0) (2025-11-18)

### Features

* Add attributes to resouce ip_dhcp_server,interface_vlan … ([#861](https://github.com/terraform-routeros/terraform-provider-routeros/issues/861)) ([a88d0e1](https://github.com/terraform-routeros/terraform-provider-routeros/commit/a88d0e1ce67ca8ca3594fdacb1b710f5f01a64ed))

## [1.91.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.90.0...v1.91.0) (2025-11-10)

### Features

* **devcontainer:** add initial setup ([b4f9922](https://github.com/terraform-routeros/terraform-provider-routeros/commit/b4f992205736fee359f057b357f9c66cee3aa23f))

## [1.90.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.89.0...v1.90.0) (2025-11-01)

### Features

* **ip nat-pmp:** add examples ([174f0d1](https://github.com/terraform-routeros/terraform-provider-routeros/commit/174f0d158ec28ade47069ffe4fefad5198f5ba8f))
* **ip nat-pmp:** add resources ([3eaf59b](https://github.com/terraform-routeros/terraform-provider-routeros/commit/3eaf59b41dd56a20e2c09c0fa27d3083b95a4fff))
* **ospf area range:** Add new resource `routeros_routing_ospf_area_range` ([a336bd3](https://github.com/terraform-routeros/terraform-provider-routeros/commit/a336bd373908ac9f977a5462245470660f085cfa)), closes [#841](https://github.com/terraform-routeros/terraform-provider-routeros/issues/841)
* update routeros_interface_ethernet to ROS 7.20 ([87dcd1f](https://github.com/terraform-routeros/terraform-provider-routeros/commit/87dcd1fb7b079f42908fe5d84320fce3aa538263))
* update routeros_ip_firewall_connection_tracking to ROS 7.20 ([d2bc43c](https://github.com/terraform-routeros/terraform-provider-routeros/commit/d2bc43c74b8bf6b5375cceedd8fb4b53e8a86e1c))
* update routeros_system_user_sshkeys to ROS 7.20 ([2f28d7b](https://github.com/terraform-routeros/terraform-provider-routeros/commit/2f28d7b5f94aa6b634db1d717e6c38ea4cb4dc27))

### Bug Fixes

* **container:** Add missing attributes ([9b78f48](https://github.com/terraform-routeros/terraform-provider-routeros/commit/9b78f48e12a4d681b3e114568a873742caa31aed)), closes [#839](https://github.com/terraform-routeros/terraform-provider-routeros/issues/839)
* **contaner envs:** Change attribute `name` to `list` ([e004f90](https://github.com/terraform-routeros/terraform-provider-routeros/commit/e004f90339deeeb691b7c5cdbf41f64689652cdb)), closes [#856](https://github.com/terraform-routeros/terraform-provider-routeros/issues/856)
* **interface l2tp client:** Add a new attribute `random_source_port` ([fc1c511](https://github.com/terraform-routeros/terraform-provider-routeros/commit/fc1c5119dcbb3335b5b329ef273d2b84b240367f)), closes [#836](https://github.com/terraform-routeros/terraform-provider-routeros/issues/836)
* **ip nat-pmp interfaces:** fix property name ([46f95cf](https://github.com/terraform-routeros/terraform-provider-routeros/commit/46f95cf80d07679b5193d4dcc83b61a9cd0a1e19))
* **ip nat-pmp:** set min version for tests ([fa74e8a](https://github.com/terraform-routeros/terraform-provider-routeros/commit/fa74e8a5dbb3a4560f2e0f2b3d76d8bbba0e5dcf))
* **ip upnp:** typos ([fca8ff7](https://github.com/terraform-routeros/terraform-provider-routeros/commit/fca8ff7c8687513df7cf35db37df1121a1204348))
* **ospf interface template:** Add a new attribute `use-bfd` ([5a9b850](https://github.com/terraform-routeros/terraform-provider-routeros/commit/5a9b850b3dba817603356691d84e0af4972e9b0c)), closes [#837](https://github.com/terraform-routeros/terraform-provider-routeros/issues/837)
* **pppoe client:** Add a new attribute `host-uniq` ([01379db](https://github.com/terraform-routeros/terraform-provider-routeros/commit/01379dbb56740dfcfd4a16b13e418008bc72b00f)), closes [#846](https://github.com/terraform-routeros/terraform-provider-routeros/issues/846)

## [1.89.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.88.1...v1.89.0) (2025-10-15)

### Features

* Add support for EVPN-related resources ([bc250cb](https://github.com/terraform-routeros/terraform-provider-routeros/commit/bc250cbcc50eb1bcfcf49e587dbf0ccf7054a55c)), closes [#825](https://github.com/terraform-routeros/terraform-provider-routeros/issues/825)
* **bfd:** Add new resource `routeros_routing_bfd_configuration` ([2d88357](https://github.com/terraform-routeros/terraform-provider-routeros/commit/2d88357667e8898cd2d5bba45b4d4abd4e77e061)), closes [#829](https://github.com/terraform-routeros/terraform-provider-routeros/issues/829)
* Update `routeros_container` to ROS 7.20 ([8f10b7e](https://github.com/terraform-routeros/terraform-provider-routeros/commit/8f10b7ec13a01538bb85ac8549cd6d5dcebf41fa))
* Update `routeros_dhcp_server_network` to ROS 7.20 ([489e7e0](https://github.com/terraform-routeros/terraform-provider-routeros/commit/489e7e04f35681ab748e48b75fee21bd9514ec0d))
* Update `routeros_interface_veth` to ROS 7.20 ([d90e281](https://github.com/terraform-routeros/terraform-provider-routeros/commit/d90e28154a9518b2f6ce6752be531b0cd7e20de9)), closes [#826](https://github.com/terraform-routeros/terraform-provider-routeros/issues/826)
* Update `routeros_interface_vrrp` to ROS 7.20 ([0065b69](https://github.com/terraform-routeros/terraform-provider-routeros/commit/0065b69ef3b141abcd7b0766971621f141da30b5))
* Update `routeros_interface_vxlan` to ROS 7.20 ([638e781](https://github.com/terraform-routeros/terraform-provider-routeros/commit/638e781630ed41e977e931e309e04253c28f8903))
* Update `routeros_ip_dhcp_client` to ROS 7.20 ([8e2f838](https://github.com/terraform-routeros/terraform-provider-routeros/commit/8e2f8381837934961a9646b67ca51bed56f407f8))
* Update `routeros_ip_firewall_connection_tracking` to ROS 7.20 ([1b3c77d](https://github.com/terraform-routeros/terraform-provider-routeros/commit/1b3c77d8181ad5746af1202b51a04a849c150194))
* Update `routeros_ip_firewall_nat` to ROS 7.20 ([ce43460](https://github.com/terraform-routeros/terraform-provider-routeros/commit/ce43460b5e13aeb0b2ae3709e4c5fbeffcf03bdf))
* Update `routeros_ipv6_dhcp_client` to ROS 7.20 ([f91acba](https://github.com/terraform-routeros/terraform-provider-routeros/commit/f91acba91e82e3ef57fd2eaf23f4cc1936405d0e))
* Update `routeros_ipv6_dhcp_server` to ROS 7.20 ([f11a08b](https://github.com/terraform-routeros/terraform-provider-routeros/commit/f11a08b2258043877db8c0c4bfd3291e1272a57a))
* Update `routeros_ppp_profile` to ROS 7.20 ([bacdbfe](https://github.com/terraform-routeros/terraform-provider-routeros/commit/bacdbfe4eb44f96bde84ef24d318137dc85f5e4c))
* Update `routeros_radius` to ROS 7.20 ([42a8c7e](https://github.com/terraform-routeros/terraform-provider-routeros/commit/42a8c7e1b523e6dbea404bff38b8255785c9ce38))
* Update `routeros_system_logging_action` to ROS 7.20 ([8d4543f](https://github.com/terraform-routeros/terraform-provider-routeros/commit/8d4543fd78edf4fae2cf411d91ad384f885171dc))
* Update `routeros_tool_netwatch` to ROS 7.20 ([837ebab](https://github.com/terraform-routeros/terraform-provider-routeros/commit/837ebab5a5ae206a3c46b920423b0b17a389cf40))
* Update `routeros_wifi_channel` to ROS 7.20 ([11ef93c](https://github.com/terraform-routeros/terraform-provider-routeros/commit/11ef93c8587aaf65bdb59296e7fd081f6d6a88e9))
* Update `routeros_wifi_configuration` to ROS 7.20 ([0a1614f](https://github.com/terraform-routeros/terraform-provider-routeros/commit/0a1614ff664f209cb882fd6507b783f544014770))
* Update `routeros_wifi` to ROS 7.20 ([61e821f](https://github.com/terraform-routeros/terraform-provider-routeros/commit/61e821f87dea7e4d5d3b403825551f1a30028b6e))

### Bug Fixes

* **capsman_interface:** Added processing of the `configuration.config` parameter and an extended resource import function ([e18ca7a](https://github.com/terraform-routeros/terraform-provider-routeros/commit/e18ca7a6bf25e8808897b54f8b801792da484c4b)), closes [#828](https://github.com/terraform-routeros/terraform-provider-routeros/issues/828)

## [1.88.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.88.0...v1.88.1) (2025-10-14)

### Bug Fixes

* **cloud:** Missing fields at the `routeros_ip_cloud resource` ([ada8ea0](https://github.com/terraform-routeros/terraform-provider-routeros/commit/ada8ea0a750366dbfc2f0a98a8ffcef9dcb476c4)), closes [#820](https://github.com/terraform-routeros/terraform-provider-routeros/issues/820)

## [1.88.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.87.0...v1.88.0) (2025-09-26)

### Features

* **firewall:** Add new resource `routeros_ip_firewall_layer7_protocol` ([6e49323](https://github.com/terraform-routeros/terraform-provider-routeros/commit/6e49323b5c9fe12f164e7d8e45873a6e4aab6cc0)), closes [#817](https://github.com/terraform-routeros/terraform-provider-routeros/issues/817)

## [1.87.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.86.3...v1.87.0) (2025-09-25)

### Features

* **detect-internet:** Add new resource `routeros_interface_detect_internet` ([92aaa08](https://github.com/terraform-routeros/terraform-provider-routeros/commit/92aaa0832a95a55cc4d3b9fb9dfcc410d09d4b95)), closes [#808](https://github.com/terraform-routeros/terraform-provider-routeros/issues/808)

### Bug Fixes

* **ipv6-firewall:** Wrong allowed values for `reject_with` field ([392ea04](https://github.com/terraform-routeros/terraform-provider-routeros/commit/392ea04e6e4894d105e043c54d9e8e4272c306ba)), closes [#810](https://github.com/terraform-routeros/terraform-provider-routeros/issues/810)

## [1.87.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.86.3...v1.87.0) (2025-09-25)

### Features

* **detect-internet:** Add new resource `routeros_interface_detect_internet` ([92aaa08](https://github.com/terraform-routeros/terraform-provider-routeros/commit/92aaa0832a95a55cc4d3b9fb9dfcc410d09d4b95)), closes [#808](https://github.com/terraform-routeros/terraform-provider-routeros/issues/808)

### Bug Fixes

* **ipv6-firewall:** Wrong allowed values for `reject_with` field ([392ea04](https://github.com/terraform-routeros/terraform-provider-routeros/commit/392ea04e6e4894d105e043c54d9e8e4272c306ba)), closes [#810](https://github.com/terraform-routeros/terraform-provider-routeros/issues/810)

## [1.86.3](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.86.2...v1.86.3) (2025-08-26)

### Bug Fixes

* **interface_veth:** Add support for multiple addresses on the interface ([c6a656f](https://github.com/terraform-routeros/terraform-provider-routeros/commit/c6a656facc05a34f4cbbbb2fbba320e465cdfdd0)), closes [#804](https://github.com/terraform-routeros/terraform-provider-routeros/issues/804)
* **wireless:** Support 2ghz-g/n band specification ([1647d25](https://github.com/terraform-routeros/terraform-provider-routeros/commit/1647d2574b8bfaf02598dde26ab470233409c8f0))

## [1.86.2](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.86.1...v1.86.2) (2025-08-04)

### Bug Fixes

* **container:** Remove mutually exclusive attributes ([d879a0c](https://github.com/terraform-routeros/terraform-provider-routeros/commit/d879a0c69b415a2e0ecc01071224f57555dba0de)), closes [#793](https://github.com/terraform-routeros/terraform-provider-routeros/issues/793)

## [1.86.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.86.0...v1.86.1) (2025-08-01)

### Bug Fixes

* **container:** Add a reaction to the start/stop of a container from outside TF ([c46d5e6](https://github.com/terraform-routeros/terraform-provider-routeros/commit/c46d5e61a40234a72e57f16e992078cd33ed10d5)), closes [#793](https://github.com/terraform-routeros/terraform-provider-routeros/issues/793)

## [1.86.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.85.4...v1.86.0) (2025-07-27)

### Features

* implement RouterOS DHCP server option matcher resource ([d459e68](https://github.com/terraform-routeros/terraform-provider-routeros/commit/d459e6891e325f320975213af72ae4971ec1f765))
* implement RouterOS TFTP resource ([2b98b84](https://github.com/terraform-routeros/terraform-provider-routeros/commit/2b98b8451f7fd5ebe3e03b4fcf338138182c52a5))
* implement RouterOS TFTP settings resource ([4533151](https://github.com/terraform-routeros/terraform-provider-routeros/commit/4533151c59044db366f8afe4bbc2d17da4fdb5c2))
* **ip_dns_forwarders:** Add new resource `routeros_ip_dns_forwarders` ([48be482](https://github.com/terraform-routeros/terraform-provider-routeros/commit/48be482cb6f562ae156404f8755c3f8aa4fb3778)), closes [#778](https://github.com/terraform-routeros/terraform-provider-routeros/issues/778)
* **routing_igmp_proxy_interface:** Add support for IGMP Proxy ([4db8ffa](https://github.com/terraform-routeros/terraform-provider-routeros/commit/4db8ffad0d8233241cfad4b9321838c811b6e3a7)), closes [#773](https://github.com/terraform-routeros/terraform-provider-routeros/issues/773)

### Bug Fixes

* `last_modified` missing on `routeros_files` datasource ([ffa72d3](https://github.com/terraform-routeros/terraform-provider-routeros/commit/ffa72d354291d50baef5ea5cf1ba0a0b70a6773e)), closes [#790](https://github.com/terraform-routeros/terraform-provider-routeros/issues/790)
* **system:** Add missing default_mount_point_template attribute ([5e5bc6e](https://github.com/terraform-routeros/terraform-provider-routeros/commit/5e5bc6ee37efb8e0fedfa25a5940841cd0a4fdeb))
* **tool_email:** Add new fields ([c15047a](https://github.com/terraform-routeros/terraform-provider-routeros/commit/c15047a3d57d60a3de891d8ad7fdaeb9f1164ee0)), closes [#779](https://github.com/terraform-routeros/terraform-provider-routeros/issues/779)

## [1.85.4](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.85.3...v1.85.4) (2025-07-25)

### Bug Fixes

* **address-list:** "context deadline exceeded" when working with firewall address-lists ([b450160](https://github.com/terraform-routeros/terraform-provider-routeros/commit/b450160ba65275aaad91815e71a5030ca29b61f6)), closes [#772](https://github.com/terraform-routeros/terraform-provider-routeros/issues/772)
* routeros_ipv6_nd_prefix data source schema is missing "invalid" field ([e76a987](https://github.com/terraform-routeros/terraform-provider-routeros/commit/e76a98774da428aa27bb968c674b0de41983264f))

## [1.85.3](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.85.2...v1.85.3) (2025-06-06)

### Bug Fixes

* **bgp:** Add new filtering attributes ([bb8878d](https://github.com/terraform-routeros/terraform-provider-routeros/commit/bb8878d02ed043ebd1980bff125782ca2bd4999b))
* **interface-bonding:** Add a `lacp_mode` attribute ([684e81b](https://github.com/terraform-routeros/terraform-provider-routeros/commit/684e81ba2e2f1e44b04d49297a6edcd5b5aeb5f7))
* **interface-gre6:** Correct local IPv6 address validation ([30da790](https://github.com/terraform-routeros/terraform-provider-routeros/commit/30da790bfc372dc5258853ecd3c281e68a08e13e)), closes [#764](https://github.com/terraform-routeros/terraform-provider-routeros/issues/764)
* **ip-dhcp-client:** Correct the `default_route_tables` attribute type to TypeSet ([93fe11c](https://github.com/terraform-routeros/terraform-provider-routeros/commit/93fe11c04a4ba27e0b50c45aec42c721aa293ff3))
* **ip-dhcp-server:** Add an `use_reconfigure` attribute ([db9bc30](https://github.com/terraform-routeros/terraform-provider-routeros/commit/db9bc308886e7cba71dbd037bda2bcf9c3c0db25)), closes [#749](https://github.com/terraform-routeros/terraform-provider-routeros/issues/749)
* **ipv6-dhcp-client:** Add new attributes ([b8a1ff8](https://github.com/terraform-routeros/terraform-provider-routeros/commit/b8a1ff8b097b99cd0e5ed1ee738bde42563862fd))
* **tool-sniffer:** Add a `max_packet_size` attribute ([c0d06e0](https://github.com/terraform-routeros/terraform-provider-routeros/commit/c0d06e07c518342724ebfc668862e974e90d868f))
* **wifi-channel:** Add a `reselect-time` attribute ([28184c9](https://github.com/terraform-routeros/terraform-provider-routeros/commit/28184c91be5c50c9f9256d34f0c6858145c6c507))
* **wifi-datapath:** Add a `traffic_processing` attribute ([1de9d37](https://github.com/terraform-routeros/terraform-provider-routeros/commit/1de9d378b20bcd25adfcb3a32921a4407aa496f4))

## [1.85.2](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.85.1...v1.85.2) (2025-06-04)

### Bug Fixes

* **bridge_port:** Add new attributes (ROS 7.19) ([98c0dbf](https://github.com/terraform-routeros/terraform-provider-routeros/commit/98c0dbf479398a6f94c56530ddc05953d1123ba1))
* **bridge_port:** Skip counters ([3839756](https://github.com/terraform-routeros/terraform-provider-routeros/commit/38397561413974fc5616053ba8bc3de31fc50201))
* **bridge-mlag:** Add missing attributes to routeros_bridge_mlag ([47e1d26](https://github.com/terraform-routeros/terraform-provider-routeros/commit/47e1d267227a45bbda5ff666da30944e40737a85)), closes [#755](https://github.com/terraform-routeros/terraform-provider-routeros/issues/755)
* **ip-service:** routeros_ip_service lead to Warning: Field 'proto' not found in the schema, Warning: Field 'dynamic' not found in the schema ([85dff9b](https://github.com/terraform-routeros/terraform-provider-routeros/commit/85dff9b58859a518d556ec2e015c613e5d2593c1)), closes [#756](https://github.com/terraform-routeros/terraform-provider-routeros/issues/756)

## [1.85.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.85.0...v1.85.1) (2025-05-27)

### Bug Fixes

* **dhcp-client:** Add new attributes ([#753](https://github.com/terraform-routeros/terraform-provider-routeros/issues/753)) ([b2b8db0](https://github.com/terraform-routeros/terraform-provider-routeros/commit/b2b8db0b74fc1bc42a241fc7a6bc05137048bdb4)), closes [#749](https://github.com/terraform-routeros/terraform-provider-routeros/issues/749)
* **firewall-v4:** Add Set/Unset attributes ([9886d39](https://github.com/terraform-routeros/terraform-provider-routeros/commit/9886d39b400b83521244f976b20fd2eb2da1844d)), closes [#748](https://github.com/terraform-routeros/terraform-provider-routeros/issues/748)
* **wg-peers:** Fix drift attributes ([ba95824](https://github.com/terraform-routeros/terraform-provider-routeros/commit/ba95824334b192768be33c5337be9e8ff104f04c)), closes [#750](https://github.com/terraform-routeros/terraform-provider-routeros/issues/750)

## [1.85.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.84.0...v1.85.0) (2025-05-15)

### Features

* **GRE v6:** Add new resource `routeros_interface_gre6` ([034cd10](https://github.com/terraform-routeros/terraform-provider-routeros/commit/034cd10e11cc6bb2f7ac5b2ff1158b8dfec55e60)), closes [#737](https://github.com/terraform-routeros/terraform-provider-routeros/issues/737)

### Bug Fixes

* Fix IPSec resource input validation ([72063f5](https://github.com/terraform-routeros/terraform-provider-routeros/commit/72063f53af6a8108719d8137ab3284439d6e3beb))
* **logging_action changed:**  Update for ROS 7.18 ([129f5fe](https://github.com/terraform-routeros/terraform-provider-routeros/commit/129f5feef2c349938cc09da5c952dca354aaf923)), closes [#735](https://github.com/terraform-routeros/terraform-provider-routeros/issues/735)

## [1.84.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.83.1...v1.84.0) (2025-05-06)

### Features

* Implement duration check ([bf1a21d](https://github.com/terraform-routeros/terraform-provider-routeros/commit/bf1a21d89a98389a040bf296a3f261bfc365aae4))
* **ipv6:** Add new resource `routeros_ipv6_nd_prefix` ([476a70b](https://github.com/terraform-routeros/terraform-provider-routeros/commit/476a70bf59079c5c79582704ed648f40b9c10c9e)), closes [#730](https://github.com/terraform-routeros/terraform-provider-routeros/issues/730)

## [1.83.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.83.0...v1.83.1) (2025-04-28)

### Bug Fixes

* **radius:** Remove server IP address validator ([6c486dd](https://github.com/terraform-routeros/terraform-provider-routeros/commit/6c486ddfaf71ee36271445d3b90adf3a6c8acafa)), closes [#726](https://github.com/terraform-routeros/terraform-provider-routeros/issues/726)
* **sshkey:** Add `ForceNew` for the `key` attribute ([23f7c82](https://github.com/terraform-routeros/terraform-provider-routeros/commit/23f7c824b8069ce7fe13844aa0dab48fe7b1bb3c)), closes [#725](https://github.com/terraform-routeros/terraform-provider-routeros/issues/725)

## [1.83.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.82.0...v1.83.0) (2025-04-16)

### Features

* **script:** Add a script start attribute ([e379996](https://github.com/terraform-routeros/terraform-provider-routeros/commit/e3799964205e4fd1623f7576604b7232ec454708))

## [1.82.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.81.2...v1.82.0) (2025-04-15)

### Features

* **queue:** Add new resource `routeros_queue_tree` ([54fe8a2](https://github.com/terraform-routeros/terraform-provider-routeros/commit/54fe8a274be5ea3ab89d28c6d1dbd624af85650a))

## [1.81.2](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.81.1...v1.81.2) (2025-04-12)

### Bug Fixes

* **ospf_area:** NSSA Translator/Translate Schema Error ([1348d32](https://github.com/terraform-routeros/terraform-provider-routeros/commit/1348d32589d3b508faa8e433ad3d925bb11cc012)), closes [#718](https://github.com/terraform-routeros/terraform-provider-routeros/issues/718)
* routeros_interfaces data source schema is missing "inactive" field ([aeb11b4](https://github.com/terraform-routeros/terraform-provider-routeros/commit/aeb11b4d599a78c2837016bca943fd92a126e93a))

## [1.81.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.81.0...v1.81.1) (2025-04-07)

### Bug Fixes

* Resource `routeros_ip_route` schema is missing 2 fields ([c01edad](https://github.com/terraform-routeros/terraform-provider-routeros/commit/c01edad8a49130f154b1c9d5431bbe53415378f8)), closes [#712](https://github.com/terraform-routeros/terraform-provider-routeros/issues/712)

## [1.81.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.80.0...v1.81.0) (2025-03-30)

### Features

* **ipv6_firewall_filter:** add action fasttrack-connection ([e0ca80c](https://github.com/terraform-routeros/terraform-provider-routeros/commit/e0ca80c93cedf8475f3816b3c52496d4ac3293f5))

## [1.80.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.79.0...v1.80.0) (2025-03-21)

### Features

* **ip-settings:** Add new resource `routeros_ip_settings` ([390f577](https://github.com/terraform-routeros/terraform-provider-routeros/commit/390f5775059a77eb7eea3e9b778251cffcf68275)), closes [#704](https://github.com/terraform-routeros/terraform-provider-routeros/issues/704)

### Bug Fixes

* routeros_ip_hotspot_walled_garden_ip action validation may be incorrect ([3529224](https://github.com/terraform-routeros/terraform-provider-routeros/commit/3529224f637dcdaf713fa55529df2b84d7467a8d)), closes [#708](https://github.com/terraform-routeros/terraform-provider-routeros/issues/708)

## [1.79.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.78.0...v1.79.0) (2025-03-18)

### Features

* **ip:** Add /ip/smb resource ([d05f4c0](https://github.com/terraform-routeros/terraform-provider-routeros/commit/d05f4c0a2c0b35e17ccca041047d5f0ce10d0b28))
* **system:** Add `/disk/settings` resource ([#697](https://github.com/terraform-routeros/terraform-provider-routeros/issues/697)) ([4193991](https://github.com/terraform-routeros/terraform-provider-routeros/commit/4193991203e54079d1dd8e0d1874965313163cc5))
* **tool:** Add /tool/graphing resources ([722108d](https://github.com/terraform-routeros/terraform-provider-routeros/commit/722108d2e16cab7e0cbcef7898899310c15f4113))
* **tool:** Add /tool/mac-server/ping resource ([31f9c50](https://github.com/terraform-routeros/terraform-provider-routeros/commit/31f9c509cfea8bf18dd852521c7f8f1297ef3c46))

### Bug Fixes

* **hotspot:** Add "radius_location_id" attribute to routeros_ip_hotspot_profile resource ([6d2dcbc](https://github.com/terraform-routeros/terraform-provider-routeros/commit/6d2dcbc768b09d68f0deb858cb39eba38cbdbcc4)), closes [#701](https://github.com/terraform-routeros/terraform-provider-routeros/issues/701)

## [1.78.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.77.3...v1.78.0) (2025-03-17)

### Features

* **system:** Add /system/note resource ([6e588b3](https://github.com/terraform-routeros/terraform-provider-routeros/commit/6e588b319ecea978178cfc3fdc10af2dd74f7b1f))

## [1.77.3](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.77.2...v1.77.3) (2025-03-12)

### Bug Fixes

* **ethernet:** Ethernet `advertise` do not allow multiple values ([73823e5](https://github.com/terraform-routeros/terraform-provider-routeros/commit/73823e50db10cab5570d5dfac7d7fe1da7bee5d5)), closes [#692](https://github.com/terraform-routeros/terraform-provider-routeros/issues/692)
* **ipv6-settings:** Add new attributes (ROS 7.18) ([32c8820](https://github.com/terraform-routeros/terraform-provider-routeros/commit/32c882048e875745e7f2c7105618933ce0d1b7f0))
* **ipv6-settings:** Skip metrics ([a8b3b33](https://github.com/terraform-routeros/terraform-provider-routeros/commit/a8b3b3388176de5189206a681516d96be28d1a3c))

## [1.77.2](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.77.1...v1.77.2) (2025-03-05)

### Bug Fixes

* **ip-cloud:** IP Cloud ddns_enabled 7.17 and newer ([d341c29](https://github.com/terraform-routeros/terraform-provider-routeros/commit/d341c29a2f400319b5dfc6891e8acbd46a0d312f)), closes [#686](https://github.com/terraform-routeros/terraform-provider-routeros/issues/686)
* **ipv6-dhcp-client:** Add a new attribute (ROS 7.18) ([bcbd262](https://github.com/terraform-routeros/terraform-provider-routeros/commit/bcbd26240fbaf461d51e33bd57d8f889c23fe2e5))
* **ipv6-dhcp-client:** replace Default parameter with DiffSuppressFunc: AlwaysPresentNotUserProvided ([d5aad3b](https://github.com/terraform-routeros/terraform-provider-routeros/commit/d5aad3b66b03c01340295b151ca76ce927e5b26d))

## [1.77.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.77.0...v1.77.1) (2025-03-04)

### Bug Fixes

* **bandwidth-server:** Add new attributes (ROS 7.18) ([83d4bda](https://github.com/terraform-routeros/terraform-provider-routeros/commit/83d4bda0f738d404af67d92ee565e7826a08b77e)), closes [#683](https://github.com/terraform-routeros/terraform-provider-routeros/issues/683)
* **dhcp-client:** Add a new attribute (ROS 7.18) ([253483c](https://github.com/terraform-routeros/terraform-provider-routeros/commit/253483cadea0cb9f7781e3c5727777d46792aa3a)), closes [#684](https://github.com/terraform-routeros/terraform-provider-routeros/issues/684)

## [1.77.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.76.7...v1.77.0) (2025-03-02)

### Features

* **system:** Add user ssh-keys resource ([7c15569](https://github.com/terraform-routeros/terraform-provider-routeros/commit/7c15569b364162a1b1f37a7212ae29aa73a5c3d3))

### Bug Fixes

* **container:** Importing containers and container mounts (inconsistency and force recreate) ([af2585a](https://github.com/terraform-routeros/terraform-provider-routeros/commit/af2585a427479967aaf183980667f656455149ef)), closes [#652](https://github.com/terraform-routeros/terraform-provider-routeros/issues/652)
* **dhcp-server:** 'Bad Request', details: 'unknown parameter src-address' ([3eb7c44](https://github.com/terraform-routeros/terraform-provider-routeros/commit/3eb7c44b07576a09bf4a164b7c46909c237c8dc1)), closes [#679](https://github.com/terraform-routeros/terraform-provider-routeros/issues/679)
* minutes short name is lower m ([0379a1a](https://github.com/terraform-routeros/terraform-provider-routeros/commit/0379a1ac2c491683d0401ef1ac8d4241807c178f))

## [1.76.7](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.76.6...v1.76.7) (2025-02-28)

### Bug Fixes

* **bridge:** Add a `dynamic` attribute (ROS 7.18) ([cccee12](https://github.com/terraform-routeros/terraform-provider-routeros/commit/cccee12d85df6dae22e5d671d7a13c14acbb3025)), closes [#671](https://github.com/terraform-routeros/terraform-provider-routeros/issues/671)
* **ip-pool:** Ignore new computed attributes (ROS 7.18) ([060d96e](https://github.com/terraform-routeros/terraform-provider-routeros/commit/060d96ecb71344797a8b0ee05a37f325bcfb7c99)), closes [#673](https://github.com/terraform-routeros/terraform-provider-routeros/issues/673)
* **ipv6:** Add a `auto_link_local` attribute (ROS 7.18) ([c162cb0](https://github.com/terraform-routeros/terraform-provider-routeros/commit/c162cb0627f4d8487d33eb9169d85b489336e8a2)), closes [#672](https://github.com/terraform-routeros/terraform-provider-routeros/issues/672)
* **logging:** Add new attributes (ROS 7.18) ([337bd74](https://github.com/terraform-routeros/terraform-provider-routeros/commit/337bd741f22d3494289471d9d4c6b1e4cd5f0224)), closes [#674](https://github.com/terraform-routeros/terraform-provider-routeros/issues/674)
* **neighbor-discovery:** Add a `lldp_dcbx` attribute. ([0f3ea69](https://github.com/terraform-routeros/terraform-provider-routeros/commit/0f3ea6902e4e27882bf590a21e00b94b35023a7a)), closes [#670](https://github.com/terraform-routeros/terraform-provider-routeros/issues/670)

## [1.76.6](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.76.5...v1.76.6) (2025-02-19)

### Bug Fixes

* Importing certificate and key without `routeros_file` resource ([8d3374e](https://github.com/terraform-routeros/terraform-provider-routeros/commit/8d3374e74fb2e25983fa58d4ef80dedda823396b)), closes [#660](https://github.com/terraform-routeros/terraform-provider-routeros/issues/660)

## [1.76.5](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.76.4...v1.76.5) (2025-02-17)

### Bug Fixes

* **API:** Error: unknown RouterOS reply word: `!empty` (ROS 7.18) ([e1660aa](https://github.com/terraform-routeros/terraform-provider-routeros/commit/e1660aa3efdaebef3b1526b2af084adc8c75f515)), closes [#661](https://github.com/terraform-routeros/terraform-provider-routeros/issues/661)

## [1.76.4](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.76.3...v1.76.4) (2025-02-16)

### Bug Fixes

* **wifi_security:** ft_mobility_domain is a (hex) string, not an integer ([497e430](https://github.com/terraform-routeros/terraform-provider-routeros/commit/497e4304192979d5f6d25b7220457e16dc3e1344)), closes [#662](https://github.com/terraform-routeros/terraform-provider-routeros/issues/662)

## [1.76.3](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.76.2...v1.76.3) (2025-02-13)

### Bug Fixes

* Modify regex for RouterOS version ([b61d1db](https://github.com/terraform-routeros/terraform-provider-routeros/commit/b61d1db79433895027ae63002311ba062099c251))
* **routeros_version:** Rename the attribute ([51a4c7c](https://github.com/terraform-routeros/terraform-provider-routeros/commit/51a4c7cbb6b198c53742b382741f74547d67866d)), closes [#658](https://github.com/terraform-routeros/terraform-provider-routeros/issues/658)

## [1.76.2](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.76.1...v1.76.2) (2025-02-12)

### Bug Fixes

* Fix the typo ([fa85c7b](https://github.com/terraform-routeros/terraform-provider-routeros/commit/fa85c7b624735dc618333bb251249d34e8bc5404))
* **queue:** Fix byte attribute comparison error ([9437e77](https://github.com/terraform-routeros/terraform-provider-routeros/commit/9437e77e8ada48393357a60ddcce4c7f064e6826)), closes [#643](https://github.com/terraform-routeros/terraform-provider-routeros/issues/643)

## [1.76.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.76.0...v1.76.1) (2025-02-11)

### Bug Fixes

* **provider:** Add attribute drift handling between ROS versions ([32141a6](https://github.com/terraform-routeros/terraform-provider-routeros/commit/32141a622f24e903fa3e8601a16498714ac9573d)), closes [#654](https://github.com/terraform-routeros/terraform-provider-routeros/issues/654)

## [1.76.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.75.0...v1.76.0) (2025-02-04)

### Features

* **queue:** Add new resource `routeros_queue_simple` ([b63640c](https://github.com/terraform-routeros/terraform-provider-routeros/commit/b63640ca0dea66e8d7863652508d6fa065f1ef64)), closes [#643](https://github.com/terraform-routeros/terraform-provider-routeros/issues/643)
* **queue:** Add new resource `routeros_queue_type` ([45aadfb](https://github.com/terraform-routeros/terraform-provider-routeros/commit/45aadfbf6ab6cffc6c960399c9bf584e975e9466)), closes [#643](https://github.com/terraform-routeros/terraform-provider-routeros/issues/643)

### Bug Fixes

* **dhcp:** Update for ROS 7.17 ([9b6a238](https://github.com/terraform-routeros/terraform-provider-routeros/commit/9b6a238b5bb9ceae1dbaeb45215a1fe804f005b8)), closes [#644](https://github.com/terraform-routeros/terraform-provider-routeros/issues/644)
* Missing fields in the schema with version 7.17.1 ([8199e43](https://github.com/terraform-routeros/terraform-provider-routeros/commit/8199e4339f7e5917c25d28c7028f940a16d52df8)), closes [#650](https://github.com/terraform-routeros/terraform-provider-routeros/issues/650)
* **w60:** Fix bugs after the first testing ([b1264e7](https://github.com/terraform-routeros/terraform-provider-routeros/commit/b1264e73cdd48d75d4a98e755fe38817eac4d6c9)), closes [#618](https://github.com/terraform-routeros/terraform-provider-routeros/issues/618)

## [1.75.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.74.0...v1.75.0) (2025-01-17)

### Features

* add `resource_interface_sstp_*` resources ([219153c](https://github.com/terraform-routeros/terraform-provider-routeros/commit/219153c1b36a268af396d122c66cde75096ce8ff))

### Bug Fixes

* datasource routeros_interface_bridge_filter ([00c3ad3](https://github.com/terraform-routeros/terraform-provider-routeros/commit/00c3ad3fc59a6e25917d59a958db01e1d83d822d))
* **ip_dhcp_server_lease:** class_id should be computed ([c11c3c4](https://github.com/terraform-routeros/terraform-provider-routeros/commit/c11c3c462f10ad308f27afb020c49a7f616e68b8))
* **ipv6_firewall_net:** ipv6_firewall_nat has to_address not to_addresses ([c77d027](https://github.com/terraform-routeros/terraform-provider-routeros/commit/c77d027ff196a8a72e0fee26a8b9b14ebd0089b0))
* Remove DiffSuppress for `multi_passphrase_group` ([367d5e4](https://github.com/terraform-routeros/terraform-provider-routeros/commit/367d5e47e3e9b1aedfa63b6e5fafc0eb440734ba)), closes [#630](https://github.com/terraform-routeros/terraform-provider-routeros/issues/630)

## [1.74.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.73.0...v1.74.0) (2025-01-03)

### Features

* **l2tp:** Add new resource `routeros_interface_l2tp_client` ([f2d3f9c](https://github.com/terraform-routeros/terraform-provider-routeros/commit/f2d3f9c372774c69f6edb9e55df5699a3117390f)), closes [#629](https://github.com/terraform-routeros/terraform-provider-routeros/issues/629)

### Bug Fixes

* use DiffSuppressFunc: AlwaysPresentNotUserProvided ([931752c](https://github.com/terraform-routeros/terraform-provider-routeros/commit/931752c20b9b5caaec92391457150dbc001786c0))

## [1.73.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.72.0...v1.73.0) (2024-12-30)

### Features

* **wifi:** Add new resource `routeros_wifi_security_multi_passphrase` ([4ef710a](https://github.com/terraform-routeros/terraform-provider-routeros/commit/4ef710a25713aa9b5b6faa15eb164f8f9e152205)), closes [#621](https://github.com/terraform-routeros/terraform-provider-routeros/issues/621)
* **wifi:** Add new resources `routeros_interface_w60g`, `routeros_interface_w60g_station` ([746a223](https://github.com/terraform-routeros/terraform-provider-routeros/commit/746a2235ed90b76eab38f0a8f8ef375d96cde5c9)), closes [#618](https://github.com/terraform-routeros/terraform-provider-routeros/issues/618)

### Bug Fixes

* **bridge_filter:** Missing field in interface_bridge_filter ([63474c8](https://github.com/terraform-routeros/terraform-provider-routeros/commit/63474c834b520591602299e13ae0f7c783367dc8)), closes [#620](https://github.com/terraform-routeros/terraform-provider-routeros/issues/620)
* **ip_dhcp_relay:** Field 'dhcp_server_vrf' not found in the schema ([f2f12b0](https://github.com/terraform-routeros/terraform-provider-routeros/commit/f2f12b0603685795097b32aa9c4bb4344e44f200)), closes [#622](https://github.com/terraform-routeros/terraform-provider-routeros/issues/622)
* **wifi:** Using `5ghz-an` for `routeros_wifi_channel` returns status code `400` ([d603682](https://github.com/terraform-routeros/terraform-provider-routeros/commit/d603682702b60e56f079092b0cb98c1f008f38ea)), closes [#619](https://github.com/terraform-routeros/terraform-provider-routeros/issues/619)

## [1.72.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.71.0...v1.72.0) (2024-12-22)

### Features

* **ipv6:** new resource routeros_ipv6_firewall_mangle ([e6a0c1e](https://github.com/terraform-routeros/terraform-provider-routeros/commit/e6a0c1e663f6a6e0ed447e6d2d46167562afe9a0))
* **ipv6:** new resource routeros_ipv6_firewall_nat ([03c2587](https://github.com/terraform-routeros/terraform-provider-routeros/commit/03c2587c8be7583ff122fafeb96edfe6be489d35))
* **pppoe:** Add new resource `routeros_interface_pppoe_server` ([f04afd1](https://github.com/terraform-routeros/terraform-provider-routeros/commit/f04afd1993c22a0010f60d2cd1dbd7150d854f84)), closes [#617](https://github.com/terraform-routeros/terraform-provider-routeros/issues/617)

### Bug Fixes

* **ipv6_dhcp_server_option:** Add Hotspot 2.0 fields ([1d64d8c](https://github.com/terraform-routeros/terraform-provider-routeros/commit/1d64d8c04adf9204bb554da2ec8bec9521203db0)), closes [#605](https://github.com/terraform-routeros/terraform-provider-routeros/issues/605)

## [1.71.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.70.0...v1.71.0) (2024-12-02)

### Features

* Add interface bridge filter ([#608](https://github.com/terraform-routeros/terraform-provider-routeros/issues/608)) ([441f11c](https://github.com/terraform-routeros/terraform-provider-routeros/commit/441f11cc21f0637096f68bc1c22c90afb01dfb08))

### Bug Fixes

* routeros_ipv6_dhcp_client prefix_hint validation issue ([#611](https://github.com/terraform-routeros/terraform-provider-routeros/issues/611)) ([fb51382](https://github.com/terraform-routeros/terraform-provider-routeros/commit/fb51382092de81ad6846c980403a314cc479fbf7)), closes [#609](https://github.com/terraform-routeros/terraform-provider-routeros/issues/609)

## [1.70.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.69.2...v1.70.0) (2024-11-24)

### Features

* **wave2:** Add new resource `routeros_interface_wireless_connect_list` ([10078df](https://github.com/terraform-routeros/terraform-provider-routeros/commit/10078df6c341d20e2d7e94344bafbff224312828)), closes [#605](https://github.com/terraform-routeros/terraform-provider-routeros/issues/605)

## [1.69.2](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.69.1...v1.69.2) (2024-11-21)

### Bug Fixes

* **ds:** warning  Field 'revision' not found in the schema in routeros_system_routerboard datasource ([1b6f86b](https://github.com/terraform-routeros/terraform-provider-routeros/commit/1b6f86b28b5d566dbda32e6f0ed3cf95cacd251c)), closes [#602](https://github.com/terraform-routeros/terraform-provider-routeros/issues/602)

## [1.69.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.69.0...v1.69.1) (2024-11-18)

### Bug Fixes

* **ospf:** Change the “network” attribute ([f057f18](https://github.com/terraform-routeros/terraform-provider-routeros/commit/f057f184608e9de9c08e2f5b773e0ed4f4f740af)), closes [#598](https://github.com/terraform-routeros/terraform-provider-routeros/issues/598)

## [1.69.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.68.0...v1.69.0) (2024-11-14)

### Features

* **ipv6:** Add new resource `routeros_ipv6_settings` ([4410ddc](https://github.com/terraform-routeros/terraform-provider-routeros/commit/4410ddc9941257e9ff58b40d434d7b59b6ea886b)), closes [#596](https://github.com/terraform-routeros/terraform-provider-routeros/issues/596)

## [1.68.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.67.0...v1.68.0) (2024-11-13)

### Features

* **ipv6:** Add new resource `routeros_interface_6to4` ([742e6bd](https://github.com/terraform-routeros/terraform-provider-routeros/commit/742e6bd55fd4682e952bd35a7acece9968951978)), closes [#593](https://github.com/terraform-routeros/terraform-provider-routeros/issues/593)

### Bug Fixes

* **helpers:** Fix PropKeepaliveRw ([32ee5bc](https://github.com/terraform-routeros/terraform-provider-routeros/commit/32ee5bceffdd4c4ed66461ddfc7e469b95068de8))
* **helpers:** Fix the plan was not empty ([31d420f](https://github.com/terraform-routeros/terraform-provider-routeros/commit/31d420f4f60998830bc860ba1e0f38025cd241d1))
* **mangle-connection-state:** typo in `established` value ([#595](https://github.com/terraform-routeros/terraform-provider-routeros/issues/595)) ([19c6c97](https://github.com/terraform-routeros/terraform-provider-routeros/commit/19c6c974a9f03f3848e5662ccb73401a72d0f4cf))

## [1.67.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.66.0...v1.67.0) (2024-11-10)

### Features

* **ds:** Add new datasource `routeros_system_routerboard` ([2f78e94](https://github.com/terraform-routeros/terraform-provider-routeros/commit/2f78e94ceaf4c484376a5909159675371f3da90b)), closes [#588](https://github.com/terraform-routeros/terraform-provider-routeros/issues/588)
* **vxlan:** Add new resource `routeros_interface_vxlan_vteps` ([1e72222](https://github.com/terraform-routeros/terraform-provider-routeros/commit/1e72222b890de8ef8a72d2adad7d46138835a0c8)), closes [#590](https://github.com/terraform-routeros/terraform-provider-routeros/issues/590)
* **vxlan:** Add new resource `routeros_interface_vxlan` ([7ccb49f](https://github.com/terraform-routeros/terraform-provider-routeros/commit/7ccb49f5418b545afda724f933993e94f134324b))

## [1.66.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.65.1...v1.66.0) (2024-10-11)

### Features

* **ipv6:** Add new resource `routeros_ipv6_dhcp_server_option` ([c224bb1](https://github.com/terraform-routeros/terraform-provider-routeros/commit/c224bb192b090ac553cf67aade61e81aab9db93a))
* **ipv6:** Add new resource `routeros_ipv6_dhcp_server` ([23ec941](https://github.com/terraform-routeros/terraform-provider-routeros/commit/23ec9414b9c17a6acf483a602ad741135104b012))
* **ipv6:** Add new resource `routeros_ipv6_pool` ([96d1c48](https://github.com/terraform-routeros/terraform-provider-routeros/commit/96d1c48ed30efe0ff0d36e716bbac84282431e33))
* **ipv6:** New resource `routeros_ipv6_dhcp_server_option_sets` ([b92458b](https://github.com/terraform-routeros/terraform-provider-routeros/commit/b92458bf79cc9e987beb137df85e40d2b76c436b))

## [1.65.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.65.0...v1.65.1) (2024-10-09)

### Bug Fixes

* **ipsec:** Fix `auth_method` attribute validation ([#582](https://github.com/terraform-routeros/terraform-provider-routeros/issues/582)) ([98516b0](https://github.com/terraform-routeros/terraform-provider-routeros/commit/98516b07f1dbf99b9fa444f95b4a04ed61b5dcec)), closes [#581](https://github.com/terraform-routeros/terraform-provider-routeros/issues/581)

## [1.65.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.64.2...v1.65.0) (2024-10-07)

### Features

* Disable warning output on system resources ([c799f29](https://github.com/terraform-routeros/terraform-provider-routeros/commit/c799f299629fd82c3004dd1903487d694aeffbbf))
* **ipsec:** Add new resource `routeros_ip_ipsec_identity` ([afdbadb](https://github.com/terraform-routeros/terraform-provider-routeros/commit/afdbadb7297560dc97d07cbd0abe8d1f7fa0fb9f))
* **ipsec:** Add new resource `routeros_ip_ipsec_key` ([14865b9](https://github.com/terraform-routeros/terraform-provider-routeros/commit/14865b9fa989e79a07a9e46857399cf560208f61))
* **ipsec:** Add new resource `routeros_ip_ipsec_mode_config` ([ca88a77](https://github.com/terraform-routeros/terraform-provider-routeros/commit/ca88a77f14bd5177938d20930da4ccd515e647fa))
* **ipsec:** Add new resource `routeros_ip_ipsec_policy_group` ([d4c0817](https://github.com/terraform-routeros/terraform-provider-routeros/commit/d4c0817f6af31a3b5d1fa766bf9c4c415b43aefd))
* **ipsec:** Add new resource `routeros_ip_ipsec_policy` ([9ba2bf9](https://github.com/terraform-routeros/terraform-provider-routeros/commit/9ba2bf961a0d7dafaf5d778a10f0a2e153f0d666))
* **ipsec:** Add new resource `routeros_ip_ipsec_profile` ([66aa2f8](https://github.com/terraform-routeros/terraform-provider-routeros/commit/66aa2f8c830a4cfe925a315c36114ded9842264d))
* **ipsec:** Add new resource `routeros_ip_ipsec_proposal` ([9fee803](https://github.com/terraform-routeros/terraform-provider-routeros/commit/9fee803c5a19ad244d63ff8f052dc6c9de90c47f))
* **ipsec:** Add new resource `routeros_ip_ipsec_settings` ([7388cae](https://github.com/terraform-routeros/terraform-provider-routeros/commit/7388cae3a5753ef9c0dfbf5c2423497ee9eba838))
* **ipsec:** New resource `routeros_ip_ipsec_peer` ([7600d45](https://github.com/terraform-routeros/terraform-provider-routeros/commit/7600d457e5e4245c9c743f2f28d42194c2a0f0e8))

### Bug Fixes

* Adding fields to skip for rx and tx on resource_interface_ethernet.go ([#573](https://github.com/terraform-routeros/terraform-provider-routeros/issues/573)) ([1a905f5](https://github.com/terraform-routeros/terraform-provider-routeros/commit/1a905f5e26afbd782bbd9f67c18e49abe6937ea1))
* **file:** Field 'last_modified' not found in the schema ([#580](https://github.com/terraform-routeros/terraform-provider-routeros/issues/580)) ([cb4635a](https://github.com/terraform-routeros/terraform-provider-routeros/commit/cb4635a97b50b1e33c82622ee3f3a3defe1e6b69)), closes [#579](https://github.com/terraform-routeros/terraform-provider-routeros/issues/579)
* **ipsec:** Add the lost attributes ([6f61879](https://github.com/terraform-routeros/terraform-provider-routeros/commit/6f61879ef8e84254fbf0a602ec8a09b714bc58be))
* validation for routeros_wifi_security.wps ([#578](https://github.com/terraform-routeros/terraform-provider-routeros/issues/578)) ([770bfe3](https://github.com/terraform-routeros/terraform-provider-routeros/commit/770bfe37c5708227daa3a891581d86f09c9c4b54))
* **wireless:** Delete required parameters ([616049e](https://github.com/terraform-routeros/terraform-provider-routeros/commit/616049e04bd351111c7515e801b3e9e229dc3da8))
* **wireless:** Delete required parameters ([33793fd](https://github.com/terraform-routeros/terraform-provider-routeros/commit/33793fd70aef7580862b577cd6f001cf9d954306))

## [1.64.2](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.64.1...v1.64.2) (2024-09-29)

### Bug Fixes

* Unable to remove routeros_interface_wireless from config ([#571](https://github.com/terraform-routeros/terraform-provider-routeros/issues/571)) ([8990eb0](https://github.com/terraform-routeros/terraform-provider-routeros/commit/8990eb05c63859b0d58be996d7d2a886868ab13a)), closes [#570](https://github.com/terraform-routeros/terraform-provider-routeros/issues/570)

## [1.64.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.64.0...v1.64.1) (2024-09-28)

### Bug Fixes

* **no-release:** Update resource_ipv6_address.go ([#568](https://github.com/terraform-routeros/terraform-provider-routeros/issues/568)) ([587f527](https://github.com/terraform-routeros/terraform-provider-routeros/commit/587f527c532ffabfe404508512ab5b9185f68695))
* Update resource_ip_address.go ([#569](https://github.com/terraform-routeros/terraform-provider-routeros/issues/569)) ([25a6496](https://github.com/terraform-routeros/terraform-provider-routeros/commit/25a64968c0bb8279f599859af371309fcbca1871))

## [1.64.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.63.1...v1.64.0) (2024-09-27)

### Features

* **wireless:** Add new resource `routeros_interface_wireless_access_list` ([cd82592](https://github.com/terraform-routeros/terraform-provider-routeros/commit/cd825923390b489e41463b71e6ea253286f28c25))
* **wireless:** Add new resource `routeros_interface_wireless_security_profiles` ([ef40fd2](https://github.com/terraform-routeros/terraform-provider-routeros/commit/ef40fd25f448c3ae1d92d2c2a60fcdde782f5df4))
* **wireless:** Add new resource `routeros_interface_wireless` ([15c2650](https://github.com/terraform-routeros/terraform-provider-routeros/commit/15c265088c024a5b770fda1990ea3713c724d253))

### Bug Fixes

* **serialize:** Fix `PropTransformSet` ([8baebae](https://github.com/terraform-routeros/terraform-provider-routeros/commit/8baebae070e047a3dad3f86aa72cf502f053d781))
* **serialize:** Fix the transformation of attribute names. ([9796820](https://github.com/terraform-routeros/terraform-provider-routeros/commit/9796820aa413ad6f45731b613a1cfa88fc124e0a))
* **tool_sniffer:** Add resource state control ([eb53e45](https://github.com/terraform-routeros/terraform-provider-routeros/commit/eb53e456ffce679464bb9023df59056663b9c6e9))
* **wireless_security_profile:** Add `Sensitive` flag to attributes ([909b4c7](https://github.com/terraform-routeros/terraform-provider-routeros/commit/909b4c792e169eec5c60619156be8bd10dfab7df))
* **wireless:** Add missing field ([c112740](https://github.com/terraform-routeros/terraform-provider-routeros/commit/c112740a660362528f9d53037e97e59c7138441f))

## [1.63.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.63.0...v1.63.1) (2024-09-26)

### Features

* **no-release:** Add QR code generation for WiFi ([#564](https://github.com/terraform-routeros/terraform-provider-routeros/issues/564)) ([9c39ae2](https://github.com/terraform-routeros/terraform-provider-routeros/commit/9c39ae2b52e62dfc7b3451887ee726a137b1c64c))

### Bug Fixes

* Missing fields in `routeros_system_user ` (introduced in 7.16) ([3ae2e10](https://github.com/terraform-routeros/terraform-provider-routeros/commit/3ae2e1030a7e30e935da4cdc316028c2abd3c91b)), closes [#560](https://github.com/terraform-routeros/terraform-provider-routeros/issues/560)
* **no-release:** Add field introduced in 7.16 `routeros_ip_address` ([de72d8e](https://github.com/terraform-routeros/terraform-provider-routeros/commit/de72d8e6c421dc4c4dee36c0a45d207a8d232c9c))
* **no-release:** Update `datasource_ipv6_addresses.go` ROS 7.16 ([0c0306b](https://github.com/terraform-routeros/terraform-provider-routeros/commit/0c0306b975c6be0314ea773b6ccc6f4acbfbe04d))
* **no-release:** Update `resource_system_logging.go` ROS 7.16 ([884addc](https://github.com/terraform-routeros/terraform-provider-routeros/commit/884addca655d60e91b74bd9cc507bb17cc55e6ca))

## [1.63.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.62.0...v1.63.0) (2024-09-24)


### Features

* Add new resource `routeros_tool_sniffer` ([f046966](https://github.com/terraform-routeros/terraform-provider-routeros/commit/f0469663fe7ea395825dc92c6a246e09dcd0f81c))
* **hotspot:** Add new resource `routeros_ip_hotspot_ip_binding` ([f2e27b4](https://github.com/terraform-routeros/terraform-provider-routeros/commit/f2e27b4732f863400c6f6c7f8341315019623f47))
* **hotspot:** Add new resource `routeros_ip_hotspot_profile` ([06b974b](https://github.com/terraform-routeros/terraform-provider-routeros/commit/06b974bdfbd6508a3f32a6b52456f0c8a6ba10b0))
* **hotspot:** Add new resource `routeros_ip_hotspot_service_port` ([153bf68](https://github.com/terraform-routeros/terraform-provider-routeros/commit/153bf68723a1def279431593ed5f5cc2bf2f1ddc))
* **hotspot:** Add new resource `routeros_ip_hotspot_user_profile` ([4de2db9](https://github.com/terraform-routeros/terraform-provider-routeros/commit/4de2db943c82de89e2081b8b32ffed8621f42ab9))
* **hotspot:** Add new resource `routeros_ip_hotspot_user` ([b897532](https://github.com/terraform-routeros/terraform-provider-routeros/commit/b897532232c12245b4c92f27ca129ebe8f6c8d31))
* **hotspot:** Add new resource `routeros_ip_hotspot_walled_garden_ip` ([92778ff](https://github.com/terraform-routeros/terraform-provider-routeros/commit/92778ff184124ce8d2d7dc0dfc266b29b6f306ef))
* **hotspot:** Add new resource `routeros_ip_hotspot_walled_garden` ([9c111ee](https://github.com/terraform-routeros/terraform-provider-routeros/commit/9c111eeefa0ce90e0bc2f02c9c2d9a4794487cbe))
* **hotspot:** Add new resource `routeros_ip_hotspot` ([1da8f3d](https://github.com/terraform-routeros/terraform-provider-routeros/commit/1da8f3d755ec0a35836b67037d57e40d27d3453f))


### Bug Fixes

* **dns_adlist:** Change an invalid resource name. ([e77bcbf](https://github.com/terraform-routeros/terraform-provider-routeros/commit/e77bcbf6f3bccfc3271269a4ad1fbff00038cfaf)), closes [#554](https://github.com/terraform-routeros/terraform-provider-routeros/issues/554)

## [1.62.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.61.2...v1.62.0) (2024-09-19)


### Features

* **resource_dns_adlist:** Add dns ad list support ([b1ca164](https://github.com/terraform-routeros/terraform-provider-routeros/commit/b1ca164c8e1b2b5f6e9ecca90abe62f84c5a488c)), closes [#554](https://github.com/terraform-routeros/terraform-provider-routeros/issues/554)

## [1.61.2](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.61.1...v1.61.2) (2024-09-02)


### Bug Fixes

* Add a validator for map name keys. ([059de30](https://github.com/terraform-routeros/terraform-provider-routeros/commit/059de305bcc829fb9be16c49c9903049f074cb13)), closes [#552](https://github.com/terraform-routeros/terraform-provider-routeros/issues/552)
* **data.routeros_interfaces:** Fix the schema ([f9489a2](https://github.com/terraform-routeros/terraform-provider-routeros/commit/f9489a23148c4bb96a1b8db76cc409a4a9bda00a)), closes [#550](https://github.com/terraform-routeros/terraform-provider-routeros/issues/550)
* **routeros_capsman_provisioning :** Change attributes type ([b9967d9](https://github.com/terraform-routeros/terraform-provider-routeros/commit/b9967d9c65f20da871cd6bc60597b49bd9d99d14)), closes [#551](https://github.com/terraform-routeros/terraform-provider-routeros/issues/551) [#551](https://github.com/terraform-routeros/terraform-provider-routeros/issues/551)
* **system_user_group:** Fix `policy` validator ([6a0a0bd](https://github.com/terraform-routeros/terraform-provider-routeros/commit/6a0a0bdd370680b0ac19172ff40bc555a08c82cc))

## [1.61.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.61.0...v1.61.1) (2024-08-27)


### Bug Fixes

* **routeros_system_user_group:** Fix change detection in sets with `SetUnset` values ([#547](https://github.com/terraform-routeros/terraform-provider-routeros/issues/547)) ([d51c4c2](https://github.com/terraform-routeros/terraform-provider-routeros/commit/d51c4c207e9f5df3979c0b090a34f24ad713c062)), closes [#544](https://github.com/terraform-routeros/terraform-provider-routeros/issues/544)

## [1.61.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.60.0...v1.61.0) (2024-08-23)


### Features

* Add `class_id` property support to the `routeros_ip_dhcp_server_lease` resource ([90033af](https://github.com/terraform-routeros/terraform-provider-routeros/commit/90033afb96247b4f9f70ea5974259cdcc16d9a08))
* Add `discover_interval` property support to the `routeros_ip_neighbor_discovery_settings` resource ([8233ace](https://github.com/terraform-routeros/terraform-provider-routeros/commit/8233ace372ea9dd9bbb0b9ad83dfbaf58c27a912))
* Add `forward_reserved_addresses` property support to the `routeros_interface_bridge` resource ([600e2af](https://github.com/terraform-routeros/terraform-provider-routeros/commit/600e2af88f0a8e6ed9c4fe0e8e3c0a73a7359f0b))
* Add `lldp_vlan_info` property support to the `routeros_ip_neighbor_discovery_settings` resource ([c2b07af](https://github.com/terraform-routeros/terraform-provider-routeros/commit/c2b07aff724758326f5551aa1d7efa87b72d16fa))
* Add `max_learned_entries` property support to the `routeros_interface_bridge` resource ([a8b4191](https://github.com/terraform-routeros/terraform-provider-routeros/commit/a8b4191e67ad2e399f18729352244898caebdbf3))
* Add `max_sessions` property support to the `routeros_ip_service` resource ([9007755](https://github.com/terraform-routeros/terraform-provider-routeros/commit/9007755b50236cb8676e191c9af4e6db8629f976))
* Add `mdns_repeat_ifaces` property support to the `routeros_ip_dns` resource ([fb71ff7](https://github.com/terraform-routeros/terraform-provider-routeros/commit/fb71ff71ad7392088f4ee21434e9b0fc833ed252))
* Add `slave_name_format` property support to the `routeros_wifi_provisioning` resource ([2ad0de6](https://github.com/terraform-routeros/terraform-provider-routeros/commit/2ad0de6ac0e10f431000022badec0c6595c8938c))
* Add the `comment` property support to the DHCP resources ([0d326e9](https://github.com/terraform-routeros/terraform-provider-routeros/commit/0d326e93ea28e4f7bf2d71a0399a116d759033e0))


### Bug Fixes

* Remove the default value for the `radio_mac` property in the `routeros_wifi_provisioning` resource ([a35d307](https://github.com/terraform-routeros/terraform-provider-routeros/commit/a35d307c659d973c5a7ab0e48f70a38d0e70d17e))
* **routeros_system_resource:** Add `bad_blocks` attribute to the ignored ones ([7b03141](https://github.com/terraform-routeros/terraform-provider-routeros/commit/7b03141f5fbbcbf56301149610dea93d1debdd82)), closes [#541](https://github.com/terraform-routeros/terraform-provider-routeros/issues/541)
* supress hw offload writable ([6c358eb](https://github.com/terraform-routeros/terraform-provider-routeros/commit/6c358eba75de36baf5653f79da653cd10ab26982)), closes [#540](https://github.com/terraform-routeros/terraform-provider-routeros/issues/540)

## [1.60.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.59.4...v1.60.0) (2024-08-19)


### Features

* Add the `routeros_system_led_settings` resource to manage the LED settings ([82434f8](https://github.com/terraform-routeros/terraform-provider-routeros/commit/82434f8c3d7f266dbd6624860d23019f29907975))
* Add the `routeros_system_led` resource to manage LEDs ([5131156](https://github.com/terraform-routeros/terraform-provider-routeros/commit/5131156a03608b531813cb7f340a1e5201d56187))
* Add the `routeros_system_routerboard_button_mode` resource to manage the mode button ([243dd9d](https://github.com/terraform-routeros/terraform-provider-routeros/commit/243dd9d0686643b1e02ba15699897228cc08fa7f))
* Add the `routeros_system_routerboard_button_reset` resource to manage the reset button ([358f640](https://github.com/terraform-routeros/terraform-provider-routeros/commit/358f640dd1f246c8a67af35fc5f1bdd01a8d7230))
* Add the `routeros_system_routerboard_button_wps` resource to manage the WPS button ([85db3b5](https://github.com/terraform-routeros/terraform-provider-routeros/commit/85db3b566a8641ac0256344e554a79ddb2fb39d0))
* Add the `routeros_system_routerboard_settings` resource to manage the RouterBOARD settings ([39a7e39](https://github.com/terraform-routeros/terraform-provider-routeros/commit/39a7e3957b46e47e4fd0c83da252eb7039c84f34))
* Add the `routeros_system_routerboard_usb` resource to manage the USB port ([4251b6a](https://github.com/terraform-routeros/terraform-provider-routeros/commit/4251b6a6552bd8a07edebab046f5f97a0d85300a))

## [1.59.4](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.59.3...v1.59.4) (2024-08-13)


### Bug Fixes

* **datasources:** Datasource int overflow ([#507](https://github.com/terraform-routeros/terraform-provider-routeros/issues/507)) ([aea6028](https://github.com/terraform-routeros/terraform-provider-routeros/commit/aea60289833ae74e15740447d36020d4485f72b2))

## [1.59.3](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.59.2...v1.59.3) (2024-08-12)


### Bug Fixes

* Add missing attributes, add fields to ignore. Fix wrong `new_dst_ports` type ([c570002](https://github.com/terraform-routeros/terraform-provider-routeros/commit/c570002b40171054b85b62a1e59b2dc907791100)), closes [#538](https://github.com/terraform-routeros/terraform-provider-routeros/issues/538)
* **netwatch:** Add `http_codes` attribute to skip fields. ([e7d0356](https://github.com/terraform-routeros/terraform-provider-routeros/commit/e7d03563143c4a70e7052cb8652513f375e4bfb3))

## [1.59.2](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.59.1...v1.59.2) (2024-08-07)


### Bug Fixes

* Can not set rtt parameters in `routeros_tool_netwatch` ([afa971e](https://github.com/terraform-routeros/terraform-provider-routeros/commit/afa971e337f6141c4dfa717c791807a52ffc0692)), closes [#535](https://github.com/terraform-routeros/terraform-provider-routeros/issues/535)

## [1.59.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.59.0...v1.59.1) (2024-08-07)


### Bug Fixes

* **bridge_port:** Fix the `priority` attribute type. ([4f342fb](https://github.com/terraform-routeros/terraform-provider-routeros/commit/4f342fb0fcedce97f0dbff806f0c60a6ed072b8b)), closes [#528](https://github.com/terraform-routeros/terraform-provider-routeros/issues/528)
* **firewall_raw:** "no track" action in routeros_ip_firewall_raw needs tweaking ([9d46a55](https://github.com/terraform-routeros/terraform-provider-routeros/commit/9d46a5503163b60534d63f414ddfd1b1a50ea631)), closes [#529](https://github.com/terraform-routeros/terraform-provider-routeros/issues/529)
* **wireguard_peer:** Need new filed 'is-responder' in resource 'routeros_interface_wireguard_peer' ([a31a394](https://github.com/terraform-routeros/terraform-provider-routeros/commit/a31a394e3abd8bcab3af06d6ab35d4b4f0abc89a)), closes [#530](https://github.com/terraform-routeros/terraform-provider-routeros/issues/530)

## [1.59.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.58.2...v1.59.0) (2024-08-05)


### Features

* Add `routeros_routing_rule` resource ([a263193](https://github.com/terraform-routeros/terraform-provider-routeros/commit/a26319308ae96b41fd0ffe291fb62073459378f7)), closes [#524](https://github.com/terraform-routeros/terraform-provider-routeros/issues/524)
* Add `routeros_tool_netwatch` resource ([2150241](https://github.com/terraform-routeros/terraform-provider-routeros/commit/215024137033233ab13a22ae731ef8a72f4a3612)), closes [#487](https://github.com/terraform-routeros/terraform-provider-routeros/issues/487)

## [1.58.2](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.58.1...v1.58.2) (2024-08-04)


### Bug Fixes

* **interface_lte:** Add missing attributes ([f67e8d9](https://github.com/terraform-routeros/terraform-provider-routeros/commit/f67e8d90242b0b1e0f9c630cf0ffd37a819bd049)), closes [#522](https://github.com/terraform-routeros/terraform-provider-routeros/issues/522)
* **ip_firewall:** Deleting routeros_ip_firewall_filter.in_interface tries to PATCH ([2189054](https://github.com/terraform-routeros/terraform-provider-routeros/commit/21890540901501e7f8da61eb23a452ad52791f44)), closes [#521](https://github.com/terraform-routeros/terraform-provider-routeros/issues/521)
* **ipv4_nat:** Need new filed 'randomise_ports' in action 'endpoint-independent-nat' ([51e161d](https://github.com/terraform-routeros/terraform-provider-routeros/commit/51e161d5b090762fed7bd6f638d36f9d31564ddc)), closes [#520](https://github.com/terraform-routeros/terraform-provider-routeros/issues/520)

## [1.58.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.58.0...v1.58.1) (2024-08-02)


### Bug Fixes

* resource_system_logging missing topics and excluding Fixes [#518](https://github.com/terraform-routeros/terraform-provider-routeros/issues/518) ([2805dc3](https://github.com/terraform-routeros/terraform-provider-routeros/commit/2805dc35453f8b8051bb9b612814d5fef34a31df))
* routeros_system_logging_action and default settings ([1a5d02c](https://github.com/terraform-routeros/terraform-provider-routeros/commit/1a5d02ce5a8b27bdb97ba38f241db24627621882)), closes [#517](https://github.com/terraform-routeros/terraform-provider-routeros/issues/517)

## [1.58.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.57.2...v1.58.0) (2024-07-31)


### Features

* **lte:** Add LTE resources ([fcaf349](https://github.com/terraform-routeros/terraform-provider-routeros/commit/fcaf3495a50f3cabed77c240de9604d4d5347b3e)), closes [#464](https://github.com/terraform-routeros/terraform-provider-routeros/issues/464)


### Bug Fixes

* **ethernet:** Schema Error on SFP Ports ([01e4c11](https://github.com/terraform-routeros/terraform-provider-routeros/commit/01e4c119d86ade3cfaf040e3deb68749f14282b5)), closes [#514](https://github.com/terraform-routeros/terraform-provider-routeros/issues/514)
* **firewall_nat:** add NAT firewall action "endpoint-independent-nat" ([d55b30d](https://github.com/terraform-routeros/terraform-provider-routeros/commit/d55b30d07ef8f9c2d7c3b53088c8224e819907d9)), closes [#516](https://github.com/terraform-routeros/terraform-provider-routeros/issues/516)
* **system_logging:** Bugfix in routeros_system_logging ([aa1d376](https://github.com/terraform-routeros/terraform-provider-routeros/commit/aa1d3762000006aac00d13a7ae74ee5330c0245f)), closes [#515](https://github.com/terraform-routeros/terraform-provider-routeros/issues/515)

## [1.57.2](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.57.1...v1.57.2) (2024-07-30)


### Bug Fixes

* **certificate:** Certificate import ambiguous value of file-name ([01f4294](https://github.com/terraform-routeros/terraform-provider-routeros/commit/01f4294a2898ecdceb4772db048dd620fa9a221e)), closes [#511](https://github.com/terraform-routeros/terraform-provider-routeros/issues/511)

## [1.57.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.57.0...v1.57.1) (2024-07-30)


### Bug Fixes

* **ospf_interface_template:** RouterOS 7.x & OSPF Interface Template Auth Key ([995ba46](https://github.com/terraform-routeros/terraform-provider-routeros/commit/995ba46e8d8f8fc11c9d47a300f36879876c285f)), closes [#510](https://github.com/terraform-routeros/terraform-provider-routeros/issues/510)

## [1.57.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.56.0...v1.57.0) (2024-07-23)


### Features

* Update the `routeros_ovpn_server` resource to support multiple values ([668ef09](https://github.com/terraform-routeros/terraform-provider-routeros/commit/668ef0986404bea864150b64509b86f43e5b4494))


### Bug Fixes

* **ipv6_neighbor_discovery:** Change the schema attributes ([9688bab](https://github.com/terraform-routeros/terraform-provider-routeros/commit/9688babe485e51b2acc981b6e8a8bb8edd26e98a)), closes [#509](https://github.com/terraform-routeros/terraform-provider-routeros/issues/509)
* Update the `routeros_ovpn_server` resource to handle default values correctly ([f377a26](https://github.com/terraform-routeros/terraform-provider-routeros/commit/f377a2615219baed70ffc779f9be0fe022e42712))

## [1.56.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.55.0...v1.56.0) (2024-07-04)


### Features

* Add `enable_ipv6_accounting` introduced in 7.15 to the `routeros_ppp_aaa` resource ([4c792c7](https://github.com/terraform-routeros/terraform-provider-routeros/commit/4c792c74a4d5e5de43fc5c62176ea6f74925336c))
* Add `require_message_auth` introduced in 7.15 to the `routeros_radius` resource ([6946d95](https://github.com/terraform-routeros/terraform-provider-routeros/commit/6946d95654b504dbc214d046dbab0723dbdb4e0d))
* Add `require_message_auth` introduced in 7.15 to the `routeros_user_manager_settings` resource ([91b41b7](https://github.com/terraform-routeros/terraform-provider-routeros/commit/91b41b7b3c184ec1f0c2cede4e61fadc437171a2))
* Add `reselect_interval` introduced in 7.15 to the `routeros_wifi_channel` resource ([77186c0](https://github.com/terraform-routeros/terraform-provider-routeros/commit/77186c0fbc478c2e8c4a06bc1b5e330e5db402ac))
* Add `sfp_ignore_rx_los` introduced in 7.15 to the `routeros_interface_ethernet` resource ([a361715](https://github.com/terraform-routeros/terraform-provider-routeros/commit/a36171577d57a5af9050cf8e57639f962c5f4ae8))
* Add current CAPsMAN identity computed properties introduced in 7.15 to the `routeros_wifi_cap` resource ([f0fae7b](https://github.com/terraform-routeros/terraform-provider-routeros/commit/f0fae7baf1771b6f65665269794bc68a25b40e23))


### Bug Fixes

* Fix the `AlwaysPresentNotUserProvided` helper to ignore stored empty values as RouterOS started returning such in 7.15 ([c22e017](https://github.com/terraform-routeros/terraform-provider-routeros/commit/c22e017534a6024a35d530d8da689a69ce07f360))

## [1.55.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.54.3...v1.55.0) (2024-07-02)


### Features

* Add state migrator helper to convert scalar values to lists ([858ecab](https://github.com/terraform-routeros/terraform-provider-routeros/commit/858ecab2b52c8933051c2ae992d272969b90c6be))
* Update properties in `routeros_interface_dot1x_client` and `routeros_interface_dot1x_server` to support multiple values ([66ab1d2](https://github.com/terraform-routeros/terraform-provider-routeros/commit/66ab1d2cec80d8c677190315d258b47ca9fe50fa))
* Update properties in `routeros_ip_dhcp_server_network` to support multiple values ([bc1cf68](https://github.com/terraform-routeros/terraform-provider-routeros/commit/bc1cf68b5f3126f0f19bd3f4b4288cc269f25f67))
* Update the `servers` property in `routeros_dns` to support multiple values ([ea25f00](https://github.com/terraform-routeros/terraform-provider-routeros/commit/ea25f0088d2da7dfaa001fab0b06c0bc67bbc847))
* Update the `service` property in `routeros_radius` to support multiple values ([a1becca](https://github.com/terraform-routeros/terraform-provider-routeros/commit/a1beccacd612960abb27b4c66c4941e0e3f0cf4c))
* Update the `vlan_ids` property in `routeros_interface_bridge_vlan` to support multiple values ([3ec3d34](https://github.com/terraform-routeros/terraform-provider-routeros/commit/3ec3d343d9ed0bf501624fae89aae5840da8fa8e))
* Update the frequency properties in `routeros_capsman_channel` to support multiple values ([14a6e4e](https://github.com/terraform-routeros/terraform-provider-routeros/commit/14a6e4e45a91be4f4811f97a7696c25460e1313b))


### Bug Fixes

* Add state migrator to `routeros_snmp_community` to fix backward compatibility ([27c8588](https://github.com/terraform-routeros/terraform-provider-routeros/commit/27c858806bdf9f79bf96a4cd5e85746abc2b2e99))
* Fix `tagged` and `untagged` properties in `routeros_interface_bridge_vlan` to ignore values order ([eb8ff28](https://github.com/terraform-routeros/terraform-provider-routeros/commit/eb8ff28b6c125f20c45ebb783ed248a28f72b935))
* Fix `topics` property type in `routeros_system_logging` to ignore values order ([0e3c29f](https://github.com/terraform-routeros/terraform-provider-routeros/commit/0e3c29f62f587db7b340fee8299c6ec0a3622898))

## [1.54.3](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.54.2...v1.54.3) (2024-06-27)


### Features

* **no-release:** Added ipv6 filter data ([#496](https://github.com/terraform-routeros/terraform-provider-routeros/issues/496)) ([6d45e88](https://github.com/terraform-routeros/terraform-provider-routeros/commit/6d45e88d0af4e99fd552754835d865a034bb54ac))
* **no-release:** Support Multiple VLAN Registration protocol (MVRP). ([#497](https://github.com/terraform-routeros/terraform-provider-routeros/issues/497)) ([0dc994a](https://github.com/terraform-routeros/terraform-provider-routeros/commit/0dc994aed90ed6453dbe422b0ede093b1fae08c7)), closes [#492](https://github.com/terraform-routeros/terraform-provider-routeros/issues/492) [#493](https://github.com/terraform-routeros/terraform-provider-routeros/issues/493)


### Bug Fixes

* Field 'vrf' not found in the schema (introduced in 7.15) ([563401b](https://github.com/terraform-routeros/terraform-provider-routeros/commit/563401bc8c8b8b980afdbd2ad36c424ad8134ebb)), closes [#490](https://github.com/terraform-routeros/terraform-provider-routeros/issues/490)
* **no-release:** Allow a set of `addresses` ([#498](https://github.com/terraform-routeros/terraform-provider-routeros/issues/498)) ([bcf417f](https://github.com/terraform-routeros/terraform-provider-routeros/commit/bcf417f491090bad60cb4f2a9fe313c146259d19)), closes [#495](https://github.com/terraform-routeros/terraform-provider-routeros/issues/495)
* **no-release:** nil resources ([#486](https://github.com/terraform-routeros/terraform-provider-routeros/issues/486)) ([8571dea](https://github.com/terraform-routeros/terraform-provider-routeros/commit/8571dea493a5a22d167c083e65e744f97a50c05b))
* **routeros_interface_wireguard_peer:** Field 'name' not found in the schema (introduced in 7.15) ([9fb13ad](https://github.com/terraform-routeros/terraform-provider-routeros/commit/9fb13ad7523be815ae41cd5b35c5d7e889e02a9e)), closes [#494](https://github.com/terraform-routeros/terraform-provider-routeros/issues/494)
* **routeros_ip_neighbor_discovery_settings:** Multiple fields not found in schema (introduced in 7.15) ([7f44443](https://github.com/terraform-routeros/terraform-provider-routeros/commit/7f44443784cfe5b4372c7af7aee1da94acc0d1c1)), closes [#491](https://github.com/terraform-routeros/terraform-provider-routeros/issues/491)

## [1.54.2](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.54.1...v1.54.2) (2024-06-04)


### Bug Fixes

* **fw-mangle:** Fix `dst_address_list` validation ([0eb25c9](https://github.com/terraform-routeros/terraform-provider-routeros/commit/0eb25c973c2022704a1f85670f125fb5c772388b)), closes [#480](https://github.com/terraform-routeros/terraform-provider-routeros/issues/480)

## [1.54.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.54.0...v1.54.1) (2024-05-30)


### Bug Fixes

* **vrrp:** Add `group-authority` attribute handling to the `group-master` replacement ([58cf139](https://github.com/terraform-routeros/terraform-provider-routeros/commit/58cf139015490181973170d97d1e3931919b1af0)), closes [#446](https://github.com/terraform-routeros/terraform-provider-routeros/issues/446)

## [1.54.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.53.0...v1.54.0) (2024-05-29)


### Features

* Add bandwidth server resource ([#475](https://github.com/terraform-routeros/terraform-provider-routeros/issues/475)) ([d98ce0f](https://github.com/terraform-routeros/terraform-provider-routeros/commit/d98ce0f3aeb918ca02cbfc42eef6b9ba98ef7382)), closes [#474](https://github.com/terraform-routeros/terraform-provider-routeros/issues/474)

## [1.53.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.52.1...v1.53.0) (2024-05-29)


### Features

* **logging-action:** Support for logging actions setup ([f8b9824](https://github.com/terraform-routeros/terraform-provider-routeros/commit/f8b9824a1f00ee456ab3a87793900922942daf42)), closes [#477](https://github.com/terraform-routeros/terraform-provider-routeros/issues/477)


### Bug Fixes

* **certificate-scep-server:** Rename the resource from "routeros_certificate_scep_server" to "routeros_system_certificate_scep_server" ([a9a8138](https://github.com/terraform-routeros/terraform-provider-routeros/commit/a9a81385f2280fa1485d142b7c1e1d9dde541aff)), closes [#473](https://github.com/terraform-routeros/terraform-provider-routeros/issues/473)

## [1.52.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.52.0...v1.52.1) (2024-05-28)


### Bug Fixes

* **firewall-raw:** Fix resource name ([e956f9f](https://github.com/terraform-routeros/terraform-provider-routeros/commit/e956f9ff5d72478f65e4db3e86808f689d0b3a40))

## [1.52.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.51.0...v1.52.0) (2024-05-28)


### Features

* **firewall/raw:** Add new resource ([90eb2fa](https://github.com/terraform-routeros/terraform-provider-routeros/commit/90eb2fa762a92e61bdb408095f7c2ef5a1c03e8e)), closes [#462](https://github.com/terraform-routeros/terraform-provider-routeros/issues/462)


### Bug Fixes

* **dhcp-server:** Remove default values ([#470](https://github.com/terraform-routeros/terraform-provider-routeros/issues/470)) ([884e464](https://github.com/terraform-routeros/terraform-provider-routeros/commit/884e464d7f16f016b99c12371c2cbfca84a149fb)), closes [#466](https://github.com/terraform-routeros/terraform-provider-routeros/issues/466)

## [1.51.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.50.0...v1.51.0) (2024-05-21)


### Features

* **x509:** Datasource for PEM data normalization and common_name extraction ([5f29176](https://github.com/terraform-routeros/terraform-provider-routeros/commit/5f29176d8109379bea87eeb65e8b49cbbc0ceffb))
* **x509:** Import certificates ([5a3bf8e](https://github.com/terraform-routeros/terraform-provider-routeros/commit/5a3bf8ed177e984a7b52322bd70a25431bfb42cd)), closes [#448](https://github.com/terraform-routeros/terraform-provider-routeros/issues/448)

## [1.50.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.49.0...v1.50.0) (2024-05-17)


### Features

* **ovpn:** Add routeros_interface_ovpn_client ([85fd6be](https://github.com/terraform-routeros/terraform-provider-routeros/commit/85fd6be76bfb5a474f419ff226a969a40bc90c92)), closes [#452](https://github.com/terraform-routeros/terraform-provider-routeros/issues/452)

## [1.49.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.48.3...v1.49.0) (2024-05-17)


### Features

* **clock:** Add routeros_system_clock ([e7b3131](https://github.com/terraform-routeros/terraform-provider-routeros/commit/e7b3131606bb1fb6e957c3e4ed68f352f674fa23)), closes [#453](https://github.com/terraform-routeros/terraform-provider-routeros/issues/453)


### Bug Fixes

* **vrf:** Change import method ([915df28](https://github.com/terraform-routeros/terraform-provider-routeros/commit/915df28f7489e35b97c401f1ff4b8fbe1a223826))
* Warnings ([7fe815b](https://github.com/terraform-routeros/terraform-provider-routeros/commit/7fe815b91311a4b62ca0f42940b8cbdca938fbea))

## [1.48.3](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.48.2...v1.48.3) (2024-05-14)


### Bug Fixes

* Add comment to routeros_wifi resource ([#455](https://github.com/terraform-routeros/terraform-provider-routeros/issues/455)) ([5a2782c](https://github.com/terraform-routeros/terraform-provider-routeros/commit/5a2782c30e01ce2e7fa17944058a68eed03e9a2b))

## [1.48.2](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.48.1...v1.48.2) (2024-05-13)


### Bug Fixes

* **veth:** Remove deprecated options ([8565c5b](https://github.com/terraform-routeros/terraform-provider-routeros/commit/8565c5b5325c5b6fd61bcf8a2e76aa53e586f4fa))

## [1.48.1](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.48.0...v1.48.1) (2024-05-09)


### Bug Fixes

* **dns-record:** resource "routeros_ip_dns_record" keeps updating ([a755ab0](https://github.com/terraform-routeros/terraform-provider-routeros/commit/a755ab090746e1456b1bb1b4060f240d4420de5e)), closes [#445](https://github.com/terraform-routeros/terraform-provider-routeros/issues/445)
* **TimeEquall:** Crash when using store_leases_disk in routeros_ip_dhcp_server_config ([ba8c5a9](https://github.com/terraform-routeros/terraform-provider-routeros/commit/ba8c5a92b31e25a134d369e400b2a883e774d692)), closes [#447](https://github.com/terraform-routeros/terraform-provider-routeros/issues/447)

## [1.48.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.47.0...v1.48.0) (2024-05-07)


### Features

* **vrf:** Added routeros_ip_vrf resource ([#443](https://github.com/terraform-routeros/terraform-provider-routeros/issues/443)) ([a091b7d](https://github.com/terraform-routeros/terraform-provider-routeros/commit/a091b7da83aa9a70d04e9c50c8e81c2b6286e8d3))

## [1.47.0](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.46.3...v1.47.0) (2024-05-04)


### Features

* add ability to sign certificates with scep ([1dce5af](https://github.com/terraform-routeros/terraform-provider-routeros/commit/1dce5af9a010463974df011e8d530a792e29f8f2))


### Bug Fixes

* add challenge_password parameter for system_certificate resource ([d933589](https://github.com/terraform-routeros/terraform-provider-routeros/commit/d933589a186f86afad0ba609c05101ce26111149))
* fix changing sign data in routeros_system_certificate resource ([83848f9](https://github.com/terraform-routeros/terraform-provider-routeros/commit/83848f9dac2247c8b60ace492f62d448b85950cb))

## [1.46.3](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.46.2...v1.46.3) (2024-04-27)


### Bug Fixes

* properly unset firewall filter rule `protocol` field when removed ([#435](https://github.com/terraform-routeros/terraform-provider-routeros/issues/435)) ([03a017d](https://github.com/terraform-routeros/terraform-provider-routeros/commit/03a017d40e59f7ba94b6ad1e8b0f5f3236f28f9b))

## [1.46.2](https://github.com/terraform-routeros/terraform-provider-routeros/compare/v1.46.1...v1.46.2) (2024-04-26)


### Bug Fixes

* **bridge_port:** [Backward compatibility] routeros_interface_bridge_port  ([#436](https://github.com/terraform-routeros/terraform-provider-routeros/issues/436)) ([ad64040](https://github.com/terraform-routeros/terraform-provider-routeros/commit/ad64040358e5da70e326ab70d353008f263eb7fc)), closes [#419](https://github.com/terraform-routeros/terraform-provider-routeros/issues/419) [#419](https://github.com/terraform-routeros/terraform-provider-routeros/issues/419)

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
