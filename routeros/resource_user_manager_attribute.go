package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
    ".id": "*1",
    "default": "true",
    "default-name": "Framed-IP-Address",
    "name": "Framed-IP-Address",
    "packet-types": "access-accept",
    "standard-name": "Framed-IP-Address",
    "type-id": "8",
    "value-type": "ip-address",
    "vendor-id": "standard"
}
*/

// https://help.mikrotik.com/docs/display/ROS/User+Manager#UserManager-Attributes
func ResourceUserManagerAttribute() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/user-manager/attribute"),
		MetaId:           PropId(Id),

		KeyDefault:     PropDefaultRo,
		KeyDefaultName: PropDefaultNameRo("The attribute's default name."),
		KeyName:        PropName("The attribute's name."),
		"packet_types": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"access-accept", "access-challenge"}, false),
			},
			Description:      "A set of `access-accept` and `access-challenge`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"standard_name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"type_id": {
			Type:        schema.TypeInt,
			Required:    true,
			Description: "Attribute identification number from the specific vendor's attribute database.",
		},
		"value_type": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "hex",
			Description:  "The attribute's value type.",
			ValidateFunc: validation.StringInSlice([]string{"hex", "ip-address", "ip6-prefix", "macro", "string", "uint32"}, false),
		},
		"vendor_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "standard",
			Description: "IANA allocated a specific enterprise identification number.",
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

		Schema: resSchema,
	}
}
