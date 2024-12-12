package routeros

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*54",
    "arp": "enabled",
    "arp-timeout": "auto",
    "authentication": "none",
    "disabled": "false",
    "group-master": "",
    "interface": "vlan55",
    "interval": "1s",
    "invalid": "false",
    "mac-address": "00:00:5E:00:01:01",
    "master": "true",
    "mtu": "1500",
    "name": "vrrp1",
    "on-backup": "",
    "on-fail": "",
    "on-master": "",
    "password": "",
    "preemption-mode": "true",
    "priority": "100",
    "running": "true",
    "sync-connection-tracking": "false",
    "v3-protocol": "ipv4",
    "version": "3",
    "vrid": "1"
  }
*/

// ResourceInterfaceVrrp https://help.mikrotik.com/docs/display/ROS/VRRP
func ResourceInterfaceVrrp() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/vrrp"),
		MetaId:           PropId(Id),

		KeyArp:        PropArpRw,
		KeyArpTimeout: PropArpTimeoutRw,
		"authentication": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Authentication method to use for VRRP advertisement packets.",
			ValidateFunc:     validation.StringInSlice([]string{"ah", "none", "simple"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"group_authority": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Allows combining multiple VRRP interfaces to maintain the same VRRP status within the group. " +
				"`group_authority` was previously called `group_master`, `group_master` is kept for compatibility with " +
				"scripts, but if both are set only `group_authority` will be taken into account.",
			// This resource behavior has not changed in ROS 7.14.2.
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				if old == new {
					return true
				}

				if new == "none" && old == "" {
					return true
				}

				return AlwaysPresentNotUserProvided(k, old, new, d)
			},
		},
		"group_master": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Allows combining multiple VRRP interfaces to maintain the same VRRP status within the group. " +
				"`group_authority` was previously called `group_master`, `group_master` is kept for compatibility with " +
				"scripts, but if both are set only `group_authority` will be taken into account.",
			// Maybe this is a bug, but for the 'none' value, the Mikrotik ROS 7.5 returns an empty string.
			// This resource behavior has not changed in ROS 7.14.2.
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				if old == new {
					return true
				}

				if new == "none" && old == "" {
					return true
				}

				return AlwaysPresentNotUserProvided(k, old, new, d)
			},
		},
		KeyInterface: PropInterfaceRw,
		"interval": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "VRRP update interval in seconds. Defines how often master sends advertisement packets.",
			ValidateFunc: validation.StringMatch(regexp.MustCompile(`^(\d+(ms|s|M)?)+$`),
				"expected hello interval 10ms..4m15s"),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
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
			Description: "Whether the master node always has the priority. When set to `no` the backup node will not " +
				"be elected to be a master until the current master fails, even if the backup node has higher priority " +
				"than the current master. This setting is ignored if the owner router becomes available",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"priority": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Priority of VRRP node used in Master election algorithm. A higher number means higher " +
				"priority. `255` is reserved for the router that owns VR IP and `0` is reserved for the Master router " +
				"to indicate that it is releasing responsibility.",
			ValidateFunc:     validation.IntBetween(1, 254),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"remote_address": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Specifies the remote address of the other VRRP router for syncing connection tracking. " +
				"If not set, the system autodetects the remote address via VRRP. The remote address is used only if " +
				"`sync_connection_tracking = true`.Sync connection tracking uses UDP port 8275.",
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
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "A protocol that will be used by VRRPv3. Valid only if the version is 3.",
			ValidateFunc:     validation.StringInSlice([]string{"ipv4", "ipv6"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"version": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Which VRRP version to use.",
			ValidateFunc:     validation.IntBetween(2, 3),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"vrid": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Virtual Router identifier. Each Virtual router must have a unique id number.",
			ValidateFunc:     validation.IntBetween(1, 255),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
	}

	return &schema.Resource{
		CreateContext: DefaultValidateCreate(resSchema, func(d *schema.ResourceData) diag.Diagnostics {
			if d.Get("remote_address").(string) != "" && !d.Get("sync_connection_tracking").(bool) {
				return diag.Diagnostics{
					{
						Severity: diag.Warning,
						Summary:  "sync_connection_tracking not enabled",
						Detail: "The remote address is used only if sync_connection_tracking=true. " +
							"The field will be omitted in the returned response.",
					},
				}
			}
			return nil
		}),
		ReadContext: DefaultRead(resSchema),
		UpdateContext: DefaultValidateUpdate(resSchema, func(d *schema.ResourceData) diag.Diagnostics {
			if d.Get("remote_address").(string) != "" && !d.Get("sync_connection_tracking").(bool) {
				return diag.Diagnostics{
					{
						Severity: diag.Warning,
						Summary:  "sync_connection_tracking not enabled",
						Detail: "The remote address is used only if sync_connection_tracking=true. " +
							"The field will be omitted in the returned response.",
					},
				}
			}
			return nil
		}),
		DeleteContext: DefaultDelete(resSchema),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    ResourceInterfaceVrrpV0().CoreConfigSchema().ImpliedType(),
				Upgrade: stateMigrationNameToId(resSchema[MetaResourcePath].Default.(string)),
				Version: 0,
			},
		},

		Schema: resSchema,
	}
}
