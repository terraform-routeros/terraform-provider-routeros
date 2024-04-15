package routeros

import (
	"testing"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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

func TestTimeEquall(t *testing.T) {
	type args struct {
		old string
		new string
		d   *schema.ResourceData
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"TimeEquall #1", args{"", "", nil}, true},
		{"TimeEquall #2", args{"none", "none", nil}, true},
		{"TimeEquall #3", args{"none-dynamic", "none-dynamic", nil}, true},
		{"TimeEquall #4", args{"none", "", nil}, true},
		{"TimeEquall #5", args{"", "none", nil}, true},
		{"TimeEquall #6", args{"none-dynamic", "", nil}, true},
		{"TimeEquall #7", args{"", "none-dynamic", nil}, true},
		{"TimeEquall #8", args{"", "1m20s", nil}, false},
		{"TimeEquall #9", args{"1m20s", "", nil}, false},
		{"TimeEquall #10", args{"none", "1m20s", nil}, false},
		{"TimeEquall #11", args{"1m20s", "none", nil}, false},
		{"TimeEquall #12", args{"none-dynamic", "1m20s", nil}, false},
		{"TimeEquall #13", args{"1m20s", "none-dynamic", nil}, false},
		{"TimeEquall #14", args{"1m20s", "1h30m", nil}, false},
		{"TimeEquall #15", args{"1m20s", "80s", nil}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeEquall("", tt.args.old, tt.args.new, tt.args.d); got != tt.want {
				t.Errorf("TimeEquall() = %v, want %v", got, tt.want)
			}
		})
	}
}
