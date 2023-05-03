package routeros

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/go-cty/cty"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ResourceInterfaceGre https://wiki.mikrotik.com/wiki/Manual:Interface/Gre
func ResourceInterfaceGre() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/gre"),
		MetaId:           PropId(Name),

		KeyActualMtu: PropActualMtuRo,
		"allow_fast_path": {
			Type:        schema.TypeBool,
			Optional:    true, // Must be present in the request so that the IPSEC PSK can be set correctly.
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
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
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
			ValidateDiagFunc: func(v interface{}, p cty.Path) (diags diag.Diagnostics) {
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
						"expected %s to be in the range 0 - 63, got %d", value, i)
					return
				}
				return
			},
			Description: "Set dscp value in GRE header to a fixed value '0..63' or 'inherit' from dscp value taken " +
				"from tunnelled traffic.",
		},
		"ipsec_secret": {
			Type:      schema.TypeString,
			Optional:  true,
			Default:   "",
			Sensitive: true,
			Description: "When secret is specified, router adds dynamic IPsec peer to remote-address with " +
				"pre-shared key and policy (by default phase2 uses sha1/aes128cbc).",
		},
		"keepalive": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "10s,10",
			ValidateFunc: validation.StringMatch(regexp.MustCompile(`^(\d+[smhdw]?)+(,\d+)?$`),
				"value must be integer[/time],integer 0..4294967295 (https://help.mikrotik.com/docs/display/ROS/GRE)"),
			Description: "Tunnel keepalive parameter sets the time interval in which the tunnel running flag will " +
				"remain even if the remote end of tunnel goes down. If configured time,retries fail, interface " +
				"running flag is removed. Parameters are written in following format: " +
				"KeepaliveInterval,KeepaliveRetries where KeepaliveInterval is time interval and " +
				"KeepaliveRetries - number of retry attempts. KeepaliveInterval is integer 0..4294967295",
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				if old == new {
					return true
				}

				if old == "" || new == "" {
					return false
				}

				o := strings.Split(old, ",")
				n := strings.Split(new, ",")
				if len(o) != 2 || len(n) != 2 {
					panic(fmt.Sprintf("[GRE keepalive] wrong keepalive format, old: '%v', new: '%v'", old, new))
				}

				// Compare keepalive retries.
				if o[1] != n[1] {
					return false
				}

				// Compare keepalive intervals.
				oDuration, err := ParseDuration(o[0])
				if err != nil {
					panic("[GRE keepalive] parse 'old' duration error: " + err.Error())
				}

				nDuration, err := ParseDuration(n[0])
				if err != nil {
					panic("[GRE keepalive] parse 'new' duration error: " + err.Error())
				}

				return oDuration.Seconds() == nDuration.Seconds()
			},
		},
		KeyL2Mtu: PropL2MtuRo,
		"local_address": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0.0.0.0",
			ValidateFunc: validation.IsIPv4Address,
		},
		KeyMtu:  PropMtuRw(),
		KeyName: PropNameForceNewRw,
		"remote_address": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.IsIPv4Address,
		},
		KeyRunning: PropRunningRo,
	}

	return &schema.Resource{
		CreateContext: DefaultValidateCreate(resSchema, func(d *schema.ResourceData) diag.Diagnostics {
			if d.Get("allow_fast_path").(bool) && d.Get("ipsec_secret").(string) != "" {
				return diag.Errorf("can't enable fastpath together with ipsec")
			}
			return nil
		}),
		ReadContext: DefaultRead(resSchema),
		UpdateContext: DefaultValidateUpdate(resSchema, func(d *schema.ResourceData) diag.Diagnostics {
			if d.Get("allow_fast_path").(bool) && d.Get("ipsec_secret").(string) != "" {
				return diag.Errorf("can't enable fastpath together with ipsec")
			}
			return nil
		}),
		DeleteContext: DefaultDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
