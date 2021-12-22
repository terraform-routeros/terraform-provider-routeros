package routeros

import (
	"log"
	"strconv"

	roscl "github.com/gnewbury1/terraform-provider-routeros/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDhcpServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceDhcpServerCreate,
		Read:   resourceDhcpServerRead,
		Update: resourceDhcpServerUpdate,
		Delete: resourceDhcpServerDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"address_pool": {
				Type:     schema.TypeString,
				Required: true,
			},
			"authoritative": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"dynamic": {
				Type:     schema.TypeBool,
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
			"lease_script": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"lease_time": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "10m",
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"use_radius": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceDhcpServerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	dhcp_server := new(roscl.DhcpServer)

	dhcp_server.AddressPool = d.Get("address_pool").(string)
	dhcp_server.Authoritative = BoolStringYesNo(strconv.FormatBool(d.Get("authoritative").(bool)))
	dhcp_server.Disabled = strconv.FormatBool(d.Get("disabled").(bool))
	dhcp_server.Interface = d.Get("interface").(string)
	dhcp_server.LeaseScript = d.Get("lease_script").(string)
	dhcp_server.LeaseTime = d.Get("lease_time").(string)
	dhcp_server.Name = d.Get("name").(string)
	dhcp_server.UseRadius = BoolStringYesNo(strconv.FormatBool(d.Get("use_radius").(bool)))

	res, err := c.CreateDhcpServer(dhcp_server)
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PUT request to the API")
		log.Fatal(err.Error())
		return err
	}

	d.SetId(res.ID)
	return nil
}

func resourceDhcpServerRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	res, err := c.ReadDhcpServer(d.Id())

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a GET request to the API")
		log.Fatal(err.Error())
		return err
	}

	authoritative, _ := strconv.ParseBool(BoolStringTrueFalse(res.Authoritative))
	disabled, _ := strconv.ParseBool(res.Disabled)
	dynamic, _ := strconv.ParseBool(res.Dynamic)
	use_radius, _ := strconv.ParseBool(BoolStringTrueFalse(res.UseRadius))

	d.SetId(res.ID)
	d.Set("address_pool", res.AddressPool)
	d.Set("authoritative", authoritative)
	d.Set("disabled", disabled)
	d.Set("dynamic", dynamic)
	d.Set("interface", res.Interface)
	d.Set("lease_script", res.LeaseScript)
	d.Set("lease_time", res.LeaseTime)
	d.Set("name", res.Name)
	d.Set("use_radius", use_radius)

	return nil

}

func resourceDhcpServerUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	dhcp_server := new(roscl.DhcpServer)

	dhcp_server.AddressPool = d.Get("address_pool").(string)
	dhcp_server.Authoritative = strconv.FormatBool(d.Get("authoritative").(bool))
	dhcp_server.Disabled = strconv.FormatBool(d.Get("disabled").(bool))
	dhcp_server.Interface = d.Get("interface").(string)
	dhcp_server.LeaseScript = d.Get("lease_script").(string)
	dhcp_server.LeaseTime = d.Get("lease_time").(string)
	dhcp_server.Name = d.Get("name").(string)
	dhcp_server.UseRadius = strconv.FormatBool(d.Get("use_radius").(bool))

	res, err := c.UpdateDhcpServer(d.Id(), dhcp_server)

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PATCH request to the API")
		log.Fatal(err.Error())
		return err
	}

	d.SetId(res.ID)

	return nil
}

func resourceDhcpServerDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	dhcp_server, _ := c.ReadDhcpServer(d.Id())
	err := c.DeleteDhcpServer(dhcp_server)
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a DELETE request to the API")
		log.Fatal(err.Error())
		return err
	}
	d.SetId("")
	return nil
}
