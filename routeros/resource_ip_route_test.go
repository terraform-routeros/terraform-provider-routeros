package routeros

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gnewbury1/terraform-provider-routeros/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testIpRouteAddress = "routeros_ip_route.test_route"

func TestAccIpRouteTest_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIpRouteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIpRouteConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpRouteExists(testIpRouteAddress),
					resource.TestCheckResourceAttr(testIpRouteAddress, "distance", "1"),
					resource.TestCheckResourceAttr(testIpRouteAddress, "dst_address", "10.0.0.0/24"),
					resource.TestCheckResourceAttr(testIpRouteAddress, "gateway", "10.1.99.1"),
				),
			},
		},
	})
}

func testAccCheckIpRouteExists(name string) resource.TestCheckFunc {
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

func testAccIpRouteConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_ip_route" "test_route" {
	distance    = 1
	dst_address = "10.0.0.0/24"
	gateway		= "10.1.99.1"
  }

`
}

func testAccCheckIpRouteDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "routeros_ip_route" {
			continue
		}
		id := rs.Primary.ID
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/ip/route/%s", c.HostURL, id), nil)
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
			return fmt.Errorf("ip route %s has been found", id)
		}
		return nil
	}

	return nil
}
