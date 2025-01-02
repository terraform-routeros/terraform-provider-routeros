package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfaceL2tpClient = "routeros_interface_l2tp_client.test"

func TestAccInterfaceL2tpClientTest_basic(t *testing.T) {
	t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/l2tp-client", "routeros_interface_l2tp_client"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceL2tpClientConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceL2tpClient),
							resource.TestCheckResourceAttr(testInterfaceL2tpClient, "name", "l2tp-test-out"),
							resource.TestCheckResourceAttr(testInterfaceL2tpClient, "connect_to", "127.0.0.1"),
							resource.TestCheckResourceAttr(testInterfaceL2tpClient, "user", "aaa"),
							resource.TestCheckResourceAttr(testInterfaceL2tpClient, "password", "bbb"),
						),
					},
				},
			})

		})
	}
}

func testAccInterfaceL2tpClientConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_interface_l2tp_client" "test" {
  name       = "l2tp-test-out"
  connect_to = "127.0.0.1"
  user       = "aaa"
  password   = "bbb"
}
`, providerConfig)
}
