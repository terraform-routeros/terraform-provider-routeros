package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testSystemNtpClient = "routeros_system_ntp_client.test"

func TestAccSystemNtpClientTest_basic(t *testing.T) {
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
						Config: testAccSystemNtpClientConfig("manycast", "10.10.10.10"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testSystemNtpClient),
							resource.TestCheckResourceAttr(testSystemNtpClient, "mode", "manycast"),
						),
					},
					{
						Config: testAccSystemNtpClientConfig("unicast", "168.119.4.163"),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(testSystemNtpClient, "mode", "unicast"),
						),
					},
				},
			})

		})
	}
}

func testAccSystemNtpClientConfig(mode, servers string) string {
	return fmt.Sprintf(`%v

resource "routeros_system_ntp_client" "test" {
	enabled  = true
	mode     = "%v"
	servers  = ["%v"]
}
`, providerConfig, mode, servers)
}
