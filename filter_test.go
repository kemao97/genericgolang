package generics

import (
	"testing"
)

func TestFilter(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []int{2}
	fx := func(a int) bool {
		return a%2 == 0
	}

	r := Filter(fx, input)

	if !IsSliceEqual(r, expected) {
		t.Errorf("Fail")
	}
}
