package routeros

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gnewbury1/terraform-provider-routeros/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testCapsManConfigurationAddress = "routeros_capsman_configuration.test_configuration"

func TestAccCapsManConfigurationTest_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCapsManConfigurationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCapsManConfigurationConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCapsManConfigurationExists(testCapsManConfigurationAddress),
					resource.TestCheckResourceAttr(testCapsManConfigurationAddress, "name", "test_configuration"),
				),
			},
		},
	})
}

func testAccCheckCapsManConfigurationExists(name string) resource.TestCheckFunc {
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

func testAccCapsManConfigurationConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_capsman_configuration" "test_configuration" {
	name   = "test_configuration"
  }

`
}

func testAccCheckCapsManConfigurationDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "routeros_capsman_configuration" {
			continue
		}
		id := rs.Primary.ID
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/caps-man/configuration/%s", c.HostURL, id), nil)
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
			return fmt.Errorf("capsman configuration id %s has been found", id)
		}
		return nil
	}

	return nil
}
