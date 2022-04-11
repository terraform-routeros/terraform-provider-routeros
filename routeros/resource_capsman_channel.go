package routeros

import (
	"log"
	"strconv"

	roscl "github.com/gnewbury1/terraform-provider-routeros/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCapsManChannel() *schema.Resource {
	return &schema.Resource{
		Create: resourceCapsManChannelCreate,
		Read:   resourceCapsManChannelRead,
		Update: resourceCapsManChannelUpdate,
		Delete: resourceCapsManChannelDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"save_selected": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"width": {
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
			"band": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"extension_channel": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"frequency": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tx_power": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceCapsManChannelCreate(d *schema.ResourceData, m interface{}) error {

	c := m.(*roscl.Client)
	channel_obj := new(roscl.CapsManChannel)

	channel_obj.Name = d.Get("name").(string)
	channel_obj.SaveSelected = strconv.FormatBool(d.Get("save_selected").(bool))
	channel_obj.Width = d.Get("width").(string)
	channel_obj.Comment = d.Get("comment").(string)
	channel_obj.Band = d.Get("band").(string)
	channel_obj.ExtensionChannel = d.Get("extension_channel").(string)
	channel_obj.Frequency = strconv.Itoa(d.Get("frequency").(int))
	channel_obj.TXPower = strconv.Itoa(d.Get("tx_power").(int))

	res, err := c.CreateCapsManChannel(channel_obj)
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PUT request to the API")
		log.Fatal(err.Error())
		return err
	}

	d.SetId(res.ID)
	return resourceCapsManChannelRead(d, m)
}

func resourceCapsManChannelRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	channel, err := c.ReadCapsManChannel(d.Id())

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a GET request to the API")
		log.Fatal(err.Error())
		return err
	}

	frequency, _ := strconv.Atoi(channel.Frequency)
	save_selected, _ := strconv.ParseBool(channel.SaveSelected)
	tx_power, _ := strconv.Atoi(channel.TXPower)

	d.SetId(channel.ID)
	d.Set("name", channel.Name)
	d.Set("width", channel.Width)
	d.Set("save_selected", save_selected)
	d.Set("comment", channel.Comment)
	d.Set("band", channel.Band)
	d.Set("extension_channel", channel.ExtensionChannel)
	d.Set("frequency", frequency)
	d.Set("tx_power", tx_power)

	return nil

}

func resourceCapsManChannelUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	channel_obj := new(roscl.CapsManChannel)

	channel_obj.Name = d.Get("name").(string)
	channel_obj.SaveSelected = strconv.FormatBool(d.Get("save_selected").(bool))
	channel_obj.Width = d.Get("width").(string)
	channel_obj.Comment = d.Get("comment").(string)
	channel_obj.Band = d.Get("band").(string)
	channel_obj.ExtensionChannel = d.Get("extension_channel").(string)
	channel_obj.Frequency = strconv.Itoa(d.Get("frequency").(int))
	channel_obj.TXPower = strconv.Itoa(d.Get("tx_power").(int))

	res, err := c.UpdateCapsManChannel(d.Id(), channel_obj)

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PATCH request to the API")
		log.Fatal(err.Error())
		return err
	}

	d.SetId(res.ID)

	return nil
}

func resourceCapsManChannelDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	err := c.DeleteCapsManChannel(d.Id())
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a DELETE request to the API")
		log.Fatal(err.Error())
		return err
	}
	d.SetId("")
	return nil
}
