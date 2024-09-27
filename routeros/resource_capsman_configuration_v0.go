package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceCapsManConfigurationV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			MetaResourcePath: PropResourcePath("/caps-man/configuration"),
			MetaId:           PropId(Name),
			MetaTransformSet: PropTransformSet("channel.config: channel", "datapath.config: datapath",
				"rates.config: rates", "security.config: security"),

			"channel": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Channel inline settings.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				ValidateDiagFunc: ValidationMapKeyNames,
			},
			KeyComment: PropCommentRw,
			"country": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "Limits available bands, frequencies and maximum transmit power for each frequency. Also " +
					"specifies default value of scan-list. Value no_country_set is an FCC compliant set of channels.",
			},
			"datapath": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Datapath inline settings.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				ValidateDiagFunc: ValidationMapKeyNames,
			},
			"disconnect_timeout": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "This interval is measured from third sending failure on the lowest data rate. At this point " +
					"3 * (hw-retries + 1) frame transmits on the lowest data rate had failed. During disconnect-timeout packet " +
					"transmission will be retried with on-fail-retry-time interval. If no frame can be transmitted successfully " +
					`during disconnect-timeout, the connection is closed, and this event is logged as "extensive data loss". ` +
					"Successful frame transmission resets this timer.",
				DiffSuppressFunc: TimeEquall,
			},
			"distance": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "How long to wait for confirmation of unicast frames (ACKs) before considering transmission " +
					"unsuccessful, or in short ACK-Timeout.",
			},
			"frame_lifetime": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "Discard frames that have been queued for sending longer than frame-lifetime. By default, when " +
					"value of this property is 0, frames are discarded only after connection is closed (format: 0.00 sec).",
				DiffSuppressFunc: TimeEquall,
			},
			"guard_interval": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "Whether to allow use of short guard interval (refer to 802.11n MCS specification to see how " +
					`this may affect throughput). "any" will use either short or long, depending on data rate, "long" will ` +
					"use long.",
				ValidateFunc: validation.StringInSlice([]string{"any ", "long"}, false),
			},
			"hide_ssid": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				Description: "This property has effect only in AP mode. Setting it to yes can remove this network from " +
					"the list of wireless networks that are shown by some client software. Changing this setting does not " +
					"improve the security of the wireless network, because SSID is included in other frames sent by the AP.",
			},
			"hw_protection_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "Frame protection support property. " +
					"[See docs](https://wiki.mikrotik.com/wiki/Manual:Interface/Wireless#Frame_protection_support_(RTS/CTS)).",
				ValidateFunc: validation.StringInSlice([]string{"cts-to-self", "none", "rts-cts"}, false),
			},
			"hw_retries": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: "Number of times sending frame is retried without considering it a transmission failure. " +
					"[See docs](https://wiki.mikrotik.com/wiki/Manual:Interface/Wireless)",
				ValidateFunc: validation.IntBetween(0, 15),
			},
			"installation": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "Adjusts scan-list to use indoor, outdoor or all frequencies for the country that is set.",
				ValidateFunc: validation.StringInSlice([]string{"any", "indoor", "outdoor"}, false),
			},
			"keepalive_frames": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  `If a client has not communicated for around 20 seconds, AP sends a "keepalive-frame".`,
				ValidateFunc: validation.StringInSlice([]string{"enabled", "disabled"}, false),
			},
			"load_balancing_group": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "Tags the interface to the load balancing group. For a client to connect to interface in this " +
					"group, the interface should have the same number of already connected clients as all other interfaces " +
					"in the group or smaller. Useful in setups where ranges of CAPs mostly overlap.",
			},
			"max_sta_count": {
				Type:         schema.TypeInt,
				Optional:     true,
				Description:  "Maximum number of associated clients.",
				ValidateFunc: validation.IntBetween(1, 2007),
			},
			"mode": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "Set operational mode. Only **ap** currently supported.",
				ValidateFunc: validation.StringInSlice([]string{"ap"}, false),
			},
			KeyName: PropNameForceNewRw,
			"multicast_helper": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "When set to full multicast packets will be sent with unicast destination MAC address, " +
					"resolving multicast problem on a wireless link. This option should be enabled only on the access " +
					"point, clients should be configured in station-bridge mode.",
				ValidateFunc: validation.StringInSlice([]string{"default", "disabled", "full"}, false),
			},
			"rates": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Rates inline settings.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				ValidateDiagFunc: ValidationMapKeyNames,
			},
			"rx_chains": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Which antennas to use for receive.",
				Elem: &schema.Schema{
					Type:         schema.TypeInt,
					ValidateFunc: validation.IntBetween(0, 3),
				},
			},
			"security": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Security inline settings.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				ValidateDiagFunc: ValidationMapKeyNames,
			},
			"ssid": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "SSID (service set identifier) is a name broadcast in the beacons that identifies " +
					"wireless network.",
				ValidateFunc: validation.StringLenBetween(0, 32),
			},
			"tx_chains": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Which antennas to use for transmit.",
				Elem: &schema.Schema{
					Type:         schema.TypeInt,
					ValidateFunc: validation.IntBetween(0, 3),
				},
			},
		},
	}
}
