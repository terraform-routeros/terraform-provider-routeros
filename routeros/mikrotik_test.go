package routeros

import (
	"testing"
)

func TestBoolFromMikrotikJSON(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			`Go bool from Mikrotik JSON - "true"`,
			args{"true"},
			true,
		},
		{
			`Go bool from Mikrotik JSON - "false"`,
			args{"false"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BoolFromMikrotikJSON(tt.args.s); got != tt.want {
				t.Errorf("BoolFromMikrotikJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoolToMikrotikJSON(t *testing.T) {
	t.Parallel()

	type args struct {
		b bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			`Go bool to Mikrotik JSON - "true"`,
			args{true},
			"yes",
		},
		{
			`Go bool to Mikrotik JSON - "false"`,
			args{false},
			"no",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BoolToMikrotikJSON(tt.args.b); got != tt.want {
				t.Errorf("BoolToMikrotikJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMikrotikItem_GetID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		mi   MikrotikItem
		want string
	}{
		{
			"Get Mikrotik Item ID",
			MikrotikItem{".id": "*39"},
			"*39",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mi.GetID(Id); got != tt.want {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_kebabToSnake(t *testing.T) {
	t.Parallel()

	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Kebab to snake case",
			args{"kebab-to-snake"},
			"kebab_to_snake",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KebabToSnake(tt.args.name); got != tt.want {
				t.Errorf("KebabToSnake() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_snakeToKebab(t *testing.T) {
	t.Parallel()

	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Snake to kebab case",
			args{"snake_to_kebab"},
			"snake-to-kebab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SnakeToKebab(tt.args.name); got != tt.want {
				t.Errorf("SnakeToKebab() = %v, want %v", got, tt.want)
			}
		})
	}
}
