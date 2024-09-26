package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testDatasourceWiFiEasyConnect = "data.routeros_wifi_easy_connect.test"

func TestAccDatasourceWiFiEasyConnectTest_basic(t *testing.T) {
	t.Run("QR Code", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProviderFactories: testAccProviderFactories,
			Steps: []resource.TestStep{
				{
					Config: testAccDatasourceWiFiEasyConnectConfig(),
					Check: resource.ComposeTestCheckFunc(
						testResourcePrimaryInstanceId(testDatasourceWiFiEasyConnect),
					),
				},
			},
		})

	})
}

func testAccDatasourceWiFiEasyConnectConfig() string {
	return providerConfig + `

data "routeros_wifi_easy_connect" "test" {
  type     = "WPA2"
  ssid     = "test"
  password = "password12345"
}
`
}
