package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfacePppoeServer = "routeros_interface_pppoe_server.test"

func TestAccInterfacePppoeServerTest_basic(t *testing.T) {
	t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/pppoe-server", "routeros_interface_pppoe_server"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfacePppoeServerConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfacePppoeServer),
							resource.TestCheckResourceAttr(testInterfacePppoeServer, "comment", "comment"),
							resource.TestCheckResourceAttr(testInterfacePppoeServer, "disabled", "true"),
							resource.TestCheckResourceAttr(testInterfacePppoeServer, "name", "pppoe-in1"),
							resource.TestCheckResourceAttr(testInterfacePppoeServer, "user", ""),
							resource.TestCheckResourceAttr(testInterfacePppoeServer, "service", ""),
						),
					},
				},
			})

		})
	}
}

func testAccInterfacePppoeServerConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_interface_pppoe_server" "test" {
  comment  = "comment"
  disabled = true
  name     = "pppoe-in1"
  user     = ""
  service  = ""
}
`, providerConfig)
}
