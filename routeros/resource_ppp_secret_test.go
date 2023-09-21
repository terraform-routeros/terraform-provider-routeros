package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testPPPSecret = "routeros_ppp_secret.test"

func TestAccPPPSecretTest_basic(t *testing.T) {
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
							testResourcePrimaryInstanceId(testPPPSecret),
							resource.TestCheckResourceAttr(testPPPSecret, "name", "user-test"),
						),
					},
				},
			})

		})
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
