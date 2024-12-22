package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testUserSettings = "routeros_system_user_settings.test"

func TestAccUserSettingsTest_basic(t *testing.T) {
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
						Config: testAccUserSettingsConfig("1", "5"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testUserSettings),
							resource.TestCheckResourceAttr(testUserSettings, "minimum_categories", "1"),
							resource.TestCheckResourceAttr(testUserSettings, "minimum_password_length", "5"),
						),
					},
					{
						Config: testAccUserSettingsConfig("0", "0"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testUserSettings),
							resource.TestCheckResourceAttr(testUserSettings, "minimum_categories", "0"),
							resource.TestCheckResourceAttr(testUserSettings, "minimum_password_length", "0"),
						),
					},
				},
			})

		})
	}
}

func testAccUserSettingsConfig(p1, p2 string) string {
	return fmt.Sprintf(`%v

resource "routeros_system_user_settings" "test" {
	minimum_categories      = %v
	minimum_password_length = %v
}
`, providerConfig, p1, p2)
}
