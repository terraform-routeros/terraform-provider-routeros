package routeros

import (
	"log"

	roscl "github.com/gnewbury1/terraform-provider-routeros/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceListMember() *schema.Resource {
	return &schema.Resource{
		Create: resourceInterfaceListMemberCreate,
		Read:   resourceInterfaceListMemberRead,
		Update: resourceInterfaceListMemberUpdate,
		Delete: resourceInterfaceListMemberDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"interface": {
				Type:     schema.TypeString,
				Required: true,
			},
			"list": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceInterfaceListMemberCreate(d *schema.ResourceData, m interface{}) error {

	c := m.(*roscl.Client)
	member_obj := new(roscl.InterfaceListMember)

	member_obj.Interface = d.Get("interface").(string)
	member_obj.List = d.Get("list").(string)

	res, err := c.CreateInterfaceListMember(member_obj)
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PUT request to the API")
		log.Fatal(err.Error())
		return err
	}

	d.SetId(res.ID)
	return resourceInterfaceListMemberRead(d, m)
}

func resourceInterfaceListMemberRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	member, err := c.ReadInterfaceListMember(d.Id())

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a GET request to the API")
		log.Fatal(err.Error())
		return err
	}

	d.SetId(member.ID)
	d.Set("interface", member.Interface)
	d.Set("list", member.List)
	return nil

}

func resourceInterfaceListMemberUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	member_obj := new(roscl.InterfaceListMember)
	member_obj.Interface = d.Get("interface").(string)
	member_obj.List = d.Get("list").(string)

	res, err := c.UpdateInterfaceListMember(d.Id(), member_obj)

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PATCH request to the API")
		log.Fatal(err.Error())
		return err
	}

	d.SetId(res.ID)

	return nil
}

func resourceInterfaceListMemberDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	err := c.DeleteInterfaceListMember(d.Id())
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a DELETE request to the API")
		log.Fatal(err.Error())
		return err
	}
	d.SetId("")
	return nil
}
