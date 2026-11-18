package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestSshKeyDiffSuppress(t *testing.T) {
	cases := []struct {
		name string
		old  string
		new  string
		want bool
	}{
		{
			name: "same key without comment (RouterOS 7.25+)",
			old:  "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIFSFXHv7eSqu8G9TPobMwVsEsaBRT6x6XDzPHWBsw9Ht",
			new:  "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIFSFXHv7eSqu8G9TPobMwVsEsaBRT6x6XDzPHWBsw9Ht oxidized@fog.k8s",
			want: true,
		},
		{
			name: "same key with different comment",
			old:  "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIFSFXHv7eSqu8G9TPobMwVsEsaBRT6x6XDzPHWBsw9Ht old@host",
			new:  "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIFSFXHv7eSqu8G9TPobMwVsEsaBRT6x6XDzPHWBsw9Ht new@host",
			want: true,
		},
		{
			name: "different key material",
			old:  "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIFSFXHv7eSqu8G9TPobMwVsEsaBRT6x6XDzPHWBsw9Ht",
			new:  "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIDifferentKeyMaterialHere oxidized@fog.k8s",
			want: false,
		},
		{
			name: "different key type",
			old:  "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQCy",
			new:  "ssh-ed25519 AAAAB3NzaC1yc2EAAAADAQABAAABgQCy",
			want: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := sshKeyDiffSuppress("key", tc.old, tc.new, nil); got != tc.want {
				t.Fatalf("sshKeyDiffSuppress() = %v, want %v", got, tc.want)
			}
		})
	}
}

const testUserAddress2 = "routeros_system_user_sshkeys.test"

func TestUserSshKeyTest_basic(t *testing.T) {
	// t.Parallel()
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/user/ssh-keys", "routeros_system_user_sshkeys"),
				Steps: []resource.TestStep{
					{
						Config: testUserSshKeyConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testUserAddress2),
							resource.TestCheckResourceAttr(testUserAddress2, "key_type", "rsa"),
						),
					},
				},
			})

		})
	}
}

func testUserSshKeyConfig() string {
	return providerConfig + `
resource "routeros_system_user" "test" {
  name     = "test-user-1"
  address  = "0.0.0.0/0"
  group    = "read"
  password = "secret"
  comment  = "Test User"
}

resource "routeros_system_user_sshkeys" "test" {
  depends_on  = [routeros_system_user.test]
  user        = "test-user-1"
  key         = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQCyJ1EvW98veNVzR3VamNgmu0xOd/JK9YNvP/pa4WC5eT90UbX4TN7dKEK/x2FCwnnG9u0FQhzG2qa/Cg8meUvlfydn6uxc0/WCeXTKSu6sT63noPO6m4fHY7gu3Zt+fOc/WYGch9sBeWjZlCS1mA2lajkWhM3J8TFWCFm2Zk4/S3s5mt6VLbwpQnH2LhE41+azzDEVhcR6i3FfdgOF/J+j2fYYHJsBEKoQA5zUac2zWmz7X4Rv/g11ZBRqdMpHSD58o5F9lBb13antu5GcEs5RXpXp08OyXuRV9qhFpDBC8DOMALSOgT3vnu8uJLgo8QIulERofj/cRXbLCsmvMbpioBuGFXWx3ha4Ntd6z07kUh2KVbaIQLd/629UHNvgIhoBLlREJ8E5vllsX+jh8hRITHcCiEwXcDO+gG3hvJt0+jm8S8SObE/IHk8VuwWdhIsSku5vd+wVlxm8VeJzjc0cjdIiytvsq8VpLudKEUiqR0f2tHcoq8H+xcJv3Ycu1i8="
  comment     = "Test User"
}
`
}
