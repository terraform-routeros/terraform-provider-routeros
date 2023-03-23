package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testResourceDnsTask = "routeros_dns.test"

func TestAccResourceDnsTest_basic(t *testing.T) {
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
						Config: testAccResourceDnsConfig(0),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckResourceDnsExists(testResourceDnsTask),
							resource.TestCheckResourceAttr(testResourceDnsTask, "allow_remote_requests", "true"),
						),
					},
					{
						Config: testAccResourceDnsConfig(1),
					},
				},
			})

		})
	}
}

func testAccCheckResourceDnsExists(name string) resource.TestCheckFunc {
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

func testAccResourceDnsConfig(n int) string {
	provider := `
provider "routeros" {
	insecure = true
}

`

	tests := []string{`
resource "routeros_dns" "test" {
	allow_remote_requests = true
	cache_max_ttl = "3d"
	cache_size = 4096
	max_concurrent_queries = 200
	max_concurrent_tcp_sessions = 40
	max_udp_packet_size = 8192
	query_server_timeout = "500ms"
	query_total_timeout = "15"
	servers = "1.1.1.1"
	use_doh_server = "https://cloudflare-dns.com/dns-query"
	verify_doh_cert = true
}`,
		`
resource "routeros_ip_dns" "dns-server" {
	allow_remote_requests = true
	servers = "2606:4700:4700::1112,1.1.1.2,2606:4700:4700::1002,1.0.0.2"
}`,
	}
	return provider + tests[n]
}
