package routeros

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	dhcpServerAuthoritative = "authoritative"
	dhcpServerBootpLease    = "bootp_lease_time"
)

func ReadDhcpServer(ctx context.Context, s map[string]*schema.Schema, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	res, err := ReadItems(&ItemId{Id, d.Id()}, s[MetaResourcePath].Default.(string), m.(Client))
	if err != nil {
		return diag.FromErr(err)
	}
	if len(*res) == 0 {
		d.SetId("")
		return nil
	}

	item := (*res)[0]
	if err := validateDhcpServerRESTFields(d, item, m.(Client)); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(item.GetID(Id))
	return MikrotikResourceDataToTerraform(item, s, d)
}

func UpdateDhcpServer(s map[string]*schema.Schema) schema.UpdateContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		item, metadata := TerraformResourceDataToMikrotik(s, d)
		id, err := dynamicIdLookup(metadata.IdType, metadata.Path, m.(Client), d)
		if err != nil {
			return diag.FromErr(err)
		}

		res, err := UpdateItem(&ItemId{Id, id}, metadata.Path, item, m.(Client))
		if err != nil {
			return diag.FromErr(err)
		}

		if m.(Client).GetTransport() == TransportREST {
			refreshed, err := ReadItems(&ItemId{Id, id}, metadata.Path, m.(Client))
			if err != nil {
				return diag.FromErr(err)
			}
			if len(*refreshed) == 0 {
				return diag.Errorf("Mikrotik resource path='%v' id='%v' not found after update", metadata.Path, id)
			}
			res = (*refreshed)[0]
		}

		if err := validateDhcpServerRESTFields(d, res, m.(Client)); err != nil {
			return diag.FromErr(err)
		}
		return MikrotikResourceDataToTerraform(res, s, d)
	}
}

func validateDhcpServerRESTFields(d *schema.ResourceData, item MikrotikItem, client Client) error {
	if client.GetTransport() != TransportREST {
		return nil
	}

	for _, field := range []string{dhcpServerAuthoritative, dhcpServerBootpLease} {
		if _, configured := d.GetOk(field); !configured {
			continue
		}
		if _, present := item[SnakeToKebab(field)]; !present {
			return fmt.Errorf("RouterOS REST response omitted configured DHCP server field %q", field)
		}
	}
	return nil
}
