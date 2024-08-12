package routeros

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  "allow-remote-requests": "true",
  "cache-max-ttl": "1w",
  "cache-size": "2048",
  "cache-used": "99", RO
  "dynamic-servers": "", RO
  "max-concurrent-queries": "100",
  "max-concurrent-tcp-sessions": "20",
  "max-udp-packet-size": "4096",
  "query-server-timeout": "2s",
  "query-total-timeout": "10s",
  "servers": "192.168.1.1",
  "use-doh-server": "",
  "verify-doh-cert": "false"
}
*/

// ResourceDns https://wiki.mikrotik.com/wiki/Manual:IP/DNS
func ResourceDns() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/dns"),
		MetaId:           PropId(Name),

		"address_list_extra_time": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "",
			ValidateFunc: ValidationTime,
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				return AlwaysPresentNotUserProvided(k, old, new, d) || TimeEquall(k, old, new, d)
			},
		},
		"allow_remote_requests": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Specifies whether to allow network requests.",
		},
		"cache_max_ttl": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "Maximum time-to-live for cache records. In other words, cache records will expire " +
				"unconditionally after cache-max-ttl time. Shorter TTL received from DNS servers are respected. " +
				"*Default: 1w*",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEquall,
		},
		"cache_size": {
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Specifies the size of DNS cache in KiB (64..4294967295). *Default: 2048*",
		},
		"cache_used": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Shows the currently used cache size in KiB.",
		},
		"doh_max_concurrent_queries": {
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Specifies how many DoH concurrent queries are allowed.",
		},
		"doh_max_server_connections": {
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Specifies how many concurrent connections to the DoH server are allowed.",
		},
		"doh_timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Computed:         true,
			Description:      "Specifies how long to wait for query response from the DoH server.",
			DiffSuppressFunc: TimeEquall,
		},
		"dynamic_servers": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "List of dynamically added DNS server from different services, for example, DHCP.",
		},
		"max_concurrent_queries": {
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Specifies how much concurrent queries are allowed. *Default: 100*",
		},
		"max_concurrent_tcp_sessions": {
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Specifies how much concurrent TCP sessions are allowed. *Default: 20*",
		},
		"max_udp_packet_size": {
			Type:         schema.TypeInt,
			Optional:     true,
			Computed:     true,
			Description:  "Maximum size of allowed UDP packet. *Default: 4096*",
			ValidateFunc: validation.IntBetween(50, 65507),
		},
		"mdns_repeat_ifaces": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "An option to enable mDNS repeater on specified interfaces. This option is available in RouterOS starting from version 7.16.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"query_server_timeout": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "Specifies how long to wait for query response from one server. " +
				"Time can be specified in milliseconds. *Default: 2s*",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEquall,
		},
		"query_total_timeout": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "Specifies how long to wait for query response in total. Note that this setting must be " +
				"configured taking into account query_server_timeout and number of used DNS server. " +
				"Time can be specified in milliseconds. *Default: 10s*",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEquall,
		},
		"servers": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "List of DNS server IPv4/IPv6 addresses.",
		},
		KeyVrf: PropVrfRw,
		"use_doh_server": {
			Type:     schema.TypeString,
			Optional: true,
			Description: `DNS over HTTPS (DoH) server URL.
	> Mikrotik strongly suggest not use third-party download links for certificate fetching. 
	Use the Certificate Authority's own website.

	> RouterOS prioritize DoH over DNS server if both are configured on the device.`,
		},
		"verify_doh_cert": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "DoH certificate verification. [See docs](https://wiki.mikrotik.com/wiki/Manual:IP/DNS#DNS_over_HTTPS).",
		},
	}

	return &schema.Resource{
		Description: "A MikroTik router with DNS feature enabled can be set as a DNS server for any DNS-compliant client.",

		CreateContext: DefaultSystemCreate(resSchema),
		ReadContext:   DefaultSystemRead(resSchema),
		UpdateContext: DefaultSystemUpdate(resSchema),
		// This behavior when deleting a system resource is the exception rather than the rule.
		// With existing serialization logic, the best way to avoid undefined DNS service state
		// is to clear the main fields.
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			// Values in the Mikrotik notation!
			resetFileds := map[string]string{
				"allow-remote-requests": "no",
				"servers":               "",
				"use-doh-server":        "",
				"verify-doh-cert":       "no",
			}

			var resUrl string
			if m.(Client).GetTransport() == TransportREST {
				// https://router/rest/ip/dns/set
				resUrl = "/set"
			}

			// Used POST request!
			err := m.(Client).SendRequest(crudPost, &URL{Path: resSchema[MetaResourcePath].Default.(string) + resUrl},
				resetFileds, nil)
			if err != nil {
				return diag.FromErr(err)
			}

			d.SetId("")
			return nil
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type: ResourceDnsV0().CoreConfigSchema().ImpliedType(),
				Upgrade: stateMigrationScalarToList("servers"),
				Version: 0,
			},
		},
	}
}
