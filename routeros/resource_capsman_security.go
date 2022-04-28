package routeros

import (
	"log"

	roscl "github.com/gnewbury1/terraform-provider-routeros/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCapsManSecurity() *schema.Resource {
	return &schema.Resource{
		Create: resourceCapsManSecurityCreate,
		Read:   resourceCapsManSecurityRead,
		Update: resourceCapsManSecurityUpdate,
		Delete: resourceCapsManSecurityDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"group_encryption": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"authentication_types": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"eap_methods": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"eap_radius_accounting": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"encryption": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"passphrase": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_key_update": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tls_certificate": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tls_mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceCapsManSecurityCreate(d *schema.ResourceData, m interface{}) error {

	c := m.(*roscl.Client)
	security_obj := new(roscl.CapsManSecurity)

	security_obj.Name = d.Get("name").(string)
	security_obj.GroupEncryption = d.Get("group_encryption").(string)
	security_obj.AuthenticationTypes = d.Get("authentication_types").(string)
	security_obj.EapMethods = d.Get("eap_methods").(string)
	security_obj.EapRadiusAccounting = d.Get("eap_radius_accounting").(string)
	security_obj.Comment = d.Get("comment").(string)
	security_obj.Encryption = d.Get("encryption").(string)
	security_obj.Passphrase = d.Get("passphrase").(string)
	security_obj.GroupKeyUpdate = d.Get("group_key_update").(string)
	security_obj.TlsCertificate = d.Get("tls_certificate").(string)
	security_obj.TlsMode = d.Get("tls_mode").(string)

	res, err := c.CreateCapsManSecurity(security_obj)
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PUT request to the API")
		log.Fatal(err.Error())
		return err
	}

	d.SetId(res.ID)
	return resourceCapsManSecurityRead(d, m)
}

func resourceCapsManSecurityRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	security, err := c.ReadCapsManSecurity(d.Id())

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a GET request to the API")
		log.Fatal(err.Error())
		return err
	}

	d.SetId(security.ID)
	d.Set("name", security.Name)
	d.Set("authentication_types", security.AuthenticationTypes)
	d.Set("eap_methods", security.EapMethods)
	d.Set("eap_radius_accounting", security.EapRadiusAccounting)
	d.Set("group_encryption", security.GroupEncryption)
	d.Set("comment", security.Comment)
	d.Set("encryption", security.Encryption)
	d.Set("group_key_update", security.GroupKeyUpdate)
	d.Set("tls_certificate", security.TlsCertificate)
	d.Set("tls_mode", security.TlsMode)

	return nil

}

func resourceCapsManSecurityUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	security_obj := new(roscl.CapsManSecurity)

	security_obj.Name = d.Get("name").(string)
	security_obj.GroupEncryption = d.Get("group_encryption").(string)
	security_obj.AuthenticationTypes = d.Get("authentication_types").(string)
	security_obj.EapMethods = d.Get("eap_methods").(string)
	security_obj.EapRadiusAccounting = d.Get("eap_radius_accounting").(string)
	security_obj.Comment = d.Get("comment").(string)
	security_obj.Encryption = d.Get("encryption").(string)
	security_obj.Passphrase = d.Get("passphrase").(string)
	security_obj.GroupKeyUpdate = d.Get("group_key_update").(string)
	security_obj.TlsCertificate = d.Get("tls_certificate").(string)
	security_obj.TlsMode = d.Get("tls_mode").(string)

	res, err := c.UpdateCapsManSecurity(d.Id(), security_obj)

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PATCH request to the API")
		log.Fatal(err.Error())
		return err
	}

	d.SetId(res.ID)

	return nil
}

func resourceCapsManSecurityDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	err := c.DeleteCapsManSecurity(d.Id())
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a DELETE request to the API")
		log.Fatal(err.Error())
		return err
	}
	d.SetId("")
	return nil
}
