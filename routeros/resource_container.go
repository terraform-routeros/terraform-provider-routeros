package routeros

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  {
    ".id": "*3",
    "arch": "amd64",
    "interface": "veth1",
    "mounts": "",
    "name": "76e7fc0c-e2c0-4b8c-b3b0-4c985c496a90",
    "os": "linux",
    "repo": "registry-1.docker.io/adguard/adguardhome:latest",
    "root-dir": "sata4-part1/docker/adg",
    "status": "stopped",
    "workdir": "/opt/adguardhome/work"
  }

  RouterOS 7.2x has no "status"; the state comes as flags that are only present while true:
  "stopped", "starting", "running", "stopping". While the image is pulled and extracted REST sends
  none of them, only ".about": "extracting layer 2/5 66%/9.00M" (observed on 7.25).
*/

// https://help.mikrotik.com/docs/display/ROS/Container#Container-Properties
func ResourceContainer() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/container"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("running", "starting", "stopped", "passed_devs", "stopping", "extracting", "config_json"),

		"arch": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The architecture of the container image",
		},
		"auto_restart_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Specify an interval at which Container will be restarted on Container failure.",
			DiffSuppressFunc: TimeEqual,
		},
		"check_certificate": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Enables trust chain validation from local certificate store.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"cmd": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The main purpose of a CMD is to provide defaults for an executing container. These defaults can include an executable, or they can omit the executable, in which case you must specify an ENTRYPOINT instruction as well.",
		},
		KeyComment: PropCommentRw,
		"devices": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Passes through physical device to the container.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
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
		"memory_high": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "RAM usage limit in bytes for a specific container (string value).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
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
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				if old == "" {
					return false
				}

				if AlwaysPresentNotUserProvided(k, old, new, d) {
					return true
				}

				// Checking the presence of a tag:
				// ~ remote_image = "traefik/whoami:latest" -> "traefik/whoami" # forces replacement
				return old == new || old == new+":latest"
			},
		},
		"root_dir": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Used to save container store outside main memory",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"running": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Container state.",
			Default:     true,
		},
		"start_on_boot": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Start the container on boot",
		},
		"status": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The status of the container: `pulling`, `extracting`, `starting`, `running`, `stopping` or `stopped`.",
		},
		"stop_signal": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "15-SIGTERM",
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

	resRead := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		metadata := GetMetadata(resSchema)

		res, err := ReadItems(&ItemId{metadata.IdType, d.Id()}, metadata.Path, m.(Client))
		if err != nil {
			ColorizedDebug(ctx, fmt.Sprintf(ErrorMsgGet, err))
			return diag.FromErr(err)
		}

		// Resource not found.
		if len(*res) == 0 {
			d.SetId("")
			return nil
		}

		item := (*res)[0]
		d.SetId(item.GetID(metadata.IdType))

		diags := MikrotikResourceDataToTerraform(item, resSchema, d)
		if diags.HasError() {
			return diags
		}

		// The `running` field is excluded from the default read (MetaSkipFields), so the state is
		// synchronized here. RouterOS returns the container state in the `running` field ("true"/"false"),
		// older versions returned a `status` field instead.
		if v, ok := item["running"]; ok {
			d.Set("running", BoolFromMikrotikJSON(v))
		} else {
			d.Set("running", item["status"] == "running")
		}
		d.Set("status", containerState(item))

		tag, ok := d.Get("tag").(string)
		if ok && tag != "" {
			// Get Registry URL.
			res, err := ReadItems(nil, "/container/config", m.(Client))
			if err != nil {
				ColorizedDebug(ctx, fmt.Sprintf(ErrorMsgPut, err))
				return diag.FromErr(err)
			}

			if len(*res) == 0 {
				return diag.Diagnostics{
					diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "Failed to retrieve the URL of the container registry, the response is empty",
					},
				}
			}

			registryUrl, ok := (*res)[0]["registry-url"]
			if !ok {
				// Key name for an unspecified registry.
				registryUrl, ok = (*res)[0]["assumed-registry-url"]
				if !ok {
					return diag.Diagnostics{
						diag.Diagnostic{
							Severity: diag.Error,
							Summary:  "The `registry-url` was not found in the response",
						},
					}
				}
			}

			u, err := url.Parse(registryUrl)
			if err != nil {
				return diag.FromErr(err)
			}

			// Remove http(s); host:port; path...
			for _, item := range []string{u.Scheme, u.Host, u.Path} {
				tag = strings.TrimPrefix(tag, item)
			}
			// Remove (:////)adguard/adguardhome:latest
			tag = strings.TrimLeft(tag, ":/")

			d.Set("remote_image", strings.TrimPrefix(tag, registryUrl))
		}

		return nil
	}

	resCreate := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		// Run DefaultCreate.
		diags := ResourceCreate(ctx, resSchema, d, m)
		if diags.HasError() {
			return diags
		}

		if d.Get("running").(bool) {
			if diags := startContainer(ctx, resSchema, d, m, d.Timeout(schema.TimeoutCreate)); diags.HasError() {
				return diags
			}
		}

		return resRead(ctx, d, m)
	}

	resUpdate := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		timeout := d.Timeout(schema.TimeoutUpdate)

		// RouterOS only accepts changes on a stopped container, so every update restarts it.
		if diags := stopContainer(ctx, resSchema, d, m, timeout); diags.HasError() {
			return diags
		}

		// Run DefaultUpdate.
		diags := ResourceUpdate(ctx, resSchema, d, m)
		if diags.HasError() {
			return diags
		}
		if d.Get("running").(bool) {
			if diags := startContainer(ctx, resSchema, d, m, timeout); diags.HasError() {
				return diags
			}
		}

		return resRead(ctx, d, m)
	}

	resDelete := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		// RouterOS refuses to remove a running container.
		if diags := stopContainer(ctx, resSchema, d, m, d.Timeout(schema.TimeoutDelete)); diags.HasError() {
			return diags
		}

		// Run DefaultDelete.
		return ResourceDelete(ctx, resSchema, d, m)
	}

	return &schema.Resource{
		CreateContext: resCreate,
		ReadContext:   resRead,
		UpdateContext: resUpdate,
		DeleteContext: resDelete,

		Importer: &schema.ResourceImporter{
			StateContext: ImportStateCustomContext(resSchema),
		},

		Schema: resSchema,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(1 * time.Minute),
		},
	}
}

// containerState maps a container item to pulling, extracting, starting, running, stopping or
// stopped, from `status` (7.1x) or the state flags (7.2x, see the sample at the top of the file).
func containerState(item MikrotikItem) string {
	if status, ok := item["status"]; ok && status != "" {
		return status
	}
	for _, flag := range []string{"extracting", "starting", "stopping", "running", "stopped"} {
		if BoolFromMikrotikJSON(item[flag]) {
			return flag
		}
	}
	return "pulling"
}

var containerTransientStates = []string{"pulling", "extracting", "starting", "stopping"}

// waitForContainerState polls the container until its state is one of targets and returns that
// state. States in pending keep the poll going, anything else is an error.
func waitForContainerState(ctx context.Context, s map[string]*schema.Schema, d *schema.ResourceData, m interface{},
	pending, targets []string, timeout time.Duration) (string, error) {
	metadata := GetMetadata(s)
	conf := &retry.StateChangeConf{
		Pending: pending,
		Target:  targets,
		Refresh: func() (interface{}, string, error) {
			res, err := ReadItems(&ItemId{metadata.IdType, d.Id()}, metadata.Path, m.(Client))
			if err != nil {
				return nil, "", err
			}
			if len(*res) == 0 {
				return nil, "", fmt.Errorf("container %s not found", d.Id())
			}
			return res, containerState((*res)[0]), nil
		},
		Timeout: timeout,
	}
	res, err := conf.WaitForStateContext(ctx)
	if err != nil {
		return "", err
	}
	return containerState((*res.(*[]MikrotikItem))[0]), nil
}

// sendContainerCommand runs /container/<command> for this container. The item is addressed by
// `.id` on both transports, like /system/script/run; RouterOS 7.2x REST rejects `number`.
func sendContainerCommand(s map[string]*schema.Schema, d *schema.ResourceData, m interface{}, crud crudMethod, command string) error {
	resUrl := &URL{Path: s[MetaResourcePath].Default.(string)}
	if m.(Client).GetTransport() == TransportREST {
		resUrl.Path += "/" + command
	}
	return m.(Client).SendRequest(crud, resUrl, MikrotikItem{Id.String(): d.Id()}, nil)
}

// startContainer waits for the image to be pulled and extracted, then starts the container
// unless it is already running.
func startContainer(ctx context.Context, s map[string]*schema.Schema, d *schema.ResourceData, m interface{}, timeout time.Duration) diag.Diagnostics {
	state, err := waitForContainerState(ctx, s, d, m, containerTransientStates, []string{"stopped", "running"}, timeout)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error waiting for container instance (%s) to be pulled: %s", d.Id(), err))
	}
	if state == "running" {
		return nil
	}

	if err := sendContainerCommand(s, d, m, crudStart, "start"); err != nil {
		return diag.FromErr(err)
	}

	if _, err := waitForContainerState(ctx, s, d, m, []string{"stopped", "starting"}, []string{"running"}, timeout); err != nil {
		return diag.FromErr(fmt.Errorf("error waiting for container instance (%s) to be started: %s", d.Id(), err))
	}
	return nil
}

// stopContainer stops the container and waits until RouterOS reports it stopped. A container that
// is not running is left alone, RouterOS rejects stop for it.
func stopContainer(ctx context.Context, s map[string]*schema.Schema, d *schema.ResourceData, m interface{}, timeout time.Duration) diag.Diagnostics {
	state, err := waitForContainerState(ctx, s, d, m, containerTransientStates, []string{"stopped", "running"}, timeout)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error waiting for container instance (%s) to settle: %s", d.Id(), err))
	}
	if state != "running" {
		return nil
	}

	if err := sendContainerCommand(s, d, m, crudStop, "stop"); err != nil {
		return diag.FromErr(err)
	}

	if _, err := waitForContainerState(ctx, s, d, m, []string{"running", "stopping"}, []string{"stopped"}, timeout); err != nil {
		return diag.FromErr(fmt.Errorf("error waiting for container instance (%s) to be stopped: %s", d.Id(), err))
	}
	return nil
}
