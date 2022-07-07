package generics

import "testing"

func TestReduce(t *testing.T) {
	input := []int{1, 2, 3}
	fx := func(collect int, item int) int {
		return collect + item
	}

	expected := 6

	r := Reduce(fx, input)
	if expected != r {
		t.Errorf("Fail")
	}
	t.Log("input: ", input)
	t.Log("expected: ", expected)
	t.Log("TestReduce Success")
}
