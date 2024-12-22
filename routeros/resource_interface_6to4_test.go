package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterface6to4 = "routeros_interface_6to4.test"

func TestAccInterface6to4Test_basic(t *testing.T) {
	// t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/6to4", "routeros_interface_6to4"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterface6to4Config(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterface6to4),
							resource.TestCheckResourceAttr(testInterface6to4, "name", "6to4-tunnel1"),
							resource.TestCheckResourceAttr(testInterface6to4, "keepalive", "10s,10"),
						),
					},
				},
			})

		})
	}
}

func testAccInterface6to4Config() string {
	return fmt.Sprintf(`%v

resource "routeros_interface_6to4" "test" {
  name      = "6to4-tunnel1"
  keepalive = "10,10"
}
`, providerConfig)
}
