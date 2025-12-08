package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*1",
    "disabled": "false",
    "dst-address": "0.0.0.0",
    "port": "2055",
    "src-address": "0.0.0.0",
    "v9-template-refresh": "20",
    "v9-template-timeout": "30m",
    "version": "9"
  }
*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/21102653/Traffic+Flow#TrafficFlow-Targets
func ResourceIpTrafficFlowTarget() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/traffic-flow/target"),
		MetaId:           PropId(Id),

		KeyDisabled: PropDisabledRw,
		"dst_address": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "IP address of the host which receives Traffic-Flow statistic packets from the router.",
			ValidateFunc: validation.IsIPAddress,
		},
		"port": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Port (UDP) of the host which receives Traffic-Flow statistic packets from the router.",
			ValidateFunc:     Validation64k,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"src_address": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "IP address used as source when sending Traffic-Flow statistics.",
			ValidateFunc:     validation.IsIPAddress,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"v9_template_refresh": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Number of packets after which the template is sent to the receiving host (only for NetFlow " +
				"version 9 and IPFIX).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"v9_template_timeout": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "After how long to send the template, if it has not been sent. (only for NetFlow version 9 " +
				"and IPFIX).",
			DiffSuppressFunc: TimeEqual,
		},
		"version": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Which version format of NetFlow to use.",
			ValidateFunc: validation.StringInSlice([]string{"1", "5", "9", "IPFIX"}, false),
		},
	}

	return &schema.Resource{
		CreateContext: DefaultCreate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultUpdate(resSchema),
		DeleteContext: DefaultDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: ImportStateCustomContext(resSchema),
		},

		Schema: resSchema,
	}
}
