package routeros

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gnewbury1/terraform-provider-routeros/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testIPFirewallFilterAddress = "routeros_ip_firewall_filter.rule"

func TestAccIPFirewallFilterTest_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIPFirewallFilterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIPFirewallFilterConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPFirewallFilterExists(testIPFirewallFilterAddress),
					resource.TestCheckResourceAttr(testIPFirewallFilterAddress, "action", "accept"),
				),
			},
		},
	})
}

func testAccCheckIPFirewallFilterExists(name string) resource.TestCheckFunc {
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

func testAccIPFirewallFilterConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_ip_firewall_filter" "rule" {
	action 		= "accept"
	chain   	= "forward"
	src_address = "10.0.0.1"
	dst_address = "10.0.1.1"
	dst_port 	= "443"
	protocol 	= "tcp"
  }

`
}

func testAccCheckIPFirewallFilterDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "routeros_ip_firewall_filter" {
			continue
		}
		id := rs.Primary.ID
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/ip/firewall/filter/%s", c.HostURL, id), nil)
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/json")
		req.SetBasicAuth(c.Username, c.Password)

		res, err := c.HTTPClient.Do(req)
		if err != nil {
			return nil
		}
		if res.StatusCode != 404 {
			return fmt.Errorf("firewall filter %s has been found", id)
		}
		return nil
	}

	return nil
}
