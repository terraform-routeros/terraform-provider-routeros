package routeros

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/crypto/ssh"
)

// Internal RouterOS IDs can refer to a different object after an out-of-band
// restore or recreation. Check the identity before refreshing state or using
// an ID for a write. Never adopt the different object or guess a replacement ID.
func validateUserResourceIdentity(path string, items []MikrotikItem, d *schema.ResourceData) error {
	if path != "/user" && path != "/user/ssh-keys" {
		return nil
	}
	if len(items) == 0 {
		// Preserve the usual missing-resource handling in the caller.
		return nil
	}
	if len(items) != 1 {
		return userResourceIdentityError(path, d.Id(), "the ID returned multiple objects")
	}
	item := items[0]
	field := "name"
	if path == "/user/ssh-keys" {
		field = "user"
	}
	expected := priorIdentityString(d, field)
	// An initial import has no prior name, owner or key. Its first read
	// establishes the identity; subsequent reads must preserve it.
	if expected != "" && item[field] != expected {
		return userResourceIdentityError(path, d.Id(), "the "+field+" does not match the saved identity")
	}
	if path == "/user" {
		return nil
	}
	expectedFingerprint := priorIdentityString(d, "fingerprint")
	actualFingerprint := item["fingerprint"]
	if expectedFingerprint != "" && normalizeUserKeyFingerprint(expectedFingerprint) == normalizeUserKeyFingerprint(actualFingerprint) {
		return nil
	}
	key := priorIdentityString(d, "key")
	if key != "" {
		publicKey, _, _, _, err := ssh.ParseAuthorizedKey([]byte(key))
		if err != nil {
			return userResourceIdentityError(path, d.Id(), "the saved public key cannot be parsed to verify its fingerprint")
		}
		// RouterOS versions may expose SHA256 or legacy MD5 fingerprints.
		for _, candidate := range []string{ssh.FingerprintSHA256(publicKey), ssh.FingerprintLegacyMD5(publicKey)} {
			if normalizeUserKeyFingerprint(candidate) == normalizeUserKeyFingerprint(actualFingerprint) {
				return nil
			}
		}
		return userResourceIdentityError(path, d.Id(), "the SSH key fingerprint does not match the saved identity")
	}
	if expectedFingerprint != "" {
		return userResourceIdentityError(path, d.Id(), "the SSH key fingerprint does not match the saved identity")
	}
	return nil
}

// Updates may intentionally rename a user or reassign/replace a key. Validate
// the old identity, not the planned new value, before performing that change.
func priorIdentityString(d *schema.ResourceData, field string) string {
	old, _ := d.GetChange(field)
	if value, ok := old.(string); ok && value != "" {
		return value
	}
	value, _ := d.Get(field).(string)
	return value
}

func normalizeUserKeyFingerprint(value string) string {
	value = strings.TrimPrefix(value, "SHA256:")
	value = strings.TrimPrefix(value, "MD5:")
	return strings.TrimRight(value, "=")
}

func userResourceIdentityError(path, id, reason string) error {
	return fmt.Errorf("resource identity mismatch for %s at ID %q: %s; refusing to adopt or modify a different object. Verify the object on the router and re-import the correct resource before retrying", path, id, reason)
}
