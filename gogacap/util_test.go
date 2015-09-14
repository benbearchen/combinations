package gogacap

import (
	"testing"
)

import (
	"fmt"
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

func TestUpperBoundInts(t *testing.T) {
	test := func(a []int, x, i int) {
		p := upperBoundInts(a, x)
		if p != i {
			t.Errorf("upperBoundInts(%v, %d) => %d != %d", a, x, p, i)
		}
	}

	test([]int{}, 1, 0)
	test([]int{0}, -1, 0)
	test([]int{0}, 0, 1)
	test([]int{0}, 1, 1)
	test([]int{0, 0}, -1, 0)
	test([]int{0, 0}, 0, 2)
	test([]int{0, 0}, 1, 2)
	test([]int{0, 0, 1}, -1, 0)
	test([]int{0, 0, 1}, 0, 2)
	test([]int{0, 0, 1}, 1, 3)
	test([]int{0, 0, 1}, 2, 3)
	test([]int{0, 1, 1}, -1, 0)
	test([]int{0, 1, 1}, 0, 1)
	test([]int{0, 1, 1}, 1, 3)
	test([]int{0, 1, 1}, 2, 3)
	test([]int{0, 1, 2}, -1, 0)
	test([]int{0, 1, 2}, 0, 1)
	test([]int{0, 1, 2}, 1, 2)
	test([]int{0, 1, 2}, 2, 3)
	test([]int{0, 1, 2}, 3, 3)
}

func TestLowerBoundInts(t *testing.T) {
	test := func(a []int, x, i int) {
		p := lowerBoundInts(a, x)
		if p != i {
			t.Errorf("lowerBoundInts(%v, %d) => %d != %d", a, x, p, i)
		}
	}

	test([]int{}, 1, 0)
	test([]int{0}, -1, 0)
	test([]int{0}, 0, 0)
	test([]int{0}, 1, 1)
	test([]int{0, 0}, -1, 0)
	test([]int{0, 0}, 0, 0)
	test([]int{0, 0}, 1, 2)
	test([]int{0, 0, 1}, -1, 0)
	test([]int{0, 0, 1}, 0, 0)
	test([]int{0, 0, 1}, 1, 2)
	test([]int{0, 0, 1}, 2, 3)
	test([]int{0, 1, 1}, -1, 0)
	test([]int{0, 1, 1}, 0, 0)
	test([]int{0, 1, 1}, 1, 1)
	test([]int{0, 1, 1}, 2, 3)
	test([]int{0, 1, 2}, -1, 0)
	test([]int{0, 1, 2}, 0, 0)
	test([]int{0, 1, 2}, 1, 1)
	test([]int{0, 1, 2}, 2, 2)
	test([]int{0, 1, 2}, 3, 3)
}

func TestInplaceMergeInts(t *testing.T) {
	test := func(a, b, c, d []int) {
		input := fmt.Sprintf("%v-%v", a, b)
		inplaceMergeInts(a, b)
		if !sliceEq(a, c) || !sliceEq(b, d) {
			t.Errorf("%s => %v-%v != %v-%v", input, a, b, c, d)
		}
	}

	test([]int{}, []int{}, []int{}, []int{})
	test([]int{0}, []int{1}, []int{0}, []int{1})
	test([]int{1}, []int{0}, []int{0}, []int{1})
	test([]int{1, 2}, []int{0}, []int{0, 1}, []int{2})
	test([]int{2}, []int{0, 1}, []int{0}, []int{1, 2})
	test([]int{0, 1}, []int{1}, []int{0, 1}, []int{1})
	test([]int{0, 2}, []int{1}, []int{0, 1}, []int{2})
	test([]int{0, 2}, []int{1, 2, 2}, []int{0, 1}, []int{2, 2, 2})
	test([]int{0, 2}, []int{1, 2, 3}, []int{0, 1}, []int{2, 2, 3})
	test([]int{0, 4}, []int{1, 2, 3}, []int{0, 1}, []int{2, 3, 4})
}
