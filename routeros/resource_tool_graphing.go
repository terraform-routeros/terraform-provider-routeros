package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
[
  {
    ".id": "*0",
    "disabled": "false",
    "interface": "all",
    "allow-address": "0.0.0.0/0",
    "store-on-disk": "true"
  },
  {...}
*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/22773810/Graphing#Graphing-Interfacegraphing
func ResourceToolGraphingInterface() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/tool/graphing/interface"),
		MetaId:           PropId(Id),

		KeyDisabled:  PropDisabledRw,
		KeyInterface: PropInterfaceRw,
		"allow_address": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "IP address range from which is allowed to access graphing information.",
			ValidateFunc:     ValidationIpAddress,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"store_on_disk": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Defines whether to store collected information on system drive.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
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

/*
[
  {
    ".id": "*0",
    "disabled": "false",
    "simple-queue": "all",
    "allow-address": "0.0.0.0/0",
    "allow-target": "true",
    "store-on-disk": "true"
  },
  {...}
*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/22773810/Graphing#Graphing-Queuegraphing
func ResourceToolGraphingQueue() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/tool/graphing/queue"),
		MetaId:           PropId(Id),

		KeyDisabled: PropDisabledRw,
		"simple_queue": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Defines which queues will be monitored. all means that all queues on router will be monitored.",
		},
		"allow_address": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "IP address range from which is allowed to access graphing information.",
			ValidateFunc:     ValidationIpAddress,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"allow_target": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether to allow access to graphs from queue's target-address.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"store_on_disk": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Defines whether to store collected information on system drive.",
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

/*
[
  {
    ".id": "*0",
    "disabled": "false",
    "allow-address": "0.0.0.0/0",
    "store-on-disk": "true"
  },
  {...}
*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/22773810/Graphing#Graphing-Resourcegraphing
func ResourceToolGraphingResource() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/tool/graphing/resource"),
		MetaId:           PropId(Id),

		KeyDisabled: PropDisabledRw,
		"allow_address": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "IP address range from which is allowed to access graphing information.",
			ValidateFunc:     ValidationIpAddress,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"store_on_disk": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Defines whether to store collected information on system drive.",
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
