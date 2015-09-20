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

func sliceLt(a, b []int) bool {
	i, j := 0, 0
	for i < len(a) && j < len(b) && a[i] == b[j] {
		i++
		j++
	}

	al := i < len(a)
	bl := j < len(b)
	if al && bl {
		return a[i] < b[j]
	} else {
		return !al && bl
	}
}

func sliceGt(a, b []int) bool {
	return sliceLt(b, a)
}

func sliceCp(a []int) []int {
	c := make([]int, len(a))
	copy(c, a)
	return c
}

func TestSliceF(t *testing.T) {
	test := func(f func(a, b []int) bool, a, b []int, r bool) {
		if f(a, b) != r {
			t.Errorf("slice fail: %v(%v, %v) != %v", f, a, b, r)
		}
	}

	test(sliceEq, []int{}, []int{}, true)
	test(sliceEq, []int{}, []int{1}, false)
	test(sliceEq, []int{2}, []int{1}, false)
	test(sliceEq, []int{2}, []int{}, false)

	test(sliceLt, []int{}, []int{}, false)
	test(sliceLt, []int{}, []int{1}, true)
	test(sliceLt, []int{1}, []int{}, false)
	test(sliceLt, []int{1}, []int{1}, false)
	test(sliceLt, []int{2}, []int{1}, false)
	test(sliceLt, []int{1}, []int{2}, true)
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

func TestRotateShiftRightOneInts(t *testing.T) {
	test := func(a, b []int) {
		input := fmt.Sprintf("%v", a)
		rotateShiftRightOneInts(a)
		if !sliceEq(a, b) {
			t.Errorf("rotateShiftRightOneInts(%s) => %v != %v", input, a, b)
		}
	}

	test([]int{}, []int{})
	test([]int{1}, []int{1})
	test([]int{1, 2}, []int{2, 1})
	test([]int{1, 2, 3}, []int{3, 1, 2})
}

func TestRotateShiftLeftOneInts(t *testing.T) {
	test := func(a, b []int) {
		input := fmt.Sprintf("%v", a)
		rotateShiftLeftOneInts(a)
		if !sliceEq(a, b) {
			t.Errorf("rotateShiftLeftOneInts(%s) => %v != %v", input, a, b)
		}
	}

	test([]int{}, []int{})
	test([]int{1}, []int{1})
	test([]int{1, 2}, []int{2, 1})
	test([]int{1, 2, 3}, []int{2, 3, 1})
}

func TestRotateInts(t *testing.T) {
	test := func(a []int, c int, b []int) {
		d := make([]int, len(a))
		copy(d, a)

		rotateInts(a, c)
		if !sliceEq(a, b) {
			t.Errorf("rotateInts(%v/%d) => %v != %v", d, c, a, b)
		}

		rotateBackInts(a, c)
		if !sliceEq(a, d) {
			t.Errorf("rotateBackInts(%v/%d) => %v != %v", b, c, a, d)
		}

	}

	test([]int{}, 0, []int{})
	test([]int{1}, 0, []int{1})
	test([]int{1}, 1, []int{1})
	test([]int{1, 2}, 0, []int{1, 2})
	test([]int{1, 2}, 1, []int{2, 1})
	test([]int{1, 2}, 2, []int{1, 2})
	test([]int{1, 2, 3}, 0, []int{1, 2, 3})
	test([]int{1, 2, 3}, 1, []int{2, 3, 1})
	test([]int{1, 2, 3}, 2, []int{3, 1, 2})
	test([]int{1, 2, 3}, 3, []int{1, 2, 3})
}
