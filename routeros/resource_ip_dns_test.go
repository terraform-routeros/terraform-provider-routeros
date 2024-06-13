package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
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
							testResourcePrimaryInstanceId(testResourceDnsTask),
							resource.TestCheckResourceAttr(testResourceDnsTask, "allow_remote_requests", "true"),
						),
					},
					{
						Config: testAccResourceDnsConfig(),
					},
				},
			})

		})
	}
}

func testAccResourceDnsConfig() string {
	return providerConfig + `
resource "routeros_dns" "test" {
	allow_remote_requests = true
	cache_max_ttl = "3d"
	cache_size = 4096
	max_concurrent_queries = 200
	max_concurrent_tcp_sessions = 40
	max_udp_packet_size = 8192
	query_server_timeout = "500ms"
	query_total_timeout = "15"
	servers = [
		"2606:4700:4700::1112",
		"1.1.1.2",
		"2606:4700:4700::1002",
		"1.0.0.2",
	]
	use_doh_server = "https://cloudflare-dns.com/dns-query"
	verify_doh_cert = true
}`
}
