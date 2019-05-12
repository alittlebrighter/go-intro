package calculator

import (
	"testing"
)

func TestAdd(t *testing.T) {
	want := 7
	got := Add.IntApply(4, 1, 2)

	if got != want {
		t.Errorf("Add.IntApply(4, 1, 2) got %d want %d", got, want)
	}
}

// $ cd calculator
// $ go test -bench . -benchmem
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add.IntApply(i, i+1, i+2)
	}
}
