package routeros

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gnewbury1/terraform-provider-routeros/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testListMemberAddress = "routeros_interface_list_member.test_list_member"
const testListMemberInterface = "ether1"

func TestAccInterfaceListMemberTest_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckInterfaceListMemberDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccInterfaceListMemberConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckInterfaceListMemberExists(testListMemberAddress),
					resource.TestCheckResourceAttr(testListMemberAddress, "interface", "ether1"),
				),
			},
		},
	})
}

func testAccCheckInterfaceListMemberExists(name string) resource.TestCheckFunc {
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

func testAccInterfaceListMemberConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_interface_list_member" "test_list_member" {
	interface      = "ether1"
	list           = "list"
}
`
}

func testAccCheckInterfaceListMemberDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "routeros_interface_list_member" {
			continue
		}
		id := rs.Primary.ID
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/interface/list_member/member/%s", c.HostURL, id), nil)
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
			return fmt.Errorf("list_member id %s has been found", id)
		}
		return nil
	}

	return nil
}
