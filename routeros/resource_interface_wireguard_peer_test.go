package routeros

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfaceWireguardPeerMinVersion = "7.12"
const testInterfaceWireguardPeerAddress = "routeros_interface_wireguard_peer.wg_peer"

func TestAccInterfaceWireguardPeerTest_basic(t *testing.T) {
	if !testCheckMinVersion(t, testInterfaceWireguardPeerMinVersion) {
		t.Logf("Test skipped, the minimum required version is %v", testInterfaceWireguardPeerMinVersion)
		return
	}

	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/wireguard/peers", "routeros_interface_wireguard_peer"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceWireguardPeerConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceWireguardPeerAddress),
							resource.TestCheckResourceAttr(testInterfaceWireguardPeerAddress, "interface", "wg1"),
						),
					},
				},
			})

		})
	}
}

func testAccInterfaceWireguardPeerConfig() string {
	return providerConfig + `

resource "routeros_interface_wireguard_peer" "wg_peer" {
	allowed_address  = ["1.2.3.0/30"]
	interface        = "wg1"
	public_key       = "QxC+CTcrDdU5+ny0+2ChUH3NegTrwoVCv53TllI5T0I="
	client_keepalive = "85s"
  }

`
}

func Test_resourceInterfaceWireguardPeerAllowedIPRegexp(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		args
		want bool
	}{
		{
			"empty",
			args{""},
			true,
		}, {
			"10.0.0.1",
			args{"10.0.0.1"},
			true,
		},
		{
			"10.0.0.1/32",
			args{"10.0.0.1/32"},
			true,
		},
		{
			"10.0.0.1/33",
			args{"10.0.0.1/33"},
			false,
		},
	}

	re := regexp.MustCompile(`^$|^(\b(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(/([1-2][0-9]|3[0-2]))?)$`)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := re.MatchString(tt.args.ip); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_resourceInterfaceWireguardPeerAllowedIPRegexp() = %v, want %v", got, tt.want)
			}
		})
	}
}
