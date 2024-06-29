package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
[]
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
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"nr_band": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "5G NR Frequency band used in communication [5G NR Bands and bandwidths](https://en.wikipedia.org/wiki/5G_NR_frequency_bands).",
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"modem_init": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Modem init string (AT command that will be executed at modem startup).",
		},
		KeyMtu:  PropMtuRw(),
		KeyName: PropName("Descriptive name of the interface."),
		"network_mode": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Select/force mode for LTE interface to operate with.",
			ValidateFunc:     validation.StringInSlice([]string{"3g", "gsm", "lte", "5g"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"operator": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Used to lock the device to a specific operator full PLMN number is used for the lock " +
				"consisting of MCC+MNC. [PLMN codes](https://en.wikipedia.org/wiki/Public_land_mobile_network).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"pin": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "SIM Card's PIN code.",
		},
	}

	return &schema.Resource{
		// FIXME
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
