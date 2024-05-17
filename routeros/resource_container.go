package routeros

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// https://help.mikrotik.com/docs/display/ROS/Container#Container-Properties
func ResourceContainer() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/container"),
		MetaId:           PropId(Id),

		"arch": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The architecture of the container image",
		},
		"cmd": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The main purpose of a CMD is to provide defaults for an executing container. These defaults can include an executable, or they can omit the executable, in which case you must specify an ENTRYPOINT instruction as well.",
		},
		KeyComment: PropCommentRw,
		"dns": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Set custom DNS servers",
		},
		"domain_name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Container NIS domain name",
		},
		"entrypoint": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "An ENTRYPOINT allows to specify executable to run when starting container. Example: /bin/sh",
		},
		"envlist": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "list of environmental variables (configured under /container envs ) to be used with container",
		},
		"file": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "container *tar.gz tarball if the container is imported from a file",
			ExactlyOneOf: []string{"file", "remote_image"},
		},
		"hostname": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Container host name",
		},
		"interface": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "veth interface to be used with the container",
		},
		"logging": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "if set to yes, all container-generated output will be shown in the RouterOS log",
		},
		"mounts": {
			Type:        schema.TypeSet,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "Mounts from /container/mounts/ sub-menu to be used with this container",
		},
		"name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Assign a name to the container",
		},
		"os": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The OS of the container image",
		},
		"remote_image": {
			Type:         schema.TypeString,
			Optional:     true,
			ForceNew:     true,
			Description:  "The container image name to be installed if an external registry is used (configured under /container/config set registry-url=...)",
			ExactlyOneOf: []string{"file", "remote_image"},
		},
		"root_dir": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Used to save container store outside main memory",
		},
		"start_on_boot": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Start the container on boot",
		},
		"status": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The status of the container",
		},
		"stop_signal": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Signal to stop the container.",
		},
		"tag": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The tag of the container image",
		},
		"user": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Sets the username used",
		},
		"workdir": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The working directory for cmd entrypoint",
		},
	}

	resCreate := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		// Run DefaultCreate.
		diags := ResourceCreate(ctx, resSchema, d, m)
		if diags.HasError() {
			return diags
		}

		startContainer(ctx, resSchema, d, m)

		return ResourceRead(ctx, resSchema, d, m)
	}

	resUpdate := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		stopContainer(ctx, resSchema, d, m)

		// Run DefaultUpdate.
		diags := ResourceUpdate(ctx, resSchema, d, m)
		if diags.HasError() {
			return diags
		}
		startContainer(ctx, resSchema, d, m)

		return ResourceRead(ctx, resSchema, d, m)
	}

	resDelete := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		// Stop container
		stopContainer(ctx, resSchema, d, m)

		// Run DefaultDelete.
		return ResourceDelete(ctx, resSchema, d, m)
	}

	return &schema.Resource{
		CreateContext: resCreate,
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: resUpdate,
		DeleteContext: resDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(1 * time.Minute),
		},
	}
}

func startContainer(ctx context.Context, s map[string]*schema.Schema, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	stopStateConf := &retry.StateChangeConf{
		Pending: []string{"pulling", "extracting"},
		Target:  []string{"stopped"},
		Refresh: func() (result interface{}, state string, err error) {
			metadata := GetMetadata(s)

			res, err := ReadItems(&ItemId{metadata.IdType, d.Id()}, metadata.Path, m.(Client))
			if err != nil {
				return res, (*res)[0]["status"], err
			}

			return res, (*res)[0]["status"], nil
		},
		Timeout: d.Timeout(schema.TimeoutCreate),
	}
	_, err := stopStateConf.WaitForStateContext(ctx)
	if err != nil {
		err = fmt.Errorf("error waiting for container instance (%s) to be pulled: %s", d.Id(), err)
		return diag.FromErr(err)
	}

	item := MikrotikItem{"number": d.Id()}

	// Start container
	var resUrl = &URL{
		Path: s[MetaResourcePath].Default.(string),
	}
	if m.(Client).GetTransport() == TransportREST {
		resUrl.Path += "/start"
	}

	err = m.(Client).SendRequest(crudStart, resUrl, item, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	startStateConf := &retry.StateChangeConf{
		Pending: []string{"stopped"},
		Target:  []string{"running"},
		Refresh: func() (result interface{}, state string, err error) {
			metadata := GetMetadata(s)

			res, err := ReadItems(&ItemId{metadata.IdType, d.Id()}, metadata.Path, m.(Client))
			if err != nil {
				return res, (*res)[0]["status"], err
			}

			return res, (*res)[0]["status"], nil
		},
		Timeout: d.Timeout(schema.TimeoutCreate),
	}
	_, err = startStateConf.WaitForStateContext(ctx)
	if err != nil {
		err = fmt.Errorf("error waiting for container instance (%s) to be started: %s", d.Id(), err)
		return diag.FromErr(err)
	}

	return nil
}

func stopContainer(ctx context.Context, s map[string]*schema.Schema, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	item := MikrotikItem{"number": d.Id()}

	var resUrl = &URL{
		Path: s[MetaResourcePath].Default.(string),
	}
	if m.(Client).GetTransport() == TransportREST {
		resUrl.Path += "/stop"
	}

	err := m.(Client).SendRequest(crudStop, resUrl, item, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	stopStateConf := &retry.StateChangeConf{
		Pending: []string{"stopping"},
		Target:  []string{"stopped"},
		Refresh: func() (result interface{}, state string, err error) {
			metadata := GetMetadata(s)

			res, err := ReadItems(&ItemId{metadata.IdType, d.Id()}, metadata.Path, m.(Client))
			if err != nil {
				return res, (*res)[0]["status"], err
			}

			return res, (*res)[0]["status"], nil
		},
		Timeout: d.Timeout(schema.TimeoutDelete),
	}
	_, err = stopStateConf.WaitForStateContext(ctx)
	if err != nil {
		err = fmt.Errorf("error waiting for container instance (%s) to be stopped: %s", d.Id(), err)
		return diag.FromErr(err)
	}
	return nil
}
