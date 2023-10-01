package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testBGPConnectionAddress = "routeros_routing_bgp_connection.test"

func TestAccBGPConnectionTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/routing/bgp/connection", "routeros_routing_bgp_connection"),
				Steps: []resource.TestStep{
					{
						Config: testAccBGPConnectionConfig(ROSVersion),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testBGPConnectionAddress),
							resource.TestCheckResourceAttr(testBGPConnectionAddress, "name", "neighbor-test"),
						),
					},
				},
			})
		})
	}
}

func testAccBGPConnectionConfig(ver rosVersionType) string {
	switch {
	case ver < v7_10:
		return providerConfig + `

resource "routeros_routing_bgp_connection" "test" {
	name         = "neighbor-test"
	as           = "65550/5"
	as_override  = true
	add_path_out = "none"
	remote {
	  address = "172.17.0.1"
	  as      = "12345/5"
	}
	local {
	  role = "ebgp"
	}
}
`
	default:
		// ver >= v7_10
		return providerConfig + `

resource "routeros_routing_bgp_connection" "test" {
	name         = "neighbor-test"
	as           = "65550/5"
	output {
		as_override  = true
	}
	add_path_out = "none"
	remote {
	  address = "172.17.0.1"
	  as      = "12345/5"
	}
	local {
	  role = "ebgp"
	}
}
`
	}
}
