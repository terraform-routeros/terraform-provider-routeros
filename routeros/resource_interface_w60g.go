package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
".id":"*2",
"arp":"enabled",
"arp-timeout":"auto",
"beamforming-event":"1858",
"default-scan-list":"58320,60480,62640,64800",
"disabled":"false",
"frequency":"60480",
"isolate-stations":"false",
"l2mtu":"1600",
"mac-address":"48:8F:5A:F3:34:E0",
"mode":"ap-bridge",
"mtu":"1500",
"name":"wlan60-1",
"password":"password",
"put-stations-in-bridge":"bridge",
"region":"eu",
"running":"false",
"rx-mpdu-crc-err":"20354232",
"rx-mpdu-crc-ok":"1401642596",
"rx-ppdu":"1403448812",
"ssid":"ptmp01",
"tx-fw-msdu":"52",
"tx-io-msdu":"0",
"tx-mpdu-new":"202533008",
"tx-mpdu-retry":"3570458",
"tx-mpdu-total":"206103466",
"tx-ppdu":"3230579961",
"tx-ppdu-from-q":"2031518609",
"tx-sector":"auto",
"tx-sw-msdu":"202533128"
}
*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/39059501/W60G#W60G-Generalinterfaceproperties
func ResourceInterfaceW60g() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/w60g"),
		MetaId:           PropId(Id),
		MetaSkipFields: PropSkipFields("beamforming_event", "rx_mpdu_crc_err", "rx_mpdu_crc_ok", "rx_ppdu", "tx_fw_msdu",
			"tx_io_msdu", "tx_mpdu_new", "tx_mpdu_retry", "tx_mpdu_total", "tx_ppdu", "tx_ppdu_from_q", "tx_sector",
			"tx_sw_msdu"),

		KeyArp:        PropArpRw,
		KeyArpTimeout: PropArpTimeoutRw,
		KeyComment:    PropCommentRw,
		KeyDisabled:   PropDisabledRw,
		"frequency": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Frequency used in communication (Only active on bridge device).",
			ValidateFunc:     validation.StringInSlice([]string{"58320", "60480", "62640", "64800", "66000", "auto"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"isolate_stations": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Don't allow communication between connected clients (from RouterOS 6.41).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyL2Mtu:      PropL2MtuRw,
		KeyMacAddress: PropMacAddressRw("MAC address of the radio interface.", false),
		"mdmg_fix": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Experimental feature working only on wAP60Gx3 devices, providing better point to multi point " +
				"stability in some cases.",
		},
		"mode": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Operation mode.",
			ValidateFunc:     validation.StringInSlice([]string{"ap-bridge", "bridge", "sniff", "station-bridge"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyMtu:  PropMtuRw(),
		KeyName: PropName("Name of the interface."),
		"password": {
			Type:             schema.TypeString,
			Optional:         true,
			Sensitive:        true,
			Description:      "Password used for AES encryption.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"put_stations_in_bridge": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Put newly created station device interfaces in this bridge.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"region": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Parameter to limit frequency use.",
			ValidateFunc:     validation.StringInSlice([]string{"asia", "australia", "canada", "china", "eu", "japan", "no-region-set", "usa"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyRunning: PropRunningRo,
		"scan_list": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Scan list to limit connectivity over frequencies in station mode.",
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"58320", "60480", "62640", "64800", "66000", "auto"}, false),
			},
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ssid": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "SSID (service set identifier) is a name that identifies wireless network (0..32 char).",
			ValidateFunc:     validation.StringLenBetween(0, 32),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tx_sector": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Disables beamforming and locks to selected radiation pattern.",
			ValidateFunc:     validation.IntBetween(0, 63),
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
