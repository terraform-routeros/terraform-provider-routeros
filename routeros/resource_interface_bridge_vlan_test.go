package routeros

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gnewbury1/terraform-provider-routeros/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testInterfaceBridgeVlanAddress = "routeros_interface_bridge_vlan.test_vlan"

func TestAccInterfaceBridgeVlanTest_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckInterfaceBridgeVlanDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccInterfaceBridgeVlanConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckInterfaceBridgeVlanExists(testInterfaceBridgeVlanAddress),
					resource.TestCheckResourceAttr(testInterfaceBridgeVlanAddress, "bridge", "bridge"),
				),
			},
		},
	})
}

func testAccCheckInterfaceBridgeVlanExists(name string) resource.TestCheckFunc {
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

func testAccInterfaceBridgeVlanConfig() string {
	return `

provider "routeros" {
	insecure = true
}

resource "routeros_interface_bridge_vlan" "test_vlan" {
	bridge   = "bridge"
	untagged = ["sfp-sfpplus1"]
	tagged	 = ["bridge","ether1"]
	vlan_id  = 200
  }

`
}

func testAccCheckInterfaceBridgeVlanDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "routeros_ip_route" {
			continue
		}
		id := rs.Primary.ID
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/interface/bridge/vlan/%s", c.HostURL, id), nil)
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
