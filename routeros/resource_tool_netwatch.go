package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*1",
    "disabled": "true",
    "down-script": "",
    "host": "192.168.180.1",
    "http-codes": "",
    "name": "111",
    "status": "unknown",
    "test-script": "",
    "type": "simple",
    "up-script": ""
  }
*/

// https://help.mikrotik.com/docs/display/ROS/Netwatch
func ResourceToolNetwatch() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/tool/netwatch"),
		MetaId:           PropId(Id),
		MetaSkipFields: PropSkipFields(
			// Generic
			"since", "status", "done_tests", "failed_tests",
			// ICMP
			"sent_count", "response_count", "loss_count", "loss_percent", "rtt_avg", "rtt_min", "rtt_max", "rtt_jitter",
			"rtt_stdev",
			// TCP
			"tcp_connect_time",
			// HTTP, HTTPS
			"http_status_code", "http_codes",
			// DNS
			"ip", "ip6", "mail_servers", "name_servers",
		),

		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"down_script": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Script to execute on the event of probe state change `OK` --> `fail`.",
		},
		"host": {
			Type:     schema.TypeString,
			Required: true,
			Description: "The IP address of the server to be probed. Formats:" +
				"\n  * ipv4" +
				"\n  * ipv4@vrf" +
				"\n  * ipv6 " +
				"\n  * ipv6@vrf" +
				"\n  * ipv6-linklocal%interface",
		},
		"interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The time interval between probe tests.",
			DiffSuppressFunc: TimeEquall,
		},
		KeyName: PropName("Task name."),
		"src_address": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Source IP address which the Netwatch will try to use in order to reach the host. If address " +
				"is not present, then the host will be considered as `down`.",
		},
		"start_delay": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Time to wait before starting probe (on add, enable, or system start).",
			DiffSuppressFunc: TimeEquall,
		},
		"startup_delay": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Time to wait until starting Netwatch probe after system startup.",
			DiffSuppressFunc: TimeEquall,
		},
		"test_script": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Script to execute at the end of every probe test.",
		},
		"timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Max time limit to wait for a response.",
			DiffSuppressFunc: TimeEquall,
		},
		"type": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Type of the probe:" +
				"\n  *  icmp - (ping-style) series of ICMP request-response with statistics" +
				"\n  *  tcp-conn - test TCP connection (3-way handshake) to a server specified by IP and port" +
				"\n  *  http-get - do an HTTP Get request and test for a range of correct replies" +
				"\n  *  simple - simplified ICMP probe, with fewer options than **ICMP** type, used for backward " +
				"compatibility with the older Netwatch version",
			ValidateFunc:     validation.StringInSlice([]string{"icmp", "tcp-conn", "http-get", "simple"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"up_script": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Script to execute on the event of probe state change `fail` --> `OK`.",
		},

		// ICMP probe options
		"accept_icmp_time_exceeded": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "If the ICMP `time exceeded` message should be considered a valid response.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"packet_count": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Total count of ICMP packets to send out within a single test.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"packet_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The time between ICMP-request packet send.",
			DiffSuppressFunc: TimeEquall,
		},
		"packet_size": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "The total size of the IP ICMP packet.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"thr_loss_count": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Fail threshold for loss-count.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"thr_loss_percent": {
			Type:             schema.TypeFloat,
			Optional:         true,
			Description:      "Fail threshold for loss-percent.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"thr_avg": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Fail threshold for rtt-avg.",
			DiffSuppressFunc: TimeEquall,
		},
		"thr_jitter": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Fail threshold for rtt-jitter.",
			DiffSuppressFunc: TimeEquall,
		},
		"thr_max": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Fail threshold for rtt-max (a value above thr-max is a probe fail).",
			DiffSuppressFunc: TimeEquall,
		},
		"thr_stdev": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Fail threshold for rtt-stdev.",
			DiffSuppressFunc: TimeEquall,
		},
		"ttl": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Manually set time to live value for ICMP packet.",
		},

		// TCP-CONNECT/HTTP-GET probe options
		"port": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "TCP port (for both tcp-conn and http-get probes)",
			ValidateFunc:     Validation64k,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},

		// TCP-CONNECT pass-fail criteria
		"thr_tcp_conn_time": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Fail threshold for tcp-connect-time, the configuration uses microseconds, if the time " +
				"unit is not specified (s/m/h), log and status pages display the same value in milliseconds.",
			DiffSuppressFunc: TimeEquall,
		},

		// HTTP-GET probe pass/fail criteria
		"thr_http_time": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Fail threshold for http-resp-time.",
			DiffSuppressFunc: TimeEquall,
		},
		"http_code_min": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "OK/fail criteria for HTTP response code.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"http_code_max": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Response in the range [http-code-min , http-code-max] is a probe pass/OK; outside - a " +
				"probe fail. See [mozilla-http-status](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status) or " +
				"[rfc7231](https://datatracker.ietf.org/doc/html/rfc7231#section-6).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},

		// DNS probe options
		"record_type": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Record type that will be used for DNS probe.",
			ValidateFunc:     validation.StringInSlice([]string{"A", "AAAA", "MX", "NS"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"dns_server": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The DNS server that the probe should send its requests to, if not specified it will use the value from `/ip dns`.",
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
