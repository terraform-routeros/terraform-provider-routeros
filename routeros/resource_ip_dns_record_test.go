package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testIpDhsServer = "routeros_dns_record.test_dns_record"

func TestAccIpDnsRecordTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/dns/static", "routeros_dns_record"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpDnsRecordConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckIpDnsRecordExists(testIpDhsServer),
							resource.TestCheckResourceAttr(testIpDhsServer, "address", "192.168.88.33"),
							resource.TestCheckResourceAttr(testIpDhsServer, "name", "resource.tf"),
							resource.TestCheckResourceAttr(testIpDhsServer, "ttl", "8h"),
						),
					},
				},
			})

		})
	}
}

func testAccCheckIpDnsRecordExists(name string) resource.TestCheckFunc {
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

func testAccIpDnsRecordConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_dns_record" "test_dns_record" {
	address 	 = "192.168.88.33"
	name		 = "resource.tf"
    ttl			 = "8h"
  }

`
}
