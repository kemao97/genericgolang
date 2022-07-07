package generics

import (
	"testing"
)

func TestMap(t *testing.T) {
	input := []int{1, 2, 3}
	fx := func(a int) int {
		return a * 2
	}
	expected := []int{2, 4, 6}

	r := Map(fx, input)
	if !IsSliceEqual(expected, r) {
		t.Errorf("Fail")
	}
}
