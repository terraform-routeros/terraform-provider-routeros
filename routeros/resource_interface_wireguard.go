package routeros

import (
	"log"
	"strconv"

	roscl "github.com/gnewbury1/terraform-provider-routeros/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceWireguard() *schema.Resource {
	return &schema.Resource{
		Create: resourceInterfaceWireguardCreate,
		Read:   resourceInterfaceWireguardRead,
		Update: resourceInterfaceWireguardUpdate,
		Delete: resourceInterfaceWireguardDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"listen_port": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"mtu": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1420,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"private_key": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"public_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"running": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func resourceInterfaceWireguardCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)

	interface_wireguard := new(roscl.InterfaceWireguard)
	interface_wireguard.Name = d.Get("name").(string)
	interface_wireguard.Disabled = strconv.FormatBool(d.Get("disabled").(bool))
	interface_wireguard.ListenPort = strconv.Itoa(d.Get("listen_port").(int))
	interface_wireguard.Mtu = strconv.Itoa(d.Get("mtu").(int))

	res, err := c.CreateInterfaceWireguard(interface_wireguard)
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PUT request to the API")
		log.Fatal(err.Error())
		return err
	}

	disabled, _ := strconv.ParseBool(res.Disabled)
	listen_port, _ := strconv.Atoi(res.ListenPort)
	mtu, _ := strconv.Atoi(res.Mtu)
	running, _ := strconv.ParseBool(res.Running)

	d.SetId(res.ID)
	d.Set("disabled", disabled)
	d.Set("listen_port", listen_port)
	d.Set("mtu", mtu)
	d.Set("running", running)
	d.Set("name", res.Name)
	d.Set("private_key", res.PrivateKey)
	d.Set("public_key", res.PublicKey)

	return nil
}

func resourceInterfaceWireguardRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	res, err := c.ReadInterfaceWireguard(d.Id())

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a GET request to the API")
		log.Fatal(err.Error())
		return err
	}

	disabled, _ := strconv.ParseBool(res.Disabled)
	listen_port, _ := strconv.Atoi(res.ListenPort)
	mtu, _ := strconv.Atoi(res.Mtu)
	running, _ := strconv.ParseBool(res.Running)

	d.SetId(res.ID)
	d.Set("disabled", disabled)
	d.Set("listen_port", listen_port)
	d.Set("mtu", mtu)
	d.Set("running", running)
	d.Set("name", res.Name)
	d.Set("private_key", res.PrivateKey)
	d.Set("public_key", res.PublicKey)

	return nil

}

func resourceInterfaceWireguardUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	interface_wireguard := new(roscl.InterfaceWireguard)
	interface_wireguard.Name = d.Get("name").(string)
	interface_wireguard.Disabled = strconv.FormatBool(d.Get("disabled").(bool))
	interface_wireguard.ListenPort = strconv.Itoa(d.Get("listen_port").(int))
	interface_wireguard.Mtu = strconv.Itoa(d.Get("mtu").(int))

	res, err := c.UpdateInterfaceWireguard(d.Id(), interface_wireguard)

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PATCH request to the API")
		log.Fatal(err.Error())
		return err
	}

	disabled, _ := strconv.ParseBool(res.Disabled)
	listen_port, _ := strconv.Atoi(res.ListenPort)
	mtu, _ := strconv.Atoi(res.Mtu)
	running, _ := strconv.ParseBool(res.Running)

	d.SetId(res.ID)
	d.Set("disabled", disabled)
	d.Set("listen_port", listen_port)
	d.Set("mtu", mtu)
	d.Set("running", running)
	d.Set("name", res.Name)
	d.Set("private_key", res.PrivateKey)
	d.Set("public_key", res.PublicKey)

	return nil
}

func resourceInterfaceWireguardDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	interface_wireguard, _ := c.ReadInterfaceWireguard(d.Id())
	err := c.DeleteInterfaceWireguard(interface_wireguard)
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a DELETE request to the API")
		log.Fatal(err.Error())
		return err
	}
	d.SetId("")
	return nil
}
