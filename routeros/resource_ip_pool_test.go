package routeros

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gnewbury1/terraform-provider-routeros/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testIpPoolAddress = "routeros_ip_pool.test_pool"

func TestAccIpPoolTest_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIpPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIpPoolConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpPoolExists(testIpPoolAddress),
					resource.TestCheckResourceAttr(testIpPoolAddress, "name", "test_pool"),
				),
			},
		},
	})
}

func testAccCheckIpPoolExists(name string) resource.TestCheckFunc {
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

func testAccIpPoolConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_ip_pool" "test_pool" {
	name   = "test_pool"
	ranges = "10.0.0.100-10.0.0.200"
  }

`
}

func testAccCheckIpPoolDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "routeros_ip_pool" {
			continue
		}
		id := rs.Primary.ID
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/ip/pool/%s", c.HostURL, id), nil)
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
			return fmt.Errorf("dhcp client %s has been found", id)
		}
		return nil
	}

	return nil
}
