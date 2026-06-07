package tdd

import "testing"

func TestAdd(t *testing.T) {
	got := Add(2, 3)
	want := 5
	if got != want {
		t.Errorf("Add(2, 3) = %d; want %d", got, want)
	}
}

func TestSubtract(t *testing.T) {
	got := Subtract(5, 3)
	want := 2
	if got != want {
		t.Errorf("Subtract(5, 3) = %d; want %d", got, want)
	}
}

func TestMultiply(t *testing.T) {
	got := Multiply(4, 3)
	want := 12
	if got != want {
		t.Errorf("Multiply(4, 3) = %d; want %d", got, want)
	}
}

func TestDivide(t *testing.T) {
	t.Run("valid division", func(t *testing.T) {
		got, err := Divide(6, 3)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := 2.0
		if got != want {
			t.Errorf("Divide(6, 3) = %f; want %f", got, want)
		}
	})

	t.Run("divide by zero", func(t *testing.T) {
		_, err := Divide(6, 0)
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}
