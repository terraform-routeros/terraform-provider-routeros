package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  "auto-upgrade": "false",
  "boot-device": "nand-if-fail-then-ethernet",
  "boot-protocol": "bootp",
  "cpu-frequency": "800MHz",
  "force-backup-booter": "false",
  "preboot-etherboot": "disabled",
  "preboot-etherboot-server": "any",
  "protected-routerboot": "disabled",
  "reformat-hold-button": "20s",
  "reformat-hold-button-max": "10m",
  "silent-boot": "false"
}
*/

// https://help.mikrotik.com/docs/display/ROS/RouterBOARD#RouterBOARD-Settings
func ResourceSystemRouterboardSettings() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/routerboard/settings"),
		MetaId:           PropId(Id),

		"auto_upgrade": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "An option to enable firmware upgrade automatically after the RouterOS upgrade.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"baud_rate": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "An option to choose the onboard RS232 speed in bits per second.",
			ValidateFunc:     validation.IntAtLeast(0),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"boot_delay": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "A delay for a keystroke while booting.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEquall,
		},
		"boot_device": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "An option to choose the way RouterBOOT loads the operating system. Possible values: " +
				"`ethernet`, `flash-boot`, `flash-boot-once-then-nand`, `nand-if-fail-then-ethernet`, `nand-only`, `try-ethernet-once-then-nand`.",
			ValidateFunc:     validation.StringInSlice([]string{"ethernet", "flash-boot", "flash-boot-once-then-nand", "nand-if-fail-then-ethernet", "nand-only", "try-ethernet-once-then-nand"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"boot_os": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "An option to choose the booting operating system for CRS3xx series switches. Possible values: `router-os`, `swos`.",
			ValidateFunc:     validation.StringInSlice([]string{"router-os", "swos"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"boot_protocol": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Boot protocol to use. Possible values: `bootp`, `dhcp`.",
			ValidateFunc:     validation.StringInSlice([]string{"bootp", "dhcp"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"cpu_frequency": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "An option to change the CPU frequency of the device.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"cpu_mode": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "An option whether to enter CPU suspend mode in HTL instruction. Possible values: `power-save`, `regular`.",
			ValidateFunc:     validation.StringInSlice([]string{"power-save", "regular"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"enable_jumper_reset": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "An option to enable reset via the onboard jumper.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"enter_setup_on": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "An option to set which key will cause the BIOS to enter configuration mode during boot delay. Possible values: `any-key`, `delete-key`.",
			ValidateFunc:     validation.StringInSlice([]string{"any-key", "delete-key"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"force_backup_booter": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "An option to use the backup RouterBOOT.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"init_delay": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "An option to set a delay before the USB port is initialized. Used for mPCIe modems with RB9xx series devices only.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEquall,
		},
		"memory_frequency": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "An option to change the memory frequency of the device. Values depend on the model.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"memory_data_rate": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "An option to change the memory data rate of the device. Values depend on the model.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"preboot_etherboot": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "An option to enable preboot `etherboot`, which runs before the regular boot device.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"preboot_etherboot_server": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "An option to instruct `preboot-etherboot` to accept only from the specified Netinstall server.",
			ValidateFunc:     validation.IsIPv4Address,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"protected_routerboot": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "An option to disable any access to the RouterBOOT configuration settings over a console cable and disables the operation of the reset button to change the boot mode (Netinstall will be disabled). Possible values: `disabled`, `enabled`.",
			ValidateFunc:     validation.StringInSlice([]string{"disabled", "enabled"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"reformat_hold_button": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "An option to enable resetting everything by pressing the button at power-on for longer than the specified time but less than `reformat_hold_button_max.`",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEquall,
		},
		"reformat_hold_button_max": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "See `reformat_hold_button`.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEquall,
		},
		"regulatory_domain_ce": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "An option to enable extra-low TX power for high antenna gain devices.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"silent_boot": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "An option to turn off output on the serial console and beeping sounds during booting.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
	}

	return &schema.Resource{
		CreateContext: DefaultSystemCreate(resSchema),
		ReadContext:   DefaultSystemRead(resSchema),
		UpdateContext: DefaultSystemUpdate(resSchema),
		DeleteContext: DefaultSystemDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
