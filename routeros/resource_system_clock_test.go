package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testSystemClockTask = "routeros_system_clock.set"

func TestAccSystemClockTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				Steps:             makeSteps(name),
			})

		})
	}
}

func makeSteps(name string) (res []resource.TestStep) {
	params := map[string]map[string]string{
		"API": {
			"date":           `2024-05-15`,
			"time":           `17:58:11`,
			"time_zone_name": `EST`,
		},
		"REST": {
			"date":           `2024-05-17`,
			"time":           `18:58:11`,
			"time_zone_name": `UTC`,
		},
	}

	for k, v := range params[name] {
		res = append(res, resource.TestStep{
			Config: fmt.Sprintf(`%v
			resource "routeros_system_clock" "set" {
				%v = "%v"
			}`, providerConfig, k, v),
			Check: resource.ComposeTestCheckFunc(
				testResourcePrimaryInstanceId(testSystemClockTask),
				resource.TestCheckResourceAttr(testSystemClockTask, k, v),
			),
		})

	}
	return
}
