package routeros

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// providerConfigBulkRead enables the opt-in bulk_read_refresh flag so the
// provider under test exercises the cache path on every Read.
const providerConfigBulkRead = `
provider "routeros" {
	insecure          = true
	bulk_read_refresh = true
}
`

func testAccBulkReadDnsRecordsConfig(providerBlock string, count int) string {
	var b strings.Builder
	b.WriteString(providerBlock)
	for i := 0; i < count; i++ {
		fmt.Fprintf(&b, `
resource "routeros_dns_record" "bulk_%d" {
	address = "127.0.0.1"
	name    = "bulk-test-%d.example"
	type    = "A"
	ttl     = "1m"
}
`, i, i)
	}
	return b.String()
}

func TestAccDnsRecord_BulkReadRefresh_Roundtrip(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/dns/static", "routeros_dns_record"),
				Steps: []resource.TestStep{
					{
						Config: testAccBulkReadDnsRecordsConfig(providerConfigBulkRead, 5),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr("routeros_dns_record.bulk_0", "name", "bulk-test-0.example"),
							resource.TestCheckResourceAttr("routeros_dns_record.bulk_4", "name", "bulk-test-4.example"),
						),
					},
					// Second step with identical config: the refresh goes through the cache.
					// Any drift reported by Terraform here would be a correctness regression.
					{
						Config:             testAccBulkReadDnsRecordsConfig(providerConfigBulkRead, 5),
						PlanOnly:           true,
						ExpectNonEmptyPlan: false,
					},
				},
			})
		})
	}
}

func TestAccDnsRecord_BulkReadRefresh_DetectsExternalDelete(t *testing.T) {
	// Confirms the cache does not mask out-of-band deletions. The test
	// deletes a record directly via the REST API between steps and asserts
	// that the next plan surfaces the drift so Terraform recreates it.
	for _, name := range testNames {
		if !strings.Contains(name, "REST") {
			// The out-of-band delete helper uses the REST API.
			continue
		}
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/dns/static", "routeros_dns_record"),
				Steps: []resource.TestStep{
					{
						Config: testAccBulkReadDnsRecordsConfig(providerConfigBulkRead, 3),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr("routeros_dns_record.bulk_1", "name", "bulk-test-1.example"),
							deleteDnsRecordOutOfBand("routeros_dns_record.bulk_1"),
						),
					},
					{
						Config:             testAccBulkReadDnsRecordsConfig(providerConfigBulkRead, 3),
						PlanOnly:           true,
						ExpectNonEmptyPlan: true, // the out-of-band delete must produce a plan
					},
				},
			})
		})
	}
}

// deleteDnsRecordOutOfBand issues a DELETE directly against the REST API,
// bypassing Terraform. Used to simulate external drift.
func deleteDnsRecordOutOfBand(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("%s: resource not found in state", resourceName)
		}
		id := rs.Primary.ID
		if id == "" {
			return fmt.Errorf("%s: empty id", resourceName)
		}
		client, ok := testAccProvider.Meta().(*RestClient)
		if !ok {
			return fmt.Errorf("out-of-band delete only supports REST transport")
		}
		url := fmt.Sprintf("%s/rest/ip/dns/static/%s", client.HostURL, id)
		req, err := http.NewRequest(http.MethodDelete, url, nil)
		if err != nil {
			return err
		}
		req.SetBasicAuth(client.Username, client.Password)
		res, err := client.Do(req)
		if err != nil {
			return err
		}
		defer func() { _ = res.Body.Close() }()
		body, _ := io.ReadAll(res.Body)
		if res.StatusCode >= 400 {
			return fmt.Errorf("out-of-band DELETE %s: status=%d body=%s", url, res.StatusCode, body)
		}
		return nil
	}
}

