package gogacap

import (
	"testing"
)

func sliceEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i != len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestReverseInts(t *testing.T) {
	test := func(a, b []int) {
		reverseInts(a)
		if !sliceEq(a, b) {
			t.Errorf("result(%v) != %v", a, b)
		}
	}

	test([]int{}, []int{})
	test([]int{0}, []int{0})
	test([]int{0, 1}, []int{1, 0})
	test([]int{1, 1}, []int{1, 1})
	test([]int{1, 0}, []int{0, 1})
	test([]int{1, 0, 2}, []int{2, 0, 1})
	test([]int{1, 0, 2, 0}, []int{0, 2, 0, 1})
	test([]int{1, 0, 2, 0, 4}, []int{4, 0, 2, 0, 1})
}
