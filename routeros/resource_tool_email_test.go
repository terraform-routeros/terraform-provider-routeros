package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testToolEmail = "routeros_tool_email.email"

func TestAccToolEmailTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/tool/e-mail", "routeros_email"),
				Steps: []resource.TestStep{
					{
						Config: testAccToolEmailConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testToolEmail),
							resource.TestCheckResourceAttr(testToolEmail, "from", "John Doe <jdoe@example.com>"),
							resource.TestCheckResourceAttr(testToolEmail, "password", "password"),
							resource.TestCheckResourceAttr(testToolEmail, "port", "25"),
							resource.TestCheckResourceAttr(testToolEmail, "server", "smtp.example.com"),
							resource.TestCheckResourceAttr(testToolEmail, "tls", "yes"),
							resource.TestCheckResourceAttr(testToolEmail, "user", "jdoe"),
							resource.TestCheckResourceAttr(testToolEmail, "vrf", "main"),
						),
					},
				},
			})

		})
	}
}

func testAccToolEmailConfig() string {
	return providerConfig + `
resource "routeros_tool_email" "email" {
	from     = "John Doe <jdoe@example.com>"
	password = "password"
	port     = "25"
	server   = "smtp.example.com"
	tls      = "yes"
	user     = "jdoe"
	vrf      = "main"
}
`
}
