package routeros

import "testing"

func TestContainerState(t *testing.T) {
	tests := []struct {
		name string
		item MikrotikItem
		want string
	}{
		{"status field wins (RouterOS 7.1x)", MikrotikItem{"status": "pulling", "running": "false"}, "pulling"},
		{"empty status falls back to the flags", MikrotikItem{"status": "", "running": "true"}, "running"},
		// RouterOS 7.2x: a flag is only present while it is true.
		{"extracting", MikrotikItem{"extracting": "true"}, "extracting"},
		{"stopped after extraction", MikrotikItem{"stopped": "true"}, "stopped"},
		{"starting", MikrotikItem{"starting": "true"}, "starting"},
		{"running", MikrotikItem{"running": "true"}, "running"},
		{"stopping", MikrotikItem{"stopping": "true"}, "stopping"},
		{"all flags false reads as stopped", MikrotikItem{"extracting": "false", "starting": "false", "running": "false", "stopping": "false", "stopped": "true"}, "stopped"},
		{"no flag at all means the image is still being pulled", MikrotikItem{"name": "x", "running": "false", "stopped": "false"}, "pulling"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containerState(tt.item); got != tt.want {
				t.Errorf("containerState() = %q, want %q", got, tt.want)
			}
		})
	}
}
