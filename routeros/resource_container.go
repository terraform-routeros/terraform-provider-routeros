package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceContainer https://help.mikrotik.com/docs/display/ROS/Container
func ResourceContainer() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/dns/static"),
		MetaId:           PropId(Id),

		"id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The ID of the container",
		},
		"arch": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The CPU architecture the container is designed to run on.",
		},
		"cmd": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The command to run in the container",
		},
		KeyComment: PropCommentRw,
		"dns": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The DNS server the container should use",
		},
		"domain_name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The domain name of the container",
		},
		"entrypoint": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The entrypoint to use in the container",
		},
		"envlist": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The list of environment variables to use",
		},
		"file": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The local .tar.gz file if the container is imported from a file",
		},
		"hostname": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The container hostname",
		},
		"interface": {
			Type:        schema.TypeString,
			Description: "The name of the veth interface the container should use",
		},
		"logging": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "If set to yes, all container-generated output will be shown in the RouterOS log",
		},
		"mounts": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Mounts from /container/mounts to be used with this container",
		},
		"name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The name of the container",
		},
		"os": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The OS of the container",
		},
		"remote-image": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The container image to be installed if an external registry is used",
		},
		"root_dir": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Used to save the container store outside main memory",
		},
		"status": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The current state the container is in (stopped/running)",
		},
		"stop-signal": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The stop signal to give the container when the stop action is triggered",
		},
		"tag": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The container tag",
		},
		"workdir": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The working directory for cmd entrypoint",
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
