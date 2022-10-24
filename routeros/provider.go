package routeros

import (
	"context"
	"fmt"

	roscl "github.com/gnewbury1/terraform-provider-routeros/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"hosturl": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ROS_HOSTURL", nil),
				Description: "URL of the ROS router. Include the scheme (http/https)",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ROS_USERNAME", nil),
				Description: "Username for the ROS user",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ROS_PASSWORD", nil),
				Description: "Password for the ROS user",
				Sensitive:   true,
			},
			"insecure": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				DefaultFunc: schema.EnvDefaultFunc("ROS_INSECURE", false),
				Description: "Whether to verify the SSL certificate or not",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"routeros_ip_address":               resourceIPAddress(),
			"routeros_ip_dhcp_client":           resourceDhcpClient(),
			"routeros_ip_dhcp_server":           resourceDhcpServer(),
			"routeros_ip_dhcp_server_network":   resourceDhcpServerNetwork(),
			"routeros_ip_pool":                  resourceIPPool(),
			"routeros_ip_route":                 resourceIPRoute(),
			"routeros_ip_firewall_filter":       resourceIPFirewallFilter(),
			"routeros_interface_vlan":           resourceInterfaceVlan(),
			"routeros_interface_bridge_vlan":    resourceInterfaceBridgeVlan(),
			"routeros_interface_bridge_port":    resourceInterfaceBridgePort(),
			"routeros_interface_bridge":         resourceInterfaceBridge(),
			"routeros_interface_wireguard":      resourceInterfaceWireguard(),
			"routeros_interface_wireguard_peer": resourceInterfaceWireguardPeer(),
			"routeros_capsman_channel":          resourceCapsManChannel(),
			"routeros_capsman_datapath":         resourceCapsManDatapath(),
			"routeros_capsman_security":         resourceCapsManSecurity(),
			"routeros_capsman_manager":          resourceCapsManManager(),
			"routeros_capsman_provisioning":     resourceCapsManProvisioning(),
			"routeros_capsman_configuration":    resourceCapsManConfiguration(),
			"routeros_interface_list":           resourceInterfaceList(),
			"routeros_interface_list_member":    resourceInterfaceListMember(),
			"routeros_interface_vrrp":           resourceInterfaceVrrp(),
			"routeros_system_identity":          resourceSystemIdentity(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"routeros_ip_addresses": datasourceIPAddresses(),
			"routeros_ip_routes":    datasourceIPRoutes(),
			"routeros_interfaces":   datasourceInterfaces(),
		},
	}

	provider.ConfigureContextFunc = func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {

		hosturl := d.Get("hosturl").(string)
		username := d.Get("username").(string)
		password := d.Get("password").(string)
		insecure := d.Get("insecure").(bool)

		return roscl.NewClient(hosturl, username, password, insecure), nil
	}

	return provider
}

func NewProvider() *schema.Provider {
	return Provider()
}

// Add these two functions as some of the RouterOS API treat bool string as "yes" and "no" rather than "true" and "false" (Eg, /ip/dhcp-client/add-default-route).
// Seems to be an API inconsistentcy, as others accept "true" and "false"
func BoolStringYesNo(boolstring string) string {
	var new_bool_string string
	if boolstring == "true" {
		new_bool_string = "yes"
	} else if boolstring == "false" {
		new_bool_string = "no"
	}
	return new_bool_string
}

func BoolStringTrueFalse(boolstring string) string {
	var new_bool_string string
	if boolstring == "yes" {
		new_bool_string = "true"
	} else if boolstring == "no" {
		new_bool_string = "false"
	}
	return new_bool_string
}

func ConvSInterfaceToSString(slice []interface{}) []string {
	string_slice := make([]string, len(slice))
	for k, v := range slice {
		string_slice[k] = fmt.Sprint(v)
	}
	return string_slice
}

func ConvSStringToSInterface(slice []string) []interface{} {
	interface_slice := make([]interface{}, len(slice))
	for k, v := range slice {
		interface_slice[k] = v
	}
	return interface_slice
}
