package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testFileMinVersion = "7.13"
const testFile = "routeros_file.test"

func TestAccFileTest_basic(t *testing.T) {
	if !testCheckMinVersion(t, testFileMinVersion) {
		t.Logf("Test skipped, the minimum required version is %v", testFileMinVersion)
		return
	}

	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/file", "routeros_file"),
				Steps: []resource.TestStep{
					{
						Config: testAccFileConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testFile),
							resource.TestCheckResourceAttr(testFile, "name", "test"),
						),
					},
				},
			})
		})
	}
}

func testAccFileConfig() string {
	return providerConfig + `

	resource "routeros_file" "test" {
		name     = "test"
		contents = "This is a test"
	}  
`
}
