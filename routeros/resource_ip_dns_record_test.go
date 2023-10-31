package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

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
							// A
							testResourcePrimaryInstanceId("routeros_dns_record.test_dns_a"),
							resource.TestCheckResourceAttr("routeros_dns_record.test_dns_a", "name", "ipv4"),
							resource.TestCheckResourceAttr("routeros_dns_record.test_dns_a", "address", "127.0.0.1"),
							resource.TestCheckResourceAttr("routeros_dns_record.test_dns_a", "type", "A"),
							// A regexp
							testResourcePrimaryInstanceId("routeros_dns_record.test_dns_a_regexp"),
							resource.TestCheckResourceAttr("routeros_dns_record.test_dns_a", "address", "127.0.0.1"),
							resource.TestCheckResourceAttr("routeros_dns_record.test_dns_a_regexp", "regexp", "regexp"),
							resource.TestCheckResourceAttr("routeros_dns_record.test_dns_a_regexp", "type", "A"),
							// AAAA
							testResourcePrimaryInstanceId("routeros_dns_record.test_dns_aaaa"),
							resource.TestCheckResourceAttr("routeros_dns_record.test_dns_aaaa", "name", "ipv6"),
							resource.TestCheckResourceAttr("routeros_dns_record.test_dns_aaaa", "address", "ff00::1"),
							resource.TestCheckResourceAttr("routeros_dns_record.test_dns_aaaa", "type", "AAAA"),
							// CNAME
							testResourcePrimaryInstanceId("routeros_dns_record.test_dns_cname"),
							resource.TestCheckResourceAttr("routeros_dns_record.test_dns_cname", "cname", "ipv4"),
							resource.TestCheckResourceAttr("routeros_dns_record.test_dns_cname", "name", "cname"),
							resource.TestCheckResourceAttr("routeros_dns_record.test_dns_cname", "type", "CNAME"),
							// FWD
							testResourcePrimaryInstanceId("routeros_dns_record.test_dns_fwd"),
							resource.TestCheckResourceAttr("routeros_dns_record.test_dns_fwd", "name", "fwd"),
							resource.TestCheckResourceAttr("routeros_dns_record.test_dns_fwd", "forward_to", "127.0.0.1"),
							resource.TestCheckResourceAttr("routeros_dns_record.test_dns_fwd", "type", "FWD"),
							// MX
							testResourcePrimaryInstanceId("routeros_dns_record.test_dns_mx"),
							resource.TestCheckResourceAttr("routeros_dns_record.test_dns_mx", "name", "mx"),
							resource.TestCheckResourceAttr("routeros_dns_record.test_dns_mx", "mx_exchange", "127.0.0.1"),
							resource.TestCheckResourceAttr("routeros_dns_record.test_dns_mx", "mx_preference", "10"),
							resource.TestCheckResourceAttr("routeros_dns_record.test_dns_mx", "type", "MX"),
							// NS
							testResourcePrimaryInstanceId("routeros_dns_record.test_dns_ns"),
							resource.TestCheckResourceAttr("routeros_dns_record.test_dns_ns", "name", "ns"),
							resource.TestCheckResourceAttr("routeros_dns_record.test_dns_ns", "ns", "127.0.0.1"),
							resource.TestCheckResourceAttr("routeros_dns_record.test_dns_ns", "type", "NS"),
							// NXDOMAIN
							testResourcePrimaryInstanceId("routeros_dns_record.test_dns_nxdomain"),
							resource.TestCheckResourceAttr("routeros_dns_record.test_dns_nxdomain", "name", "nxdomain"),
							resource.TestCheckResourceAttr("routeros_dns_record.test_dns_nxdomain", "type", "NXDOMAIN"),
							// SRV
							testResourcePrimaryInstanceId("routeros_dns_record.test_dns_srv"),
							resource.TestCheckResourceAttr("routeros_dns_record.test_dns_srv", "name", "srv"),
							resource.TestCheckResourceAttr("routeros_dns_record.test_dns_srv", "srv_port", "8080"),
							resource.TestCheckResourceAttr("routeros_dns_record.test_dns_srv", "srv_priority", "10"),
							resource.TestCheckResourceAttr("routeros_dns_record.test_dns_srv", "srv_target", "127.0.0.1"),
							resource.TestCheckResourceAttr("routeros_dns_record.test_dns_srv", "srv_weight", "100"),
							resource.TestCheckResourceAttr("routeros_dns_record.test_dns_srv", "type", "SRV"),
							// TXT
							testResourcePrimaryInstanceId("routeros_dns_record.test_dns_txt"),
							resource.TestCheckResourceAttr("routeros_dns_record.test_dns_txt", "name", "_acme-challenge.yourwebsite.com"),
							resource.TestCheckResourceAttr("routeros_dns_record.test_dns_txt", "text", "dW6MrI3nBy3eJgYWH3QAg1Cwk_TvjFESOuKo+mp6nm1"),
							resource.TestCheckResourceAttr("routeros_dns_record.test_dns_txt", "type", "TXT"),
						),
					},
				},
			})

		})
	}
}

func testAccIpDnsRecordConfig() string {
	return providerConfig + `

resource "routeros_dns_record" "test_dns_a" {
	address         = "127.0.0.1"
	address_list    = "subdomain"
	disabled        = false
	match_subdomain = true
	name            = "ipv4"
	ttl             = "8m"
	type            = "A"
}

resource "routeros_dns_record" "test_dns_a_regexp" {
	address         = "127.0.0.1"
	disabled        = true
	regexp          = "regexp"
	type            = "A"
}

resource "routeros_dns_record" "test_dns_aaaa" {
	address         = "ff00::1"
	address_list    = "subdomain"
	disabled        = false
	match_subdomain = true
	name            = "ipv6"
	ttl             = "8m"
	type            = "AAAA"
}
  
resource "routeros_dns_record" "test_dns_cname" {
	address_list    = "subdomain"
	cname           = "ipv4"
	disabled        = false
	match_subdomain = true
	name            = "cname"
	ttl             = "8m"
	type            = "CNAME"
}
  
resource "routeros_dns_record" "test_dns_fwd" {
	address_list    = "subdomain"
	disabled        = false
	forward_to      = "127.0.0.1"
	match_subdomain = true
	name            = "fwd"
	ttl             = "8m"
	type            = "FWD"
}
  
resource "routeros_dns_record" "test_dns_mx" {
	address_list    = "subdomain"
	disabled        = false
	match_subdomain = true
	mx_exchange     = "127.0.0.1"
	mx_preference   = 10
	name            = "mx"
	ttl             = "8m"
	type            = "MX"
}
  
resource "routeros_dns_record" "test_dns_ns" {
	address_list    = "subdomain"
	disabled        = false
	match_subdomain = true
	name            = "ns"
	ns              = "127.0.0.1"
	ttl             = "8m"
	type            = "NS"
}
  
resource "routeros_dns_record" "test_dns_nxdomain" {
	address_list    = "subdomain"
	disabled        = false
	match_subdomain = true
	name            = "nxdomain"
	ttl             = "8m"
	type            = "NXDOMAIN"
}
  
resource "routeros_dns_record" "test_dns_srv" {
	address_list    = "subdomain"
	disabled        = false
	match_subdomain = true
	name            = "srv"
	srv_port        = 8080
	srv_priority    = 10
	srv_target      = "127.0.0.1"
	srv_weight      = 100
	ttl             = "8m"
	type            = "SRV"
}

resource "routeros_dns_record" "test_dns_txt" {
	address_list    = "subdomain"
	disabled        = false
	match_subdomain = true
	name            = "_acme-challenge.yourwebsite.com"
	text            = "dW6MrI3nBy3eJgYWH3QAg1Cwk_TvjFESOuKo+mp6nm1"
	ttl             = "8m"
	type            = "TXT"
}
`
}
