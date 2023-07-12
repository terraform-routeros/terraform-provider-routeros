package routeros

import (
	"testing"
	"time"
)

func TestParseDuration(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Duration
		wantErr bool
	}{
		{name: "300ms", args: args{"300ms"}, want: time.Duration(300 * time.Millisecond), wantErr: false},
		{name: "300", args: args{"300"}, want: time.Duration(300 * time.Second), wantErr: false},
		{name: "300s", args: args{"300s"}, want: time.Duration(300 * time.Second), wantErr: false},
		{name: "00:00:10", args: args{"00:00:10"}, want: time.Duration(10 * time.Second), wantErr: false},
		{name: "2h45m", args: args{"2h45m"}, want: time.Duration(2*time.Hour + 45*time.Minute), wantErr: false},
		{name: "163w4d9h", args: args{"163w4d9h"}, want: time.Duration(98960400 * time.Second), wantErr: false},
		{name: "27489h", args: args{"27489h"}, want: time.Duration(98960400 * time.Second), wantErr: false},
		{name: "120ms", args: args{"0.12"}, want: time.Duration(120 * time.Millisecond), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDuration(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDuration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseDuration() got = %v, want %v", got, tt.want)
			}
		})
	}
}
