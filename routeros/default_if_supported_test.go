package routeros

import (
	"testing"
)

func TestDefaultIfSupported(t *testing.T) {
	t.Run("returns nil if not enabled", func(t *testing.T) {
		defaultFunc := DefaultIfSupported("something")
		if value, _ := defaultFunc(); value != nil {
			t.Fatalf("expected nil, got %v", value)
		}
	})

	t.Run("returns value if enabled", func(t *testing.T) {
		defaultFunc := DefaultIfSupported("something")
		SetSupported(defaultFunc, true)
		if value, _ := defaultFunc(); value != "something" {
			t.Fatalf(`expected "something", got %v`, value)
		}
	})

	t.Run("does not overlap with other functions", func(t *testing.T) {
		defaultFunc1 := DefaultIfSupported("something1")
		defaultFunc2 := DefaultIfSupported("something2")
		SetSupported(defaultFunc1, true)
		if value, _ := defaultFunc1(); value != "something1" {
			t.Fatalf(`expected "something1", got %v`, value)
		}
		if value, _ := defaultFunc2(); value != nil {
			t.Fatalf(`expected nil, got %v`, value)
		}
	})

	t.Run("disables function back", func(t *testing.T) {
		defaultFunc := DefaultIfSupported("something")
		SetSupported(defaultFunc, true)
		if value, _ := defaultFunc(); value == nil {
			t.Fatal(`expected "something", got nil`)
		}

		SetSupported(defaultFunc, false)
		if value, _ := defaultFunc(); value != nil {
			t.Fatalf("expected nil, got %v", value)
		}
	})
}
