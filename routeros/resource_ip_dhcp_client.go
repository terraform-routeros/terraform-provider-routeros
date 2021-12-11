package routeros

import (
	"fmt"
	"strconv"

	roscl "github.com/gnewbury1/terraform-provider-routeros/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDhcpClient() *schema.Resource {
	return &schema.Resource{
		Create: resourceDhcpClientCreate,
		Read:   resourceDhcpClientRead,
		Update: resourceDhcpClientUpdate,
		Delete: resourceDhcpClientDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"add_default_route": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_route_distance": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dhcp_options": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_server": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"dynamic": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"expires_after": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"gateway": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": {
				Type:     schema.TypeString,
				Required: true,
			},
			"invalid": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"primary_dns": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secondary_dns": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"use_peer_dns": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"use_peer_ntp": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func resourceDhcpClientCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	dhcp_client := new(roscl.DhcpClient)

	dhcp_client.AddDefaultRoute = strconv.FormatBool(d.Get("add_default_route").(bool))
	dhcp_client.DefaultRouteDistance = strconv.Itoa(d.Get("default_route_distance").(int))
	dhcp_client.DhcpOptions = d.Get("dhcp_options").(string)
	dhcp_client.DhcpServer = d.Get("dhcp_server").(string)
	dhcp_client.Disabled = strconv.FormatBool(d.Get("disabled").(bool))
	dhcp_client.Dynamic = strconv.FormatBool(d.Get("dynamic").(bool))
	dhcp_client.Gateway = d.Get("gateway").(string)
	dhcp_client.Interface = d.Get("interface").(string)
	dhcp_client.PrimaryDNS = d.Get("primary_dns").(string)
	dhcp_client.SecondaryDNS = d.Get("secondary_dns").(string)
	dhcp_client.UsePeerDNS = strconv.FormatBool(d.Get("use_peer_dns").(bool))
	dhcp_client.UsePeerNtp = strconv.FormatBool(d.Get("use_peer_ntp").(bool))

	res, err := c.CreateDhcpClient(dhcp_client)
	if err != nil {
		return fmt.Errorf("error creating dhcp client: %s", err.Error())
	}

	d.SetId(res.ID)
	return nil
}

func resourceDhcpClientRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	res, err := c.ReadDhcpClient(d.Id())

	if err != nil {
		return fmt.Errorf("error fetching dhcp_client: %s", err.Error())
	}

	add_default_route, _ := strconv.ParseBool(res.AddDefaultRoute)
	default_route_distance, _ := strconv.Atoi(res.DefaultRouteDistance)
	disabled, _ := strconv.ParseBool(res.Disabled)
	dynamic, _ := strconv.ParseBool(res.Dynamic)
	invalid, _ := strconv.ParseBool(res.Invalid)
	use_peer_dns, _ := strconv.ParseBool(res.UsePeerDNS)
	use_peer_ntp, _ := strconv.ParseBool(res.UsePeerNtp)

	d.SetId(res.ID)
	d.Set("add_default_route", add_default_route)
	d.Set("address", res.Address)
	d.Set("default_route_distance", default_route_distance)
	d.Set("dhcp_options", res.DhcpOptions)
	d.Set("dhcp_server", res.DhcpServer)
	d.Set("disabled", disabled)
	d.Set("dynamic", dynamic)
	d.Set("expires_after", res.ExpiresAfter)
	d.Set("gateway", res.Gateway)
	d.Set("interface", res.Interface)
	d.Set("invalid", invalid)
	d.Set("primary_dns", res.PrimaryDNS)
	d.Set("secondary_dns", res.SecondaryDNS)
	d.Set("status", res.Status)
	d.Set("use_peer_dns", use_peer_dns)
	d.Set("use_peer_ntp", use_peer_ntp)

	return nil

}

func resourceDhcpClientUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	dhcp_client := new(roscl.DhcpClient)

	dhcp_client.AddDefaultRoute = strconv.FormatBool(d.Get("add_default_route").(bool))
	dhcp_client.DefaultRouteDistance = strconv.FormatInt(d.Get("default_route_distance").(int64), 10)
	dhcp_client.DhcpOptions = d.Get("dhcp_options").(string)
	dhcp_client.DhcpServer = d.Get("dhcp_server").(string)
	dhcp_client.Disabled = strconv.FormatBool(d.Get("disabled").(bool))
	dhcp_client.Dynamic = strconv.FormatBool(d.Get("dynamic").(bool))
	dhcp_client.Gateway = d.Get("gateway").(string)
	dhcp_client.Interface = d.Get("interface").(string)
	dhcp_client.PrimaryDNS = d.Get("primary_dns").(string)
	dhcp_client.SecondaryDNS = d.Get("secondary_dns").(string)
	dhcp_client.UsePeerDNS = strconv.FormatBool(d.Get("use_peer_dns").(bool))
	dhcp_client.UsePeerNtp = strconv.FormatBool(d.Get("use_peer_ntp").(bool))

	res, err := c.UpdateDhcpClient(d.Id(), dhcp_client)

	if err != nil {
		return fmt.Errorf("error updating dhcp client: %s", err.Error())
	}

	d.SetId(res.ID)

	return nil
}

func resourceDhcpClientDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	dhcp_client, _ := c.ReadDhcpClient(d.Id())
	err := c.DeleteDhcpClient(dhcp_client)
	if err != nil {
		return fmt.Errorf("error deleting ip address: %s", err.Error())
	}
	d.SetId("")
	return nil
}
