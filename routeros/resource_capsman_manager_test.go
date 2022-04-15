package routeros

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gnewbury1/terraform-provider-routeros/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testCapsManManagerAddress = "routeros_capsman_manager.test_manager"

func TestAccCapsManManagerTest_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCapsManManagerConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCapsManManagerExists(testCapsManManagerAddress),
					resource.TestCheckResourceAttr(testCapsManManagerAddress, "enabled", "true"),
				),
			},
		},
	})
}

func testAccCheckCapsManManagerExists(address string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[address]
		if !ok {
			return fmt.Errorf("not found: %s", address)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no id is set")
		}

		return nil
	}
}

func testAccCapsManManagerConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_capsman_manager" "test_manager" {
	enabled   = "true"
  }

`
}

func testAccCheckCapsManManagerDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "routeros_capsman_manager" {
			continue
		}
		id := rs.Primary.ID
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/caps-man/manager/%s", c.HostURL, id), nil)
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
			return fmt.Errorf("capsman manager id %s has been found", id)
		}
		return nil
	}

	return nil
}
