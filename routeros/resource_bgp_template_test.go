package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testBGPTemplateAddress = "routeros_bgp_template.test"

func TestAccBGPTemplateTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/routing/bgp/template", "routeros_bgp_template"),
				Steps: []resource.TestStep{
					{
						Config: testAccBGPTemplateConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckBGPTemplateExists(testBGPTemplateAddress),
							resource.TestCheckResourceAttr(testBGPTemplateAddress, "name", "test-template"),
						),
					},
				},
			})
		})
	}
}

func testAccCheckBGPTemplateExists(name string) resource.TestCheckFunc {
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

func testAccBGPTemplateConfig() string {
	return providerConfig + `

resource "routeros_bgp_template" "test" {
  name = "test-template"
  as   = 65521
  input {
	limit_process_routes_ipv4 = 5
	limit_process_routes_ipv6 = 5
  }
  output {
    affinity             = "alone"
    keep_sent_attributes = true
    default_originate    = "never"
  }
  // save_to = "bgp.dump"
}
`
}
