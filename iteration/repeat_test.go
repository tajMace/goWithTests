package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q, got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 3)
	fmt.Printf("%q", repeated)
	// Output: "aaa"
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 1)
	}
}
