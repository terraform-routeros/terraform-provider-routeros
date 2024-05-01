package routeros

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

var testAccProvider *schema.Provider
var testAccProviderFactories map[string]func() (*schema.Provider, error)
var testNames = []string{"API", "REST"}

var reHost = regexp.MustCompile(`^(?:\S+://)?(\S+?)(?::\d+)*$`)
var reVersion = regexp.MustCompile(`\d+`)

var providerConfig = `
provider "routeros" {
	insecure = true
}
`

func init() {
	testAccProvider = Provider()
	testAccProviderFactories = map[string]func() (*schema.Provider, error){
		"routeros": func() (*schema.Provider, error) {
			return testAccProvider, nil
		},
	}
}

func testCheckMinVersion(t *testing.T, version string) bool {
	// version: 6.39.1
	var current, min uint64
	for pos, s := range reVersion.FindAllString(os.Getenv("ROS_VERSION"), -1) {
		if pos > 2 {
			t.Fatal("The version does not match the format x[.y[.z]]")
		}

		i, err := strconv.ParseUint(s, 10, 16)
		if err != nil {
			t.Error(err)
		}
		current += i << ((2 - pos) * 10)
	}

	for pos, s := range reVersion.FindAllString(version, -1) {
		if pos > 2 {
			t.Fatal("The version does not match the format x[.y[.z]]")
		}

		i, err := strconv.ParseUint(s, 10, 16)
		if err != nil {
			t.Error(err)
		}
		min += i << ((2 - pos) * 10)
	}

	return current >= min
}

func testCheckMaxVersion(t *testing.T, version string) bool {
	// version: 6.39.1
	var current, max uint64
	for pos, s := range reVersion.FindAllString(os.Getenv("ROS_VERSION"), -1) {
		if pos > 2 {
			t.Fatal("The version does not match the format x[.y[.z]]")
		}

		i, err := strconv.ParseUint(s, 10, 16)
		if err != nil {
			t.Error(err)
		}
		current += i << ((2 - pos) * 10)
	}

	for pos, s := range reVersion.FindAllString(version, -1) {
		if pos > 2 {
			t.Fatal("The version does not match the format x[.y[.z]]")
		}

		i, err := strconv.ParseUint(s, 10, 16)
		if err != nil {
			t.Error(err)
		}
		max += i << ((2 - pos) * 10)
	}

	return current <= max
}

func TestCheckMinVersion(t *testing.T) {
	originalVersion := os.Getenv("ROS_VERSION")
	defer func() {
		os.Setenv("ROS_VERSION", originalVersion)
	}()

	type args struct {
		current string
		min     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Positive #1", args{"7", "6.2.53"}, true},
		{"Positive #2", args{"7.1", "6.2.53"}, true},
		{"Positive #3", args{"7.1.35", "6.2.53"}, true},
		{"Positive #4", args{"7.1.35", "6.2"}, true},
		{"Positive #5", args{"7.1.35", "6"}, true},
		{"Positive #6", args{"7", "7"}, true},
		{"Positive #7", args{"7.1", "7.1"}, true},
		{"Positive #8", args{"7.1.53", "7.1.53"}, true},
		{"Negative #1", args{"6", "7.1.35"}, false},
		{"Negative #2", args{"6.2", "7.1.35"}, false},
		{"Negative #3", args{"6.2.53", "7.1.35"}, false},
		{"Negative #4", args{"6.2.53", "7.1"}, false},
		{"Negative #5", args{"6.2.53", "7"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := os.Setenv("ROS_VERSION", tt.args.current); err != nil {
				t.Error(err)
			}
			if got := testCheckMinVersion(t, tt.args.min); got != tt.want {
				t.Errorf("TestCheckMinVersion() diag got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func testSetTransportEnv(t *testing.T, testName string) {
	host := reHost.FindStringSubmatch(os.Getenv("ROS_HOSTURL"))
	switch {
	case strings.Contains(testName, "API"):
		if err := os.Setenv("ROS_HOSTURL", "apis://"+host[1]); err != nil {
			t.Error(err)
		}
	case strings.Contains(testName, "REST"):
		if err := os.Setenv("ROS_HOSTURL", "https://"+host[1]); err != nil {
			t.Error(err)
		}
	default:
		t.Fatal("Unsupported test name format. The test must have the suffix API or REST.")
	}

}

func testAccPreCheck(t *testing.T) {
	if os.Getenv("ROS_HOSTURL") == "" ||
		os.Getenv("ROS_USERNAME") == "" {
		t.Fatal("Environment variables (ROS_HOSTURL & ROS_USERNAME) must be set for testing")
	}

	for _, v := range Provider().ResourcesMap {
		checkResourceSchema(v.Schema, t)
	}
}

func checkResourceSchema(s map[string]*schema.Schema, t *testing.T) {
	f, ok := s[MetaResourcePath]
	if !ok {
		t.Fatalf("the schema does not contain field '%v'", MetaResourcePath)
	}
	if f.Default.(string) == "" {
		t.Fatalf("the field '%v', contains no data", MetaResourcePath)
	}
	f, ok = s[MetaId]
	if !ok {
		t.Fatalf("the schema does not contain field '%v'", MetaId)
	}
	if f.Default.(int) < 1 {
		t.Fatalf("the field '%v' is not defined", MetaId)
	}
}

func testCheckResourceDestroy(resourcePath, resourceType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		cApi, _ := testAccProvider.Meta().(*ApiClient)
		cRest, _ := testAccProvider.Meta().(*RestClient)
		var testTransport TransportType

		switch testAccProvider.Meta().(type) {
		case *ApiClient:
			testTransport = TransportAPI
		case *RestClient:
			testTransport = TransportREST
		default:
			panic("[testCheckResourceDestroy] wrong transport type")
		}

		for _, rs := range s.RootModule().Resources {
			if rs.Type != resourceType {
				continue
			}
			id := rs.Primary.ID
			idName := IdType(Provider().ResourcesMap[resourceType].Schema[MetaId].Default.(int)).String()

			switch testTransport {
			case TransportAPI:
				cmd := []string{resourcePath + "/print", "?" + idName + "=" + id}
				res, err := cApi.RunArgs(cmd)
				if err != nil {
					return nil
				}

				if len(res.Re) > 0 {
					return fmt.Errorf("resource %v %s has been found", resourceType, id)
				}
			case TransportREST:
				// Escaping spaces!
				req, err := http.NewRequest("GET",
					fmt.Sprintf("%s/rest%s?%s=%s", cRest.HostURL, resourcePath, idName, strings.Replace(id, " ", "%20", -1)), nil)
				if err != nil {
					return err
				}
				req.Header.Set("Content-Type", "application/json")
				req.SetBasicAuth(cRest.Username, cRest.Password)

				res, err := cRest.Do(req)

				if err != nil {
					return err
				}

				if res == nil {
					return fmt.Errorf("the response body is empty")
				}

				if buf, _ := io.ReadAll(res.Body); string(buf) != "[]" {
					return fmt.Errorf("resource %v %s has been found", resourceType, id)
				}
			}

			return nil
		}

		return nil
	}
}

// testCheckResourceExists queries the MikroTik API and retrieves the matching resource parameters
func testCheckResourceExists(name string, resourcePath string, resource *MikrotikItem) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("%s: resource not found in terraform state", name)
		}

		id := rs.Primary.ID
		if id == "" {
			return fmt.Errorf("%s: no id is set", name)
		}

		resourceType := strings.Split(name, ".")[0]
		resourceSchema, ok := testAccProvider.ResourcesMap[resourceType]
		if !ok {
			return fmt.Errorf("%s: schema for '%s' resource not found", name, resourceType)
		}
		idType := IdType(resourceSchema.Schema[MetaId].Default.(int))

		client := testAccProvider.Meta().(Client)
		resources, err := ReadItems(&ItemId{Type: idType, Value: id}, resourcePath, client)
		if err != nil {
			return err
		}

		if len(*resources) == 0 {
			return fmt.Errorf("%s: resource not found", name)
		}

		if resource != nil {
			*resource = (*resources)[0]
		}

		return nil
	}
}
