package routeros

import (
	"log"
	"strconv"

	roscl "github.com/gnewbury1/terraform-provider-routeros/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIPv6Address() *schema.Resource {
	return &schema.Resource{
		Create: resourceIPv6AddressCreate,
		Read:   resourceIPv6AddressRead,
		Update: resourceIPv6AddressUpdate,
		Delete: resourceIPv6AddressDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"actual_interface": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"advertise": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"eui64": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"from_pool": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"interface": {
				Type:     schema.TypeString,
				Required: true,
			},
			"invalid": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"no_dad": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceIPv6AddressCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	ip_addr := readFromTerraform(d, m)

	res, err := c.CreateIPv6Address(ip_addr)
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PUT request to the API")
		log.Fatal(err.Error())
		return err
	}

	err = writeToTerraform(d, m, res)

	if err != nil {
		return err
	}

	return nil

}

func resourceIPv6AddressRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	res, err := c.GetIPv6Address(d.Id())

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a GET request to the API")
		log.Fatal(err.Error())
		return err
	}

	err = writeToTerraform(d, m, res)

	if err != nil {
		return err
	}

	return nil

}

func resourceIPv6AddressUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	ip_addr := readFromTerraform(d, m)

	res, err := c.UpdateIPv6Address(d.Id(), ip_addr)

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PATCH request to the API")
		log.Fatal(err.Error())
		return err
	}

	err = writeToTerraform(d, m, res)

	if err != nil {
		return err
	}

	return nil
}

func resourceIPv6AddressDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	err := c.DeleteIPv6Address(d.Id())
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a DELETE request to the API")
		log.Fatal(err.Error())
		return err
	}
	d.SetId("")
	return nil
}

func readFromTerraform(d *schema.ResourceData, m interface{}) *roscl.IPv6Address {

	ip_addr := new(roscl.IPv6Address)

	ip_addr.ActualInterface = d.Get("actual_interface").(string)
	ip_addr.Address = d.Get("address").(string)
	ip_addr.Advertise = strconv.FormatBool(d.Get("advertise").(bool))
	ip_addr.Comment = d.Get("comment").(string)
	ip_addr.Disabled = strconv.FormatBool(d.Get("disabled").(bool))
	ip_addr.Eui64 = strconv.FormatBool(d.Get("eui64").(bool))
	ip_addr.FromPool = d.Get("from_pool").(string)
	ip_addr.Interface = d.Get("interface").(string)
	ip_addr.NoDad = strconv.FormatBool(d.Get("no_dad").(bool))

	return ip_addr
}

func writeToTerraform(d *schema.ResourceData, m interface{}, res *roscl.IPv6Address) error {

	disabled, _ := strconv.ParseBool(res.Disabled)
	invalid, _ := strconv.ParseBool(res.Invalid)
	eui64, _ := strconv.ParseBool(res.Eui64)
	no_dad, _ := strconv.ParseBool(res.NoDad)

	d.SetId(res.ID)
	d.Set("actual_interface", res.ActualInterface)
	d.Set("address", res.Address)
	d.Set("comment", res.Comment)
	d.Set("disabled", disabled)
	d.Set("eui64", eui64)
	d.Set("from_pool", res.FromPool)
	d.Set("interface", res.Interface)
	d.Set("invalid", invalid)
	d.Set("no_dad", no_dad)

	return nil
}
