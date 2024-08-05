package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {									  {								  {
    ".id": "*3",					    ".id": "*5",				    ".id": "*4",
    "address": "127.0.0.1",			    "address": "ff00::1",		    "cname": "ipv4",
    "address-list": "subdomain",	    "disabled": "false",		    "disabled": "false",
    "disabled": "false",			    "dynamic": "false",			    "dynamic": "false",
    "dynamic": "false",				    "name": "ipv6",				    "name": "cname",
    "match-subdomain": "true",		    "ttl": "1d",				    "ttl": "1d",
    "name": "ipv4",					    "type": "AAAA"				    "type": "CNAME"
    "ttl": "1d"						  },							  },
  },
  {									  {								  {
    ".id": "*6",					    ".id": "*7",				    ".id": "*8",
    "disabled": "false",			    "disabled": "true",			    "disabled": "false",
    "dynamic": "false",				    "dynamic": "false",			    "dynamic": "false",
    "forward-to": "127.0.0.1",		    "mx-exchange": "127.0.0.1",	    "name": "ns",
    "name": "fwd",					    "mx-preference": "10",		    "ns": "127.0.0.1",
    "ttl": "1d",					    "name": "mx",				    "ttl": "1d",
    "type": "FWD"					    "ttl": "1d",				    "type": "NS"
  },								    "type": "MX"				  },
                                      },
  {									  {								  {
    ".id": "*9",					    ".id": "*A",				    ".id": "*B",
    "disabled": "false",			    "disabled": "false",		    "disabled": "false",
    "dynamic": "false",				    "dynamic": "false",			    "dynamic": "false",
    "name": "nxdomain",				    "name": "srv",				    "name": "txt",
    "ttl": "1d",					    "srv-port": "8080",			    "text": "DNSSEC",
    "type": "NXDOMAIN"				    "srv-priority": "10",		    "ttl": "1d",
  },								    "srv-target": "127.0.0.1",	    "type": "TXT"
                                        "srv-weight": "100",		  }
                                        "ttl": "1d",
                                        "type": "SRV"
                                      },
*/

// ResourceDnsRecord https://wiki.mikrotik.com/wiki/Manual:IP/DNS https://help.mikrotik.com/docs/display/ROS/DNS
func ResourceDnsRecord() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/dns/static"),
		MetaId:           PropId(Id),

		"address": {
			Type:          schema.TypeString,
			Optional:      true,
			Description:   "The A record to be returend from the DNS hostname.",
			ConflictsWith: []string{"cname", "forward_to", "mx_exchange", "ns", "srv_target", "text"},
			ValidateFunc:  validation.IsIPAddress,
		},
		"address_list": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Name of the Firewall address list to which address must be dynamically added when some " +
				"request matches the entry.",
		},
		KeyComment: PropCommentRw,
		"cname": {
			Type:          schema.TypeString,
			Optional:      true,
			Description:   "Alias name for a domain name.",
			ConflictsWith: []string{"address", "forward_to", "mx_exchange", "ns", "srv_target", "text"},
		},
		KeyDisabled: PropDisabledRw,
		KeyDynamic:  PropDynamicRo,
		"forward_to": {
			Type:          schema.TypeString,
			Optional:      true,
			Description:   "The IP address of a domain name server to which a particular DNS request must be forwarded.",
			ConflictsWith: []string{"address", "cname", "mx_exchange", "ns", "srv_target", "text"},
		},
		"match_subdomain": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Whether the record will match requests for subdomains.",
		},
		"mx_exchange": {
			Type:          schema.TypeString,
			Optional:      true,
			Description:   "The domain name of the MX server.",
			ConflictsWith: []string{"address", "cname", "forward_to", "ns", "srv_target", "text"},
		},
		"mx_preference": {
			Type:         schema.TypeInt,
			Optional:     true,
			Computed:     true,
			Description:  "Preference of the particular MX record.",
			RequiredWith: []string{"mx_exchange"},
		},
		"name": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "The name of the DNS hostname to be created.",
			ExactlyOneOf: []string{"name", "regexp"},
		},
		"ns": {
			Type:          schema.TypeString,
			Optional:      true,
			Description:   "Name of the authoritative domain name server for the particular record.",
			ConflictsWith: []string{"address", "cname", "forward_to", "mx_exchange", "srv_target", "text"},
		},
		"regexp": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "DNS regexp. Regexp entries are case sensitive, but since DNS requests are not case sensitive, " +
				"RouterOS converts DNS names to lowercase, you should write regex only with lowercase letters.",
			ExactlyOneOf: []string{"name", "regexp"},
		},
		"srv_port": {
			Type:         schema.TypeInt,
			Optional:     true,
			Computed:     true,
			Description:  "The TCP or UDP port on which the service is to be found.",
			ValidateFunc: Validation64k,
			RequiredWith: []string{"srv_target"},
		},
		"srv_priority": {
			Type:         schema.TypeInt,
			Optional:     true,
			Computed:     true,
			Description:  "Priority of the particular SRV record.",
			RequiredWith: []string{"srv_target"},
		},
		"srv_target": {
			Type:          schema.TypeString,
			Optional:      true,
			Description:   "The canonical hostname of the machine providing the service ends in a dot.",
			ConflictsWith: []string{"address", "cname", "forward_to", "mx_exchange", "ns", "text"},
		},
		"srv_weight": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			Description:  "Weight of the particular SRC record.",
			RequiredWith: []string{"srv_target"},
		},
		"text": {
			Type:          schema.TypeString,
			Optional:      true,
			Description:   "Textual information about the domain name.",
			ConflictsWith: []string{"address", "cname", "forward_to", "mx_exchange", "ns", "srv_target"},
		},
		"ttl": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The ttl of the DNS record.",
			DiffSuppressFunc: TimeEquall,
		},
		"type": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "Type of the DNS record. Available values are: A, AAAA, CNAME, FWD, MX, NS, NXDOMAIN, SRV, TXT",
			ValidateFunc: validation.StringInSlice([]string{"A", "AAAA", "CNAME", "FWD", "MX", "NS", "NXDOMAIN",
				"SRV", "TXT"}, false),
		},
	}
	return &schema.Resource{
		Description: "Creates a DNS record on the MikroTik device.",

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
