package routeros

import (
	"context"
	"reflect"
	"testing"
)

func TestStateMigrationScalarToList(t *testing.T) {
	testCases := []struct {
		name  string
		state map[string]interface{}
		want  map[string]interface{}
	}{
		{
			name:  "comma separated string is split",
			state: map[string]interface{}{"servers": "192.168.1.1,1.1.1.1"},
			want:  map[string]interface{}{"servers": []interface{}{"192.168.1.1", "1.1.1.1"}},
		},
		{
			name:  "single value becomes a one element list",
			state: map[string]interface{}{"servers": "192.168.1.1"},
			want:  map[string]interface{}{"servers": []interface{}{"192.168.1.1"}},
		},
		{
			name:  "surrounding spaces are trimmed",
			state: map[string]interface{}{"servers": "192.168.1.1, 1.1.1.1"},
			want:  map[string]interface{}{"servers": []interface{}{"192.168.1.1", "1.1.1.1"}},
		},
		{
			name:  "empty string becomes an empty list",
			state: map[string]interface{}{"servers": ""},
			want:  map[string]interface{}{"servers": []interface{}{}},
		},
		{
			name:  "absent attribute is left alone",
			state: map[string]interface{}{"other": "x"},
			want:  map[string]interface{}{"other": "x"},
		},
		{
			name:  "already a list is left alone",
			state: map[string]interface{}{"servers": []interface{}{"1.1.1.1"}},
			want:  map[string]interface{}{"servers": []interface{}{"1.1.1.1"}},
		},
		{
			name:  "non string scalar is wrapped",
			state: map[string]interface{}{"servers": 5},
			want:  map[string]interface{}{"servers": []interface{}{5}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := stateMigrationScalarToList("servers")(context.Background(), tc.state, nil)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %#v, want %#v", got, tc.want)
			}
		})
	}
}

func TestStateMigrationScalarToListMultipleKeys(t *testing.T) {
	state := map[string]interface{}{
		"auth":   "sha1,md5",
		"cipher": "aes256",
	}

	got, err := stateMigrationScalarToList("auth", "cipher")(context.Background(), state, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := map[string]interface{}{
		"auth":   []interface{}{"sha1", "md5"},
		"cipher": []interface{}{"aes256"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}
