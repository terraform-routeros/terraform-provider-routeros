package routeros

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func testResourcePrimaryInstanceId(name string) resource.TestCheckFunc {
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

// testCheckMikrotikItemAttr ensures that an attribute in MikrotikItem parameters
// with the given key equals a specific value
func testCheckMikrotikItemAttr(name string, params *MikrotikItem, key string, value string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if params == nil {
			return fmt.Errorf("%s: mikrotik params is nil", name)
		}

		v, ok := (*params)[key]
		if !ok {
			return fmt.Errorf("%s: attribute '%s' not found", name, key)
		}

		if v != value {
			return fmt.Errorf("%s: attribute '%s' expected %s, got %s", name, key, value, v)
		}

		return nil
	}
}

// testCheckMikrotikItemAttrSet ensures any value exists in MikrotikItem parameters
// for the given key
func testCheckMikrotikItemAttrSet(name string, params *MikrotikItem, key string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if params == nil {
			return fmt.Errorf("%s: mikrotik params is nil", name)
		}

		_, ok := (*params)[key]
		if !ok {
			return fmt.Errorf("%s: attribute '%s' expected to be set", name, key)
		}

		return nil
	}
}

// testCheckNoMikrotikItemAttr ensures no value exists in MikrotikItem parameters
// for the given key
func testCheckNoMikrotikItemAttr(name string, params *MikrotikItem, key string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if params == nil {
			return fmt.Errorf("%s: mikrotik params is nil", name)
		}

		_, ok := (*params)[key]
		if ok {
			return fmt.Errorf("%s: attribute '%s' found when not expected", name, key)
		}

		return nil
	}
}
