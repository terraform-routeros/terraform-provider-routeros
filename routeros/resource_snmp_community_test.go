package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testSNMPCommunityAddress = "routeros_snmp_community.test"

func TestAccSNMPCommunityTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/snmp/community", "routeros_snmp_community"),
				Steps: []resource.TestStep{
					{
						Config: testAccSNMPCommunityConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckSNMPCommunityExists(testSNMPCommunityAddress),
							resource.TestCheckResourceAttr(testSNMPCommunityAddress, "name", "private"),
						),
					},
				},
			})
		})
	}
}

func testAccCheckSNMPCommunityExists(name string) resource.TestCheckFunc {
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

func testAccSNMPCommunityConfig() string {
	return providerConfig + `
resource "routeros_snmp_community" "test" {
	authentication_password = "authpasswd"
	authentication_protocol = "MD5"
	comment                 = "Comment"
	disabled                = true
	encryption_password     = "encpassword"
	encryption_protocol     = "DES"
	name                    = "private"
	read_access             = true
	security                = "private"
	write_access            = true
}`
}
