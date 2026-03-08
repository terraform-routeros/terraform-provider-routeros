package routeros

import (
	"testing"
	"time"

	"github.com/hashicorp/go-cty/cty"
)

func TestValidationDurationAtLeast(t *testing.T) {
	const minDuration = time.Minute
	cases := []struct {
		arg      string
		hasError bool
	}{
		{"1s", true},
		{"59", true},
		{"1m", false},
		{"1m1s", false},
		{"00:01:01", false},
		{"", true},
		{"invalid", true},
	}
	validator := ValidationDurationAtLeast(minDuration)
	for _, c := range cases {
		result := validator(c.arg, *new(cty.Path))
		hasError := result.HasError()
		if hasError != c.hasError {
			t.Errorf("ValidationDurationAtLeast(%v)(%q, ...).hasError() == %t, want %t. Diagnostics: %v.", minDuration, c.arg, hasError, c.hasError, result)
		}
	}
}

func TestValidationMultiValInSlice(t *testing.T) {
	type args struct {
		valid            []string
		ignoreCase       bool
		mikrotikNegative bool
		value            string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Positive #1",
			args{[]string{"a", "b", "c", "d"}, false, false, "a, b,   c"},
			0,
		},
		{
			"Positive #2",
			args{[]string{"a", "b", "c", "d"}, false, false, "a,b"},
			0,
		},
		{
			"Positive #3",
			args{[]string{"a", "b", "c", "d"}, false, false, "d"},
			0,
		},
		{
			"Positive #4",
			args{[]string{"a", "b", "c", "d"}, false, true, "a,!d"},
			0,
		},
		{
			"Negative #1",
			args{[]string{"a", "b", "c", "d"}, false, false, "a,e"},
			1,
		},
		{
			"Negative #2",
			args{[]string{"a", "b", "c", "d"}, false, false, "a,e,f,g"},
			3,
		},
		{
			"Negative #3",
			args{[]string{"a", "b", "c", "d"}, false, false, "a,b,,,"},
			3,
		},
		{
			"Positive #4",
			args{[]string{"a", "b", "c", "d"}, false, false, "a,!d"},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := ValidationMultiValInSlice(tt.args.valid, tt.args.ignoreCase, tt.args.mikrotikNegative)
			if got := f(tt.args.value, *new(cty.Path)); len(got) != tt.want {
				t.Errorf("ValidationMultiValInSlice() diag length = %v, want %v, diags: %v", len(got), tt.want, got)
			}
		})
	}
}

func TestValidationValInSlice(t *testing.T) {
	type args struct {
		valid            []string
		ignoreCase       bool
		mikrotikNegative bool
		value            string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Positive #1",
			args{[]string{"a", "b", "c", "d"}, false, false, " a "},
			0,
		},
		{
			"Positive #2",
			args{[]string{"a", "b", "c", "d"}, false, false, "b"},
			0,
		},
		{
			"Positive #3",
			args{[]string{"a", "b", "c", "d"}, false, false, "d"},
			0,
		},
		{
			"Positive #4",
			args{[]string{"a", "b", "c", "d"}, false, true, "!d"},
			0,
		},
		{
			"Positive #5",
			args{[]string{"a", "b", "c", "d"}, true, true, "!D"},
			0,
		},
		{
			"Negative #1",
			args{[]string{"a", "b", "c", "d"}, false, false, "e"},
			1,
		},
		{
			"Negative #2",
			args{[]string{"a", "b", "c", "d"}, false, false, "!a"},
			1,
		},
		{
			"Positive #3",
			args{[]string{"a", "b", "c", "d"}, false, false, "A"},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := ValidationValInSlice(tt.args.valid, tt.args.ignoreCase, tt.args.mikrotikNegative)
			if got := f(tt.args.value, *new(cty.Path)); len(got) != tt.want {
				t.Errorf("ValidationValInSlice() diag length = %v, want %v, diags: %v", len(got), tt.want, got)
			}
		})
	}
}

func Test_toQuotedCommaSeparatedString(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{
			"Concatenates many",
			[]string{"a", "b", "c"},
			`"a","b","c"`,
		},
		{
			"Can do oney",
			[]string{"a"},
			`"a"`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toQuotedCommaSeparatedString(tt.args...)
			if got != tt.want {
				t.Errorf("toQuotedCommaSeparatedString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImplicitSingleHostCIDR4(t *testing.T) {
	type args struct {
		old string
		new string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Single host pseudo-match",
			args{"192.168.1.2", "192.168.1.2/32"},
			true, // diff suppress
		},
		{
			"Subnet match",
			args{"172.34.45.67/24", "172.34.45.67/24"},
			false, // no diff suppress
		},
		{
			"Single host mismatch",
			args{"192.168.1.2", "192.168.3.4/32"},
			false, // no diff suppress
		},
		{
			"Subnet mismatch",
			args{"172.34.45.67/24", "10.1.2.3/24"},
			false, // no diff suppress
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ImplicitSingleHostCIDR4("", tt.args.old, tt.args.new, nil); got != tt.want {
				t.Errorf("ImplicitSingleHostCIDR4() suppress diff = %v, want %v", got, tt.want)
			}
		})
	}
}
