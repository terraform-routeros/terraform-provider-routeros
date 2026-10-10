package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testVlanAddress = "routeros_interface_vlan.vlan900"
const testVlanName = "VLAN_900_TEST"

func TestAccInterfaceVlanTest(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/vlan", "routeros_interface_vlan"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceVlanConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testVlanAddress),
							resource.TestCheckResourceAttr(testVlanAddress, "name", testVlanName),
						),
					},
				},
			})

		})
	}
}

func testAccInterfaceVlanConfig() string {
	return providerConfig + `

resource "routeros_interface_vlan" "vlan900" {
	name      = "VLAN_900_TEST"
	vlan_id   = 900
	disabled  = true
	interface = "bridge"
}
`
}

func TestInterfaceVlanSchema(t *testing.T) {
	RouterOSVersion = "7.24.5"
	s := ResourceInterfaceVlan().Schema

	if got := s["l3_hw_offloading"]; got == nil || !got.Optional || got.Computed || got.DiffSuppressFunc == nil {
		t.Fatalf("l3_hw_offloading must be an optional desired-state attribute")
	}
	if got := s[KeyHwOffloaded]; got == nil || !got.Computed || got.Optional {
		t.Fatalf("hw_offloaded must remain read-only")
	}

	for _, test := range []struct {
		value    bool
		expected string
	}{
		{value: true, expected: "yes"},
		{value: false, expected: "no"},
	} {
		if got := SnakeToKebab("l3_hw_offloading"); got != "l3-hw-offloading" {
			t.Fatalf("l3_hw_offloading maps to %q, want %q", got, "l3-hw-offloading")
		}
		if got := BoolToMikrotikJSON(test.value); got != test.expected {
			t.Fatalf("l3_hw_offloading=%v serializes as %q, want %q", test.value, got, test.expected)
		}
	}
}

func TestInterfaceVlanStateRefreshPreservesBooleanValues(t *testing.T) {
	RouterOSVersion = "7.24.5"
	s := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/vlan"),
		MetaId:           PropId(Id),
		"l3_hw_offloading": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		KeyHwOffloaded: PropHwOffloadedRo,
	}

	for _, test := range []struct {
		name        string
		item        MikrotikItem
		wantL3HW    bool
		wantHWState bool
	}{
		{
			name:        "enabled",
			item:        MikrotikItem{"l3-hw-offloading": "true", "hw-offloaded": "true"},
			wantL3HW:    true,
			wantHWState: true,
		},
		{
			name:        "explicitly disabled",
			item:        MikrotikItem{"l3-hw-offloading": "false", "hw-offloaded": "false"},
			wantL3HW:    false,
			wantHWState: false,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			d := schema.TestResourceDataRaw(t, s, nil)
			if diags := MikrotikResourceDataToTerraform(test.item, s, d); diags.HasError() {
				t.Fatalf("state refresh returned diagnostics: %v", diags)
			}
			if got := d.Get("l3_hw_offloading"); got != test.wantL3HW {
				t.Fatalf("l3_hw_offloading = %v, want %v", got, test.wantL3HW)
			}
			if got := d.Get(KeyHwOffloaded); got != test.wantHWState {
				t.Fatalf("hw_offloaded = %v, want %v", got, test.wantHWState)
			}
		})
	}
}

func TestInterfaceVlanStateRefreshLeavesOmittedL3HWUnset(t *testing.T) {
	RouterOSVersion = "7.24.5"
	s := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/vlan"),
		MetaId:           PropId(Id),
		"l3_hw_offloading": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		KeyHwOffloaded: PropHwOffloadedRo,
	}
	d := schema.TestResourceDataRaw(t, s, nil)
	if diags := MikrotikResourceDataToTerraform(MikrotikItem{}, s, d); diags.HasError() {
		t.Fatalf("state refresh returned diagnostics: %v", diags)
	}
	if got := d.Get("l3_hw_offloading"); got != nil && got != false {
		t.Fatalf("omitted l3_hw_offloading = %v, want unset or false", got)
	}
}

func TestInterfaceVlanImporterIsConfigured(t *testing.T) {
	r := ResourceInterfaceVlan()
	if r.Importer == nil || r.Importer.StateContext == nil {
		t.Fatal("routeros_interface_vlan must retain its importer")
	}
}

func TestInterfaceVlanStateRefreshDetectsDrift(t *testing.T) {
	RouterOSVersion = "7.24.5"
	s := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/vlan"),
		MetaId:           PropId(Id),
		"l3_hw_offloading": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		KeyHwOffloaded: PropHwOffloadedRo,
	}
	d := schema.TestResourceDataRaw(t, s, map[string]interface{}{
		"l3_hw_offloading": true,
	})
	if diags := MikrotikResourceDataToTerraform(
		MikrotikItem{"l3-hw-offloading": "false", "hw-offloaded": "false"},
		s,
		d,
	); diags.HasError() {
		t.Fatalf("state refresh returned diagnostics: %v", diags)
	}
	if got := d.Get("l3_hw_offloading"); got != false {
		t.Fatalf("drifted l3_hw_offloading = %v, want false", got)
	}
}
