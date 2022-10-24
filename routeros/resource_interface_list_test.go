package routeros

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gnewbury1/terraform-provider-routeros/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testListAddress = "routeros_interface_list.test_list"
const testListName = "List_TEST"

func TestAccInterfaceListTest_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckInterfaceListDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccInterfaceListConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckInterfaceListExists(testListAddress),
					resource.TestCheckResourceAttr(testListAddress, "name", testListName),
				),
			},
		},
	})
}

func testAccCheckInterfaceListExists(name string) resource.TestCheckFunc {
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

func testAccInterfaceListConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_interface_list" "test_list" {
	name      = "List_TEST"
}
`
}

func testAccCheckInterfaceListDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "routeros_interface_list" {
			continue
		}
		id := rs.Primary.ID
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/interface/list/%s", c.HostURL, id), nil)
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
			return fmt.Errorf("list id %s has been found", id)
		}
		return nil
	}

	return nil
}
