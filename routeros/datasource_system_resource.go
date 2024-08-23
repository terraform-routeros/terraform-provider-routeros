package routeros

// Script generated from sampled device MikroTik 7.10 (stable) on CHR QEMU-x86_64

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceSystemResource() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/resource"),
		MetaId:           PropId(Id),
		MetaSkipFields: PropSkipFields(
			"cpu_frequency", "cpu_load", "free_hdd_space", "free_memory",
			"uptime", "write_sect_since_reboot", "write_sect_total", "bad_blocks",
		),
		"architecture_name": { // Sample = architecture-name: "x86_64"
			Type:     schema.TypeString,
			Computed: true,
		},
		"board_name": { // Sample = board-name: "CHR"
			Type:     schema.TypeString,
			Computed: true,
		},
		"build_time": { // Sample = build-time: "Jun/15/2023 05:17:29"
			Type:     schema.TypeString,
			Computed: true,
		},
		"cpu": { // Sample = cpu: "QEMU"
			Type:     schema.TypeString,
			Computed: true,
		},
		"cpu_count": { // Sample = cpu-count: "4"
			Type:     schema.TypeInt,
			Computed: true,
		},
		"factory_software": { // Sample = factory-software: "7.1"
			Type:     schema.TypeString,
			Computed: true,
		},
		"platform": { // Sample = platform: "MikroTik"
			Type:     schema.TypeString,
			Computed: true,
		},
		"total_hdd_space": { // Sample = total-hdd-space: "93564928"
			Type:     schema.TypeInt,
			Computed: true,
		},
		"total_memory": { // Sample = total-memory: "469762048"
			Type:     schema.TypeInt,
			Computed: true,
		},
		"version": { // Sample = version: "7.10 (stable)"
			Type:     schema.TypeString,
			Computed: true,
		},
	}

	return &schema.Resource{
		ReadContext: DefaultSystemDatasourceRead(resSchema),
		Schema:      resSchema,
	}
}
