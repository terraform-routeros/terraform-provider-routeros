package routeros

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  ".id": "*7",
  "adaptive-noise-immunity": "none",
  "allow-sharedkey": "false",
  "ampdu-priorities": "0",
  "amsdu-limit": "8192",
  "amsdu-threshold": "8192",
  "antenna-gain": "3",
  "area": "",
  "arp": "enabled",
  "arp-timeout": "auto",
  "band": "2ghz-b/g/n",
  "basic-rates-a/g": "6Mbps",
  "basic-rates-b": "1Mbps",
  "bridge-mode": "enabled",
  "channel-width": "20/40mhz-XX",
  "compression": "false",
  "country": "etsi",
  "default-ap-tx-limit": "0",
  "default-authentication": "true",
  "default-client-tx-limit": "0",
  "default-forwarding": "true",
  "default-name": "wlan1",
  "disable-running-check": "false",
  "disabled": "true",
  "disconnect-timeout": "3s",
  "distance": "indoors",
  "frame-lifetime": "0",
  "frequency": "auto",
  "frequency-mode": "regulatory-domain",
  "frequency-offset": "0",
  "guard-interval": "any",
  "hide-ssid": "false",
  "ht-basic-mcs": "mcs-0,mcs-1,mcs-2,mcs-3,mcs-4,mcs-5,mcs-6,mcs-7",
  "ht-supported-mcs": "mcs-0,mcs-1,mcs-2,mcs-3,mcs-4,mcs-5,mcs-6,mcs-7,mcs-8,mcs-9,mcs-10,mcs-11,mcs-12,mcs-13,mcs-14,mcs-15,mcs-16,mcs-17,mcs-18,mcs-19,mcs-20,mcs-21,mcs-22,mcs-23",
  "hw-fragmentation-threshold": "disabled",
  "hw-protection-mode": "none",
  "hw-protection-threshold": "0",
  "hw-retries": "7",
  "installation": "indoor",
  "interface-type": "Atheros AR9300",
  "interworking-profile": "disabled",
  "keepalive-frames": "enabled",
  "l2mtu": "1600",
  "mac-address": "D4:01:C3:30:E2:87",
  "max-station-count": "2007",
  "mode": "ap-bridge",
  "mtu": "1500",
  "multicast-buffering": "enabled",
  "multicast-helper": "default",
  "name": "wlan1",
  "noise-floor-threshold": "default",
  "nv2-cell-radius": "30",
  "nv2-downlink-ratio": "50",
  "nv2-mode": "dynamic-downlink",
  "nv2-noise-floor-offset": "default",
  "nv2-preshared-key": "",
  "nv2-qos": "default",
  "nv2-queue-count": "2",
  "nv2-security": "disabled",
  "nv2-sync-secret": "",
  "on-fail-retry-time": "100ms",
  "preamble-mode": "both",
  "radio-name": "D401C330E287",
  "rate-selection": "advanced",
  "rate-set": "default",
  "running": "false",
  "rx-chains": "0,1,2",
  "scan-list": "default",
  "secondary-frequency": "",
  "security-profile": "default",
  "skip-dfs-channels": "disabled",
  "ssid": "MikroTik-30E287",
  "station-bridge-clone-mac": "00:00:00:00:00:00",
  "station-roaming": "disabled",
  "supported-rates-a/g": "6Mbps,9Mbps,12Mbps,18Mbps,24Mbps,36Mbps,48Mbps,54Mbps",
  "supported-rates-b": "1Mbps,2Mbps,5.5Mbps,11Mbps",
  "tdma-period-size": "2",
  "tx-chains": "0,1,2",
  "tx-power-mode": "default",
  "update-stats-interval": "disabled",
  "vlan-id": "1",
  "vlan-mode": "no-tag",
  "wds-cost-range": "50-150",
  "wds-default-bridge": "none",
  "wds-default-cost": "100",
  "wds-ignore-ssid": "false",
  "wds-mode": "disabled",
  "wireless-protocol": "802.11",
  "wmm-support": "disabled",
  "wps-mode": "push-button-virtual-only"
}
*/

// https://help.mikrotik.com/docs/display/ROS/Wireless+Interface#WirelessInterface-Overview
func ResourceInterfaceWireless() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath:   PropResourcePath("/interface/wireless"),
		MetaId:             PropId(Id),
		MetaSkipFields:     PropSkipFields(".about", "pci_info"),
		MetaTransformSet:   PropTransformSet("basic_rates_ag: basic-rates-a/g", "supported_rates_ag: supported-rates-a/g"),
		MetaSetUnsetFields: PropSetUnsetFields("secondary_frequency"),

		"adaptive_noise_immunity": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "This property is only effective for cards based on Atheros chipset.",
			ValidateFunc:     validation.StringInSlice([]string{"ap-and-client-mode", "client-mode", "none"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"allow_sharedkey": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Allow WEP Shared Key clients to connect. Note that no authentication is done for these clients " +
				"(WEP Shared keys are not compared to anything) - they are just accepted at once (if access list allows " +
				"that).",
		},
		"amsdu_limit": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Max AMSDU that device is allowed to prepare when negotiated. AMSDU aggregation may significantly " +
				"increase throughput especially for small frames, but may increase latency in case of packet loss due " +
				"to retransmission of aggregated frame. Sending and receiving AMSDUs will also increase CPU usage.",
			ValidateFunc:     validation.IntBetween(0, 8192),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ampdu_priorities": {
			Type:     schema.TypeSet,
			Optional: true,
			Description: "Frame priorities for which AMPDU sending (aggregating frames and sending using block acknowledgment) " +
				"should get negotiated and used. Using AMPDUs will increase throughput, but may increase latency, therefore, " +
				"may not be desirable for real-time traffic (voice, video). Due to this, by default AMPDUs are enabled " +
				"only for best-effort traffic.",
			Elem: &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 7),
			},
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"amsdu_threshold": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Max frame size to allow including in AMSDU.",
			ValidateFunc:     validation.IntBetween(0, 8192),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"antenna_gain": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Antenna gain in dBi, used to calculate maximum transmit power according to country regulations.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"antenna_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Select antenna to use for transmitting and for receiving: `ant-a` - use only 'a'; antenna `ant-b` " +
				"- use only 'b'; antenna `txa-rxb` - use antenna 'a' for transmitting, antenna 'b' for receiving; `rxa-txb` - " +
				"use antenna 'b' for transmitting, antenna 'a' for receiving.",
			ValidateFunc: validation.StringInSlice([]string{"ant-a", "ant-b", "rxa-txb", "txa-rxb"}, false),
		},
		"area": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Identifies group of wireless networks. This value is announced by AP, and can be matched in " +
				" connect-list  by area-prefix. This is a proprietary extension.",
		},
		"arp": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "ARP Mode.",
			ValidateFunc:     validation.StringInSlice([]string{"disabled", "enabled", "proxy-arp", "reply-only"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"arp_timeout": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "ARP timeout is time how long ARP record is kept in ARP table after no packets are received " +
				"from IP. Value auto equals to the value of arp-timeout in `/ip settings`, default is 30s.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"band": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Defines set of used data rates, channel frequencies and widths.",
			ValidateFunc: validation.StringInSlice([]string{"2ghz-b", "2ghz-b/g", "2ghz-b/g/n", "2ghz-onlyg",
				"2ghz-onlyn", "5ghz-a", "5ghz-a/n", "5ghz-onlyn", "5ghz-a/n/ac", "5ghz-onlyac", "5ghz-n/ac"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"basic_rates_ag": {
			Type:     schema.TypeSet,
			Optional: true,
			Description: "Similar to the basic-rates-b property, but used for 5ghz, 5ghz-10mhz, 5ghz-5mhz, 5ghz-turbo, " +
				"2.4ghz-b/g, 2.4ghz-onlyg, 2ghz-10mhz, 2ghz-5mhz and 2.4ghz-g-turbo bands.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
				ValidateDiagFunc: ValidationMultiValInSlice([]string{"12Mbps", "18Mbps", "24Mbps", "36Mbps", "48Mbps",
					"54Mbps", "6Mbps", "9Mbps"}, false, false),
			},
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"basic_rates_b": {
			Type:     schema.TypeSet,
			Optional: true,
			Description: "List of basic rates, used for `2.4ghz-b`, `2.4ghz-b/g` and `2.4ghz-onlyg` bands.Client will connect " +
				"to AP only if it supports all basic rates announced by the AP. AP will establish WDS link only if it " +
				"supports all basic rates of the other AP.This property has effect only in AP modes, and when value of " +
				"rate-set is configured.",
			Elem: &schema.Schema{
				Type:             schema.TypeString,
				ValidateDiagFunc: ValidationMultiValInSlice([]string{"1Mbps", "2Mbps", "5Mbps", "11Mbps"}, false, false),
			},
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"bridge_mode": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Allows to use station-bridge mode.",
			ValidateFunc:     validation.StringInSlice([]string{"disabled", "enabled"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"burst_time": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Time in microseconds which will be used to send data without stopping. Note that no other " +
				"wireless cards in that network will be able to transmit data during burst-time microseconds. This setting " +
				"is available only for AR5000, AR5001X, and AR5001X+ chipset based cards.",
		},
		"channel_width": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Use of extension channels (e.g. `C`e, `eC` etc) allows additional 20MHz extension channels and " +
				"if it should be located below or above the control (main) channel. Extension channel allows 802.11n " +
				"devices to use up to 40MHz (802.11ac up to 160MHz) of spectrum in total thus increasing max throughput. " +
				"Channel widths with `XX` and `XXXX` extensions automatically scan for a less crowded control channel frequency " +
				"based on the number of concurrent devices running in every frequency and chooses the `C` - Control channel " +
				"frequency automatically.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment: PropCommentRw,
		"compression": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Setting this property to yes will allow the use of the hardware compression. Wireless interface " +
				"must have support for hardware compression. Connections with devices that do not use compression will " +
				"still work.",
		},
		"country": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Limits available bands, frequencies and maximum transmit power for each frequency. Also specifies " +
				"default value of scan-list. Value no_country_set is an FCC compliant set of channels.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"default_ap_tx_limit": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "This is the value of ap-tx-limit for clients that do not match any entry in the  access-list. " +
				"0 means no limit.",
		},
		"default_authentication": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "For AP mode, this is the value of authentication for clients that do not match any entry in " +
				"the  access-list. For station mode, this is the value of connect for APs that do not match any entry " +
				"in the  connect-list.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"default_client_tx_limit": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "This is the value of `client-tx-limit` for clients that do not match any entry in the access-list. " +
				"0 means no limit.",
		},
		"default_forwarding": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "This is the value of forwarding for clients that do not match any entry in the access-list.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"default_name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"disable_running_check": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "When set to yes interface will always have running flag. If value is set to no', the router " +
				"determines whether the card is up and running - for AP one or more clients have to be registered to " +
				"it, for station, it should be connected to an AP.",
		},
		KeyDisabled: PropDisabledRw,
		"disconnect_timeout": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "This interval is measured from third sending failure on the lowest data rate. At this point " +
				"`3 * (hw-retries + 1)` frame transmits on the lowest data rate had failed. During disconnect-timeout packet " +
				"transmission will be retried with on-fail-retry-time interval. If no frame can be transmitted successfully " +
				"during disconnect-timeout, the connection is closed, and this event is logged as `extensive data loss`. " +
				"Successful frame transmission resets this timer.",
			DiffSuppressFunc: TimeEquall,
		},
		"distance": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "How long to wait for confirmation of unicast frames (ACKs) before considering transmission " +
				"unsuccessful, or in short ACK-Timeout. Distance value has these behaviors:\n  * Dynamic - causes AP to detect " +
				"and use the smallest timeout that works with all connected clients.\n  * Indoor - uses the default ACK timeout " +
				"value that the hardware chip manufacturer has set.\n  * Number - uses the input value in formula: `ACK-timeout " +
				"= ((distance * 1000) + 299) / 300 us`\nAcknowledgments are not used in Nstreme/NV2 protocols.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"frame_lifetime": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Discard frames that have been queued for sending longer than frame-lifetime. By default, when " +
				"value of this property is 0, frames are discarded only after connection is closed.",
		},
		"frequency": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Channel frequency value in MHz on which AP will operate. Allowed values depend on the selected " +
				"band, and are restricted by country setting and wireless card capabilities. This setting has no effect " +
				"if interface is in any of station modes, or in wds-slave mode, or if DFS is active.Note: If using mode " +
				"`superchannel`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"frequency_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Three frequency modes are available:\n  * regulatory-domain - Limit available channels and maximum " +
				"transmit power for each channel according to the value of country\n  * manual-txpower - Same as above, but " +
				"do not limit maximum transmit power\n  *`superchannel` - Conformance Testing Mode. Allow all channels supported " +
				"by the card.\nList of available channels for each band can be seen in `/interface wireless` info allowed-channels. " +
				"This mode allows you to test wireless channels outside the default scan-list and/or regulatory domain. " +
				"This mode should only be used in controlled environments, or if you have special permission to use it " +
				"in your region. Before v4.3 this was called Custom Frequency Upgrade, or Superchannel. Since RouterOS " +
				"v4.3 this mode is available without special key upgrades to all installations.",
			ValidateFunc:     validation.StringInSlice([]string{"manual-txpower", "regulatory-domain", "superchannel"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"frequency_offset": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Allows to specify offset if the used wireless card operates at a different frequency than " +
				"is shown in RouterOS, in case a frequency converter is used in the card. So if your card works at 4000MHz " +
				"but RouterOS shows 5000MHz, set offset to 1000MHz and it will be displayed correctly. The value is in " +
				"MHz and can be positive or negative.",
		},
		"guard_interval": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Whether to allow use of short guard interval (refer to 802.11n MCS specification to see how " +
				"this may affect throughput). `any` will use either short or long, depending on data rate, `long` will " +
				"use long.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"hide_ssid": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "`true` - AP does not include SSID in the beacon frames, and does not reply to probe requests " +
				"that have broadcast SSID. `false` - AP includes SSID in the beacon frames, and replies to probe requests that " +
				"have broadcast SSID.This property has an effect only in AP mode. Setting it to yes can remove this network " +
				"from the list of wireless networks that are shown by some client software. Changing this setting does " +
				"not improve the security of the wireless network, because SSID is included in other frames sent by the " +
				"AP.",
		},
		"ht_basic_mcs": {
			Type:     schema.TypeSet,
			Optional: true,
			Description: "Modulation and Coding Schemes that every connecting client must support. Refer to 802.11n " +
				"for MCS specification.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`mcs-\d+`),
					`ht_basic_mcs format is "mcs-[0..23]": mcs-"12"`),
			},
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ht_supported_mcs": {
			Type:     schema.TypeSet,
			Optional: true,
			Description: "Modulation and Coding Schemes that this device advertises as supported. Refer to 802.11n for " +
				"MCS specification.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`mcs-\d+`),
					`ht_basic_mcs format is "mcs-[0..23]": mcs-"12"`),
			},
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"hw_fragmentation_threshold": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Specifies maximum fragment size in bytes when transmitted over the wireless medium. 802.11 " +
				"standard packet (MSDU in 802.11 terminologies) fragmentation allows packets to be fragmented before " +
				"transmitting over a wireless medium to increase the probability of successful transmission (only fragments " +
				"that did not transmit correctly are retransmitted). Note that transmission of a fragmented packet is " +
				"less efficient than transmitting unfragmented packet because of protocol overhead and increased resource " +
				"usage at both - transmitting and receiving party.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"hw_protection_mode": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Frame protection support property.",
			ValidateFunc:     validation.StringInSlice([]string{"cts-to-self", "none", "rts-cts"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"hw_protection_threshold": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Frame protection support property read more >>.",
			ValidateFunc:     Validation64k,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"hw_retries": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Number of times sending frame is retried without considering it a transmission failure. Data-rate " +
				"is decreased upon failure and the frame is sent again. Three sequential failures on the lowest supported " +
				"rate suspend transmission to this destination for the duration of on-fail-retry-time. After that, the " +
				"frame is sent again. The frame is being retransmitted until transmission success, or until the client " +
				"is disconnected after disconnect-timeout. The frame can be discarded during this time if frame-lifetime " +
				"is exceeded.",
			ValidateFunc:     validation.IntBetween(0, 15),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"installation": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Adjusts scan-list to use indoor, outdoor or all frequencies for the country that is set.",
			ValidateFunc:     validation.StringInSlice([]string{"any", "indoor", "outdoor"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"interface_type": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"interworking_profile": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "",
			ValidateFunc:     validation.StringInSlice([]string{"enabled", "disabled"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"keepalive_frames": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Applies only if wireless interface is in `mode = ap-bridge`. If a client has not communicated " +
				"for around 20 seconds, AP sends a `keepalive-frame`. Note, disabling the feature can lead to `ghost` " +
				"clients in registration-table.",
			ValidateFunc:     validation.StringInSlice([]string{"enabled", "disabled"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyL2Mtu:      PropL2MtuRw,
		KeyMacAddress: PropMacAddressRw("MAC address.", false),
		"master_interface": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Name of wireless interface that has virtual-ap capability. Virtual AP interface will only " +
				"work if master interface is in ap-bridge, bridge, station or wds-slave mode. This property is only for " +
				"virtual AP interfaces.",
		},
		"max_station_count": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Maximum number of associated clients. WDS links also count toward this limit.",
			ValidateFunc:     validation.IntBetween(1, 2007),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"mode": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Selection between different station and access point (AP) modes.\n  * Station modes: `station` - Basic " +
				"station mode. Find and connect to acceptable AP. `station-wds` - Same as station, but create WDS link with " +
				"AP, using proprietary extension. AP configuration has to allow WDS links with this device. Note that " +
				"this mode does not use entries in wds. `station-pseudobridge` - Same as station, but additionally perform " +
				"MAC address translation of all traffic. Allows interface to be bridged. `station-pseudobridge-clone` - " +
				"Same as station-pseudobridge, but use station-bridge-clone-mac address to connect to AP. `station-bridge` " +
				"- Provides support for transparent protocol-independent L2 bridging on the station device. RouterOS " +
				"AP accepts clients in station-bridge mode when enabled using bridge-mode parameter. In this mode, the " +
				"AP maintains a forwarding table with information on which MAC addresses are reachable over which station " +
				"device. Only works with RouterOS APs. With station-bridge mode, it is not possible to connect to CAPsMAN " +
				"controlled CAP.\n  * AP modes: `ap-bridge` - Basic access point mode. `bridge` - Same as ap-bridge, but limited " +
				"to one associated client. `wds-slave` - Same as ap-bridge, but scan for AP with the same ssid and establishes " +
				"WDS link. If this link is lost or cannot be established, then continue scanning. If dfs-mode is radar-detect, " +
				"then APs with enabled hide-ssid will not be found during scanning.\n  * Special modes: `alignment-only` - Put " +
				"the interface in a continuous transmit mode that is used for aiming the remote antenna. `nstreme-dual-slave` " +
				"- allow this interface to be used in nstreme-dual setup. MAC address translation in pseudobridge modes " +
				"works by inspecting packets and building a table of corresponding IP and MAC addresses. All packets " +
				"are sent to AP with the MAC address used by pseudobridge, and MAC addresses of received packets are " +
				"restored from the address translation table. There is a single entry in the address translation table " +
				"for all non-IP packets, hence more than one host in the bridged network cannot reliably use non-IP protocols. " +
				"Note: Currently IPv6 doesn't work over Pseudobridge.",
			ValidateFunc: validation.StringInSlice([]string{"station", "station-wds", "ap-bridge", "bridge",
				"alignment-only", "nstreme-dual-slave", "wds-slave", "station-pseudobridge", "station-pseudobridge-clone",
				"station-bridge"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyMtu: PropMtuRw(),
		"multicast_buffering": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "For a client that has power saving, buffer multicast packets until next beacon time. A client " +
				"should wake up to receive a beacon, by receiving beacon it sees that there are multicast packets pending, " +
				"and it should wait for multicast packets to be sent.",
			ValidateFunc:     validation.StringInSlice([]string{"disabled", "enabled"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"multicast_helper": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "When set to full, multicast packets will be sent with a unicast destination MAC address, resolving " +
				" multicast problem on the wireless link. This option should be enabled only on the access point, clients " +
				"should be configured in station-bridge mode. Available starting from v5.15.disabled - disables the helper " +
				"and sends multicast packets with multicast destination MAC addressesdhcp - dhcp packet mac addresses " +
				"are changed to unicast mac addresses prior to sending them outfull - all multicast packet mac address " +
				"are changed to unicast mac addresses prior to sending them outdefault - default choice that currently " +
				"is set to dhcp. Value can be changed in future releases.",
			ValidateFunc:     validation.StringInSlice([]string{"default", "disabled", "full"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyName: PropName("Name of the interface."),
		"noise_floor_threshold": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "For advanced use only, as it can badly affect the performance of the interface. It is possible " +
				"to manually set noise floor threshold value. By default, it is dynamically calculated. This property " +
				"also affects received signal strength. This property is only effective on non-AC chips.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"nv2_cell_radius": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Setting affects the size of contention time slot that AP allocates for clients to initiate " +
				"connection and also size of time slots used for estimating distance to client. When setting is too small, " +
				"clients that are farther away may have trouble connecting and/or disconnect with `ranging timeout` error. " +
				"Although during normal operation the effect of this setting should be negligible, in order to maintain " +
				"maximum performance, it is advised to not increase this setting if not necessary, so AP is not reserving " +
				"time that is actually never used, but instead allocates it for actual data transfer.on AP: distance " +
				"to farthest client in kmon station: no effect.",
			ValidateFunc:     validation.IntBetween(10, 200),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"nv2_downlink_ratio": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Specifies the Nv2 downlink ratio. Uplink ratio is automatically calculated from the " +
				"downlink-ratio value. When using dynamic-downlink mode the downlink-ratio is also used when link get " +
				"fully saturated. Minimum value is 20 and maximum 80.",
			ValidateFunc:     validation.IntBetween(20, 80),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"nv2_mode": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Specifies to use dynamic or fixed downlink/uplink ratio.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"nv2_noise_floor_offset": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"nv2_preshared_key": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "Specifies preshared key to be used.",
		},
		"nv2_qos": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Sets the packet priority mechanism, firstly data from high priority queue is sent, then lower " +
				"queue priority data until 0 queue priority is reached. When link is full with high priority queue data, " +
				"lower priority data is not sent. Use it very carefully, setting works on APframe-priority - manual setting " +
				"that can be tuned with Mangle rules.default - default setting where small packets receive priority for " +
				"best latency.",
			ValidateFunc:     validation.StringInSlice([]string{"default", "frame-priority"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"nv2_queue_count": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Specifies how many priority queues are used in Nv2 network.",
			ValidateFunc:     validation.IntBetween(2, 8),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"nv2_security": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Specifies Nv2 security mode.",
			ValidateFunc:     validation.StringInSlice([]string{"disabled", "enabled"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"nv2_sync_secret": {
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
			Description: "Specifies secret key for use in the Nv2 synchronization. Secret should match on Master " +
				"and Slave devices in order to establish the synced state.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"on_fail_retry_time": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "After third sending failure on the lowest data rate, wait for specified time interval before " +
				"retrying.",
			DiffSuppressFunc: TimeEquall,
		},
		"periodic_calibration": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Setting default enables periodic calibration if  info  default-periodic-calibration property " +
				"is enabled. Value of that property depends on the type of wireless card. This property is only effective " +
				"for cards based on Atheros chipset.",
			ValidateFunc: validation.StringInSlice([]string{"default", "disabled", "enabled"}, false),
		},
		"periodic_calibration_interval": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "This property is only effective for cards based on Atheros chipset.",
			ValidateFunc: validation.IntBetween(1, 10000),
		},
		"preamble_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Short preamble mode is an option of 802.11b standard that reduces per-frame overhead.On AP:\n  * long " +
				"- Do not use short preamble.\n  * short - Announce short preamble capability. Do not accept connections from " +
				"clients that do not have this capability.\n  * both - Announce short preamble capability.\nOn station:\n  *long - " +
				"do not use short preamble.\n  * short - do not connect to AP if it does not support short preamble.\n  * both - " +
				"Use short preamble if AP supports it.",
			ValidateFunc:     validation.StringInSlice([]string{"both", "long", "short"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"prism_cardtype": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Specify type of the installed Prism wireless card.",
			ValidateFunc: validation.StringInSlice([]string{"100mW", "200mW", "30mW"}, false),
		},
		"proprietary_extensions": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "RouterOS includes proprietary information in an information element of management frames. " +
				"This parameter controls how this information is included. `pre-2.9.25` - This is older method. It can interoperate " +
				"with newer versions of RouterOS. This method is incompatible with some clients, for example, Centrino " +
				"based ones. `post-2.9.25` - This uses standardized way of including vendor specific information, that is " +
				"compatible with newer wireless clients.",
			ValidateFunc:     validation.StringInSlice([]string{"post-2.9.25", "pre-2.9.25"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"radio_name": {
			Type:     schema.TypeString,
			Computed: true,
			Description: "Descriptive name of the device, that is shown in registration table entries on the remote " +
				"devices. This is a proprietary extension.",
		},
		"rate_selection": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Starting from v5.9 default value is advanced since legacy mode was inefficient.",
			ValidateFunc:     validation.StringInSlice([]string{"advanced", "legacy"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"rate_set": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Two options are available: `default` - default basic and supported rate sets are used. Values " +
				"from basic-rates and supported-rates parameters have no effect. `configured` - use values from basic-rates, " +
				"supported-rates, basic-mcs, mcs.",
			ValidateFunc:     validation.StringInSlice([]string{"configured", "default"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyRunning: PropRunningRo,
		"rx_chains": {
			Type:     schema.TypeSet,
			Optional: true,
			Description: "Which antennas to use for receive. In current MikroTik routers, both RX and TX chain must " +
				"be enabled, for the chain to be enabled.",
			Elem: &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 3),
			},
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"scan_list": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The default value is all channels from selected band that are supported by card and allowed " +
				"by the country and frequency-mode settings (this list can be seen in  info). For default scan list in " +
				"5ghz band channels are taken with 20MHz step, in 5ghz-turbo band - with 40MHz step, for all other bands " +
				"- with 5MHz step. If scan-list is specified manually, then all matching channels are taken. (Example: " +
				"scan-list=default,5200-5245,2412-2427 - This will use the default value of scan list for current band, " +
				"and add to it supported frequencies from 5200-5245 or 2412-2427 range.) Since RouterOS v6.0 with Winbox " +
				"or Webfig, for inputting of multiple frequencies, add each frequency or range of frequencies into separate " +
				"multiple scan-lists. Using a comma to separate frequencies is no longer supported in Winbox/Webfig since " +
				"v6.0.Since RouterOS v6.35 (wireless-rep) scan-list support step feature where it is possible to manually " +
				"specify the scan step. Example: scan-list=5500-5600:20 will generate such scan-list values 5500,5520,5540,5560,5580,5600.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"security_profile": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Name of profile from  security-profiles.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"secondary_frequency": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Specifies secondary channel, required to enable 80+80MHz transmission. To disable 80+80MHz " +
				"functionality, set secondary-frequency to `` or unset the value via CLI/GUI.",
			ValidateFunc: validation.StringInSlice([]string{"integer"}, false),
		},
		"ssid": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "SSID (service set identifier) is a name that identifies wireless network.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"skip_dfs_channels": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "These values are used to skip all DFS channels or specifically skip DFS CAC channels in range " +
				"5600-5650MHz which detection could go up to 10min.",
			ValidateFunc:     validation.StringInSlice([]string{"string", "10min-cac", "all", "disabled"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"station_bridge_clone_mac": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "This property has effect only in the station-pseudobridge-clone mode.Use this MAC address " +
				"when connection to AP. If this value is 00:00:00:00:00:00, station will initially use MAC address of " +
				"the wireless interface.As soon as packet with MAC address of another device needs to be transmitted, " +
				"station will reconnect to AP using that address.",
			ValidateFunc:     validation.IsMACAddress,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"station_roaming": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Station Roaming feature is available only for 802.11 wireless protocol and only for station " +
				"modes.",
			ValidateFunc:     validation.StringInSlice([]string{"disabled", "enabled"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"supported_rates_ag": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "List of supported rates, used for all bands except 2ghz-b.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
				ValidateDiagFunc: ValidationMultiValInSlice([]string{"12Mbps", "18Mbps", "24Mbps", "36Mbps", "48Mbps",
					"54Mbps", "6Mbps", "9Mbps"}, false, false),
			},
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"supported_rates_b": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "List of supported rates, used for `2ghz-b, `2ghz-b/g` and `2ghz-b/g/n` bands. Two devices will " +
				"communicate only using rates that are supported by both devices. This property has effect only when " +
				"value of rate-set is configured.",
			Elem: &schema.Schema{
				Type:             schema.TypeString,
				ValidateDiagFunc: ValidationMultiValInSlice([]string{"1Mbps", "2Mbps", "5Mbps", "11Mbps"}, false, false),
			},
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tdma_period_size": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Specifies TDMA period in milliseconds. It could help on the longer distance links, it could " +
				"slightly increase bandwidth, while latency is increased too.",
			ValidateFunc:     validation.IntBetween(1, 10),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tx_chains": {
			Type:     schema.TypeSet,
			Optional: true,
			Description: "Which antennas to use for transmitting. In current MikroTik routers, both RX and TX chain " +
				"must be enabled, for the chain to be enabled.",
			Elem: &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 3),
			},
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tx_power": {
			Type:         schema.TypeInt,
			Optional:     true,
			ValidateFunc: validation.IntBetween(-30, 40),
			Description:  "For 802.11ac wireless interface it's total power but for 802.11a/b/g/n it's power per chain.",
		},
		"tx_power_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Sets up tx-power mode for wireless card `default` - use values stored in the card `all-rates-fixed` " +
				"- use same transmit power for all data rates. Can damage the card if transmit power is set above rated " +
				"value of the card for used rate. `manual-table` - define transmit power for each rate separately. Can damage " +
				"the card if transmit power is set above rated value of the card for used rate. `card-rates` - use transmit " +
				"power calculated for each rate based on value of tx-power parameter. Legacy mode only compatible with " +
				"currently discontinued products.",
			ValidateFunc:     validation.StringInSlice([]string{"default", "card-rates", "all-rates-fixed", "manual-table"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"update_stats_interval": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "How often to request update of signals strength and ccq values from clients. Access to registration-table " +
				" also triggers update of these values.This is proprietary extension.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"vlan_id": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "VLAN ID to use if doing VLAN tagging.",
			ValidateFunc:     validation.IntBetween(0, 4094),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"vlan_mode": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "VLAN tagging mode specifies if traffic coming from client should get tagged (and untagged when going to client).",
			ValidateFunc:     validation.StringInSlice([]string{"default", "no-tag", "use-service-tag", "use-tag"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"vht_basic_mcs": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Modulation and Coding Schemes that every connecting client must support. Refer to 802.11ac " +
				"for MCS specification. You can set MCS interval for each of Spatial Stream\n  * none - will not use selected;" +
				"\n  * mcs0-7 - client must support MCS-0 to MCS-7;\n  * mcs0-8 - client must support MCS-0 to MCS-8;" +
				"\n  * mcs0-9 - client must support MCS-0 to MCS-9.",
			ValidateDiagFunc: ValidationMultiValInSlice([]string{"none", "mcs0-7", "`mcs0-8`", "mcs0-9"}, false, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"vht_supported_mcs": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Modulation and Coding Schemes that this device advertises as supported. Refer to 802.11ac " +
				"for MCS specification. You can set MCS interval for each of Spatial Stream\n  * none - will not use selected; " +
				"\n  * mcs0-7 - devices will advertise as supported MCS-0 to MCS-7;\n  * mcs0-8 - devices will advertise " +
				"as supported MCS-0 to MCS-8;\n  * mcs0-9 - devices will advertise as supported MCS-0 to MCS-9.",
			ValidateDiagFunc: ValidationMultiValInSlice([]string{"none", "mcs0-7", "`mcs0-8`", "mcs0-9"}, false, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"wds_cost_range": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Bridge port cost of WDS links are automatically adjusted, depending on measured link throughput. " +
				"Port cost is recalculated and adjusted every 5 seconds if it has changed by more than 10%, or if more " +
				"than 20 seconds have passed since the last adjustment.Setting this property to 0 disables automatic " +
				"cost adjustment.Automatic adjustment does not work for WDS links that are manually configured as a bridge " +
				"port.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"wds_default_bridge": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "When WDS link is established and status of the wds interface becomes running, it will be added " +
				"as a bridge port to the bridge interface specified by this property. When WDS link is lost, wds interface " +
				"is removed from the bridge. If wds interface is already included in a bridge setup when WDS link becomes " +
				"active, it will not be added to bridge specified by , and will (needs editing).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"wds_default_cost": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Initial bridge port cost of the WDS links.",
			ValidateFunc:     validation.IntBetween(1, 200000000),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"wds_ignore_ssid": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "By default, WDS link between two APs can be created only when they work on the same frequency " +
				"and have the same SSID value. If this property is set to yes, then SSID of the remote AP will not be " +
				"checked. This property has no effect on connections from clients in station-wds mode. It also does not " +
				"work if wds-mode is static-mesh or dynamic-mesh.",
		},
		"wds_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Controls how WDS links with other devices (APs and clients in station-wds mode) are established.\n  * disabled " +
				"does not allow WDS links.\n  * static only allows WDS links that are manually configured in WDS.\n  * dynamic also " +
				"allows WDS links with devices that are not configured in WDS, by creating required entries dynamically. " +
				"Such dynamic WDS entries are removed automatically after the connection with the other AP is lost.\n  * -mesh " +
				"modes use different (better) method for establishing link between AP, that is not compatible with APs " +
				"in non-mesh mode. This method avoids one-sided WDS links that are created only by one of the two APs. " +
				"Such links cannot pass any data.When AP or station is establishing WDS connection with another AP, it " +
				"uses connect-list to check whether this connection is allowed. If station in station-wds mode is establishing " +
				"connection with AP, AP uses access-list to check whether this connection is allowed.If mode is station-wds, " +
				"then this property has no effect.",
			ValidateFunc:     validation.StringInSlice([]string{"disabled", "dynamic", "dynamic-mesh", "static", "static-mesh"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"wireless_protocol": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Specifies protocol used on wireless interface;\n  * unspecified - protocol mode used on previous " +
				"RouterOS versions (v3.x, v4.x). Nstreme is enabled by old enable-nstreme setting, Nv2 configuration " +
				"is not possible.\n  * any : on AP - regular 802.11 Access Point or Nstreme Access Point; on station - selects " +
				"Access Point without specific sequence, it could be changed by connect-list rules.\n  * nstreme - enables " +
				"Nstreme protocol (the same as old enable-nstreme setting).\n  * nv2 - enables Nv2 protocol.\n  * nv2 nstreme : on " +
				"AP - uses first wireless-protocol setting, always Nv2; on station - searches for Nv2 Access Point, then " +
				"for Nstreme Access Point.\n  * nv2 nstreme 802.11 - on AP - uses first wireless-protocol setting, always Nv2; " +
				"on station - searches for Nv2 Access Point, then for Nstreme Access Point, then for regular 802.11 Access " +
				"Point.Warning! Nv2 doesn't have support for Virtual AP.",
			ValidateFunc: validation.StringInSlice([]string{"802.11", "any", "nstreme", "nv2", "nv2-nstreme",
				"nv2-nstreme-802.11", "unspecified"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		// -----------------------
		"wmm_support": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Specifies whether to enable  WMM.  Only applies to bands `B` and `G`. Other bands will have it " +
				"enabled regardless of this setting.",
			ValidateFunc:     validation.StringInSlice([]string{"disabled", "enabled", "required"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"wps_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "WPS Server allows to connect wireless clients that support WPS to AP protected with the " +
				"Pre-Shared Key without specifying that key in the clients configuration.",
			ValidateFunc:     validation.StringInSlice([]string{"disabled", "push-button", "push-button-virtual-only"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
	}

	return &schema.Resource{
		// Interaction with mixed types of elements within a single resource.
		// In this case, there are physical and virtual interfaces that need to be created and deleted in different ways.
		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

			metadata := GetMetadata(resSchema)
			filter := buildReadFilter(map[string]interface{}{"name": d.Get("name")})
			items, err := ReadItemsFiltered(filter, metadata.Path, m.(Client))
			if err != nil {
				return diag.FromErr(err)
			}

			var diags diag.Diagnostics
			if items == nil || len(*items) == 0 {
				// No interface with the specified name was found. Adding...
				diags = ResourceCreate(ctx, resSchema, d, m)
			} else {
				// An interface with the specified name is found. There are two options:
				//		it is physical and then we will update it with existing settings,
				//		or it is virtual and needs to be imported.
				iface := (*items)[0]
				if ifType, ok := iface["interface-type"]; ok {
					if strings.ToLower(ifType) != "virtual" {
						// It's a physical interface.
						d.SetId(iface.GetID(Id))
						diags = ResourceUpdate(ctx, resSchema, d, m)
					} else {
						diags = diag.Diagnostics{diag.Diagnostic{
							Severity: diag.Error,
							Summary:  fmt.Sprintf("A virtual interface named '%v' already exists", d.Get("name")),
						}}
					}
				} else {
					diags = diag.Diagnostics{diag.Diagnostic{
						Severity: diag.Error,
						Summary: fmt.Sprintf("The Mikrotik resource (%v print where name=%v) does not contain "+
							"'interface-type' attribute in the response",
							metadata.Path, d.Get("name")),
					}}
				}
			}

			if diags.HasError() {
				return diags
			}

			return ResourceRead(ctx, resSchema, d, m)
		},

		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultUpdate(resSchema),
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			if strings.ToLower(d.Get("interface_type").(string)) != "virtual" {
				// It's a physical interface.
				return SystemResourceDelete(ctx, resSchema, d, m)
			}

			return ResourceDelete(ctx, resSchema, d, m)
		},

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
