package routeros

import "testing"

func TestRouterOSVersionAtLeast(t *testing.T) {
	originalVersion := RouterOSVersion
	defer func() {
		RouterOSVersion = originalVersion
	}()

	tests := []struct {
		name    string
		current string
		min     string
		want    bool
	}{
		{name: "exact match", current: "7.22", min: "7.22", want: true},
		{name: "patch release", current: "7.22.3", min: "7.22", want: true},
		{name: "below minimum", current: "7.21.3", min: "7.22", want: false},
		{name: "trims whitespace", current: " 7.22.3 ", min: "7.22", want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RouterOSVersion = tt.current

			got, err := routerOSVersionAtLeast(tt.min)
			if err != nil {
				t.Fatalf("routerOSVersionAtLeast() error = %v", err)
			}

			if got != tt.want {
				t.Fatalf("routerOSVersionAtLeast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouterOSVersionAtLeastWithoutVersion(t *testing.T) {
	originalVersion := RouterOSVersion
	defer func() {
		RouterOSVersion = originalVersion
	}()

	RouterOSVersion = ""

	if _, err := routerOSVersionAtLeast("7.22"); err == nil {
		t.Fatal("routerOSVersionAtLeast() expected error when RouterOSVersion is unset")
	}
}
