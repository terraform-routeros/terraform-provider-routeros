package routeros

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceInterfaceVrrpV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			MetaResourcePath: PropResourcePath("/interface/vrrp"),
			MetaId:           PropId(Name),

			KeyArp:        PropArpRw,
			KeyArpTimeout: PropArpTimeoutRw,
			"authentication": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "none",
				Description:  "Authentication method to use for VRRP advertisement packets.",
				ValidateFunc: validation.StringInSlice([]string{"ah", "none", "simple"}, false),
			},
			KeyComment:  PropCommentRw,
			KeyDisabled: PropDisabledRw,
			"group_master": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Allows combining multiple VRRP interfaces to maintain the same VRRP status within the group.",
				// Maybe this is a bug, but for the 'none' value, the Mikrotik ROS 7.5 returns an empty string.
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if old == new {
						return true
					}

					if new == "none" && old == "" {
						return true
					}
					return false
				},
			},
			KeyInterface: PropInterfaceRw,
			"interval": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "1s",
				Description: "VRRP update interval in seconds. Defines how often master sends advertisement packets.",
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`^(\d+(ms|s|m)?)+$`),
					"expected hello interval 10ms..4m15s"),
			},
			"invalid": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			KeyMacAddress: PropMacAddressRo,
			KeyMtu:        PropMtuRw(),
			KeyName:       PropNameForceNewRw,
			"on_fail": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Script to execute when the node fails.",
			},
			"on_backup": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Script to execute when the node is switched to the backup state.",
			},
			"on_master": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Script to execute when the node is switched to master state.",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "Password required for authentication. Can be ignored if authentication is not used.",
			},
			"preemption_mode": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
				Description: "Whether the master node always has the priority. When set to 'no' the backup node will not " +
					"be elected to be a master until the current master fails, even if the backup node has higher priority " +
					"than the current master. This setting is ignored if the owner router becomes available",
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
				Description: "Priority of VRRP node used in Master election algorithm. A higher number means higher " +
					"priority. '255' is reserved for the router that owns VR IP and '0' is reserved for the Master router " +
					"to indicate that it is releasing responsibility.",
				ValidateFunc: validation.IntBetween(1, 254),
			},
			"remote_address": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "Specifies the remote address of the other VRRP router for syncing connection tracking. " +
					"If not set, the system autodetects the remote address via VRRP. The remote address is used only if " +
					"sync-connection-tracking=yes.Sync connection tracking uses UDP port 8275.",
				ValidateFunc: validation.IsIPv4Address,
			},
			"running": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"sync_connection_tracking": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Synchronize connection tracking entries from Master to Backup device.",
			},
			"v3_protocol": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "ipv4",
				Description:  "A protocol that will be used by VRRPv3. Valid only if the version is 3.",
				ValidateFunc: validation.StringInSlice([]string{"ipv4", "ipv6"}, false),
			},
			"version": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      3,
				Description:  "Which VRRP version to use.",
				ValidateFunc: validation.IntBetween(2, 3),
			},
			"vrid": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      1,
				Description:  "Virtual Router identifier. Each Virtual router must have a unique id number.",
				ValidateFunc: validation.IntBetween(1, 255),
			},
		},
	}
}
