package integer

import "testing"

func TestAdd(t *testing.T) {
	got := add(2, 2)
	want := 4

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}