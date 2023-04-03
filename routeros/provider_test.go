package routeros

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
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
