package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*2",
    "dns-name": "",
    "hotspot-address": "192.168.11.1",
    "html-directory": "hotspot",
    "html-directory-override": "",
    "http-cookie-lifetime": "3d",
    "http-proxy": "0.0.0.0:0",
    "install-hotspot-queue": "false",
    "login-by": "cookie,http-chap,https",
    "name": "hsprof2",
    "smtp-server": "0.0.0.0",
    "split-user-domain": "false",
    "ssl-certificate": "tls",
    "use-radius": "false"
  }
*/

// https://wiki.mikrotik.com/wiki/Manual:IP/Hotspot/Profile
func ResourceIpHotspotProfile() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/hotspot/profile"),
		MetaId:           PropId(Id),

		"dns_name": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "DNS name of the HotSpot server (it appears as the location of the login page). This name will " +
				"automatically be added as a static DNS entry in the DNS cache. Name can affect if Hotspot is automatically " +
				"detected by client device. For example, iOS devices may not detect Hotspot that has a name which includes " +
				"`.local`.",
		},
		"hotspot_address": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "IP address of HotSpot service.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"html_directory": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Directory name in which HotSpot HTML pages are stored (by default hotspot directory). It is " +
				"possible to specify different directory with modified HTML pages. To change HotSpot login page, connect " +
				"to the router with FTP and download hotspot directory contents. v6.31 and older software builds: For " +
				"devices where `flash` directory is present, hotspot html directory must be stored there and path must " +
				"be typed in as follows: `/(hotspot_dir)`. This must be done in this order as hotspot sees `flash` " +
				"directory as root location. v6.32 and newer software builds: full path must be typed in html-directory " +
				"field, including `/flash/(hotspot_dir)`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"html_directory_override": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Alternative path for hotspot html files. It should be used only if customized hotspot html " +
				"files are stored on external storage(attached usb, hdd, etc). If configured then hotspot will switch " +
				"to this html path as soon at it becomes available and switch back to html-directory path if override " +
				"path becomes non-available for some reason.",
		},
		"http_cookie_lifetime": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "HTTP cookie validity time, the option is related to cookie HotSpot login method.",
			DiffSuppressFunc: TimeEquall,
		},
		"http_proxy": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Address and port of the proxy server for HotSpot service, when default value is used all request " +
				"are resolved by the local `/ip proxy`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"https_redirect": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether to redirect unauthenticated user to hotspot login page, if he is visiting a https:// " +
				"url. Since certificate domain name will mismatch, often this leads to errors, so you can set this parameter " +
				"to `no` and all https requests will simply be rejected and user will have to visit a http page.",
		},
		"login_by": {
			Type:     schema.TypeSet,
			Optional: true,
			Description: "Used HotSpot authentication method\n " +
				"* mac-cookie - enables login by mac cookie method.\n " +
				"* cookie - may only be used with other HTTP authentication method. HTTP cookie is generated, when user authenticates " +
				"in HotSpot for the first time. User is not asked for the login/password and authenticated automatically, " +
				"until cookie-lifetime is active.\n " +
				"* http-chap - login/password is required for the user to authenticate in HotSpot. CHAP " +
				"challenge-response method with MD5 hashing algorithm is used for protecting passwords. \n " +
				"* http-pap - login/password is required for user to authenticate in HotSpot. Username and password are " +
				"sent over network in plain text.\n " +
				"* https - login/password is required for user to authenticate in HotSpot. Client login/password " +
				"exchange between client and server is encrypted with SSL tunnel.\n " +
				"* mac - client is authenticated without asking login form. Client MAC-address is added to `/ip hotspot " +
				"user` database, client is authenticated as soon as connected to the HotSpot\n " +
				"* trial - client is allowed to use internet without HotSpot login for the specified amount of time.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"cookie", "http-chap", "http-pap", "https", "mac",
					"trial", "mac-cookie"}, false),
			},
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"mac_auth_mode": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Allows to control User-Name and User-Password RADIUS attributes when using MAC authentication.",
			ValidateFunc:     validation.StringInSlice([]string{"mac-as-username", "mac-as-username-and-password"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"mac_auth_password": {
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
			Description: "Used together with MAC authentication, field used to specify password for the users to be " +
				"authenticated by their MAC addresses. The following option is required, when specific RADIUS server " +
				"rejects authentication for the clients with blank password.",
		},
		KeyName: PropName("Descriptive name of the profile."),
		"nas_port_type": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "`NAS-Port-Type` value to be sent to RADIUS server, `NAS-Port-Type` values are described in the " +
				"RADIUS RFC 2865. This optional value attribute indicates the type of the physical port of the HotSpot " +
				"server.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"radius_accounting": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Send RADIUS server accounting information for each user, when yes is used.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"radius_default_domain": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Default domain to use for RADIUS requests. Allows to use separate RADIUS server per `/ip hotspot " +
				"profile`. If used, same domain name should be specified under `/radius domain` value.",
		},
		"radius_interim_update": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "How often to send accounting updates . When received is set, interim-time is used from RADIUS " +
				"server. 0s is the same as received.",
			DiffSuppressFunc: TimeEquall,
		},
		"radius_location_name": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "`RADIUS-Location-Id` to be sent to RADIUS server. Used to identify location of the HotSpot server " +
				"during the communication with RADIUS server. Value is optional and used together with RADIUS server.",
		},
		"radius_mac_format": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Controls how the MAC address of the client is encoded in the `User-Name` and `User-Password` " +
				"attributes when using MAC authentication.",
			ValidateFunc: validation.StringInSlice([]string{"XX XX XX XX XX XX", "XX:XX:XX:XX:XX:XX",
				"XXXXXX-XXXXXX", "XXXXXXXXXXXX", "XX-XX-XX-XX-XX-XX", "XXXX:XXXX:XXXX", "XXXXXX:XXXXXX"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"rate_limit": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Rate limitation in form of rx-rate[/tx-rate] [rx-burst-rate[/tx-burst-rate] [rx-burst-threshold[/tx-burst-threshold] " +
				"[rx-burst-time[/tx-burst-time]]]] [priority] [rx-rate-min[/tx-rate-min]] from the point of view of the " +
				"router (so `rx` is client upload, and `tx` is client download). All rates should be numbers with " +
				"optional 'k' (1,000s) or 'M' (1,000,000s). If tx-rate is not specified, rx-rate is as tx-rate too. Same " +
				"goes for tx-burst-rate and tx-burst-threshold and tx-burst-time. If both rx-burst-threshold and tx-burst-threshold " +
				"are not specified (but burst-rate is specified), rx-rate and tx-rate is used as burst thresholds. If " +
				"both rx-burst-time and tx-burst-time are not specified, 1s is used as default. rx-rate-min and tx-rate " +
				"min are the values of limit-at properties.",
		},
		"smtp_server": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "SMTP server address to be used to redirect HotSpot users SMTP requests.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"split_user_domain": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Split username from domain name when the username is given in `user@domain` or in `domain\\user` " +
				"format from RADIUS server.",
		},
		"ssl_certificate": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Name of the SSL certificate on the router to to use only for HTTPS authentication.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"trial_uptime_limit": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Used only with trial authentication method. Time value specifies, how long trial user " +
				"identified by MAC address can use access to public networks without HotSpot authentication.",
			DiffSuppressFunc: TimeEquall,
		},
		"trial_uptime_reset": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Used only with trial authentication method.",
			DiffSuppressFunc: TimeEquall,
		},
		"trial_user_profile": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Specifies hotspot user profile for trial users.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"use_radius": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Use RADIUS to authenticate HotSpot users.",
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
