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
		//fmt.Println(ints, ">>>>>")
		checkCircle(t, func() bool {
			more := PermuNextInt(ints)
			//fmt.Println(ints)
			return more
		}, c, fmt.Sprintf("permutation next(%v)", ints))

		PermuPrevInt(ints)
		//fmt.Println(ints, "<<<<<")
		checkCircle(t, func() bool {
			more := PermuPrevInt(ints)
			//fmt.Println(ints)
			return more
		}, c, fmt.Sprintf("permutation prev(%v)", ints))
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

func TestPartPermutationInts(t *testing.T) {
	test := func(ints []int, p, c int) {
		fmt.Println(ints, ">>>>>")
		checkCircle(t, func() bool {
			more := PartPermuNextInt(ints, p)
			fmt.Println(ints[:p], ints[p:])
			return more
		}, c, fmt.Sprintf("part permutation next(%v, %v)", ints[:p], ints[p:]))

		PartPermuPrevInt(ints, p)
		fmt.Println(ints, "<<<<<")
		checkCircle(t, func() bool {
			more := PartPermuPrevInt(ints, p)
			fmt.Println(ints[:p], ints[p:])
			return more
		}, c, fmt.Sprintf("part permutation prev(%v, %v)", ints[:p], ints[p:]))
	}

	test([]int{}, 0, 1)
	test([]int{0}, 0, 1)
	test([]int{0}, 1, 1)
	test([]int{0, 0}, 2, 1)
	test([]int{0, 0, 0}, 0, 1)
	test([]int{0, 0, 0}, 1, 1)
	test([]int{0, 0, 1}, 1, 2)
	test([]int{0, 0, 1}, 2, 3)
	test([]int{0, 1}, 1, 2)
	test([]int{0, 1, 2}, 1, 3)
	test([]int{0, 1, 2}, 2, 6)
	test([]int{0, 1, 2}, 3, 6)
	test([]int{0, 1, 2, 3}, 1, 4)
	test([]int{0, 1, 2, 3}, 2, 12)
	test([]int{0, 1, 2, 3}, 3, 24)
	test([]int{0, 1, 2, 3}, 4, 24)
	test([]int{0, 0, 1, 1}, 1, 2)
	test([]int{0, 0, 1, 1}, 2, 4)
	test([]int{0, 0, 1, 1}, 3, 6)
	test([]int{0, 0, 1, 1}, 4, 6)
	test([]int{0, 0, 1, 1, 1}, 1, 2)
	test([]int{0, 0, 1, 1, 1}, 2, 4)
	test([]int{0, 0, 1, 1, 1}, 3, 7)
	test([]int{0, 0, 1, 1, 1}, 4, 10)
	test([]int{0, 0, 1, 1, 1}, 5, 10)
}
