package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCapsManConfiguration() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/caps-man/configuration"),
		MetaId:           PropId(Name),

		KeyName:    PropNameRw,
		KeyComment: PropCommentRw,
		"channel": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"channel_save_selected": {
			Type:     schema.TypeBool,
			Optional: true,
			//Default:  false,
		},
		"channel_width": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"channel_control_channel_width": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"channel_band": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"channel_reselect_interval": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"channel_extension_channel": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"channel_frequency": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"channel_secondary_frequency": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"channel_tx_power": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"channel_skip_dfs_channels": {
			Type:     schema.TypeBool,
			Optional: true,
			//Default:  false,
		},
		"country": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"datapath_bridge": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"datapath_bridge_cost": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"datapath_bridge_horizon": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"datapath_interface_list": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"datapath_l2mtu": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"datapath_local_forwarding": {
			Type:     schema.TypeString,
			Optional: true,
			//Default:  "false",
		},
		"datapath_client_to_client_forwarding": {
			Type:     schema.TypeString,
			Optional: true,
			//Default:  "false",
		},
		"datapath_mtu": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"datapath_openflow_switch": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"datapath_vlan_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"datapath_vlan_id": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"disconnect_timeout": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"distance": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"frame_lifetime": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"guard_interval": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"hide_ssid": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"hw_protection_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"hw_retries": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"installation": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"keepalive_frames": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"load_balancing_group": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"max_sta_count": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  2007,
		},
		"mode": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"multicast_helper": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"rates": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"rates_basic": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"rates_supported": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"rates_ht_basic_mcs": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"rates_ht_supported_mcs": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"rates_vht_basic_mcs": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"rates_vht_supported_mcs": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"rx_chains": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "0",
		},
		"security": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"security_group_encryption": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"security_authentication_types": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"security_eap_methods": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"security_eap_radius_accounting": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"security_encryption": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"security_passphrase": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"security_group_key_update": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"security_tls_certificate": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"security_tls_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"ssid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tx_chains": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "0",
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
