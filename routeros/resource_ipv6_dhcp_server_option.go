package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*80000004",
    "code": "24",
    "name": "domain-search",
    "raw-value": "076578616d706c65056c6f63616c00",
    "value": "0x07'example'0x05'local'0x00"
  }
*/

// https://help.mikrotik.com/docs/display/ROS/DHCP#DHCP-DHCPOptions.1
// https://www.ipamworldwide.com/ipam/isc-dhcpv6-options.html
// https://jjjordan.github.io/dhcp119/
func ResourceIpv6DhcpServerOption() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ipv6/dhcp-server/option"),
		MetaId:           PropId(Id),

		"code": {
			Type:         schema.TypeInt,
			Required:     true,
			Description:  "Dhcp option [code](https://www.ipamworldwide.com/ipam/isc-dhcpv6-options.html).",
			ValidateFunc: validation.IntBetween(1, 254),
		},
		KeyComment: PropCommentRw,
		KeyName:    PropName("Descriptive name of the option."),
		"value": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Parameter's value. Available data types for options are:\n" +
				"    - `'test'` -> ASCII to Hex 0x74657374\n" +
				"    - `'10.10.10.10'` -> Unicode IP to Hex 0x0a0a0a0a\n" +
				"    - `s'10.10.10.10'` -> ASCII to Hex 0x31302e31302e31302e3130\n" +
				"    - `s'160'` -> ASCII to Hex 0x313630\n" +
				"    - `'10'` -> Decimal to Hex 0x0a\n" +
				"    - `0x0a0a` -> No conversion\n" +
				"    - `$(VARIABLE)` -> hardcoded values\n\n" +
				"	RouterOS has predefined variables that can be used:\n" +
				"    - `HOSTNAME` - client hostname\n" +
				"    - `RADIUS_MT_STR1` - from radius MT attr nr. `24`\n" +
				"    - `RADIUS_MT_STR2` - from radius MT attr nr. `25`\n" +
				"    - `REMOTE_ID` - agent remote-id\n" +
				"    - `NETWORK_GATEWAY - the first gateway from `/ip dhcp-server network`, note that this option " +
				"won't work if used from lease.\n\nNow it is also possible to combine data types into one, for example: " +
				"`0x01'vards'$(HOSTNAME)`For example if HOSTNAME is 'kvm', then raw value will be 0x0176617264736b766d.",
		},
		"raw_value": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Read-only field which shows raw DHCP option value (the format actually sent out).",
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
