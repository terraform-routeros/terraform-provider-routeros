package routeros

import "testing"

func TestParseBitValue(t *testing.T) {

	tests := []struct {
		name    string
		arg     string
		want    uint64
		wantErr bool
	}{
		{"Positive #1", "30M", 30000000, false},
		{"Positive #2", "30", 30, false},
		{"Positive #3", "3E", 3000000000000000000, false},
		{"Positive #4", "10Mbps", 10000000, false},
		{"Negative #1", "30A", 0, true},
		{"Negative #2", "30.0", 0, true},
		{"Negative #3", "30Mb", 0, true},
		{"Negative #4", "30m", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseBitValues(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseBitValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseBitValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseByteValue(t *testing.T) {

	tests := []struct {
		name    string
		arg     string
		want    uint64
		wantErr bool
	}{
		{"Positive #1", "64M", 67108864, false},
		{"Positive #2", "30", 30, false},
		{"Positive #3", "3E", 3458764513820540928, false},
		{"Negative #1", "30A", 0, true},
		{"Negative #2", "30.0", 0, true},
		{"Negative #3", "30Mb", 0, true},
		{"Negative #4", "30m", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseByteValues(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseByteValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseByteValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
