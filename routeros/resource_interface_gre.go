package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceInterfaceGre https://wiki.mikrotik.com/wiki/Manual:Interface/Gre
func ResourceInterfaceGre() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/gre"),
		MetaId:           PropId(Name),

		KeyActualMtu:     PropActualMtuRo,
		KeyAllowFastPath: PropAllowFastPathRw,
		KeyClampTcpMss:   PropClampTcpMssRw,
		KeyComment:       PropCommentRw,
		KeyDisabled:      PropDisabledRw,
		KeyDontFragment:  PropDontFragmentRw,
		KeyDscp:          PropDscpRw,
		KeyIpsecSecret:   PropIpsecSecretRw,
		KeyKeepalive:     PropKeepaliveRw,
		KeyL2Mtu:         PropL2MtuRo,
		KeyLocalAddress:  PropLocalAddressRw,
		KeyMtu:           PropMtuRw(),
		KeyName:          PropNameForceNewRw,
		KeyRemoteAddress: PropRemoteAddressRw,
		KeyRunning:       PropRunningRo,
	}

	return &schema.Resource{
		CreateContext: DefaultValidateCreate(resSchema, func(d *schema.ResourceData) diag.Diagnostics {
			if d.Get("allow_fast_path").(bool) && d.Get("ipsec_secret").(string) != "" {
				return diag.Errorf("can't enable fastpath together with ipsec")
			}
			return nil
		}),
		ReadContext: DefaultRead(resSchema),
		UpdateContext: DefaultValidateUpdate(resSchema, func(d *schema.ResourceData) diag.Diagnostics {
			if d.Get("allow_fast_path").(bool) && d.Get("ipsec_secret").(string) != "" {
				return diag.Errorf("can't enable fastpath together with ipsec")
			}
			return nil
		}),
		DeleteContext: DefaultDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
