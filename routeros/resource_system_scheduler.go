package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ResourceSystemScheduler https://wiki.mikrotik.com/wiki/Manual:System/Scheduler
// https://wiki.mikrotik.com/wiki/Manual:Scripting#Variables
// https://help.mikrotik.com/docs/display/ROS/User#User-UserGroups
func ResourceSystemScheduler() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/scheduler"),
		MetaId:           PropId(Id),

		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"interval": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "Interval between two script executions, if time interval is set to zero, the script is only " +
				"executed at its start time, otherwise it is executed repeatedly at the time interval is specified.",
		},
		KeyName: PropNameForceNewRw,
		"next_run": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"on_event": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the script to execute. It must be presented at /system script.",
		},
		"owner": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"policy": {
			Type:     schema.TypeList,
			Computed: true,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"ftp", "reboot", "read", "write", "policy", "test",
					"password", "sniff", "sensitive", "romon", "dude"}, false),
			},
			Description: `List of applicable policies:
    * dude - Policy that grants rights to log in to dude server.  
    * ftp - Policy that grants full rights to log in remotely via FTP, to read/write/erase files and to transfer files from/to the router. Should be used together with read/write policies.  
    * password - Policy that grants rights to change the password.  
    * policy - Policy that grants user management rights. Should be used together with the write policy. Allows also to see global variables created by other users (requires also 'test' policy).  
    * read - Policy that grants read access to the router's configuration. All console commands that do not alter router's configuration are allowed. Doesn't affect FTP.  
    * reboot - Policy that allows rebooting the router.  
    * romon - Policy that grants rights to connect to RoMon server.  
    * sensitive - Policy that grants rights to change "hide sensitive" option, if this policy is disabled sensitive information is not displayed.  
    * sniff - Policy that grants rights to use packet sniffer tool.  
    * test - Policy that grants rights to run ping, traceroute, bandwidth-test, wireless scan, snooper, and other test commands.  
    * write - Policy that grants write access to the router's configuration, except for user management. This policy does not allow to read the configuration, so make sure to enable read policy as well.  
policy = ["ftp", "read", "write"]
`,
		},
		"run_count": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "This counter is incremented each time the script is executed.",
		},
		"start_date": {
			Type:        schema.TypeString,
			Computed:    true,
			Optional:    true,
			Description: "Date of the first script execution.",
		},
		"start_time": {
			Type:     schema.TypeString,
			Computed: true,
			Optional: true,
			Description: "Time of the first script execution. If scheduler item has start-time set to startup, it " +
				"behaves as if start-time and start-date were set to time 3 seconds after console starts up. " +
				"It means that all scripts having start-time is startup and interval is 0 will be executed once each " +
				"time router boots. If the interval is set to value other than 0 scheduler will not run at startup.",
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

		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    ResourceSystemSchedulerV0().CoreConfigSchema().ImpliedType(),
				Upgrade: stateMigrationNameToId(resSchema[MetaResourcePath].Default.(string)),
				Version: 0,
			},
		},

		Schema: resSchema,
	}
}
