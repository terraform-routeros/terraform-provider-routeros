package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testToolNetwatch = "routeros_tool_netwatch.test"

func TestAccToolNetwatchTest_basic(t *testing.T) {
	// t.Parallel()
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
						Config: testAccToolNetwatchConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testToolNetwatch),
							resource.TestCheckResourceAttr(testToolNetwatch, "name", "watch-google-pdns"),
							resource.TestCheckResourceAttr(testToolNetwatch, "host", "8.8.8.8"),
							resource.TestCheckResourceAttr(testToolNetwatch, "interval", "30s"),
							resource.TestCheckResourceAttr(testToolNetwatch, "up_script", ":log info \"Ping to 8.8.8.8 successful\""),
						),
					},
					{
						Config: testAccToolNetwatchConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testToolNetwatch),
							resource.TestCheckResourceAttr(testToolNetwatch, "name", "watch-google-pdns"),
							resource.TestCheckResourceAttr(testToolNetwatch, "host", "8.8.8.8"),
							resource.TestCheckResourceAttr(testToolNetwatch, "interval", "30s"),
							resource.TestCheckResourceAttr(testToolNetwatch, "up_script", ":log info \"Ping to 8.8.8.8 successful\""),
						),
					},
				},
			})

		})
	}
}

func testAccToolNetwatchConfig() string {
	return fmt.Sprintf(`%v
resource "routeros_tool_netwatch" "test" {
  name      = "watch-google-pdns"
  type      = "icmp"
  host      = "8.8.8.8"
  interval  = "30s"
  up_script = ":log info \"Ping to 8.8.8.8 successful\""
  thr_max   = "400ms"
}
`, providerConfig)
}

// dns-server and record-type are unknown parameters on 7.15
const testToolNetwatchDnsMinVersion = "7.16"

func TestAccToolNetwatchTest_dns(t *testing.T) {
	if !testCheckMinVersion(t, testToolNetwatchDnsMinVersion) {
		t.Logf("Test skipped, the minimum required version is %v", testToolNetwatchDnsMinVersion)
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
						Config: testAccToolNetwatchDnsConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testToolNetwatch),
							resource.TestCheckResourceAttr(testToolNetwatch, "type", "dns"),
							resource.TestCheckResourceAttr(testToolNetwatch, "host", "google.com"),
							resource.TestCheckResourceAttr(testToolNetwatch, "dns_server", "8.8.8.8"),
							resource.TestCheckResourceAttr(testToolNetwatch, "record_type", "A"),
						),
					},
				},
			})
		})
	}
}

func TestAccToolNetwatchTest_httpsGet(t *testing.T) {
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
						Config: testAccToolNetwatchHttpsConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testToolNetwatch),
							resource.TestCheckResourceAttr(testToolNetwatch, "type", "https-get"),
							resource.TestCheckResourceAttr(testToolNetwatch, "host", "1.1.1.1"),
							resource.TestCheckResourceAttr(testToolNetwatch, "port", "443"),
						),
					},
				},
			})
		})
	}
}

func testAccToolNetwatchDnsConfig() string {
	return fmt.Sprintf(`%v
resource "routeros_tool_netwatch" "test" {
  name        = "watch-google-dns"
  type        = "dns"
  host        = "google.com"
  dns_server  = "8.8.8.8"
  record_type = "A"
  interval    = "30s"
}
`, providerConfig)
}

func testAccToolNetwatchHttpsConfig() string {
	return fmt.Sprintf(`%v
resource "routeros_tool_netwatch" "test" {
  name     = "watch-cloudflare-https"
  type     = "https-get"
  host     = "1.1.1.1"
  port     = 443
  interval = "30s"
}
`, providerConfig)
}
