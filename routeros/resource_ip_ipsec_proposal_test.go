package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpIpsecProposal = "routeros_ip_ipsec_proposal.test"

func TestAccIpIpsecProposalTest_basic(t *testing.T) {
	// t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/ipsec/proposal", "routeros_ip_ipsec_proposal"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpIpsecProposalConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpIpsecProposal),
							resource.TestCheckResourceAttr(testIpIpsecProposal, "name", "NordVPN"),
							resource.TestCheckResourceAttr(testIpIpsecProposal, "pfs_group", "none"),
							resource.TestCheckResourceAttr(testIpIpsecProposal, "lifetime", "45m10s"),
						),
					},
				},
			})

		})
	}
}

func testAccIpIpsecProposalConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_ip_ipsec_proposal" "test" {
  name      = "NordVPN"
  pfs_group = "none"
  lifetime  = "45m10s"
}
`, providerConfig)
}
