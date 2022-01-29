package routeros

import (
	"log"
	"strconv"
	"strings"

	roscl "github.com/gnewbury1/terraform-provider-routeros/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIPv6NeighborDiscovery() *schema.Resource {
	return &schema.Resource{
		Create: resourceIPv6NeighborDiscoveryCreate,
		Read:   resourceIPv6NeighborDiscoveryRead,
		Update: resourceIPv6NeighborDiscoveryUpdate,
		Delete: resourceIPv6NeighborDiscoveryDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"advertise_dns": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"advertise_mac_address": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"default": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"dns": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"hop_limit": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"interface": {
				Type:     schema.TypeString,
				Required: true,
			},
			"invalid": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"managed_address_configuration": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"mtu": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"other_configuration": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ra_delay": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ra_interval": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ra_lifetime": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"reachable_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"retransmit_interval": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceIPv6NeighborDiscoveryCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	nd := new(roscl.IPv6NeighborDiscovery)

	nd.AdvertiseDNS = strconv.FormatBool(d.Get("advertise_dns").(bool))
	nd.AdvertiseMACAddress = strconv.FormatBool(d.Get("advertise_mac_address").(bool))
	nd.Disabled = strconv.FormatBool(d.Get("disabled").(bool))
	nd.DNS = strings.Join(ConvSInterfaceToSString(d.Get("tagged").([]interface{})), ",")
	nd.HopLimit = d.Get("hop_limit").(string)
	nd.Interface = d.Get("interface").(string)

	nd.ManagedAddressConfiguration = strconv.FormatBool(d.Get("managed_address_configuration").(bool))

	nd.MTU = d.Get("mtu").(string)
	nd.OtherConfiguration = strconv.FormatBool(d.Get("other_configuration").(bool))
	nd.RADelay = d.Get("ra_delay").(string)
	nd.RAInterval = d.Get("ra_interval").(string)
	nd.RALifetime = d.Get("ra_lifetime").(string)
	nd.ReachableTime = d.Get("reachable_time").(string)
	nd.RetransmitInterval = d.Get("retransmit_interval").(string)
	res, err := c.CreateIPv6NeighborDiscovery(nd)
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PUT request to the API")
		log.Fatal(err.Error())
		return err
	}

	advertiseDNS, _ := strconv.ParseBool(res.AdvertiseDNS)
	advertiseMACAddress, _ := strconv.ParseBool(res.AdvertiseMACAddress)
	def, _ := strconv.ParseBool(res.Default)
	disabled, _ := strconv.ParseBool(res.Disabled)
	invalid, _ := strconv.ParseBool(res.Invalid)
	managedAddressConfiguration, _ := strconv.ParseBool(res.ManagedAddressConfiguration)
	otherConfiguration, _ := strconv.ParseBool(res.OtherConfiguration)

	d.SetId(res.ID)
	d.Set("advertise_dns", advertiseDNS)
	d.Set("advertise_mac_address", advertiseMACAddress)
	d.Set("default", def)
	d.Set("disabled", disabled)
	d.Set("dns", ConvSStringToSInterface(strings.Split(res.DNS, ",")))
	d.Set("hop_limit", res.HopLimit)
	d.Set("interface", res.Interface)
	d.Set("invalid", invalid)
	d.Set("managed_address_configuration", managedAddressConfiguration)
	d.Set("mtu", res.MTU)
	d.Set("other_configuration", otherConfiguration)
	d.Set("ra_delay", res.RADelay)
	d.Set("ra_interval", res.RAInterval)
	d.Set("ra_lifetime", res.RALifetime)
	d.Set("reachable_time", res.ReachableTime)
	d.Set("retransmit_interval", res.RetransmitInterval)
	return nil
}

func resourceIPv6NeighborDiscoveryRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	res, err := c.GetIPv6NeighborDiscovery(d.Id())

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a GET request to the API")
		log.Fatal(err.Error())
		return err
	}

	advertiseDNS, _ := strconv.ParseBool(res.AdvertiseDNS)
	advertiseMACAddress, _ := strconv.ParseBool(res.AdvertiseMACAddress)
	def, _ := strconv.ParseBool(res.Default)
	disabled, _ := strconv.ParseBool(res.Disabled)
	invalid, _ := strconv.ParseBool(res.Invalid)
	managedAddressConfiguration, _ := strconv.ParseBool(res.ManagedAddressConfiguration)
	otherConfiguration, _ := strconv.ParseBool(res.OtherConfiguration)

	d.SetId(res.ID)
	d.Set("advertise_dns", advertiseDNS)
	d.Set("advertise_mac_address", advertiseMACAddress)
	d.Set("default", def)
	d.Set("disabled", disabled)
	d.Set("dns", ConvSStringToSInterface(strings.Split(res.DNS, ",")))
	d.Set("hop_limit", res.HopLimit)
	d.Set("interface", res.Interface)
	d.Set("invalid", invalid)
	d.Set("managed_address_configuration", managedAddressConfiguration)
	d.Set("mtu", res.MTU)
	d.Set("other_configuration", otherConfiguration)
	d.Set("ra_delay", res.RADelay)
	d.Set("ra_interval", res.RAInterval)
	d.Set("ra_lifetime", res.RALifetime)
	d.Set("reachable_time", res.ReachableTime)
	d.Set("retransmit_interval", res.RetransmitInterval)

	return nil

}

func resourceIPv6NeighborDiscoveryUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	nd := new(roscl.IPv6NeighborDiscovery)

	nd.AdvertiseDNS = strconv.FormatBool(d.Get("advertise_dns").(bool))
	nd.AdvertiseMACAddress = strconv.FormatBool(d.Get("advertise_mac_address").(bool))
	nd.Disabled = strconv.FormatBool(d.Get("disabled").(bool))
	nd.DNS = strings.Join(ConvSInterfaceToSString(d.Get("dns").([]interface{})), ",")
	nd.HopLimit = d.Get("hop_limit").(string)
	nd.Interface = d.Get("interface").(string)
	nd.Invalid = strconv.FormatBool(d.Get("invalid").(bool))
	nd.ManagedAddressConfiguration = strconv.FormatBool(d.Get("managed_address_configuration").(bool))
	nd.MTU = d.Get("mtu").(string)
	nd.OtherConfiguration = strconv.FormatBool(d.Get("other_configuration").(bool))
	nd.RADelay = d.Get("ra_delay").(string)
	nd.RAInterval = d.Get("ra_interval").(string)
	nd.RALifetime = d.Get("ra_lifetime").(string)
	nd.ReachableTime = d.Get("reachable_time").(string)
	nd.RetransmitInterval = d.Get("retransmit_interval").(string)

	res, err := c.UpdateIPv6NeighborDiscovery(d.Id(), nd)

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PATCH request to the API")
		log.Fatal(err.Error())
		return err
	}

	advertiseDNS, _ := strconv.ParseBool(res.AdvertiseDNS)
	advertiseMACAddress, _ := strconv.ParseBool(res.AdvertiseMACAddress)
	def, _ := strconv.ParseBool(res.Default)
	disabled, _ := strconv.ParseBool(res.Disabled)
	invalid, _ := strconv.ParseBool(res.Invalid)
	managedAddressConfiguration, _ := strconv.ParseBool(res.ManagedAddressConfiguration)
	otherConfiguration, _ := strconv.ParseBool(res.OtherConfiguration)

	d.SetId(res.ID)
	d.Set("advertise_dns", advertiseDNS)
	d.Set("advertise_mac_address", advertiseMACAddress)
	d.Set("default", def)
	d.Set("disabled", disabled)
	d.Set("dns", ConvSStringToSInterface(strings.Split(res.DNS, ",")))
	d.Set("hop_limit", res.HopLimit)
	d.Set("interface", res.Interface)
	d.Set("invalid", invalid)
	d.Set("managed_address_configuration", managedAddressConfiguration)
	d.Set("mtu", res.MTU)
	d.Set("other_configuration", otherConfiguration)
	d.Set("ra_delay", res.RADelay)
	d.Set("ra_interval", res.RAInterval)
	d.Set("ra_lifetime", res.RALifetime)
	d.Set("reachable_time", res.ReachableTime)
	d.Set("retransmit_interval", res.RetransmitInterval)

	return nil
}

func resourceIPv6NeighborDiscoveryDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	err := c.DeleteIPv6NeighborDiscovery(d.Id())
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a DELETE request to the API")
		log.Fatal(err.Error())
		return err
	}
	d.SetId("")
	return nil
}
