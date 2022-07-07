package generics

import (
	"runtime"
	"testing"
)

func TestPmap(t *testing.T) {
	input := []int{1, 2, 3}
	fx := func(a int) int {
		return a * 2
	}
	expected := []int{2, 4, 6}

	r := Pmap(fx, input, runtime.GOMAXPROCS(0))
	if !IsSliceEqual(expected, r) {
		t.Errorf("Fail")
	}
}
