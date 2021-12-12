package routeros

import (
	"context"

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
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ROS_PASSWORD", nil),
				Description: "Password for the ROS user",
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
			"routeros_ip_address":     resourceIPAddress(),
			"routeros_ip_dhcp_client": resourceDhcpClient(),
			"routeros_ip_dhcp_server": resourceDhcpServer(),
			"routeros_interface_vlan": resourceInterfaceVlan(),
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
