package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*1",
    "add-mac-cookie": "true",
    "address-list": "aaa",
    "advertise": "true",
    "advertise-interval": "30m,10m",
    "advertise-timeout": "immediately",
    "advertise-url": "https://www.mikrotik.com/",
    "idle-timeout": "none",
    "incoming-filter": "bbb",
    "incoming-packet-mark": "ddd",
    "insert-queue-before": "first",
    "keepalive-timeout": "2m",
    "mac-cookie-timeout": "3d",
    "name": "uprof1",
    "on-login": "s1",
    "on-logout": "s2",
    "open-status-page": "always",
    "outgoing-filter": "ccc",
    "outgoing-packet-mark": "eee",
    "parent-queue": "none",
    "queue-type": "default-small",
    "rate-limit": "1",
    "session-timeout": "1s",
    "shared-users": "1",
    "status-autorefresh": "1m",
    "transparent-proxy": "true"
  }
*/

// https://wiki.mikrotik.com/wiki/Manual:IP/Hotspot/User#User_Profile
func ResourceIpHotspotUserProfile() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath:   PropResourcePath("/ip/hotspot/user/profile"),
		MetaId:             PropId(Id),
		MetaSetUnsetFields: PropSetUnsetFields("insert_queue_before", "parent_queue", "queue_type"),

		"add_mac_cookie": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Allows to add mac cookie for users.",
		},
		"address_list": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Name of the address list in which users IP address will be added. Useful to mark traffic per " +
				"user groups for queue tree configurations.",
		},
		"address_pool": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "IP pool name from which the user will get IP. When user has improper network settings configuration " +
				"on the computer, HotSpot server makes translation and assigns correct IP address from the pool instead " +
				"of incorrect one.",
		},
		"advertise": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Enable forced advertisement popups. After certain interval specific web-page is being displayed " +
				"for HotSpot users. Advertisement page might be blocked by browsers popup blockers.",
		},
		"advertise_interval": {
			Type:     schema.TypeSet,
			Optional: true,
			Description: "Set of intervals between advertisement popups. After the list is done, the last value is used " +
				"for all further advertisements, 10 minutes.",
			Elem: &schema.Schema{
				Type:             schema.TypeString,
				DiffSuppressFunc: TimeEquall,
			},
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"advertise_timeout": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "How long advertisement is shown, before blocking network access for HotSpot client. Connection " +
				"to Internet is not allowed, when advertisement is not shown.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"advertise_url": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "List of URLs that is show for advertisement popups. After the last URL is used, list starts " +
				"from the begining.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyDefault: PropDefaultRo,
		"idle_timeout": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Maximal period of inactivity for authorized HotSpot clients. Timer is counting, when there " +
				"is no traffic coming from that client and going through the router, for example computer is switched " +
				"off. User is logged out, dropped of the host list, the address used by the user is freed, when timeout " +
				"is reached.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"incoming_filter": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Name of the firewall chain applied to incoming packets from the users of this profile, jump " +
				"rule is required from built-in chain (input, forward, output) to chain=hotspot.",
		},
		"incoming_packet_mark": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Packet mark put on incoming packets from every user of this profile.",
		},
		"insert_queue_before": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "",
			ValidateFunc: validation.StringInSlice([]string{"first", "bottom"}, false),
		},
		"keepalive_timeout": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Keepalive timeout for authorized HotSpot clients. Used to detect, that the computer of the " +
				"client is alive and reachable. User is logged out, when timeout value is reached.",
			DiffSuppressFunc: TimeEquall,
		},
		"mac_cookie_timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Selects mac-cookie timeout from last login or logout. Read more>>.",
			DiffSuppressFunc: TimeEquall,
		},
		KeyName: PropName("Descriptive name of the profile."),
		"on_login": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Script name to be executed, when user logs in to the HotSpot from the particular profile. " +
				"It is possible to get username from internal user and interface variable. For example, :log info ``User " +
				"$user logged in!`` . If hotspot is set on bridge interface, then interface variable will show bridge " +
				"as actual interface unless use-ip-firewall' is set in bridge settings. List of available variables: " +
				"$user $username (alternative var name for $user) $address $``mac-address`` $interface.",
		},
		"on_logout": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Script name to be executed, when user logs out from the HotSpot.It is possible to get username " +
				"from internal user and interface variable. For example, :log info ``User $user logged in!`` . If hotspot " +
				"is set on bridge interface, then interface variable will show bridge as actual interface unless use-ip-firewall " +
				"is set in bridge settings. List of available variables: $user $username (alternative var name for $user) " +
				"$address $``mac-address`` $interface $cause Starting with v6.34rc11 some additional variables are available: " +
				"$uptime-secs - final session time in seconds $bytes-in - bytes uploaded $bytes-out - bytes downloaded " +
				"$bytes-total - bytes up + bytes down $packets-in - packets uploaded $packets-out - packets downloaded " +
				"$packets-total - packets up + packets down.",
		},
		"open_status_page": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Option to show status page for user authenticated with mac login method. For example to show " +
				"advertisement on status page (alogin.html) http-login - open status page only for HTTP login (includes " +
				"cookie and HTTPS) always - open HTTP status page in case of mac login as well.",
			ValidateFunc:     validation.StringInSlice([]string{"always", "http-login"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"outgoing_filter": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Name of the firewall chain applied to outgoing packets from the users of this profile, jump " +
				"rule is required from built-in chain (input, forward, output) to chain=hotspot.",
		},
		"outgoing_packet_mark": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Packet mark put on outgoing packets from every user of this profile.",
		},
		"parent_queue": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"queue_type": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "",
			ValidateFunc: validation.StringInSlice([]string{"default", "default-small", "ethernet-default",
				"hotspot-default", "multi-queue-ethernet-default", "only-hardware-queue", "pcq-download-default",
				"pcq-upload-default", "synchronous-default", "wireless-default"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"rate_limit": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Simple dynamic queue is created for user, once it logs in to the HotSpot. Rate-limitation " +
				"is configured in the following form [rx-rate[/tx-rate] [rx-burst-rate[/tx-burst-rate] [rx-burst-threshold[/tx-burst-threshold] " +
				"[rx-burst-time[/tx-burst-time] [priority] [rx-rate-min[/tx-rate-min]]]]. For example, to set 1M download, " +
				"512k upload for the client, rate-limit=512k/1M.",
		},
		"session_timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Allowed session time for client. After this time, the user is logged out unconditionally.",
			DiffSuppressFunc: TimeEquall,
		},
		"shared_users": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Allowed number of simultaneously logged in users with the same HotSpot username.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"status_autorefresh": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "HotSpot status page autorefresh interval.",
			DiffSuppressFunc: TimeEquall,
		},
		"transparent_proxy": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Use transparent HTTP proxy for the authorized users of this profile.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
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
