package routeros

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"golang.org/x/crypto/ssh"
)

// The fake exercises the same CRUD lookup and state refresh used by both
// transports, without connecting to a router or using real credentials.
type userIdentityTestClient struct {
	transport TransportType
	items     []MikrotikItem
	writes    int
	readErr   error
	lastItem  MikrotikItem
}

func (c *userIdentityTestClient) GetExtraParams() *ExtraParams { return &ExtraParams{} }
func (c *userIdentityTestClient) GetTransport() TransportType  { return c.transport }
func (c *userIdentityTestClient) SendRequest(method crudMethod, _ *URL, item MikrotikItem, result interface{}) error {
	if method == crudRead {
		if c.readErr != nil {
			return c.readErr
		}
		*result.(*[]MikrotikItem) = c.items
		return nil
	}
	c.writes++
	c.lastItem = item
	if method == crudUpdate {
		updated := MikrotikItem{".id": "*1"}
		for key, value := range c.items[0] {
			updated[key] = value
		}
		for key, value := range item {
			updated[key] = value
		}
		*result.(*MikrotikItem) = updated
	}
	return nil
}

func userIdentityTestKey(t *testing.T) (string, ssh.PublicKey) {
	t.Helper()
	public, _, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	key, err := ssh.NewPublicKey(public)
	if err != nil {
		t.Fatal(err)
	}
	return strings.TrimSpace(string(ssh.MarshalAuthorizedKey(key))), key
}

func userIdentityTestVersion(t *testing.T) {
	t.Helper()
	previous := RouterOSVersion
	RouterOSVersion = "7.16"
	t.Cleanup(func() { RouterOSVersion = previous })
}

func userIdentityTestData(t *testing.T, resource *schema.Resource, raw map[string]interface{}) *schema.ResourceData {
	t.Helper()
	data := schema.TestResourceDataRaw(t, resource.Schema, raw)
	data.SetId("*1")
	state := data.State()
	config := make(map[string]cty.Value)
	for name, field := range resource.Schema {
		switch field.Type {
		case schema.TypeString:
			config[name] = cty.NullVal(cty.String)
			if value, ok := raw[name].(string); ok {
				config[name] = cty.StringVal(value)
			}
		case schema.TypeBool:
			config[name] = cty.NullVal(cty.Bool)
		case schema.TypeInt:
			config[name] = cty.NullVal(cty.Number)
		default:
			t.Fatalf("unsupported test schema type for %s", name)
		}
	}
	state.RawConfig = cty.ObjectVal(config)
	return resource.Data(state)
}

func TestUserResourceIdentityRejectsReusedIDs(t *testing.T) {
	userIdentityTestVersion(t)
	key, public := userIdentityTestKey(t)
	_, otherPublic := userIdentityTestKey(t)
	tests := []struct {
		name     string
		resource func() *schema.Resource
		state    map[string]interface{}
		actual   MikrotikItem
	}{
		{
			"user name", ResourceUser,
			map[string]interface{}{"name": "alice", "group": "read", "password": "test-only-secret"},
			MikrotikItem{".id": "*1", "name": "bob", "group": "full"},
		},
		{
			"SSH key owner", ResourceUserSshKeys,
			map[string]interface{}{"user": "alice", "key": key, "fingerprint": ssh.FingerprintSHA256(public)},
			MikrotikItem{".id": "*1", "user": "bob", "fingerprint": ssh.FingerprintSHA256(public)},
		},
		{
			"SSH key fingerprint", ResourceUserSshKeys,
			map[string]interface{}{"user": "alice", "key": key, "fingerprint": ssh.FingerprintSHA256(public)},
			MikrotikItem{".id": "*1", "user": "alice", "fingerprint": ssh.FingerprintSHA256(otherPublic)},
		},
	}
	for _, transport := range []TransportType{TransportAPI, TransportREST} {
		for _, test := range tests {
			for _, operation := range []string{"read", "update", "delete"} {
				t.Run(fmt.Sprintf("%v/%s/%s", transport, test.name, operation), func(t *testing.T) {
					resource := test.resource()
					data := userIdentityTestData(t, resource, test.state)
					data.SetId("*1")
					before := data.State()
					client := &userIdentityTestClient{transport: transport, items: []MikrotikItem{test.actual}}
					var diagnostics diag.Diagnostics
					switch operation {
					case "read":
						diagnostics = resource.ReadContext(context.Background(), data, client)
					case "update":
						diagnostics = resource.UpdateContext(context.Background(), data, client)
					case "delete":
						diagnostics = resource.DeleteContext(context.Background(), data, client)
					}
					if !diagnostics.HasError() || !strings.Contains(diagnostics[0].Summary, "resource identity mismatch") {
						t.Fatalf("expected identity error, got %v", diagnostics)
					}
					if strings.Contains(diagnostics[0].Summary, "test-only-secret") || strings.Contains(diagnostics[0].Summary, key) {
						t.Fatal("diagnostic exposed a password or key")
					}
					if client.writes != 0 {
						t.Fatalf("unexpected writes: %d", client.writes)
					}
					if !reflect.DeepEqual(before, data.State()) {
						t.Fatal("mismatched object changed the saved resource state")
					}
				})
			}
		}
	}
}

func TestUserResourceIdentityMatchingCRUD(t *testing.T) {
	userIdentityTestVersion(t)
	key, public := userIdentityTestKey(t)
	for _, transport := range []TransportType{TransportAPI, TransportREST} {
		for _, resource := range []*schema.Resource{ResourceUser(), ResourceUserSshKeys()} {
			for _, operation := range []string{"read", "update", "delete"} {
				path := GetMetadata(resource.Schema).Path
				t.Run(fmt.Sprintf("%v/%s/%s", transport, path, operation), func(t *testing.T) {
					raw := map[string]interface{}{"name": "alice", "group": "read", "password": "test-only-secret"}
					actual := MikrotikItem{".id": "*1", "name": "alice", "group": "read"}
					if path == "/user/ssh-keys" {
						raw = map[string]interface{}{"user": "alice", "key": key, "fingerprint": ssh.FingerprintSHA256(public)}
						actual = MikrotikItem{".id": "*1", "user": "alice", "fingerprint": ssh.FingerprintSHA256(public)}
					}
					data := userIdentityTestData(t, resource, raw)
					data.SetId("*1")
					client := &userIdentityTestClient{transport: transport, items: []MikrotikItem{actual}}
					var diagnostics diag.Diagnostics
					switch operation {
					case "read":
						diagnostics = resource.ReadContext(context.Background(), data, client)
					case "update":
						diagnostics = resource.UpdateContext(context.Background(), data, client)
					case "delete":
						diagnostics = resource.DeleteContext(context.Background(), data, client)
					}
					if diagnostics.HasError() {
						t.Fatal(diagnostics)
					}
					if operation != "read" && client.writes != 1 {
						t.Fatalf("expected one write, got %d", client.writes)
					}
					if operation == "update" && path == "/user" && client.lastItem["password"] != "test-only-secret" {
						t.Fatal("identity guard changed existing password update behavior")
					}
				})
			}
		}
	}
}

func TestUserResourceIdentityMissingAndImport(t *testing.T) {
	userIdentityTestVersion(t)
	for _, resource := range []*schema.Resource{ResourceUser(), ResourceUserSshKeys()} {
		path := GetMetadata(resource.Schema).Path
		for _, operation := range []string{"read", "delete"} {
			t.Run(path+"/missing/"+operation, func(t *testing.T) {
				data := userIdentityTestData(t, resource, nil)
				data.SetId("*1")
				client := &userIdentityTestClient{transport: TransportREST}
				var diagnostics diag.Diagnostics
				if operation == "read" {
					diagnostics = resource.ReadContext(context.Background(), data, client)
				} else {
					diagnostics = resource.DeleteContext(context.Background(), data, client)
				}
				if diagnostics.HasError() || data.Id() != "" || client.writes != 0 {
					t.Fatalf("missing resource not handled normally: %v", diagnostics)
				}
			})
		}
		t.Run(path+"/import", func(t *testing.T) {
			data := userIdentityTestData(t, resource, nil)
			data.SetId("*1")
			actual := MikrotikItem{".id": "*1", "name": "alice", "group": "read"}
			if path == "/user/ssh-keys" {
				actual = MikrotikItem{".id": "*1", "user": "alice", "fingerprint": "SHA256:imported"}
			}
			client := &userIdentityTestClient{transport: TransportREST, items: []MikrotikItem{actual}}
			if diagnostics := resource.ReadContext(context.Background(), data, client); diagnostics.HasError() {
				t.Fatal(diagnostics)
			}
			if data.Id() != "*1" {
				t.Fatal("import lost its ID")
			}
			if path == "/user" {
				client.items[0]["name"] = "bob"
			} else {
				client.items[0]["fingerprint"] = "SHA256:other"
			}
			if diagnostics := resource.ReadContext(context.Background(), data, client); !diagnostics.HasError() {
				t.Fatal("identity established by import was not protected")
			}
		})
	}
}

func TestUserResourceIdentityFingerprintFallback(t *testing.T) {
	userIdentityTestVersion(t)
	key, public := userIdentityTestKey(t)
	for _, actual := range []string{ssh.FingerprintSHA256(public), ssh.FingerprintSHA256(public) + "=", ssh.FingerprintLegacyMD5(public), "MD5:" + ssh.FingerprintLegacyMD5(public)} {
		data := userIdentityTestData(t, ResourceUserSshKeys(), map[string]interface{}{"user": "alice", "key": key})
		data.SetId("*1")
		if err := validateUserResourceIdentity("/user/ssh-keys", []MikrotikItem{{"user": "alice", "fingerprint": actual}}, data); err != nil {
			t.Fatal(err)
		}
	}
	data := userIdentityTestData(t, ResourceUserSshKeys(), map[string]interface{}{"user": "alice", "key": "invalid-public-key"})
	if err := validateUserResourceIdentity("/user/ssh-keys", []MikrotikItem{{"user": "alice", "fingerprint": "SHA256:any"}}, data); err == nil {
		t.Fatal("unverifiable key was accepted")
	}
}

func TestUserResourceIdentityAllowsPlannedChanges(t *testing.T) {
	userIdentityTestVersion(t)
	resource := ResourceUser()
	state := &terraform.InstanceState{ID: "*1", Attributes: map[string]string{"name": "alice", "group": "read"}}
	state.RawConfig = userIdentityTestData(t, resource, map[string]interface{}{"name": "renamed", "group": "read"}).GetRawConfig()
	diff := &terraform.InstanceDiff{Attributes: map[string]*terraform.ResourceAttrDiff{"name": {Old: "alice", New: "renamed"}}}
	data, err := schema.InternalMap(resource.Schema).Data(state, diff)
	if err != nil {
		t.Fatal(err)
	}
	if err := validateUserResourceIdentity("/user", []MikrotikItem{{"name": "alice"}}, data); err != nil {
		t.Fatalf("planned rename rejected: %v", err)
	}
	if err := validateUserResourceIdentity("/user", []MikrotikItem{{"name": "renamed"}}, data); err == nil {
		t.Fatal("planned new name was accepted as the prior identity")
	}
	client := &userIdentityTestClient{transport: TransportREST, items: []MikrotikItem{{".id": "*1", "name": "alice", "group": "read"}}}
	if diagnostics := resource.UpdateContext(context.Background(), data, client); diagnostics.HasError() {
		t.Fatal(diagnostics)
	}
	if client.writes != 1 || client.lastItem["name"] != "renamed" {
		t.Fatal("guard blocked the planned rename or wrote the wrong name")
	}
	key, public := userIdentityTestKey(t)
	newKey, _ := userIdentityTestKey(t)
	resource = ResourceUserSshKeys()
	state = &terraform.InstanceState{ID: "*1", Attributes: map[string]string{"user": "alice", "key": key, "fingerprint": ssh.FingerprintSHA256(public)}}
	diff = &terraform.InstanceDiff{Attributes: map[string]*terraform.ResourceAttrDiff{"key": {Old: key, New: newKey}, "user": {Old: "alice", New: "bob"}}}
	data, err = schema.InternalMap(resource.Schema).Data(state, diff)
	if err != nil {
		t.Fatal(err)
	}
	if err := validateUserResourceIdentity("/user/ssh-keys", []MikrotikItem{{"user": "alice", "fingerprint": ssh.FingerprintSHA256(public)}}, data); err != nil {
		t.Fatalf("planned key replacement or reassignment rejected: %v", err)
	}
}

func TestUserResourceIdentityReadErrorsAndAmbiguity(t *testing.T) {
	userIdentityTestVersion(t)
	resource := ResourceUser()
	for _, client := range []*userIdentityTestClient{
		{transport: TransportREST, readErr: errors.New("test read failure")},
		{transport: TransportREST, items: []MikrotikItem{{"name": "alice"}, {"name": "alice"}}},
	} {
		data := userIdentityTestData(t, resource, map[string]interface{}{"name": "alice", "group": "read"})
		data.SetId("*1")
		if diagnostics := resource.UpdateContext(context.Background(), data, client); !diagnostics.HasError() || client.writes != 0 {
			t.Fatalf("unsafe update after a failed lookup: %v", diagnostics)
		}
	}
	data := userIdentityTestData(t, resource, map[string]interface{}{"name": "alice", "group": "read"})
	if err := validateUserResourceIdentity("/interface/vlan", []MikrotikItem{{"name": "bob"}}, data); err != nil {
		t.Fatal("unrelated resource behavior changed")
	}
}

func TestUserResourceIdentityRechecksAfterRefresh(t *testing.T) {
	userIdentityTestVersion(t)
	for _, transport := range []TransportType{TransportAPI, TransportREST} {
		for _, operation := range []string{"update", "delete"} {
			t.Run(fmt.Sprintf("%v/%s", transport, operation), func(t *testing.T) {
				resource := ResourceUser()
				data := userIdentityTestData(t, resource, map[string]interface{}{"name": "alice", "group": "read", "password": "test-only-secret"})
				client := &userIdentityTestClient{transport: transport, items: []MikrotikItem{{".id": "*1", "name": "alice", "group": "read"}}}
				if diagnostics := resource.ReadContext(context.Background(), data, client); diagnostics.HasError() {
					t.Fatal(diagnostics)
				}
				// A successful plan/refresh does not authorize a later write if
				// the ID resolves to another user immediately before applying.
				client.items[0]["name"] = "bob"
				var diagnostics diag.Diagnostics
				if operation == "update" {
					diagnostics = resource.UpdateContext(context.Background(), data, client)
				} else {
					diagnostics = resource.DeleteContext(context.Background(), data, client)
				}
				if !diagnostics.HasError() || client.writes != 0 {
					t.Fatalf("write was not revalidated after refresh: %v", diagnostics)
				}
			})
		}
	}
}
