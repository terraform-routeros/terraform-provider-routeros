package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testBGPConnectionAddress = "routeros_bgp_connection.test"

func TestAccBGPConnectionTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/routing/bgp/connection", "routeros_bgp_connection"),
				Steps: []resource.TestStep{
					{
						Config: testAccBGPConnectionConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckBGPConnectionExists(testBGPConnectionAddress),
							resource.TestCheckResourceAttr(testBGPConnectionAddress, "name", "neighbor-test"),
						),
					},
				},
			})
		})
	}
}

func testAccCheckBGPConnectionExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("not found: %s", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no id is set")
		}

		return nil
	}
}

func testAccBGPConnectionConfig() string {
	return providerConfig + `

resource "routeros_bgp_connection" "test" {
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
}
