package routeros

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*4",
    "allow-roaming": "true",
    "apn-profiles": "default",
    "band": "",
    "comment": "wan",
    "default-name": "lte1",
    "disabled": "false",
    "inactive": "false",
    "mtu": "1500",
    "name": "lte1",
    "network-mode": "3g,lte",
    "running": "true",
    "sms-read": "false"
  }
*/

// https://help.mikrotik.com/docs/display/ROS/LTE
func ResourceInterfaceLte() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/lte"),
		MetaId:           PropId(Id),

		"allow_roaming": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Enable data roaming for connecting to other countries data-providers. Not all LTE modems " +
				"support this feature. Some modems, that do not fully support this feature, will connect to the " +
				"network but will not establish an IP data connection with allow-roaming set to no.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"apn_profiles": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Which APN profile to use for this interface.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"band": {
			Type:     schema.TypeSet,
			Optional: true,
			Description: "LTE Frequency band used in communication " +
				"[LTE Bands and bandwidths](https://en.wikipedia.org/wiki/LTE_frequency_bands#Frequency_bands_and_channel_bandwidths).",
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
		},
		KeyComment:     PropCommentRw,
		KeyDefaultName: PropDefaultNameRo("The default name for an interface."),
		KeyDisabled:    PropDisabledRw,
		KeyInactive:    PropInactiveRo,
		"modem_init": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Modem init string (AT command that will be executed at modem startup).",
		},
		KeyMtu:  PropMtuRw(),
		KeyName: PropName("Descriptive name of the interface."),
		"network_mode": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Select/force mode for LTE interface to operate with.",
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"3g", "gsm", "lte", "5g"}, false),
			},
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"nr_band": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "5G NR Frequency band used in communication [5G NR Bands and bandwidths](https://en.wikipedia.org/wiki/5G_NR_frequency_bands).",
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
		},
		KeyRunning: PropRunningRo,
		"operator": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Used to lock the device to a specific operator full PLMN number is used for the lock " +
				"consisting of MCC+MNC. [PLMN codes](https://en.wikipedia.org/wiki/Public_land_mobile_network).",
		},
		"pin": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "SIM Card's PIN code.",
		},
		"sms_protocol": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "SMS functionality. `mbim`: uses MBIM driver. `at`: uses AT-Commands. `auto`: selects the " +
				"appropriate option depending on the modem.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"sms_read": {
			Type:             schema.TypeBool,
			Optional:         true,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
	}

	resCreateUpdate := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		res, err := ReadItems(&ItemId{Name, d.Get("name").(string)}, GetMetadata(resSchema).Path, m.(Client))
		if err != nil {
			// API/REST client error.
			ColorizedDebug(ctx, fmt.Sprintf(ErrorMsgPatch, err))
			return diag.FromErr(err)
		}

		// Resource not found.
		if len(*res) == 0 {
			d.SetId("")
			ColorizedDebug(ctx, fmt.Sprintf(ErrorMsgPatch, err))
			return diag.FromErr(errorNoLongerExists)
		}

		d.SetId((*res)[0].GetID(Id))

		if diags := ResourceUpdate(ctx, resSchema, d, m); diags.HasError() {
			return diags
		}

		return ResourceRead(ctx, resSchema, d, m)
	}

	return &schema.Resource{
		CreateContext: resCreateUpdate,
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: resCreateUpdate,
		DeleteContext: DefaultSystemDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
