package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// https://help.mikrotik.com/docs/display/ROS/RADIUS#RADIUS-RADIUSClient
func ResourceRadius() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/radius"),
		MetaId:           PropId(Id),

		"accounting_backup": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "An option whether the configuration is for the backup RADIUS server.",
		},
		"accounting_port": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      1813,
			Description:  "RADIUS server port used for accounting.",
			ValidateFunc: Validation64k,
		},
		"address": {
			Type:         schema.TypeString,
			Required:     true,
			Description:  "IPv4 or IPv6 address of RADIUS server.",
			ValidateFunc: validation.IsIPAddress,
		},
		"authentication_port": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      1812,
			Description:  "RADIUS server port used for authentication.",
			ValidateFunc: Validation64k,
		},
		"called_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "RADIUS calling station identifier.",
		},
		"certificate": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "none",
			Description: "Certificate to use for communication with RADIUS Server with RadSec enabled.",
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"domain": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Microsoft Windows domain of client passed to RADIUS servers that require domain validation.",
		},
		"protocol": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "udp",
			Description:  "An option specifies the protocol to use when communicating with the RADIUS Server.",
			ValidateFunc: validation.StringInSlice([]string{"radsec", "udp"}, false),
		},
		"realm": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Explicitly stated realm (user domain), so the users do not have to provide proper ISP domain name in the user name.",
		},
		"require_message_auth": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "An option whether to require `Message-Authenticator` in received Access-Accept/Challenge/Reject messages.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			ValidateFunc:     validation.StringInSlice([]string{"no", "yes-for-request-resp"}, false),
		},
		"secret": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "The shared secret to access the RADIUS server.",
		},
		"service": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"hotspot", "login", "ppp", "wireless",
					"dhcp", "ipsec", "dot1x"}, false),
			},
			Description: "A set of router services that will use the RADIUS server. Possible values: " +
				"`hotspot`, `login`, `ppp`, `wireless`, `dhcp`, `ipsec`, `dot1x`.",
		},
		"src_address": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Source IPv4/IPv6 address of the packets sent to the RADIUS server.",
			ValidateFunc: validation.IsIPAddress,
		},
		"timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "300ms",
			Description:      "A timeout, after which the request should be resent.",
			DiffSuppressFunc: TimeEquall,
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

		Schema:        resSchema,
		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    ResourceRadiusV0().CoreConfigSchema().ImpliedType(),
				Upgrade: stateMigrationScalarToList("service"),
				Version: 0,
			},
		},
	}
}

// https://help.mikrotik.com/docs/display/ROS/RADIUS#RADIUS-ConnectionTerminatingfromRADIUS
func ResourceRadiusIncoming() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/radius/incoming"),
		MetaId:           PropId(Name),

		"accept": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "An option whether to accept the unsolicited messages.",
		},
		"port": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      3799,
			Description:  "The port number to listen for the requests on.",
			ValidateFunc: Validation64k,
		},
		"vrf": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "VRF on which service is listening for incoming connections. This option is available in RouterOS starting from version 7.4.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
	}

	return &schema.Resource{
		CreateContext: DefaultSystemCreate(resSchema),
		ReadContext:   DefaultSystemRead(resSchema),
		UpdateContext: DefaultSystemUpdate(resSchema),
		DeleteContext: DefaultSystemDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
