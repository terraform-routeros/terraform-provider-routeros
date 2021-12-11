package routeros

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"routeros": testAccProvider,
	}
}

func testAccPreCheck(t *testing.T) {
	if os.Getenv("ROS_HOSTURL") == "" ||
		os.Getenv("ROS_USERNAME") == "" ||
		os.Getenv("ROS_PASSWORD") == "" {
		t.Fatal("Environment variables must be set for testing")
	}
}
