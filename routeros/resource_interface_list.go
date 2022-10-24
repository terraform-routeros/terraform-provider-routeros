package routeros

import (
	"log"
	"strconv"

	roscl "github.com/gnewbury1/terraform-provider-routeros/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceList() *schema.Resource {
	return &schema.Resource{
		Create: resourceInterfaceListCreate,
		Read:   resourceInterfaceListRead,
		Update: resourceInterfaceListUpdate,
		Delete: resourceInterfaceListDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"include": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"exclude": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func resourceInterfaceListCreate(d *schema.ResourceData, m interface{}) error {

	c := m.(*roscl.Client)
	list_obj := new(roscl.InterfaceList)

	list_obj.Name = d.Get("name").(string)
	list_obj.Include = d.Get("include").(string)
	list_obj.Exclude = d.Get("exclude").(string)
	list_obj.Comment = d.Get("comment").(string)

	res, err := c.CreateInterfaceList(list_obj)
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PUT request to the API")
		log.Fatal(err.Error())
		return err
	}

	d.SetId(res.ID)
	return resourceInterfaceListRead(d, m)
}

func resourceInterfaceListRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	list, err := c.ReadInterfaceList(d.Id())

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a GET request to the API")
		log.Fatal(err.Error())
		return err
	}

	dynamic, _ := strconv.ParseBool(list.Dynamic)

	d.SetId(list.ID)
	d.Set("name", list.Name)
	d.Set("include", list.Include)
	d.Set("dynamic", dynamic)
	d.Set("exclude", list.Exclude)
	d.Set("comment", list.Comment)
	return nil

}

func resourceInterfaceListUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	list_obj := new(roscl.InterfaceList)
	list_obj.Name = d.Get("name").(string)
	list_obj.Include = d.Get("include").(string)
	list_obj.Exclude = d.Get("exclude").(string)
	list_obj.Comment = d.Get("comment").(string)

	res, err := c.UpdateInterfaceList(d.Id(), list_obj)

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PATCH request to the API")
		log.Fatal(err.Error())
		return err
	}

	d.SetId(res.ID)

	return nil
}

func resourceInterfaceListDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	err := c.DeleteInterfaceList(d.Id())
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a DELETE request to the API")
		log.Fatal(err.Error())
		return err
	}
	d.SetId("")
	return nil
}
