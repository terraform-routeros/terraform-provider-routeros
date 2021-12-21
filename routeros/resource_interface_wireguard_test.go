package routeros

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gnewbury1/terraform-provider-routeros/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testInterfaceWireguardAddress = "routeros_interface_wireguard.test_wg_interface"

func TestAccInterfaceWireguardTest_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckInterfaceWireguardDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccInterfaceWireguardConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckInterfaceWireguardExists(testInterfaceWireguardAddress),
					resource.TestCheckResourceAttr(testInterfaceWireguardAddress, "listen_port", "13231"),
				),
			},
		},
	})
}

func testAccCheckInterfaceWireguardExists(name string) resource.TestCheckFunc {
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

func testAccInterfaceWireguardConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_interface_wireguard" "test_wg_interface" {
	name   		= "test_wg_interface"
	listen_port = "13231"
  }

`
}

func testAccCheckInterfaceWireguardDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "routeros_interface_wireguard" {
			continue
		}
		id := rs.Primary.ID
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/interface/wireguard/%s", c.HostURL, id), nil)
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
			return fmt.Errorf("wireguard interface %s has been found", id)
		}
		return nil
	}

	return nil
}
