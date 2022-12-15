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
						Config: testAccResourceDnsConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckResourceDnsExists(testResourceDnsTask),
							resource.TestCheckResourceAttr(testResourceDnsTask, "allow_remote_requests", "true"),
						),
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

func testAccResourceDnsConfig() string {
	return `

provider "routeros" {
	insecure = true
}

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
}
`
}
