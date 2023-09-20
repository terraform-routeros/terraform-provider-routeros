package routeros

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIPConnectionTracking = "routeros_firewall_connection_tracking.data"

func TestAccIPConnectionTrackingTest_basic(t *testing.T) {
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
						Config: testAccIPConnectionTrackingConfig(),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(testIPConnectionTracking, "active_ipv4", "true"),
							resource.TestCheckResourceAttr(testIPConnectionTracking, "active_ipv6", "true"),
							resource.TestCheckResourceAttr(testIPConnectionTracking, "enabled", "yes"),
							resource.TestCheckResourceAttr(testIPConnectionTracking, "generic_timeout", "10m"),
							resource.TestCheckResourceAttr(testIPConnectionTracking, "icmp_timeout", "10s"),
							resource.TestCheckResourceAttr(testIPConnectionTracking, "loose_tcp_tracking", "true"),
							resource.TestCheckResourceAttr(testIPConnectionTracking, "max_entries", "419840"),
							resource.TestCheckResourceAttr(testIPConnectionTracking, "tcp_close_timeout", "10s"),
							resource.TestCheckResourceAttr(testIPConnectionTracking, "tcp_close_wait_timeout", "10s"),
							resource.TestCheckResourceAttr(testIPConnectionTracking, "tcp_established_timeout", "1d"),
							resource.TestCheckResourceAttr(testIPConnectionTracking, "tcp_fin_wait_timeout", "10s"),
							resource.TestCheckResourceAttr(testIPConnectionTracking, "tcp_last_ack_timeout", "10s"),
							resource.TestCheckResourceAttr(testIPConnectionTracking, "tcp_max_retrans_timeout", "5m"),
							resource.TestCheckResourceAttr(testIPConnectionTracking, "tcp_syn_received_timeout", "5s"),
							resource.TestCheckResourceAttr(testIPConnectionTracking, "tcp_syn_sent_timeout", "5s"),
							resource.TestCheckResourceAttr(testIPConnectionTracking, "tcp_time_wait_timeout", "10s"),
							resource.TestCheckResourceAttr(testIPConnectionTracking, "tcp_unacked_timeout", "5m"),
							resource.TestCheckResourceAttrWith(testIPConnectionTracking, "total_entries", func(value string) error {
								nConn, err := strconv.Atoi(value)
								if err != nil {
									return fmt.Errorf("the total_entries was not a number %q", err)
								}
								if nConn <= 0 || nConn >= 100 {
									return errors.New("number of tcp connections (total_entries) does not seem correct")
								}
								return nil
							}),
							resource.TestCheckResourceAttr(testIPConnectionTracking, "udp_stream_timeout", "3m"),
							resource.TestCheckResourceAttr(testIPConnectionTracking, "udp_timeout", "10s"),
						),
					},
				},
			})
		})
	}
}

func testAccIPConnectionTrackingConfig() string {
	return providerConfig + `
resource "routeros_firewall_connection_tracking" "data" {
	
}

`
}
