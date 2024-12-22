package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfaceLteApn = "routeros_interface_lte_apn.test"

func TestAccInterfaceLteApnTest_basic(t *testing.T) {
	// t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceLteApnConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceLteApn),
							resource.TestCheckResourceAttr(testInterfaceLteApn, "name", "apn1"),
							resource.TestCheckResourceAttr(testInterfaceLteApn, "apn", "internet"),
							resource.TestCheckResourceAttr(testInterfaceLteApn, "authentication", "pap"),
						),
					},
				},
			})

		})
	}
}

func testAccInterfaceLteApnConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_interface_lte_apn" "test" {
  name           = "apn1"
  apn            = "internet"
  authentication = "pap"
}
`, providerConfig)
}
