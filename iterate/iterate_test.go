package iterate

import "testing"

func TestIterate(t *testing.T) {
	got:= Iterate("a",9)
	want:= "aaaaaaaaa"
	if got!=want{
		t.Errorf("got %q wanted %q ", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iterate("a",9)
	}
}