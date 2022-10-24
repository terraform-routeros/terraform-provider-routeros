package routeros

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gnewbury1/terraform-provider-routeros/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testIPv6AddressAddress = "routeros_ipv6_address.test_ipv6"

func TestAccIPv6AddressTest_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIPv6AddressDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIPv6AddressConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPv6AddressExists(testIPv6AddressAddress),
					resource.TestCheckResourceAttr(testIPv6AddressAddress, "address", "fd55::1/64"),
				),
			},
		},
	})
}

func testAccCheckIPv6AddressExists(name string) resource.TestCheckFunc {
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

func testAccIPv6AddressConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_ipv6_address" "test_ipv6" {
	address = "fd55::1/64"
	interface = "ether1"
	disabled = true
  }

`
}

func testAccCheckIPv6AddressDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "routeros_ipv6_address" {
			continue
		}
		id := rs.Primary.ID
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/ipv6/address/%s", c.HostURL, id), nil)
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
			return fmt.Errorf("IPv6 Address %s has been found", id)
		}
		return nil
	}

	return nil
}
