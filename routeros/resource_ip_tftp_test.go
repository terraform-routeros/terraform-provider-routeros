package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfaceIpTFTPAddress = "routeros_ip_tftp.test_file"

func TestAccInterfaceIpTFTPTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/tftp", "routeros_ip_tftp"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceIpTFTPConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceIpTFTPAddress),
							resource.TestCheckResourceAttr(testInterfaceIpTFTPAddress, "ip_addresses.0", "1.2.3.4/5"),
							resource.TestCheckResourceAttr(testInterfaceIpTFTPAddress, "req_filename", "file.txt"),
							resource.TestCheckResourceAttr(testInterfaceIpTFTPAddress, "real_filename", "/file.txt"),
							resource.TestCheckResourceAttr(testInterfaceIpTFTPAddress, "read_only", "true"),
						),
					},
				},
			})

		})
	}
}

func testAccInterfaceIpTFTPConfig() string {
	return providerConfig + `

resource "routeros_ip_tftp" "test_file" {
	ip_addresses  = ["10.0.0.0/24]
	req_filename  = "file.txt"
	real_filename = "/file.txt"
	read_only     = true
}
`
}
