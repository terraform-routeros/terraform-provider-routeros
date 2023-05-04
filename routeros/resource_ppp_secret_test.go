package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testPPPSecret = "routeros_ppp_secret.test"

func TestAccPPPSecretTest_basic(t *testing.T) {
	t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ppp/secret", "routeros_ppp_secret"),
				Steps: []resource.TestStep{
					{
						Config: testAccPPPSecretConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckPPPSecretExists(testPPPSecret),
							resource.TestCheckResourceAttr(testPPPSecret, "name", "user-test"),
						),
					},
				},
			})

		})
	}
}

func testAccCheckPPPSecretExists(name string) resource.TestCheckFunc {
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

func testAccPPPSecretConfig() string {
	return providerConfig + `
	resource "routeros_ppp_secret" "test" {
		name        = "user-test"
		password    = "12345678"
	}
`
}
