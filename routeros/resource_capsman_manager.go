package routeros

import (
	"log"
	"strconv"

	roscl "github.com/gnewbury1/terraform-provider-routeros/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCapsManManager() *schema.Resource {
	return &schema.Resource{
		Create: resourceCapsManManagerUpdate,
		Read:   resourceCapsManManagerRead,
		Update: resourceCapsManManagerUpdate,
		Delete: resourceCapsManManagerDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"upgrade_policy": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "none",
			},
			"certificate": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "none",
			},
			"ca_certificate": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "none",
			},
			"require_peer_certificate": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"package_path": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceCapsManManagerRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	manager, err := c.ReadCapsManManager()

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a GET request to the API")
		log.Fatal(err.Error())
		return err
	}

	require_peer_certificate, _ := strconv.ParseBool(manager.RequirePeerCertificate)
	enabled, _ := strconv.ParseBool(manager.Enabled)

	d.SetId("1")
	d.Set("upgrade_policy", manager.UpgradePolicy)
	d.Set("enabled", enabled)
	d.Set("certificate", manager.Certificate)
	d.Set("ca_certificate", manager.CaCertificate)
	d.Set("require_peer_certificate", require_peer_certificate)
	d.Set("package_path", manager.PackagePath)

	return nil

}

func resourceCapsManManagerUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	manager_obj := new(roscl.CapsManManager)

	manager_obj.UpgradePolicy = d.Get("upgrade_policy").(string)
	manager_obj.Enabled = strconv.FormatBool(d.Get("enabled").(bool))
	manager_obj.Certificate = d.Get("certificate").(string)
	manager_obj.CaCertificate = d.Get("ca_certificate").(string)
	manager_obj.RequirePeerCertificate = strconv.FormatBool(d.Get("require_peer_certificate").(bool))
	manager_obj.PackagePath = d.Get("package_path").(string)

	_, err := c.UpdateCapsManManager(manager_obj)

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a POST request to the API")
		log.Fatal(err.Error())
		return err
	}

	d.SetId("1")

	return nil
}

func resourceCapsManManagerDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	manager_obj := new(roscl.CapsManManager)

	manager_obj.Enabled = "false"

	_, err := c.UpdateCapsManManager(manager_obj)

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a POST request to the API")
		log.Fatal(err.Error())
		return err
	}

	d.SetId("")

	return nil
}
