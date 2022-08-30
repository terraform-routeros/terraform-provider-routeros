package routeros

import (
	"context"
	"github.com/hashicorp/go-cty/cty"
	"log"
	"regexp"
	"strconv"

	roscl "github.com/gnewbury1/terraform-provider-routeros/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceInterfaceGre() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInterfaceGreCreate,
		ReadContext:   resourceInterfaceGreRead,
		UpdateContext: resourceInterfaceGreUpdate,
		DeleteContext: resourceInterfaceGreDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"remote_address": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsIPv4Address,
			},
			"actual_mtu": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"allow_fast_path": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether to allow FastPath processing. Must be disabled if IPsec tunneling is used.",
			},
			"clamp_tcp_mss": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
				Description: "Controls whether to change MSS size for received TCP SYN packets. When enabled, a " +
					"router will change the MSS size for received TCP SYN packets if the current MSS size exceeds the " +
					"tunnel interface MTU (taking into account the TCP/IP overhead). The received encapsulated packet " +
					"will still contain the original MSS, and only after decapsulation the MSS is changed.",
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dont_fragment": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "no",
				ValidateFunc: validation.StringInSlice([]string{"inherit", "no"}, false),
			},
			"dscp": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "inherit",
				ValidateDiagFunc: func(v any, p cty.Path) (diags diag.Diagnostics) {
					value := v.(string)

					// dscp (inherit | integer [0-63]; Default: '')
					if value == "" || value == "inherit" {
						return
					}

					i, err := strconv.Atoi(value)
					if err != nil {
						diags = diag.Errorf(
							"expected dscp value (%s) to be empty string or 'inherit' or integer 0..63", value)
						return
					}
					if i < 0 || i > 63 {
						diags = diag.Errorf(
							"expected %s to be in the range 0 - 63, got %d", value)
						return
					}
					return
				},
				Description: "Set dscp value in GRE header to a fixed value '0..63' or 'inherit' from dscp value taken " +
					"from tunnelled traffic.",
			},
			"ipsec_secret": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
				Description: "When secret is specified, router adds dynamic IPsec peer to remote-address with " +
					"pre-shared key and policy (by default phase2 uses sha1/aes128cbc).",
			},
			"keepalive": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "10s,10",
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`^\d+[smhdw]?(,\d+)?$`),
					"value must be integer[/time],integer 0..4294967295 (https://help.mikrotik.com/docs/display/ROS/GRE)"),
				Description: "Tunnel keepalive parameter sets the time interval in which the tunnel running flag will " +
					"remain even if the remote end of tunnel goes down. If configured time,retries fail, interface " +
					"running flag is removed. Parameters are written in following format: " +
					"KeepaliveInterval,KeepaliveRetries where KeepaliveInterval is time interval and " +
					"KeepaliveRetries - number of retry attempts. KeepaliveInterval is integer 0..4294967295",
			},
			"l2mtu": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Layer2 Maximum transmission unit.",
			},
			"local_address": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "0.0.0.0",
				ValidateFunc: validation.IsIPv4Address,
			},
			"mtu": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      1476,
				ValidateFunc: validation.IntBetween(0, 65535),
				Description:  "Layer3 Maximum transmission unit.",
			},
			"running": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func resourceInterfaceGreCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*roscl.Client)
	obj := marshalGreInterface(d)

	if boolFromMikrotikJSON(obj.AllowFastPath) && obj.IpsecSecret != "" {
		return diag.Errorf("can't enable fastpath together with ipsec")
	}

	res, err := c.CreateGRE(obj)
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PUT request to the API")
		return diag.FromErr(err)
	}
	d.SetId(res.ID)

	return resourceInterfaceGreRead(ctx, d, m)
}

func resourceInterfaceGreRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*roscl.Client)

	obj, err := c.ReadGRE(d.Id())
	if err != nil {
		log.Printf("[ERROR] An error was encountered while sending a GET request to the API: %v", err)
		return diag.FromErr(err)
	}

	return unmarshalGreInterface(obj, d)
}

func resourceInterfaceGreUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*roscl.Client)
	obj := marshalGreInterface(d)

	if boolFromMikrotikJSON(obj.AllowFastPath) && obj.IpsecSecret != "" {
		return diag.Errorf("can't enable fastpath together with ipsec")
	}

	res, err := c.UpdateGRE(d.Id(), obj)
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PATCH request to the API")
		return diag.FromErr(err)
	}
	d.SetId(res.ID)

	return nil
}

func resourceInterfaceGreDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*roscl.Client)

	err := c.DeleteGRE(d.Id())
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a DELETE request to the API")
		return diag.FromErr(err)
	}
	d.SetId("")

	return nil
}

func marshalGreInterface(d *schema.ResourceData) *roscl.GRE {
	return &roscl.GRE{
		AllowFastPath: boolToMikrotikJSON(d.Get("allow_fast_path").(bool)), //bool
		ClampTcpMss:   boolToMikrotikJSON(d.Get("clamp_tcp_mss").(bool)),   //bool
		Comment:       d.Get("comment").(string),
		Disabled:      d.Get("disabled").(bool), //bool
		DontFragment:  d.Get("dont_fragment").(string),
		Dscp:          d.Get("dscp").(string),
		IpsecSecret:   d.Get("ipsec_secret").(string),
		Keepalive:     d.Get("keepalive").(string),
		LocalAddress:  d.Get("local_address").(string),
		Mtu:           d.Get("mtu").(int), //int
		Name:          d.Get("name").(string),
		RemoteAddress: d.Get("remote_address").(string),
	}
}

func unmarshalGreInterface(obj *roscl.GRE, d *schema.ResourceData) diag.Diagnostics {
	var diags diag.Diagnostics

	if err := d.Set("disabled", obj.Disabled); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("name", obj.Name); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("remote_address", obj.RemoteAddress); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("actual_mtu", obj.ActualMtu); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("allow_fast_path", boolFromMikrotikJSON(obj.AllowFastPath)); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("clamp_tcp_mss", boolFromMikrotikJSON(obj.ClampTcpMss)); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("comment", obj.Comment); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("dont_fragment", obj.DontFragment); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("dscp", obj.Dscp); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("ipsec_secret", obj.IpsecSecret); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("keepalive", obj.Keepalive); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("l2mtu", obj.L2Mtu); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("local_address", obj.LocalAddress); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("mtu", obj.Mtu); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("running", obj.Running); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	d.SetId(obj.ID)

	return diags
}
