package routeros

import (
	"context"

	roscl "github.com/gnewbury1/terraform-provider-routeros/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider(client *roscl.Client) *schema.Provider {
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
				DefaultFunc: schema.EnvDefaultFunc("ROS_USER", nil),
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
				Required:    false,
				Default:     false,
				DefaultFunc: schema.EnvDefaultFunc("ROS_INSECURE", nil),
				Description: "Whether to verify the SSL certificate or not",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"ros_ip_address": resourceIPAddress(),
		},
	}

	provider.ConfigureContextFunc = func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		if client != nil {
			return client, nil
		}

		hosturl := d.Get("hosturl").(string)
		username := d.Get("username").(string)
		password := d.Get("password").(string)
		insecure := d.Get("insecure").(bool)

		return roscl.NewClient
	}
}
