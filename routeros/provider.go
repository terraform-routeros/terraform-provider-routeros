package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	ErrorMsgPut    = "An error was encountered while sending a PUT request to the API: %v"
	ErrorMsgGet    = "An error was encountered while sending a GET request to the API: %v"
	ErrorMsgPatch  = "An error was encountered while sending a PATCH request to the API: %v"
	ErrorMsgDelete = "An error was encountered while sending a DELETE request to the API: %v"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"hosturl": {
				Type:     schema.TypeString,
				Required: true,
				DefaultFunc: schema.MultiEnvDefaultFunc(
					[]string{"ROS_HOSTURL", "MIKROTIK_HOST"},
					nil,
				),
				Description: `URL of the MikroTik router, default is TLS connection to REST.
	* API: api[s]://host[:port]
		* api://router.local
		* apis://router.local:8729
	* REST: http[s]://host
		* http://router.local
		* https://router.local
		* router.local
		* 127.0.0.1


	export ROS_HOSTURL=router.local or export MIKROTIK_HOST=router.local
`,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
				DefaultFunc: schema.MultiEnvDefaultFunc(
					[]string{"ROS_USERNAME", "MIKROTIK_USER"},
					nil,
				),
				Description: `Username for the MikroTik WEB/Winbox.


	export ROS_USERNAME=admin or export MIKROTIK_USER=admin
`,
			},
			"password": {
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc(
					[]string{"ROS_PASSWORD", "MIKROTIK_PASSWORD"},
					nil,
				),
				Description: "Password for the MikroTik user.",
				Sensitive:   true,
			},
			"ca_certificate": {
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc(
					[]string{"ROS_CA_CERTIFICATE", "MIKROTIK_CA_CERTIFICATE"},
					nil,
				),
				Description: "Path to MikroTik's certificate authority file.",
			},
			"insecure": {
				Type:     schema.TypeBool,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc(
					[]string{"ROS_INSECURE", "MIKROTIK_INSECURE"},
					false,
				),
				Description: "Whether to verify the SSL certificate or not.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{

			// IP objects
			"routeros_ip_dhcp_client":                  ResourceDhcpClient(),
			"routeros_ip_dhcp_client_option":           ResourceDhcpClientOption(),
			"routeros_ip_dhcp_relay":                   ResourceDhcpRelay(),
			"routeros_ip_dhcp_server":                  ResourceDhcpServer(),
			"routeros_ip_dhcp_server_config":           ResourceDhcpServerConfig(),
			"routeros_ip_dhcp_server_network":          ResourceDhcpServerNetwork(),
			"routeros_ip_dhcp_server_lease":            ResourceDhcpServerLease(),
			"routeros_ip_dhcp_server_option":           ResourceDhcpServerOption(),
			"routeros_ip_dhcp_server_option_set":       ResourceDhcpServerOptionSet(),
			"routeros_ip_firewall_addr_list":           ResourceIPFirewallAddrList(),
			"routeros_ip_firewall_connection_tracking": ResourceIPConnectionTracking(),
			"routeros_ip_firewall_filter":              ResourceIPFirewallFilter(),
			"routeros_ip_firewall_mangle":              ResourceIPFirewallMangle(),
			"routeros_ip_firewall_nat":                 ResourceIPFirewallNat(),
			"routeros_ip_firewall_raw":                 ResourceIPFirewallRaw(),
			"routeros_ip_address":                      ResourceIPAddress(),
			"routeros_ip_pool":                         ResourceIPPool(),
			"routeros_ip_route":                        ResourceIPRoute(),
			"routeros_ip_dns":                          ResourceDns(),
			"routeros_ip_dns_record":                   ResourceDnsRecord(),
			"routeros_ip_service":                      ResourceIpService(),
			"routeros_ip_neighbor_discovery_settings":  ResourceIpNeighborDiscoverySettings(),
			"routeros_ip_ssh_server":                   ResourceIpSSHServer(),
			"routeros_ip_upnp":                         ResourceUPNPSettings(),
			"routeros_ip_upnp_interfaces":              ResourceUPNPInterfaces(),
			"routeros_ip_vrf":                          ResourceIPVrf(),
			"routeros_ipv6_address":                    ResourceIPv6Address(),
			"routeros_ipv6_dhcp_client":                ResourceIPv6DhcpClient(),
			"routeros_ipv6_dhcp_client_option":         ResourceIPv6DhcpClientOption(),
			"routeros_ipv6_firewall_addr_list":         ResourceIPv6FirewallAddrList(),
			"routeros_ipv6_firewall_filter":            ResourceIPv6FirewallFilter(),
			"routeros_ipv6_neighbor_discovery":         ResourceIPv6NeighborDiscovery(),
			"routeros_ipv6_route":                      ResourceIPv6Route(),

			// Aliases for IP objects to retain compatibility between original and fork
			"routeros_dhcp_client":         ResourceDhcpClient(),
			"routeros_dhcp_client_option":  ResourceDhcpClientOption(),
			"routeros_dhcp_server":         ResourceDhcpServer(),
			"routeros_dhcp_server_network": ResourceDhcpServerNetwork(),
			"routeros_dhcp_server_lease":   ResourceDhcpServerLease(),
			"routeros_firewall_addr_list":  ResourceIPFirewallAddrList(),
			"routeros_firewall_filter":     ResourceIPFirewallFilter(),
			"routeros_firewall_mangle":     ResourceIPFirewallMangle(),
			"routeros_firewall_nat":        ResourceIPFirewallNat(),
			"routeros_dns":                 ResourceDns(),
			"routeros_dns_record":          ResourceDnsRecord(),

			// Interface Objects
			"routeros_interface_bridge":                         ResourceInterfaceBridge(),
			"routeros_interface_bridge_port":                    ResourceInterfaceBridgePort(),
			"routeros_interface_bridge_vlan":                    ResourceInterfaceBridgeVlan(),
			"routeros_interface_bridge_settings":                ResourceInterfaceBridgeSettings(),
			"routeros_interface_dot1x_client":                   ResourceInterfaceDot1xClient(),
			"routeros_interface_dot1x_server":                   ResourceInterfaceDot1xServer(),
			"routeros_interface_eoip":                           ResourceInterfaceEoip(),
			"routeros_interface_ethernet_switch":                ResourceInterfaceEthernetSwitch(),
			"routeros_interface_ethernet_switch_host":           ResourceInterfaceEthernetSwitchHost(),
			"routeros_interface_ethernet_switch_port":           ResourceInterfaceEthernetSwitchPort(),
			"routeros_interface_ethernet_switch_port_isolation": ResourceInterfaceEthernetSwitchPortIsolation(),
			"routeros_interface_ethernet_switch_vlan":           ResourceInterfaceEthernetSwitchVlan(),
			"routeros_interface_ethernet_switch_rule":           ResourceInterfaceEthernetSwitchRule(),
			"routeros_interface_gre":                            ResourceInterfaceGre(),
			"routeros_interface_macvlan":                        ResourceInterfaceMacVlan(),
			"routeros_interface_ipip":                           ResourceInterfaceIPIP(),
			"routeros_interface_vlan":                           ResourceInterfaceVlan(),
			"routeros_interface_vrrp":                           ResourceInterfaceVrrp(),
			"routeros_interface_wireguard":                      ResourceInterfaceWireguard(),
			"routeros_interface_wireguard_peer":                 ResourceInterfaceWireguardPeer(),
			"routeros_interface_wireless_cap":                   ResourceInterfaceWirelessCap(),
			"routeros_interface_list":                           ResourceInterfaceList(),
			"routeros_interface_list_member":                    ResourceInterfaceListMember(),
			"routeros_interface_ovpn_server":                    ResourceInterfaceOpenVPNServer(),
			"routeros_interface_ovpn_client":                    ResourceOpenVPNClient(),
			"routeros_interface_veth":                           ResourceInterfaceVeth(),
			"routeros_interface_bonding":                        ResourceInterfaceBonding(),
			"routeros_interface_pppoe_client":                   ResourceInterfacePPPoEClient(),
			"routeros_interface_ethernet":                       ResourceInterfaceEthernet(),

			// Aliases for interface objects to retain compatibility between original and fork
			"routeros_bridge":         ResourceInterfaceBridge(),
			"routeros_bridge_mlag":    ResourceInterfaceBridgeMlag(),
			"routeros_bridge_port":    ResourceInterfaceBridgePort(),
			"routeros_bridge_vlan":    ResourceInterfaceBridgeVlan(),
			"routeros_gre":            ResourceInterfaceGre(),
			"routeros_ipip":           ResourceInterfaceIPIP(),
			"routeros_vlan":           ResourceInterfaceVlan(),
			"routeros_vrrp":           ResourceInterfaceVrrp(),
			"routeros_wireguard":      ResourceInterfaceWireguard(),
			"routeros_wireguard_peer": ResourceInterfaceWireguardPeer(),

			// System Objects
			"routeros_ip_cloud":                       ResourceIpCloud(),
			"routeros_ip_cloud_advanced":              ResourceIpCloudAdvanced(),
			"routeros_system_certificate":             ResourceSystemCertificate(),
			"routeros_system_certificate_scep_server": ResourceCertificateScepServer(),
			"routeros_certificate_scep_server":        ResourceCertificateScepServer(),
			"routeros_system_clock":                   ResourceSystemClock(),
			"routeros_system_identity":                ResourceSystemIdentity(),
			"routeros_system_logging":                 ResourceSystemLogging(),
			"routeros_system_logging_action":          ResourceSystemLoggingAction(),
			"routeros_system_ntp_client":              ResourceSystemNtpClient(),
			"routeros_system_ntp_server":              ResourceSystemNtpServer(),
			"routeros_system_scheduler":               ResourceSystemScheduler(),
			"routeros_system_script":                  ResourceSystemScript(),
			"routeros_system_user":                    ResourceUser(),
			"routeros_system_user_aaa":                ResourceUserAaa(),
			"routeros_system_user_group":              ResourceUserGroup(),
			"routeros_system_user_settings":           ResourceSystemUserSettings(),

			// Aliases for system objects to retain compatibility between original and fork
			"routeros_identity":  ResourceSystemIdentity(),
			"routeros_scheduler": ResourceSystemScheduler(),

			// CAPsMAN Objects
			"routeros_capsman_aaa":               ResourceCapsManAaa(),
			"routeros_capsman_access_list":       ResourceCapsManAccessList(),
			"routeros_capsman_channel":           ResourceCapsManChannel(),
			"routeros_capsman_configuration":     ResourceCapsManConfiguration(),
			"routeros_capsman_datapath":          ResourceCapsManDatapath(),
			"routeros_capsman_interface":         ResourceCapsManInterface(),
			"routeros_capsman_manager":           ResourceCapsManManager(),
			"routeros_capsman_manager_interface": ResourceCapsManManagerInterface(),
			"routeros_capsman_provisioning":      ResourceCapsManProvisioning(),
			"routeros_capsman_rates":             ResourceCapsManRates(),
			"routeros_capsman_security":          ResourceCapsManSecurity(),

			// Container objects
			"routeros_container":        ResourceContainer(),
			"routeros_container_config": ResourceContainerConfig(),
			"routeros_container_envs":   ResourceContainerEnvs(),
			"routeros_container_mounts": ResourceContainerMounts(),

			// File objects
			"routeros_file": ResourceFile(),

			// Routing
			"routeros_routing_bgp_connection": ResourceRoutingBGPConnection(),
			"routeros_routing_bgp_template":   ResourceRoutingBGPTemplate(),
			"routeros_routing_filter_rule":    ResourceRoutingFilterRule(),
			"routeros_routing_table":          ResourceRoutingTable(),

			// OSPF
			"routeros_routing_ospf_instance":           ResourceRoutingOspfInstance(),
			"routeros_routing_ospf_area":               ResourceRoutingOspfArea(),
			"routeros_routing_ospf_interface_template": ResourceRoutingOspfInterfaceTemplate(),

			// VPN
			"routeros_ovpn_server": ResourceOpenVPNServer(),

			// PPP
			"routeros_ppp_aaa":     ResourcePppAaa(),
			"routeros_ppp_profile": ResourcePPPProfile(),
			"routeros_ppp_secret":  ResourcePPPSecret(),

			// RADIUS
			"routeros_radius":          ResourceRadius(),
			"routeros_radius_incoming": ResourceRadiusIncoming(),

			// SNMP
			"routeros_snmp":           ResourceSNMP(),
			"routeros_snmp_community": ResourceSNMPCommunity(),

			// Helpers
			"routeros_wireguard_keys": ResourceWireguardKeys(),
			"routeros_move_items":     ResourceMoveItems(),

			// Tools
			"routeros_tool_bandwidth_server":  ResourceToolBandwidthServer(),
			"routeros_tool_mac_server":        ResourceToolMacServer(),
			"routeros_tool_mac_server_winbox": ResourceToolMacServerWinBox(),

			// User Manager
			"routeros_user_manager_advanced":           ResourceUserManagerAdvanced(),
			"routeros_user_manager_attribute":          ResourceUserManagerAttribute(),
			"routeros_user_manager_database":           ResourceUserManagerDatabase(),
			"routeros_user_manager_limitation":         ResourceUserManagerLimitation(),
			"routeros_user_manager_profile":            ResourceUserManagerProfile(),
			"routeros_user_manager_profile_limitation": ResourceUserManagerProfileLimitation(),
			"routeros_user_manager_router":             ResourceUserManagerRouter(),
			"routeros_user_manager_settings":           ResourceUserManagerSettings(),
			"routeros_user_manager_user":               ResourceUserManagerUser(),
			"routeros_user_manager_user_group":         ResourceUserManagerUserGroup(),
			"routeros_user_manager_user_profile":       ResourceUserManagerUserProfile(),

			// WiFi
			"routeros_wifi":               ResourceWifi(),
			"routeros_wifi_aaa":           ResourceWifiAaa(),
			"routeros_wifi_access_list":   ResourceWifiAccessList(),
			"routeros_wifi_cap":           ResourceWifiCap(),
			"routeros_wifi_capsman":       ResourceWifiCapsman(),
			"routeros_wifi_channel":       ResourceWifiChannel(),
			"routeros_wifi_configuration": ResourceWifiConfiguration(),
			"routeros_wifi_datapath":      ResourceWifiDatapath(),
			"routeros_wifi_interworking":  ResourceWifiInterworking(),
			"routeros_wifi_provisioning":  ResourceWifiProvisioning(),
			"routeros_wifi_security":      ResourceWifiSecurity(),
			"routeros_wifi_steering":      ResourceWifiSteering(),

			// ZeroTier
			"routeros_zerotier":            ResourceZerotier(),
			"routeros_zerotier_controller": ResourceZerotierController(),
			"routeros_zerotier_interface":  ResourceZerotierInterface(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"routeros_files":                 DatasourceFiles(),
			"routeros_interfaces":            DatasourceInterfaces(),
			"routeros_ip_addresses":          DatasourceIPAddresses(),
			"routeros_ip_arp":                DatasourceIpArp(),
			"routeros_ip_dhcp_server_leases": DatasourceIpDhcpServerLeases(),
			"routeros_ip_firewall":           DatasourceIPFirewall(),
			"routeros_ip_routes":             DatasourceIPRoutes(),
			"routeros_ip_services":           DatasourceIPServices(),
			"routeros_ipv6_addresses":        DatasourceIPv6Addresses(),
			"routeros_ipv6_firewall":         DatasourceIPv6Firewall(),
			"routeros_system_resource":       DatasourceSystemResource(),
			"routeros_x509":                  DatasourceX509(),

			// Aliases for entries that have been renamed
			"routeros_firewall": DatasourceIPFirewall(),
		},
		ConfigureContextFunc: NewClient,
	}
}

func NewProvider() *schema.Provider {
	return Provider()
}
