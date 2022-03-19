package iterations

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	exp := "aaaaa"

	if repeated != exp {
		t.Errorf("expected %q but got %q", exp, repeated)
	}
}
