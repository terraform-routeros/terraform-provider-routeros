package routeros

import (
	"log"
	"strconv"

	roscl "github.com/gnewbury1/terraform-provider-routeros/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceIPv6Addresses() *schema.Resource {
	return &schema.Resource{
		Read: datasourceIPv6AddressesRead,
		Schema: map[string]*schema.Schema{
			"addresses": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"actual_interface": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"advertise": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"comment": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"disabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"dynamic": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"eui64": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"from_pool": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"interface": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"invalid": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"link_local": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"no_dad": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func datasourceIPv6AddressesRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	res, err := c.ReadIPv6Addresses()

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a GET request to the API")
		log.Fatal(err.Error())
		return err
	}

	addresses := make([]map[string]interface{}, len(res))
	for k, v := range res {
		address := make(map[string]interface{})
		address["actual_interface"] = v.ActualInterface
		address["address"] = v.Address
		address["advertise"], _ = strconv.ParseBool(v.Advertise)
		address["comment"] = v.Comment
		address["disabled"], _ = strconv.ParseBool(v.Disabled)
		address["dynamic"], _ = strconv.ParseBool(v.Dynamic)
		address["eui64"], _ = strconv.ParseBool(v.Eui64)
		address["from_pool"] = v.FromPool
		address["interface"] = v.Interface
		address["invalid"], _ = strconv.ParseBool(v.Invalid)
		address["link_local"], _ = strconv.ParseBool(v.LinkLocal)
		address["no_dad"], _ = strconv.ParseBool(v.NoDad)
		addresses[k] = address
	}

	d.SetId(resource.UniqueId())
	if err := d.Set("addresses", addresses); err != nil {
		return err
	}

	return nil

}
