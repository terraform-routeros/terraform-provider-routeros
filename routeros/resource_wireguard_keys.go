package routeros

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/crypto/curve25519"
)

func ResourceWireguardKeys() *schema.Resource {
	return &schema.Resource{
		Description: "Creating key sets for WireGuard tunnels.",
		Schema: map[string]*schema.Schema{
			MetaId: {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Default:  int(Name),
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return true
				},
			},
			MetaResourcePath: {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "local",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return true
				},
			},
			"number": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Default:     1,
				Description: "The number of key sets.",
			},
			"keys": {
				Type:      schema.TypeList,
				Computed:  true,
				Sensitive: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"preshared": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Pre-shared secret key.",
						},
						"private": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Private WG key.",
						},
						"public": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Public WG Key.",
						},
					},
				},
			},
		},
		CreateContext: wgKeysCreate,
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return nil
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			d.SetId("")
			return nil
		},
	}
}

// https://github.com/WireGuard/wgctrl-go/blob/master/wgtypes/types.go
// KeyLen is the expected key length for a WireGuard key.
const KeyLen = 32 // wgh.KeyLen

// A Key is a public, private, or pre-shared secret key.  The Key constructor
// functions in this package can be used to create Keys suitable for each of
// these applications.
type Key [KeyLen]byte

// GenerateKey generates a Key suitable for use as a pre-shared secret key from
// a cryptographically safe source.
//
// The output Key should not be used as a private key; use GeneratePrivateKey
// instead.
func GenerateKey() (Key, error) {
	b := make([]byte, KeyLen)
	if _, err := rand.Read(b); err != nil {
		return Key{}, fmt.Errorf("failed to read random bytes: %v", err)
	}

	return NewKey(b)
}

// GeneratePrivateKey generates a Key suitable for use as a private key from a
// cryptographically safe source.
func GeneratePrivateKey() (Key, error) {
	key, err := GenerateKey()
	if err != nil {
		return Key{}, err
	}

	// Modify random bytes using algorithm described at:
	// https://cr.yp.to/ecdh.html.
	key[0] &= 248
	key[31] &= 127
	key[31] |= 64

	return key, nil
}

// NewKey creates a Key from an existing byte slice.  The byte slice must be
// exactly 32 bytes in length.
func NewKey(b []byte) (Key, error) {
	if len(b) != KeyLen {
		return Key{}, fmt.Errorf("incorrect key size: %d", len(b))
	}

	var k Key
	copy(k[:], b)

	return k, nil
}

// PublicKey computes a public key from the private key k.
//
// PublicKey should only be called when k is a private key.
func (k Key) PublicKey() Key {
	var (
		pub  [KeyLen]byte
		priv = [KeyLen]byte(k)
	)

	// ScalarBaseMult uses the correct base value per https://cr.yp.to/ecdh.html,
	// so no need to specify it.
	curve25519.ScalarBaseMult(&pub, &priv)

	return Key(pub)
}

// String returns the base64-encoded string representation of a Key.
//
// ParseKey can be used to produce a new Key from this string.
func (k Key) String() string {
	return base64.StdEncoding.EncodeToString(k[:])
}

func wgKeysCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var res []map[string]any

	for i := 0; i < d.Get("number").(int); i++ {
		keys := make(map[string]any)

		key, err := GeneratePrivateKey()
		if err != nil {
			return diag.FromErr(err)
		}
		keys["private"] = key.String()
		keys["public"] = key.PublicKey().String()

		key, err = GenerateKey()
		if err != nil {
			return diag.FromErr(err)
		}
		keys["preshared"] = key.String()

		res = append(res, keys)
	}

	d.SetId("wg_keys")

	return diag.FromErr(d.Set("keys", res))
}
