package routeros

import (
	"log"
	"strconv"

	roscl "github.com/gnewbury1/terraform-provider-routeros/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDhcpServerNetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceDhcpServerNetworkCreate,
		Read:   resourceDhcpServerNetworkRead,
		Update: resourceDhcpServerNetworkUpdate,
		Delete: resourceDhcpServerNetworkDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"boot_file_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"caps_manager": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dhcp_option": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dhcp_option_set": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_none": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"dns_server": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"gateway": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"netmask": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"next_server": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ntp_server": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"wins_server": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceDhcpServerNetworkCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	dhcp_server_network := new(roscl.DhcpServerNetwork)

	dhcp_server_network.Address = d.Get("address").(string)
	dhcp_server_network.BootFileName = d.Get("boot_file_name").(string)
	dhcp_server_network.CapsManager = d.Get("caps_manager").(string)
	dhcp_server_network.DhcpOption = d.Get("dhcp_option").(string)
	dhcp_server_network.DhcpOptionSet = d.Get("dhcp_option_set").(string)
	dhcp_server_network.DnsServer = d.Get("dns_server").(string)
	dhcp_server_network.Domain = d.Get("domain").(string)
	dhcp_server_network.Gateway = d.Get("gateway").(string)
	dhcp_server_network.Netmask = strconv.Itoa(d.Get("netmask").(int))
	dhcp_server_network.NextServer = d.Get("next_server").(string)
	dhcp_server_network.NtpServer = d.Get("ntp_server").(string)
	dhcp_server_network.WinsServer = d.Get("wins_server").(string)
	dhcp_server_network.DnsNone = strconv.FormatBool(d.Get("dns_none").(bool))

	res, err := c.CreateDhcpServerNetwork(dhcp_server_network)
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PUT request to the API")
		log.Fatal(err.Error())
		return err
	}

	d.SetId(res.ID)
	return nil
}

func resourceDhcpServerNetworkRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	res, err := c.ReadDhcpServerNetwork(d.Id())

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a GET request to the API")
		log.Fatal(err.Error())
		return err
	}

	dns_none, _ := strconv.ParseBool(res.DnsNone)
	netmask, _ := strconv.Atoi(res.Netmask)

	d.SetId(res.ID)
	d.Set("address", res.Address)
	d.Set("boot_file_name", res.BootFileName)
	d.Set("caps_manager", res.CapsManager)
	d.Set("dhcp_option", res.DhcpOption)
	d.Set("dhcp_option_set", res.DhcpOptionSet)
	d.Set("dns_none", dns_none)
	d.Set("dns_server", res.DnsServer)
	d.Set("domain", res.Domain)
	d.Set("gateway", res.Gateway)
	d.Set("netmask", netmask)
	d.Set("next_server", res.NextServer)
	d.Set("ntp_server", res.NtpServer)
	d.Set("wins_server", res.WinsServer)

	return nil

}

func resourceDhcpServerNetworkUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	dhcp_server_network := new(roscl.DhcpServerNetwork)

	dhcp_server_network.Address = d.Get("address").(string)
	dhcp_server_network.BootFileName = d.Get("boot_file_name").(string)
	dhcp_server_network.CapsManager = d.Get("caps_manager").(string)
	dhcp_server_network.DhcpOption = d.Get("dhcp_option").(string)
	dhcp_server_network.DhcpOptionSet = d.Get("dhcp_option_set").(string)
	dhcp_server_network.DnsServer = d.Get("dns_server").(string)
	dhcp_server_network.Domain = d.Get("domain").(string)
	dhcp_server_network.Gateway = d.Get("gateway").(string)
	dhcp_server_network.Netmask = strconv.Itoa(d.Get("netmask").(int))
	dhcp_server_network.NextServer = d.Get("next_server").(string)
	dhcp_server_network.NtpServer = d.Get("ntp_server").(string)
	dhcp_server_network.WinsServer = d.Get("wins_server").(string)
	dhcp_server_network.DnsNone = strconv.FormatBool(d.Get("dns_none").(bool))

	res, err := c.UpdateDhcpServerNetwork(d.Id(), dhcp_server_network)

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PATCH request to the API")
		log.Fatal(err.Error())
		return err
	}

	d.SetId(res.ID)

	return nil
}

func resourceDhcpServerNetworkDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	dhcp_server_network, _ := c.ReadDhcpServerNetwork(d.Id())
	err := c.DeleteDhcpServerNetwork(dhcp_server_network)
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a DELETE request to the API")
		log.Fatal(err.Error())
		return err
	}
	d.SetId("")
	return nil
}
