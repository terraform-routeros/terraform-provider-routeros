package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"regexp"
	"strconv"
)

// All metadata fields must be present in each resource schema, and the field type must be string.
const (
	MetaId           = "___id___"
	MetaResourcePath = "___path___"
)

const (
	KeyActualMtu  = "actual_mtu"
	KeyArp        = "arp"
	KeyArpTimeout = "arp_timeout"
	KeyComment    = "comment"
	KeyDynamic    = "dynamic"
	KeyDisabled   = "disabled"
	KeyInterface  = "interface"
	KeyInvalid    = "invalid"
	KeyL2Mtu      = "l2mtu"
	KeyMtu        = "mtu"
	KeyName       = "name"
	KeyRunning    = "running"
)

// PropResourcePath Resource path property.
func PropResourcePath(p string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Default:     p,
		Description: "<em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>",
		DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
			return true
		},
	}
}

// PropId Resource ID property.
func PropId(t IdType) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     int(t),
		Description: "<em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>",
		DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
			return true
		},
	}
}

// Schema properties.
var (
	PropActualMtuRo = &schema.Schema{
		Type:     schema.TypeInt,
		Computed: true,
	}
	PropArpRw = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "enabled",
		Description: "ARP resolution protocol mode.",
		ValidateFunc: validation.StringInSlice([]string{"disabled", "enabled", "local-proxy-arp", "proxy-arp",
			"reply-only"}, false),
	}
	PropArpTimeoutRw = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Default:  "auto",
		Description: "ARP timeout is time how long ARP record is kept in ARP table after no packets are received " +
			"from IP. Value auto equals to the value of arp-timeout in IP/Settings, default is 30s. Can use postfix " +
			"ms, s, M, h, d for milliseconds, seconds, minutes, hours or days. If no postfix is set then seconds (s) is used.",
		ValidateFunc: validation.StringMatch(regexp.MustCompile(`^$|auto$|(\d+(ms|s|M|h|d)?)+$`),
			"expected arp_timout value to be 'auto' string or time value"),
	}
	PropCommentRw = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	PropDisabledRw = &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
		Default:  false,
	}
	PropDynamicRo = &schema.Schema{
		Type:     schema.TypeBool,
		Computed: true,
		Description: "Configuration item created by software, not by management interface. It is not exported, " +
			"and cannot be directly modified.",
	}
	PropInterfaceRw = &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "Name of the interface.",
	}
	PropInvalidRo = &schema.Schema{
		Type:     schema.TypeBool,
		Computed: true,
	}
	PropL2MtuRo = &schema.Schema{
		Type:        schema.TypeInt,
		Computed:    true,
		Description: "Layer2 Maximum transmission unit.",
	}
	PropNameRw = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	PropRunningRo = &schema.Schema{
		Type:     schema.TypeBool,
		Computed: true,
	}
)

func PropMtuRw(defaultMtu int) *schema.Schema {
	return &schema.Schema{
		Type:         schema.TypeInt,
		Optional:     true,
		Default:      defaultMtu,
		ValidateFunc: validation.IntBetween(0, 65535),
		Description:  "Layer3 Maximum transmission unit",
	}
}

// Properties validation.
var (
	ValidationTime = validation.StringMatch(regexp.MustCompile(`^(\d+[smhdw]?)+$`),
		"value must be integer[/time],integer 0..4294967295")
	ValidationAutoYesNo = validation.StringInSlice([]string{"auto", "yes", "no"}, false)
	ValidationIpAddress = validation.StringMatch(
		regexp.MustCompile(`^$|^(\b(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(/([1-2][0-9]|3[0-2]))?)$`),
		"Allowed addresses must be a CIDR IP address or an empty string",
	)
)

// Properties DiffSuppressFunc.
var (
	TimeEquall = func(k, old, new string, d *schema.ResourceData) bool {
		if old == new {
			return true
		}

		if old == "" || new == "" {
			return false
		}

		// Compare intervals:
		oDuration, err := ParseDuration(old)
		if err != nil {
			panic("[TimeEquall] parse 'old' duration error: " + err.Error())
		}

		nDuration, err := ParseDuration(new)
		if err != nil {
			panic("[TimeEquall] parse 'new' duration error: " + err.Error())
		}

		return oDuration.Seconds() == nDuration.Seconds()
	}

	HexEqual = func(k, old, new string, d *schema.ResourceData) bool {
		if old == new {
			return true
		}

		if old == "" || new == "" {
			return false
		}

		// Compare numbers:
		var iOld, iNew int64
		var err error

		iOld, err = strconv.ParseInt(old, 0, 64)
		if err != nil {
			panic("[HexEqual] 'old' number parse error: " + err.Error())
		}

		iNew, err = strconv.ParseInt(new, 0, 64)
		if err != nil {
			panic("[HexEqual] 'new' number parse error: " + err.Error())
		}

		return iOld == iNew
	}
)
