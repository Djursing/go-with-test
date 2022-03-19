package iterations

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	word := Repeat("a", 2)
	fmt.Println(word)
	// Output: aa
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	exp := "aaaaa"

	if repeated != exp {
		t.Errorf("expected %q but got %q", exp, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", b.N)
	}
}
