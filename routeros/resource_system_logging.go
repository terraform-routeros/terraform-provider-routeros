package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
    ".id": "*4",
    "action": "echo",
    "default": "true",
    "disabled": "false",
    "invalid": "false",
    "prefix": "",
    "topics": "critical"
}
*/

var validTopics = []string{
	"account", " bfd", "caps", "ddns", "dns", "error", "gsm", "info", "iscsi", "l2tp", "manager", "ntp", "packet",
	"pppoe", "radvd", "rip", "script", "smb", "sstp", "system", "timer", "vrrp", "web-proxy", "async", "bgp",
	"certificate", "debug", "dot1x", "dude", "event", "hotspot", "interface", "isdn", "ldp", "mme", "ospf", "pim",
	"pptp", "raw", "route", "sertcp", "snmp", "state", "telephony", "upnp", "warning", "wireless", "backup", "calc",
	"critical", "dhcp", "e-mail", "firewall", "igmp-proxy", "ipsec", "kvm", "lte", "mpls", "ovpn", "ppp", "radius",
	"read", "rsvp", "simulator", "ssh", "store", "tftp", "ups", "watchdog", "write",
}

// ResourceSystemLogging defines the resource for configuring logging rules
// https://wiki.mikrotik.com/wiki/Manual:System/Log
func ResourceSystemLogging() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/logging"),
		MetaId:           PropId(Id),
		"action": {
			Type:         schema.TypeString,
			Required:     true,
			Description:  "specifies one of the system default actions or user specified action listed in actions menu",
			ValidateFunc: validation.StringInSlice([]string{"disk", "echo", "memory", "remote"}, false),
		},
		"default": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"prefix": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "prefix added at the beginning of log messages",
			Default:     "",
		},
		KeyDisabled: {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Whether or not this logging should be disabled",
		},
		"invalid": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"topics": {
			Type:     schema.TypeSet,
			Elem:     &schema.Schema{Type: schema.TypeString, ValidateFunc: validation.StringInSlice(validTopics, false)},
			Optional: true,
			Description: `log all messages that falls into specified topic or list of topics.
						  '!' character can be used before topic to exclude messages falling under this topic. For example, we want to log NTP debug info without too much details:
						  /system logging add topics=ntp,debug,!packet`,
		},
	}

	return &schema.Resource{
		CreateContext: DefaultCreate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultUpdate(resSchema),
		DeleteContext: DefaultDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
