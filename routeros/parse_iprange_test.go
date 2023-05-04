package routeros

import "testing"

func TestIpRangeToCIDR(t *testing.T) {
	t.Parallel()

	type args struct {
		ip1 string
		ip2 string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Solid range", args{"192.168.0.0", "192.168.1.255"}, "192.168.0.0/23", false},
		{"Non-solid range", args{"192.168.0.0", "192.168.1.254"}, "192.168.0.0-192.168.1.254", false},
		{"Wrong range", args{"192.168.2.0", "192.168.1.255"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IpRangeToCIDR(tt.args.ip1, tt.args.ip2)
			if (err != nil) != tt.wantErr {
				t.Errorf("IpRangeToCIDR() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IpRangeToCIDR() got = %v, want %v", got, tt.want)
			}
		})
	}
}
