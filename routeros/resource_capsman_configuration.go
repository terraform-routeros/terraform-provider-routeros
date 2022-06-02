package routeros

import (
	"log"
	"strconv"

	roscl "github.com/gnewbury1/terraform-provider-routeros/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCapsManConfiguration() *schema.Resource {
	return &schema.Resource{
		Create: resourceCapsManConfigurationCreate,
		Read:   resourceCapsManConfigurationRead,
		Update: resourceCapsManConfigurationUpdate,
		Delete: resourceCapsManConfigurationDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
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
		},
	}
}

func resourceCapsManConfigurationCreate(d *schema.ResourceData, m interface{}) error {

	c := m.(*roscl.Client)
	configuration_obj := new(roscl.CapsManConfiguration)

	configuration_obj.Name = d.Get("name").(string)
	configuration_obj.Comment = d.Get("comment").(string)
	configuration_obj.Country = d.Get("country").(string)
	configuration_obj.DisconnectTimeout = d.Get("disconnect_timeout").(string)
	configuration_obj.Distance = d.Get("distance").(string)
	configuration_obj.FrameLifetime = d.Get("frame_lifetime").(string)
	configuration_obj.GuardInterval = d.Get("guard_interval").(string)
	configuration_obj.HideSsid = strconv.FormatBool(d.Get("hide_ssid").(bool))
	configuration_obj.HwProtectionMode = d.Get("hw_protection_mode").(string)
	configuration_obj.HwRetries = d.Get("hw_retries").(string)
	configuration_obj.Installation = d.Get("installation").(string)
	configuration_obj.KeepAliveFrames = d.Get("keepalive_frames").(string)
	configuration_obj.LoadBalancingGroup = d.Get("load_balancing_group").(string)
	configuration_obj.MaxStaCount = strconv.Itoa(d.Get("max_sta_count").(int))

	configuration_obj.Mode = d.Get("mode").(string)
	configuration_obj.MulticastHelper = d.Get("multicast_helper").(string)
	configuration_obj.Rates = d.Get("rates").(string)
	configuration_obj.RatesBasic = d.Get("rates_basic").(string)
	configuration_obj.RatesSupported = d.Get("rates_supported").(string)
	configuration_obj.RatesHtBasicMcs = d.Get("rates_ht_basic_mcs").(string)
	configuration_obj.RatesHtSupportedMcs = d.Get("rates_ht_supported_mcs").(string)
	configuration_obj.RatesVhtBasicMcs = d.Get("rates_vht_basic_mcs").(string)
	configuration_obj.RatesVhtSupportedMcs = d.Get("rates_vht_supported_mcs").(string)
	configuration_obj.RxChains = d.Get("rx_chains").(string)
	configuration_obj.Ssid = d.Get("ssid").(string)
	configuration_obj.TxChains = d.Get("tx_chains").(string)

	configuration_obj.DatapathBridge = d.Get("datapath_bridge").(string)
	configuration_obj.DatapathBridgeCost = d.Get("datapath_bridge_cost").(string)
	configuration_obj.DatapathBridgeHorizon = d.Get("datapath_bridge_horizon").(string)
	configuration_obj.DatapathInterfaceList = d.Get("datapath_interface_list").(string)
	configuration_obj.DatapathL2MTU = d.Get("datapath_l2mtu").(string)
	configuration_obj.DatapathMTU = d.Get("datapath_mtu").(string)
	configuration_obj.DatapathLocalForwarding = d.Get("datapath_local_forwarding").(string)
	configuration_obj.DatapathClientToClientForwarding = d.Get("datapath_client_to_client_forwarding").(string)
	configuration_obj.DatapathOpenFlowSwitch = d.Get("datapath_openflow_switch").(string)
	configuration_obj.DatapathVlanMode = d.Get("datapath_vlan_mode").(string)
	vlan_id, is_set := d.GetOk("datapath_vlan_id")
	if is_set {
		configuration_obj.DatapathVlanID = strconv.Itoa(vlan_id.(int))
	}

	configuration_obj.SecurityGroupEncryption = d.Get("security_group_encryption").(string)
	configuration_obj.SecurityAuthenticationTypes = d.Get("security_authentication_types").(string)
	configuration_obj.SecurityEapMethods = d.Get("security_eap_methods").(string)
	configuration_obj.SecurityEapRadiusAccounting = d.Get("security_eap_radius_accounting").(string)
	configuration_obj.SecurityEncryption = d.Get("security_encryption").(string)
	configuration_obj.SecurityPassphrase = d.Get("security_passphrase").(string)
	configuration_obj.SecurityGroupKeyUpdate = d.Get("security_group_key_update").(string)
	configuration_obj.SecurityTlsCertificate = d.Get("security_tls_certificate").(string)
	configuration_obj.SecurityTlsMode = d.Get("security_tls_mode").(string)

	channel_save_selected, is_set := d.GetOk("channel_save_selected")
	if is_set {
		configuration_obj.ChannelSaveSelected = strconv.FormatBool(channel_save_selected.(bool))
	}
	channel_skip_dfs_channels, is_set := d.GetOk("channel_skip_dfs_channels")
	if is_set {
		configuration_obj.ChannelSkipDfsChannels = strconv.FormatBool(channel_skip_dfs_channels.(bool))
	}
	configuration_obj.ChannelWidth = d.Get("channel_width").(string)
	configuration_obj.ChannelControlChannelWidth = d.Get("channel_control_channel_width").(string)
	configuration_obj.ChannelBand = d.Get("channel_band").(string)
	configuration_obj.ChannelExtensionChannel = d.Get("channel_extension_channel").(string)
	configuration_obj.ChannelReselectInterval = d.Get("channel_reselect_interval").(string)
	channel_frequency, is_set := d.GetOk("channel_frequency")
	if is_set {
		configuration_obj.ChannelFrequency = strconv.FormatBool(channel_frequency.(bool))
	}
	configuration_obj.ChannelSecondaryFrequency = d.Get("channel_secondary_frequency").(string)
	channel_tx_power, is_set := d.GetOk("channel_tx_power")
	if is_set {
		configuration_obj.ChannelTXPower = strconv.FormatBool(channel_tx_power.(bool))
	}

	res, err := c.CreateCapsManConfiguration(configuration_obj)
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PUT request to the API")
		log.Fatal(err.Error())
		return err
	}

	d.SetId(res.ID)
	return nil
}

func resourceCapsManConfigurationRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	configuration, err := c.ReadCapsManConfiguration(d.Id())

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a GET request to the API")
		log.Fatal(err.Error())
		return err
	}
	vlan_id, _ := strconv.Atoi(configuration.DatapathVlanID)

	frequency, _ := strconv.Atoi(configuration.ChannelFrequency)
	max_sta_count, _ := strconv.Atoi(configuration.MaxStaCount)
	save_selected, _ := strconv.ParseBool(configuration.ChannelSaveSelected)
	skip_dfs_channels, _ := strconv.ParseBool(configuration.ChannelSkipDfsChannels)
	tx_power, _ := strconv.Atoi(configuration.ChannelTXPower)
	hide_ssid, _ := strconv.ParseBool(configuration.HideSsid)

	d.SetId(configuration.ID)
	d.Set("name", configuration.Name)
	d.Set("comment", configuration.Comment)
	d.Set("country", configuration.Country)
	d.Set("disconnect_timeout", configuration.DisconnectTimeout)
	d.Set("distance", configuration.Distance)
	d.Set("frame_lifetime", configuration.FrameLifetime)
	d.Set("guard_interval", configuration.GuardInterval)
	d.Set("hide_ssid", hide_ssid)
	d.Set("hw_protection_mode", configuration.HwProtectionMode)
	d.Set("hw_retries", configuration.HwRetries)
	d.Set("installation", configuration.Installation)
	d.Set("keepalive_frames", configuration.KeepAliveFrames)
	d.Set("load_balancing_group", configuration.LoadBalancingGroup)
	d.Set("max_sta_count", max_sta_count)
	d.Set("mode", configuration.Mode)
	d.Set("multicast_helper", configuration.MulticastHelper)
	d.Set("rates", configuration.Rates)
	d.Set("rates_basic", configuration.RatesBasic)
	d.Set("rates_supported", configuration.RatesSupported)
	d.Set("rates_ht_basic_mcs", configuration.RatesHtBasicMcs)
	d.Set("rates_ht_supported_mcs", configuration.RatesHtSupportedMcs)
	d.Set("rates_vht_basic_mcs", configuration.RatesVhtBasicMcs)
	d.Set("rates_vht_supported_mcs", configuration.RatesVhtSupportedMcs)
	d.Set("rx_chains", configuration.RxChains)
	d.Set("ssid", configuration.Ssid)
	d.Set("tx_chains", configuration.TxChains)

	d.Set("datapath_bridge_cost", configuration.DatapathBridgeCost)
	d.Set("datapath_bridge", configuration.DatapathBridge)
	d.Set("datapath_bridge_horizon", configuration.DatapathBridgeHorizon)
	d.Set("datapath_interface_list", configuration.DatapathBridgeHorizon)
	d.Set("datapath_l2mtu", configuration.DatapathBridgeHorizon)
	d.Set("datapath_mtu", configuration.DatapathBridgeHorizon)
	d.Set("datapath_client_to_client_forwarding", configuration.DatapathClientToClientForwarding)
	d.Set("datapath_openflow_switch", configuration.DatapathOpenFlowSwitch)
	d.Set("datapath_vlan_mode", configuration.DatapathVlanMode)
	d.Set("datapath_vlan_id", vlan_id)

	d.Set("security_authentication_types", configuration.SecurityAuthenticationTypes)
	d.Set("security_eap_methods", configuration.SecurityEapMethods)
	d.Set("security_eap_radius_accounting", configuration.SecurityEapRadiusAccounting)
	d.Set("security_group_encryption", configuration.SecurityGroupEncryption)
	d.Set("security_passphrase", configuration.SecurityPassphrase)
	d.Set("security_encryption", configuration.SecurityEncryption)
	d.Set("security_group_key_update", configuration.SecurityGroupKeyUpdate)
	d.Set("security_tls_certificate", configuration.SecurityTlsCertificate)
	d.Set("security_tls_mode", configuration.SecurityTlsMode)

	d.Set("channel_width", configuration.ChannelWidth)
	d.Set("channel_save_selected", save_selected)
	d.Set("channel_skip_dfs_channels", skip_dfs_channels)
	d.Set("channel_reselect_interval", configuration.ChannelReselectInterval)
	d.Set("channel_band", configuration.ChannelBand)
	d.Set("channel_extension_channel", configuration.ChannelExtensionChannel)
	d.Set("channel_frequency", frequency)
	d.Set("channel_secondary_frequency", configuration.ChannelSecondaryFrequency)
	d.Set("channel_tx_power", tx_power)

	return nil

}

func resourceCapsManConfigurationUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	configuration_obj := new(roscl.CapsManConfiguration)

	configuration_obj.Name = d.Get("name").(string)
	configuration_obj.Comment = d.Get("comment").(string)
	configuration_obj.Country = d.Get("country").(string)
	configuration_obj.DisconnectTimeout = d.Get("disconnect_timeout").(string)
	configuration_obj.Distance = d.Get("distance").(string)
	configuration_obj.FrameLifetime = d.Get("frame_lifetime").(string)
	configuration_obj.GuardInterval = d.Get("guard_interval").(string)
	configuration_obj.HideSsid = strconv.FormatBool(d.Get("hide_ssid").(bool))
	configuration_obj.HwProtectionMode = d.Get("hw_protection_mode").(string)
	configuration_obj.HwRetries = d.Get("hw_retries").(string)
	configuration_obj.Installation = d.Get("installation").(string)
	configuration_obj.KeepAliveFrames = d.Get("keepalive_frames").(string)
	configuration_obj.LoadBalancingGroup = d.Get("load_balancing_group").(string)
	configuration_obj.MaxStaCount = strconv.Itoa(d.Get("max_sta_count").(int))

	configuration_obj.Mode = d.Get("mode").(string)
	configuration_obj.MulticastHelper = d.Get("multicast_helper").(string)
	configuration_obj.Rates = d.Get("rates").(string)
	configuration_obj.RatesBasic = d.Get("rates_basic").(string)
	configuration_obj.RatesSupported = d.Get("rates_supported").(string)
	configuration_obj.RatesHtBasicMcs = d.Get("rates_ht_basic_mcs").(string)
	configuration_obj.RatesHtSupportedMcs = d.Get("rates_ht_supported_mcs").(string)
	configuration_obj.RatesVhtBasicMcs = d.Get("rates_vht_basic_mcs").(string)
	configuration_obj.RatesVhtSupportedMcs = d.Get("rates_vht_supported_mcs").(string)
	configuration_obj.RxChains = d.Get("rx_chains").(string)
	configuration_obj.Ssid = d.Get("ssid").(string)
	configuration_obj.TxChains = d.Get("tx_chains").(string)

	configuration_obj.DatapathBridge = d.Get("datapath_bridge").(string)
	configuration_obj.DatapathBridgeCost = d.Get("datapath_bridge_cost").(string)
	configuration_obj.DatapathBridgeHorizon = d.Get("datapath_bridge_horizon").(string)
	configuration_obj.DatapathInterfaceList = d.Get("datapath_interface_list").(string)
	configuration_obj.DatapathL2MTU = d.Get("datapath_l2mtu").(string)
	configuration_obj.DatapathMTU = d.Get("datapath_mtu").(string)
	configuration_obj.DatapathLocalForwarding = d.Get("datapath_local_forwarding").(string)
	configuration_obj.DatapathClientToClientForwarding = d.Get("datapath_client_to_client_forwarding").(string)
	configuration_obj.DatapathOpenFlowSwitch = d.Get("datapath_openflow_switch").(string)
	configuration_obj.DatapathVlanMode = d.Get("datapath_vlan_mode").(string)
	vlan_id, is_set := d.GetOk("datapath_vlan_id")
	if is_set {
		configuration_obj.DatapathVlanID = strconv.Itoa(vlan_id.(int))
	}

	configuration_obj.SecurityGroupEncryption = d.Get("security_group_encryption").(string)
	configuration_obj.SecurityAuthenticationTypes = d.Get("security_authentication_types").(string)
	configuration_obj.SecurityEapMethods = d.Get("security_eap_methods").(string)
	configuration_obj.SecurityEapRadiusAccounting = d.Get("security_eap_radius_accounting").(string)
	configuration_obj.SecurityEncryption = d.Get("security_encryption").(string)
	configuration_obj.SecurityPassphrase = d.Get("security_passphrase").(string)
	configuration_obj.SecurityGroupKeyUpdate = d.Get("security_group_key_update").(string)
	configuration_obj.SecurityTlsCertificate = d.Get("security_tls_certificate").(string)
	configuration_obj.SecurityTlsMode = d.Get("security_tls_mode").(string)

	channel_save_selected, is_set := d.GetOk("channel_save_selected")
	if is_set {
		configuration_obj.ChannelSaveSelected = strconv.FormatBool(channel_save_selected.(bool))
	}
	channel_skip_dfs_channels, is_set := d.GetOk("channel_skip_dfs_channels")
	if is_set {
		configuration_obj.ChannelSkipDfsChannels = strconv.FormatBool(channel_skip_dfs_channels.(bool))
	}
	configuration_obj.ChannelWidth = d.Get("channel_width").(string)
	configuration_obj.ChannelControlChannelWidth = d.Get("channel_control_channel_width").(string)
	configuration_obj.ChannelBand = d.Get("channel_band").(string)
	configuration_obj.ChannelExtensionChannel = d.Get("channel_extension_channel").(string)
	configuration_obj.ChannelReselectInterval = d.Get("channel_reselect_interval").(string)
	channel_frequency, is_set := d.GetOk("channel_frequency")
	if is_set {
		configuration_obj.ChannelFrequency = strconv.Itoa(channel_frequency.(int))
	}
	configuration_obj.ChannelSecondaryFrequency = d.Get("channel_secondary_frequency").(string)
	channel_tx_power, is_set := d.GetOk("channel_tx_power")
	if is_set {
		configuration_obj.ChannelTXPower = strconv.Itoa(channel_tx_power.(int))
	}

	res, err := c.UpdateCapsManConfiguration(d.Id(), configuration_obj)

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PATCH request to the API")
		log.Fatal(err.Error())
		return err
	}

	d.SetId(res.ID)

	return nil
}

func resourceCapsManConfigurationDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	err := c.DeleteCapsManConfiguration(d.Id())
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a DELETE request to the API")
		log.Fatal(err.Error())
		return err
	}
	d.SetId("")
	return nil
}
