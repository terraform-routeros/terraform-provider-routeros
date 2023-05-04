package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testCapsManRatesAddress = "routeros_capsman_rates.test_rates"

func TestAccCapsManRatesTest_basic(t *testing.T) {
	t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/caps-man/rates", "routeros_capsman_rates"),
				Steps: []resource.TestStep{
					{
						Config: testAccCapsManRatesConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckCapsManRatesExists(testCapsManRatesAddress),
							resource.TestCheckResourceAttr(testCapsManRatesAddress, "name", "test_rates"),
						),
					},
				},
			})

		})
	}
}

func testAccCheckCapsManRatesExists(name string) resource.TestCheckFunc {
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

func testAccCapsManRatesConfig() string {
	return providerConfig + `

resource "routeros_capsman_rates" "test_rates" {
	name              = "test_rates"
	comment           = "test_rates"
	basic             = ["1Mbps", "5.5Mbps", "6Mbps", "18Mbps", "36Mbps", "54Mbps"]
	ht_basic_mcs      = ["mcs-0", "mcs-7", "mcs-11", "mcs-14", "mcs-16", "mcs-21"]
	ht_supported_mcs  = ["mcs-3", "mcs-8", "mcs-10", "mcs-13", "mcs-17", "mcs-18"]
	supported         = ["2Mbps", "11Mbps", "9Mbps", "12Mbps", "24Mbps", "48Mbps"]
	vht_basic_mcs     = "none"
	vht_supported_mcs = "mcs0-9,mcs0-7"
}
`
}
