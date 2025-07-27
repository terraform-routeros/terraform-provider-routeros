package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testDiskSettingsTaskMinVersion = "7.17"
const testDiskSettingsTask = "routeros_disk_settings.test"

func TestAccDiskSettingsTest_basic(t *testing.T) {
	if !testCheckMinVersion(t, testDiskSettingsTaskMinVersion) {
		t.Logf("Test skipped, the minimum required version is %v", testDiskSettingsTaskMinVersion)
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
				Steps: []resource.TestStep{
					{
						Config: testAccDiskSettingsConfig(false, "guest", false, "lo", "[slot]"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testDiskSettingsTask),
							resource.TestCheckResourceAttr(testDiskSettingsTask, "auto_smb_sharing", "false"),
							resource.TestCheckResourceAttr(testDiskSettingsTask, "auto_smb_user", "guest"),
							resource.TestCheckResourceAttr(testDiskSettingsTask, "auto_media_sharing", "false"),
							resource.TestCheckResourceAttr(testDiskSettingsTask, "auto_media_interface", "lo"),
							resource.TestCheckResourceAttr(testDiskSettingsTask, "default_mount_point_template", "[slot]"),
						),
					},
				},
			})

		})
	}
}

func testAccDiskSettingsConfig(autoSmbSharing bool, autoSmbUser string, autoMediaSharing bool, autoMediaInterface string, defaultMountPointTemplate string) string {
	return fmt.Sprintf(`%v
resource "routeros_disk_settings" "test" {
  auto_smb_sharing             = %v
  auto_smb_user                = "%v"
  auto_media_sharing           = %v
  auto_media_interface         = "%v"
  default_mount_point_template = "%v"
}
`, providerConfig, autoSmbSharing, autoSmbUser, autoMediaSharing, autoMediaInterface, defaultMountPointTemplate)
}
