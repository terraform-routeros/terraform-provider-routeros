package routeros

import (
	"log"

	roscl "github.com/gnewbury1/terraform-provider-routeros/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemIdentity() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemIdentityCreate,
		Read:   resourceSystemIdentityRead,
		Update: resourceSystemIdentityUpdate,
		Delete: resourceSystemIdentityDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceSystemIdentityCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)

	system_identity := new(roscl.SystemIdentity)
	system_identity.Name = d.Get("name").(string)

	res, err := c.CreateSystemIdentity(system_identity)
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PUT request to the API")
		log.Fatal(err.Error())
		return err
	}

	d.SetId(res.Name)
	return nil
}

func resourceSystemIdentityRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	system_identity, err := c.ReadSystemIdentity()

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a GET request to the API")
		log.Fatal(err.Error())
		return err
	}

	d.SetId(system_identity.Name)
	d.Set("name", system_identity.Name)

	return nil
}

func resourceSystemIdentityUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)

	system_identity := new(roscl.SystemIdentity)
	system_identity.Name = d.Get("name").(string)

	res, err := c.UpdateSystemIdentity(system_identity)

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PATCH request to the API")
		log.Fatal(err.Error())
		return err
	}

	d.SetId(res.Name)

	return nil
}

func resourceSystemIdentityDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)

	system_identity, _ := c.ReadSystemIdentity()
	err := c.DeleteSystemIdentity(system_identity)

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a DELETE request to the API")
		log.Fatal(err.Error())
		return err
	}

	d.SetId("")

	return nil
}
