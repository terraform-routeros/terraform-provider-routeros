package routeros

import (
	"log"
	"strconv"

	roscl "github.com/gnewbury1/terraform-provider-routeros/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceIPAddresses() *schema.Resource {
	return &schema.Resource{
		Read: datasourceIPAddressesRead,
		Schema: map[string]*schema.Schema{
			"addresses": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"actual_interface": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"address": {
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
						"interface": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"invalid": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"network": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func datasourceIPAddressesRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	res, err := c.ReadIPAddresses()

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a GET request to the API")
		log.Fatal(err.Error())
		return err
	}

	addresses := make([]map[string]interface{}, len(res))
	for k, v := range res {
		address := make(map[string]interface{})
		address["id"] = v.ID
		address["actual_interface"] = v.ActualInterface
		address["address"] = v.Address
		address["disabled"], _ = strconv.ParseBool(v.Disabled)
		address["dynamic"], _ = strconv.ParseBool(v.Dynamic)
		address["interface"] = v.Interface
		address["invalid"], _ = strconv.ParseBool(v.Invalid)
		address["network"] = v.Network
		addresses[k] = address
	}

	d.SetId(resource.UniqueId())
	if err := d.Set("addresses", addresses); err != nil {
		return err
	}

	return nil

}
