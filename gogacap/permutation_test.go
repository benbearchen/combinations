package gogacap

import (
	"testing"
)

import (
	"fmt"
)

func checkCircle(t *testing.T, f func() bool, c int, tag string) {
	r := 0
	for f() {
		r++
	}

	r++
	if r != c {
		t.Errorf("%s: circle(%d) != %d", tag, r, c)
	}
}

func TestPermutationInts(t *testing.T) {
	test := func(ints []int, c int) {
		fmt.Println(ints, ">>>>>")
		checkCircle(t, func() bool {
			more := PermuNextInt(ints)
			fmt.Println(ints)
			return more
		}, c, fmt.Sprintf("nex(%v)", ints))

		PermuPrevInt(ints)
		fmt.Println(ints, "<<<<<")
		checkCircle(t, func() bool {
			more := PermuPrevInt(ints)
			fmt.Println(ints)
			return more
		}, c, fmt.Sprintf("prev(%v)", ints))
	}

	test([]int{}, 1)
	test([]int{0}, 1)
	test([]int{0, 0}, 1)
	test([]int{0, 0, 0}, 1)
	test([]int{0, 0, 1}, 3)
	test([]int{0, 1}, 2)
	test([]int{0, 1, 2}, 6)
	test([]int{0, 1, 2, 3}, 24)
	test([]int{0, 0, 1, 1}, 6)
	test([]int{0, 0, 1, 1, 1}, 10)
}
