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

// All metadata fields must be present in each resource schema, and the field type must be string.
const (
	MetaId             = "___id___"
	MetaResourcePath   = "___path___"
	MetaTransformSet   = "___ts___"
	MetaSkipFields     = "___skip___"
	MetaSetUnsetFields = "___unset___"
)

const (
	KeyActualMtu               = "actual_mtu"
	KeyAllowFastPath           = "allow_fast_path"
	KeyArp                     = "arp"
	KeyArpTimeout              = "arp_timeout"
	KeyClampTcpMss             = "clamp_tcp_mss"
	KeyComment                 = "comment"
	KeyDynamic                 = "dynamic"
	KeyDisabled                = "disabled"
	KeyDontFragment            = "dont_fragment"
	KeyDscp                    = "dscp"
	KeyFilter                  = "filter"
	KeyInactive                = "inactive"
	KeyInterface               = "interface"
	KeyInvalid                 = "invalid"
	KeyIpsecSecret             = "ipsec_secret"
	KeyKeepalive               = "keepalive"
	KeyL2Mtu                   = "l2mtu"
	KeyLocalAddress            = "local_address"
	KeyLoopProtect             = "loop_protect"
	KeyLoopProtectDisableTime  = "loop_protect_disable_time"
	KeyLoopProtectSendInterval = "loop_protect_send_interval"
	KeyLoopProtectStatus       = "loop_protect_status"
	KeyMacAddress              = "mac_address"
	KeyMtu                     = "mtu"
	KeyName                    = "name"
	KeyPlaceBefore             = "place_before"
	KeyRemoteAddress           = "remote_address"
	KeyRunning                 = "running"
	KeyVrf                     = "vrf"
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

// PropTransformSet
func PropTransformSet(s string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Default:     s,
		Description: "<em>A set of transformations for field names. This is an internal service field, setting a value is not required.</em>",
		DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
			return true
		},
	}
}

// PropSkipFields
func PropSkipFields(s string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Default:     s,
		Description: "<em>A set of transformations for field names. This is an internal service field, setting a value is not required.</em>",
		DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
			return true
		},
	}
}

// PropSetUnsetFields
func PropSetUnsetFields(s string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Default:     s,
		Description: "<em>A set of fields that require setting/unsetting. This is an internal service field, setting a value is not required.</em>",
		DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
			return true
		},
	}
}

// PropName
func PropName(description string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: description,
	}
}

// Schema properties.
var (
	PropActualMtuRo = &schema.Schema{
		Type:     schema.TypeInt,
		Computed: true,
	}
	PropAllowFastPathRw = &schema.Schema{
		Type:        schema.TypeBool,
		Optional:    true, // Must be present in the request so that the IPSEC PSK can be set correctly.
		Default:     true,
		Description: "Whether to allow FastPath processing. Must be disabled if IPsec tunneling is used.",
	}
	PropArpRw = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Default:  "enabled",
		Description: `Address Resolution Protocol mode:
	* disabled - the interface will not use ARP
	* enabled - the interface will use ARP
	* local-proxy-arp - the router performs proxy ARP on the interface and sends replies to the same interface
	* proxy-arp - the router performs proxy ARP on the interface and sends replies to other interfaces
	* reply-only - the interface will only reply to requests originated from matching IP address/MAC address combinations which are entered as static entries in the ARP table. No dynamic entries will be automatically stored in the ARP table. Therefore for communications to be successful, a valid static entry must already exist.`,
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
	PropClampTcpMssRw = &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
		Default:  true,
		Description: "Controls whether to change MSS size for received TCP SYN packets. When enabled, a " +
			"router will change the MSS size for received TCP SYN packets if the current MSS size exceeds the " +
			"tunnel interface MTU (taking into account the TCP/IP overhead). The received encapsulated packet " +
			"will still contain the original MSS, and only after decapsulation the MSS is changed.",
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
	PropDontFragmentRw = &schema.Schema{
		Type:         schema.TypeString,
		Optional:     true,
		Default:      "no",
		ValidateFunc: validation.StringInSlice([]string{"inherit", "no"}, false),
	}
	PropDscpRw = &schema.Schema{
		// dscp (inherit | integer [0-63]; Default: '')
		Type:     schema.TypeString,
		Optional: true,
		Default:  "inherit",
		ValidateDiagFunc: func(v interface{}, p cty.Path) (diags diag.Diagnostics) {
			value := v.(string)

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
	}
	PropDynamicRo = &schema.Schema{
		Type:     schema.TypeBool,
		Computed: true,
		Description: "Configuration item created by software, not by management interface. It is not exported, " +
			"and cannot be directly modified.",
	}
	PropFilterRw = &schema.Schema{
		Type:        schema.TypeMap,
		Optional:    true,
		Elem:        schema.TypeString,
		Description: "Additional request filtering options.",
	}
	PropInactiveRo = &schema.Schema{
		Type:     schema.TypeBool,
		Computed: true,
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
	PropIpsecSecretRw = &schema.Schema{
		Type:      schema.TypeString,
		Optional:  true,
		Default:   "",
		Sensitive: true,
		Description: "When secret is specified, router adds dynamic IPsec peer to remote-address with " +
			"pre-shared key and policy (by default phase2 uses sha1/aes128cbc).",
	}
	PropKeepaliveRw = &schema.Schema{
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
	}
	PropL2MtuRo = &schema.Schema{
		Type:     schema.TypeInt,
		Computed: true,
		Description: "Layer2 Maximum transmission unit. " +
			"[See](https://wiki.mikrotik.com/wiki/Maximum_Transmission_Unit_on_RouterBoards).",
	}
	PropLocalAddressRw = &schema.Schema{
		Type:         schema.TypeString,
		Optional:     true,
		Default:      "0.0.0.0",
		Description:  "Source address of the tunnel packets, local on the router.",
		ValidateFunc: validation.IsIPv4Address,
	}
	PropLoopProtectRw = &schema.Schema{
		Type:         schema.TypeString,
		Optional:     true,
		Default:      "default",
		ValidateFunc: validation.StringInSlice([]string{"default", "on", "off"}, false),
	}
	PropLoopProtectDisableTimeRw = &schema.Schema{
		Type:             schema.TypeString,
		Optional:         true,
		Default:          "5m",
		ValidateFunc:     ValidationTime,
		DiffSuppressFunc: TimeEquall,
	}
	PropLoopProtectSendIntervalRw = &schema.Schema{
		Type:             schema.TypeString,
		Optional:         true,
		Default:          "5s",
		ValidateFunc:     ValidationTime,
		DiffSuppressFunc: TimeEquall,
	}
	PropLoopProtectStatusRo = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
	PropMacAddressRo = &schema.Schema{
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Current mac address.",
	}
	// TODO: Replace in all possible resources with a property without 'ForceNew'.
	// https://github.com/orgs/terraform-routeros/discussions/192#discussioncomment-5929999
	PropNameForceNewRw = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
		Description: `Changing the name of this resource will force it to be recreated.
	> The links of other configuration properties to this resource may be lost!
	> Changing the name of the resource outside of a Terraform will result in a loss of control integrity for that resource!
`,
	}
	PropPlaceBefore = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		ForceNew: true,
		Description: `Before which position the rule will be inserted.  
	> Please check the effect of this option, as it does not work as you think!  
	> Best way to use in conjunction with a data source. See [example](../data-sources/firewall.md#example-usage).  
`,
	}
	PropRemoteAddressRw = &schema.Schema{
		Type:         schema.TypeString,
		Optional:     true,
		Default:      "0.0.0.0",
		Description:  "IP address of the remote end of the tunnel.",
		ValidateFunc: validation.IsIPv4Address,
	}
	PropRunningRo = &schema.Schema{
		Type:     schema.TypeBool,
		Computed: true,
	}
	PropVrfRw = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "main",
		Description: "The VRF table this resource operates on.",
	}
)

// PropMtuRw MTU value can be integer or 'auto'.
func PropMtuRw() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
			v := i.(string)
			if v == "auto" {
				return nil
			}

			mtu, err := strconv.ParseInt(v, 0, 64)
			if err != nil {
				return diag.Diagnostics{
					{
						Severity: diag.Error,
						Summary:  "Expected MTU value to be integer or 'auto'",
					},
				}
			}

			if mtu < 0 || mtu > 65535 {
				return diag.Diagnostics{
					{
						Severity: diag.Error,
						Summary:  "Expected MTU value to be in the range (0 - 65535), got " + v,
					},
				}
			}

			return nil
		},
		Description: "Layer3 Maximum transmission unit ('auto', 0 .. 65535)",
	}
}

// Properties validation.
var (
	ValidationTime = validation.StringMatch(regexp.MustCompile(`^(\d+([smhdw]|ms)?)+$`),
		"value should be an integer or a time interval: 0..4294967295 (seconds) or 500ms, 2d, 1w")
	ValidationAutoYesNo = validation.StringInSlice([]string{"auto", "yes", "no"}, false)
	ValidationIpAddress = validation.StringMatch(
		regexp.MustCompile(`^$|^!?(\b(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(/([0-9]|[0-9]|[1-2][0-9]|3[0-2]))?)$`),
		"Allowed addresses should be a CIDR IP address or an empty string",
	)
	ValidationMacAddress = validation.StringMatch(
		regexp.MustCompile(`^!?\b(?:[0-9A-F]{2}\:){5}(?:[0-9A-F]{2})$`),
		"Allowed MAC addresses should be [!]AA:BB:CC:DD:EE:FF",
	)

	// ValidationMultiValInSlice returns a SchemaValidateDiagFunc which works like the StringInSlice function,
	// but the provided value can be a single value or a comma-separated list of values.
	// The negative indication of the parameter is also supported by adding "!" before value if mikrotikNegative is true.
	ValidationMultiValInSlice = func(valid []string, ignoreCase, mikrotikNegative bool) schema.SchemaValidateDiagFunc {
		return func(v interface{}, path cty.Path) (diags diag.Diagnostics) {
			val, ok := v.(string)

			if !ok {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Bad value type",
					Detail:   fmt.Sprintf("Value should be a string: %v (type = %T)", val, val),
				})

				return
			}

			if mikrotikNegative {
				for _, v := range valid {
					valid = append(valid, "!"+v)
				}
			}

			for _, sValue := range strings.Split(val, ",") {
				ok := false
				sValue = strings.TrimSpace(sValue)

				for _, sValid := range valid {
					if sValue == sValid || (ignoreCase && strings.EqualFold(sValue, sValid)) {
						ok = true
						break
					}
				}

				if !ok {
					diags = append(diags, diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "Bad value",
						Detail:   fmt.Sprintf("Unexpected value: %v", sValue),
					})
				}
			}

			return
		}
	}
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

	// AlwaysPresentNotUserProvided is a SupressDiff function that prevents values not provided by users to get updated.
	// This is necessary in some system-wide fields that are present regardless if the users provides any values.
	// Prevents the need of hardcode values for default values, as those are harder to track over time/versions of
	// routeros
	AlwaysPresentNotUserProvided = func(k, old, new string, d *schema.ResourceData) bool {
		if old != "" && d.GetRawConfig().GetAttr(k).IsNull() {
			return true
		}
		return false
	}
)

func buildReadFilter(m map[string]interface{}) []string {
	var res []string

	for fieldName, fieldValue := range m {
		res = append(res, fmt.Sprintf("%v=%v", fieldName, fieldValue))
	}

	return res
}

// Diagnostics
var DeleteSystemObject = []diag.Diagnostic{{
	Severity: diag.Warning,
	Summary:  "Delete operation on a system object.",
	Detail: "This resource contains system settings and cannot be deleted or reset. " +
		"This action will remove the object from the Terraform state. " +
		"See also: 'terraform state rm' https://developer.hashicorp.com/terraform/cli/commands/state/rm",
}}
