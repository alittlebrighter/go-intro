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
