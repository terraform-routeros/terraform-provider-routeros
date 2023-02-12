package routeros

import (
	"testing"

	"github.com/hashicorp/go-cty/cty"
)

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
