package routeros

import (
	"fmt"
	"strconv"
	"strings"

	roscl "github.com/gnewbury1/terraform-provider-routeros/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceWireguardPeer() *schema.Resource {
	return &schema.Resource{
		Create: resourceInterfaceWireguardPeerCreate,
		Read:   resourceInterfaceWireguardPeerRead,
		Update: resourceInterfaceWireguardPeerUpdate,
		Delete: resourceInterfaceWireguardPeerDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"allowed_addresses": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"current_endpoint_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"current_endpoint_port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"endpoint_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"endpoint_port": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": {
				Type:     schema.TypeString,
				Required: true,
			},
			"last_handshake": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"rx": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tx": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceInterfaceWireguardPeerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)

	allowed_addresses_input := strings.Join(ConvSInterfaceToSString(d.Get("allowed_addresses").([]interface{})), ",")

	interface_wireguard_peer := new(roscl.InterfaceWireguardPeer)
	interface_wireguard_peer.AllowedAddress = allowed_addresses_input
	interface_wireguard_peer.Disabled = strconv.FormatBool(d.Get("disabled").(bool))
	interface_wireguard_peer.EndpointAddress = d.Get("endpoint_address").(string)
	interface_wireguard_peer.EndpointPort = d.Get("endpoint_port").(string)
	interface_wireguard_peer.Interface = d.Get("interface").(string)
	interface_wireguard_peer.PublicKey = d.Get("public_key").(string)

	res, err := c.CreateInterfaceWireguardPeer(interface_wireguard_peer)
	if err != nil {
		return fmt.Errorf("error creating ip pool: %s", err.Error())
	}

	allowed_addresses_output := ConvSStringToSInterface(strings.Split(res.AllowedAddress, ","))
	current_endpoint_port, _ := strconv.Atoi(res.CurrentEndpointPort)
	disabled, _ := strconv.ParseBool(res.Disabled)

	d.SetId(res.ID)
	d.Set("allowed_addresses", allowed_addresses_output)
	d.Set("current_endpoint_address", res.CurrentEndpointAddress)
	d.Set("current_endpoint_port", current_endpoint_port)
	d.Set("disabled", disabled)
	d.Set("endpoint_address", res.EndpointAddress)
	d.Set("endpoint_port", res.EndpointPort)
	d.Set("interface", res.Interface)
	d.Set("last_handshake", res.LastHandshake)
	d.Set("public_key", res.PublicKey)
	d.Set("rx", res.Rx)
	d.Set("tx", res.Tx)

	return nil
}

func resourceInterfaceWireguardPeerRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	res, err := c.ReadInterfaceWireguardPeer(d.Id())

	if err != nil {
		return fmt.Errorf("error fetching ip pool: %s", err.Error())
	}

	allowed_addresses_output := ConvSStringToSInterface(strings.Split(res.AllowedAddress, ","))
	current_endpoint_port, _ := strconv.Atoi(res.CurrentEndpointPort)
	disabled, _ := strconv.ParseBool(res.Disabled)

	d.SetId(res.ID)
	d.Set("allowed_addresses", allowed_addresses_output)
	d.Set("current_endpoint_address", res.CurrentEndpointAddress)
	d.Set("current_endpoint_port", current_endpoint_port)
	d.Set("disabled", disabled)
	d.Set("endpoint_address", res.EndpointAddress)
	d.Set("endpoint_port", res.EndpointPort)
	d.Set("interface", res.Interface)
	d.Set("last_handshake", res.LastHandshake)
	d.Set("public_key", res.PublicKey)
	d.Set("rx", res.Rx)
	d.Set("tx", res.Tx)

	return nil

}

func resourceInterfaceWireguardPeerUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	allowed_addresses_input := strings.Join(ConvSInterfaceToSString(d.Get("allowed_addresses").([]interface{})), ",")

	interface_wireguard_peer := new(roscl.InterfaceWireguardPeer)
	interface_wireguard_peer.AllowedAddress = allowed_addresses_input
	interface_wireguard_peer.Disabled = strconv.FormatBool(d.Get("disabled").(bool))
	interface_wireguard_peer.EndpointAddress = d.Get("endpoint_address").(string)
	interface_wireguard_peer.EndpointPort = d.Get("endpoint_port").(string)
	interface_wireguard_peer.Interface = d.Get("interface").(string)
	interface_wireguard_peer.PublicKey = d.Get("public_key").(string)

	res, err := c.UpdateInterfaceWireguardPeer(d.Id(), interface_wireguard_peer)

	if err != nil {
		return fmt.Errorf("error updating ip address: %s", err.Error())
	}

	allowed_addresses_output := ConvSStringToSInterface(strings.Split(res.AllowedAddress, ","))
	current_endpoint_port, _ := strconv.Atoi(res.CurrentEndpointPort)
	disabled, _ := strconv.ParseBool(res.Disabled)

	d.SetId(res.ID)
	d.Set("allowed_addresses", allowed_addresses_output)
	d.Set("current_endpoint_address", res.CurrentEndpointAddress)
	d.Set("current_endpoint_port", current_endpoint_port)
	d.Set("disabled", disabled)
	d.Set("endpoint_address", res.EndpointAddress)
	d.Set("endpoint_port", res.EndpointPort)
	d.Set("interface", res.Interface)
	d.Set("last_handshake", res.LastHandshake)
	d.Set("public_key", res.PublicKey)
	d.Set("rx", res.Rx)
	d.Set("tx", res.Tx)

	return nil
}

func resourceInterfaceWireguardPeerDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	interface_wireguard_peer, _ := c.ReadInterfaceWireguardPeer(d.Id())
	err := c.DeleteInterfaceWireguardPeer(interface_wireguard_peer)
	if err != nil {
		return fmt.Errorf("error deleting ip address: %s", err.Error())
	}
	d.SetId("")
	return nil
}
