package routeros

import (
	"log"

	roscl "github.com/gnewbury1/terraform-provider-routeros/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCapsManProvisioning() *schema.Resource {
	return &schema.Resource{
		Create: resourceCapsManProvisioningCreate,
		Read:   resourceCapsManProvisioningRead,
		Update: resourceCapsManProvisioningUpdate,
		Delete: resourceCapsManProvisioningDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"common_name_regexp": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name_prefix": {
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"hw_supported_modes": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_address_ranges": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"identity_regexp": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"master_configuration": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name_format": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radio_mac": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"slave_configurations": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceCapsManProvisioningCreate(d *schema.ResourceData, m interface{}) error {

	c := m.(*roscl.Client)
	provisioning_obj := new(roscl.CapsManProvisioning)

	provisioning_obj.NamePrefix = d.Get("name_prefix").(string)
	provisioning_obj.Action = d.Get("action").(string)
	provisioning_obj.CommonNameRegexp = d.Get("common_name_regexp").(string)
	provisioning_obj.Comment = d.Get("comment").(string)
	provisioning_obj.HwSupportedModes = d.Get("hw_supported_modes").(string)
	provisioning_obj.IpAddressRanges = d.Get("ip_address_ranges").(string)
	provisioning_obj.IdentityRegexp = d.Get("identity_regexp").(string)
	provisioning_obj.MasterConfiguration = d.Get("master_configuration").(string)
	provisioning_obj.NameFormat = d.Get("name_format").(string)
	provisioning_obj.RadioMAC = d.Get("radio_mac").(string)
	provisioning_obj.SlaveConfigurations = d.Get("slave_configurations").(string)

	res, err := c.CreateCapsManProvisioning(provisioning_obj)
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PUT request to the API")
		log.Fatal(err.Error())
		return err
	}

	d.SetId(res.ID)
	return resourceCapsManProvisioningRead(d, m)
}

func resourceCapsManProvisioningRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	provisioning, err := c.ReadCapsManProvisioning(d.Id())

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a GET request to the API")
		log.Fatal(err.Error())
		return err
	}

	d.SetId(provisioning.ID)
	d.Set("name_prefix", provisioning.NamePrefix)
	d.Set("common_name_regexp", provisioning.CommonNameRegexp)
	d.Set("action", provisioning.Action)
	d.Set("comment", provisioning.Comment)
	d.Set("hw_supported_modes", provisioning.HwSupportedModes)
	d.Set("identity_regexp", provisioning.IdentityRegexp)
	d.Set("master_configuration", provisioning.MasterConfiguration)
	d.Set("name_format", provisioning.NameFormat)
	d.Set("radio_mac", provisioning.RadioMAC)
	d.Set("slave_configurations", provisioning.SlaveConfigurations)

	return nil

}

func resourceCapsManProvisioningUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	provisioning_obj := new(roscl.CapsManProvisioning)

	provisioning_obj.NamePrefix = d.Get("name_prefix").(string)
	provisioning_obj.Action = d.Get("action").(string)
	provisioning_obj.CommonNameRegexp = d.Get("common_name_regexp").(string)
	provisioning_obj.Comment = d.Get("comment").(string)
	provisioning_obj.HwSupportedModes = d.Get("hw_supported_modes").(string)
	provisioning_obj.IpAddressRanges = d.Get("ip_address_ranges").(string)
	provisioning_obj.IdentityRegexp = d.Get("identity_regexp").(string)
	provisioning_obj.MasterConfiguration = d.Get("master_configuration").(string)
	provisioning_obj.NameFormat = d.Get("name_format").(string)
	provisioning_obj.RadioMAC = d.Get("radio_mac").(string)
	provisioning_obj.SlaveConfigurations = d.Get("slave_configurations").(string)

	res, err := c.UpdateCapsManProvisioning(d.Id(), provisioning_obj)

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PATCH request to the API")
		log.Fatal(err.Error())
		return err
	}

	d.SetId(res.ID)

	return nil
}

func resourceCapsManProvisioningDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	err := c.DeleteCapsManProvisioning(d.Id())
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a DELETE request to the API")
		log.Fatal(err.Error())
		return err
	}
	d.SetId("")
	return nil
}
