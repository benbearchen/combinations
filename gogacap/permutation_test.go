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
			more := PermuNextInts(ints)
			//fmt.Println(ints)
			return more
		}, c, fmt.Sprintf("permutation next(%v)", ints))

		PermuPrevInts(ints)
		//fmt.Println(ints, "<<<<<")
		checkCircle(t, func() bool {
			more := PermuPrevInts(ints)
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
			more := PartPermuNextInts(ints, p)
			fmt.Println(ints[:p], ints[p:])
			return more
		}, c, fmt.Sprintf("part permutation next(%v, %v)", ints[:p], ints[p:]))

		PartPermuPrevInts(ints, p)
		fmt.Println(ints, "<<<<<")
		checkCircle(t, func() bool {
			more := PartPermuPrevInts(ints, p)
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

func TestSlidePermuInts(t *testing.T) {
	test := func(min, max int, ints []int, c int, b bool, s, r []int) {
		prev := make([]int, 0, len(ints))
		prev = append(prev, s...)
		prev = append(prev, r...)
		rn, rb := SlidePermuPrevInts(min, max, prev, len(s))
		if rn != c || rb != b || !sliceEq(ints, prev) {
			t.Errorf("SlidePermuPrevInts(%d, %d, %v-%v, %d) => (%d, %v)/%v-%v != (%d, %v)/%v-%v", min, max, s, r, len(s), rn, rb, prev[:rn], prev[rn:], c, b, ints[:c], ints[c:])
		}

		input := fmt.Sprintf("%v-%v", ints[:c], ints[c:])
		rn, rb = SlidePermuNextInts(min, max, ints, c)
		if rn != len(s) || rb != b || !sliceEq(ints[:rn], s) || !sliceEq(ints[rn:], r) {
			t.Errorf("SlidePermuNextInts(%d, %d, %s, %d) => (%d, %v)/%v-%v != (%d, %v)/%v-%v", min, max, input, c, rn, rb, ints[:rn], ints[rn:], len(s), b, s, r)
		}
	}

	test(0, 0, []int{}, 0, false, []int{}, []int{})

	test(0, 0, []int{1}, 0, false, []int{}, []int{1})

	test(0, 1, []int{1}, 0, true, []int{1}, []int{})
	test(0, 1, []int{1}, 1, false, []int{}, []int{1})

	test(0, 0, []int{1, 2}, 0, false, []int{}, []int{1, 2})

	test(0, 1, []int{1, 2}, 0, true, []int{1}, []int{2})
	test(0, 1, []int{1, 2}, 1, true, []int{2}, []int{1})
	test(0, 1, []int{2, 1}, 1, false, []int{}, []int{1, 2})

	test(0, 2, []int{1, 2}, 0, true, []int{1}, []int{2})
	test(0, 2, []int{1, 2}, 1, true, []int{1, 2}, []int{})
	test(0, 2, []int{1, 2}, 2, true, []int{2}, []int{1})
	test(0, 2, []int{2, 1}, 1, true, []int{2, 1}, []int{})
	test(0, 2, []int{2, 1}, 2, false, []int{}, []int{1, 2})

	test(0, 0, []int{1, 1}, 0, false, []int{}, []int{1, 1})

	test(0, 1, []int{1, 1}, 0, true, []int{1}, []int{1})
	test(0, 1, []int{1, 1}, 1, false, []int{}, []int{1, 1})

	test(0, 2, []int{1, 1}, 0, true, []int{1}, []int{1})
	test(0, 2, []int{1, 1}, 1, true, []int{1, 1}, []int{})
	test(0, 2, []int{1, 1}, 2, false, []int{}, []int{1, 1})

	test(0, 4, []int{1, 2, 3, 4}, 0, true, []int{1}, []int{2, 3, 4})
	test(0, 4, []int{1, 2, 3, 4}, 1, true, []int{1, 2}, []int{3, 4})
	test(0, 4, []int{1, 2, 3, 4}, 2, true, []int{1, 2, 3}, []int{4})
	test(0, 4, []int{1, 2, 3, 4}, 3, true, []int{1, 2, 3, 4}, []int{})
	test(0, 4, []int{1, 2, 3, 4}, 4, true, []int{1, 2, 4}, []int{3})
	test(0, 4, []int{1, 2, 4, 3}, 3, true, []int{1, 2, 4, 3}, []int{})
	test(0, 4, []int{1, 2, 4, 3}, 4, true, []int{1, 3}, []int{2, 4})
	test(0, 4, []int{1, 3, 2, 4}, 2, true, []int{1, 3, 2}, []int{4})
	test(0, 4, []int{1, 3, 2, 4}, 3, true, []int{1, 3, 2, 4}, []int{})
	test(0, 4, []int{1, 3, 2, 4}, 4, true, []int{1, 3, 4}, []int{2})
	test(0, 4, []int{1, 3, 4, 2}, 3, true, []int{1, 3, 4, 2}, []int{})
	test(0, 4, []int{1, 3, 4, 2}, 4, true, []int{1, 4}, []int{2, 3})
	test(0, 4, []int{1, 4, 2, 3}, 2, true, []int{1, 4, 2}, []int{3})
	test(0, 4, []int{1, 4, 2, 3}, 3, true, []int{1, 4, 2, 3}, []int{})
	test(0, 4, []int{1, 4, 2, 3}, 4, true, []int{1, 4, 3}, []int{2})
	test(0, 4, []int{1, 4, 3, 2}, 3, true, []int{1, 4, 3, 2}, []int{})
	test(0, 4, []int{1, 4, 3, 2}, 4, true, []int{2}, []int{1, 3, 4})
	test(0, 4, []int{2, 1, 3, 4}, 1, true, []int{2, 1}, []int{3, 4})
	test(0, 4, []int{2, 1, 3, 4}, 2, true, []int{2, 1, 3}, []int{4})
	test(0, 4, []int{2, 1, 3, 4}, 3, true, []int{2, 1, 3, 4}, []int{})
	test(0, 4, []int{2, 1, 3, 4}, 4, true, []int{2, 1, 4}, []int{3})
	test(0, 4, []int{2, 1, 4, 3}, 3, true, []int{2, 1, 4, 3}, []int{})
	test(0, 4, []int{2, 1, 4, 3}, 4, true, []int{2, 3}, []int{1, 4})
	test(0, 4, []int{2, 3, 1, 4}, 2, true, []int{2, 3, 1}, []int{4})
	test(0, 4, []int{2, 3, 1, 4}, 3, true, []int{2, 3, 1, 4}, []int{})
	test(0, 4, []int{2, 3, 1, 4}, 4, true, []int{2, 3, 4}, []int{1})
	test(0, 4, []int{2, 3, 4, 1}, 3, true, []int{2, 3, 4, 1}, []int{})
	test(0, 4, []int{2, 3, 4, 1}, 4, true, []int{2, 4}, []int{1, 3})
	test(0, 4, []int{2, 4, 1, 3}, 2, true, []int{2, 4, 1}, []int{3})
	test(0, 4, []int{2, 4, 1, 3}, 3, true, []int{2, 4, 1, 3}, []int{})
	test(0, 4, []int{2, 4, 1, 3}, 4, true, []int{2, 4, 3}, []int{1})
	test(0, 4, []int{2, 4, 3, 1}, 3, true, []int{2, 4, 3, 1}, []int{})
	test(0, 4, []int{2, 4, 3, 1}, 4, true, []int{3}, []int{1, 2, 4})
	test(0, 4, []int{3, 1, 2, 4}, 1, true, []int{3, 1}, []int{2, 4})
	test(0, 4, []int{3, 1, 2, 4}, 2, true, []int{3, 1, 2}, []int{4})
	test(0, 4, []int{3, 1, 2, 4}, 3, true, []int{3, 1, 2, 4}, []int{})
	test(0, 4, []int{3, 1, 2, 4}, 4, true, []int{3, 1, 4}, []int{2})
	test(0, 4, []int{3, 1, 4, 2}, 3, true, []int{3, 1, 4, 2}, []int{})
	test(0, 4, []int{3, 1, 4, 2}, 4, true, []int{3, 2}, []int{1, 4})
	test(0, 4, []int{3, 2, 1, 4}, 2, true, []int{3, 2, 1}, []int{4})
	test(0, 4, []int{3, 2, 1, 4}, 3, true, []int{3, 2, 1, 4}, []int{})
	test(0, 4, []int{3, 2, 1, 4}, 4, true, []int{3, 2, 4}, []int{1})
	test(0, 4, []int{3, 2, 4, 1}, 3, true, []int{3, 2, 4, 1}, []int{})
	test(0, 4, []int{3, 2, 4, 1}, 4, true, []int{3, 4}, []int{1, 2})
	test(0, 4, []int{3, 4, 1, 2}, 2, true, []int{3, 4, 1}, []int{2})
	test(0, 4, []int{3, 4, 1, 2}, 3, true, []int{3, 4, 1, 2}, []int{})
	test(0, 4, []int{3, 4, 1, 2}, 4, true, []int{3, 4, 2}, []int{1})
	test(0, 4, []int{3, 4, 2, 1}, 3, true, []int{3, 4, 2, 1}, []int{})
	test(0, 4, []int{3, 4, 2, 1}, 4, true, []int{4}, []int{1, 2, 3})
	test(0, 4, []int{4, 1, 2, 3}, 1, true, []int{4, 1}, []int{2, 3})
	test(0, 4, []int{4, 1, 2, 3}, 2, true, []int{4, 1, 2}, []int{3})
	test(0, 4, []int{4, 1, 2, 3}, 3, true, []int{4, 1, 2, 3}, []int{})
	test(0, 4, []int{4, 1, 2, 3}, 4, true, []int{4, 1, 3}, []int{2})
	test(0, 4, []int{4, 1, 3, 2}, 3, true, []int{4, 1, 3, 2}, []int{})
	test(0, 4, []int{4, 1, 3, 2}, 4, true, []int{4, 2}, []int{1, 3})
	test(0, 4, []int{4, 2, 1, 3}, 2, true, []int{4, 2, 1}, []int{3})
	test(0, 4, []int{4, 2, 1, 3}, 3, true, []int{4, 2, 1, 3}, []int{})
	test(0, 4, []int{4, 2, 1, 3}, 4, true, []int{4, 2, 3}, []int{1})
	test(0, 4, []int{4, 2, 3, 1}, 3, true, []int{4, 2, 3, 1}, []int{})
	test(0, 4, []int{4, 2, 3, 1}, 4, true, []int{4, 3}, []int{1, 2})
	test(0, 4, []int{4, 3, 1, 2}, 2, true, []int{4, 3, 1}, []int{2})
	test(0, 4, []int{4, 3, 1, 2}, 3, true, []int{4, 3, 1, 2}, []int{})
	test(0, 4, []int{4, 3, 1, 2}, 4, true, []int{4, 3, 2}, []int{1})
	test(0, 4, []int{4, 3, 2, 1}, 3, true, []int{4, 3, 2, 1}, []int{})
	test(0, 4, []int{4, 3, 2, 1}, 4, false, []int{}, []int{1, 2, 3, 4})

	test(0, 4, []int{1, 1, 3, 4}, 0, true, []int{1}, []int{1, 3, 4})
	test(0, 4, []int{1, 1, 3, 4}, 1, true, []int{1, 1}, []int{3, 4})
	test(0, 4, []int{1, 1, 3, 4}, 2, true, []int{1, 1, 3}, []int{4})
	test(0, 4, []int{1, 1, 3, 4}, 3, true, []int{1, 1, 3, 4}, []int{})
	test(0, 4, []int{1, 1, 3, 4}, 4, true, []int{1, 1, 4}, []int{3})
	test(0, 4, []int{1, 1, 4, 3}, 3, true, []int{1, 1, 4, 3}, []int{})
	test(0, 4, []int{1, 1, 4, 3}, 4, true, []int{1, 3}, []int{1, 4})
	test(0, 4, []int{1, 3, 1, 4}, 2, true, []int{1, 3, 1}, []int{4})
	test(0, 4, []int{1, 3, 1, 4}, 3, true, []int{1, 3, 1, 4}, []int{})
	test(0, 4, []int{1, 3, 1, 4}, 4, true, []int{1, 3, 4}, []int{1})
	test(0, 4, []int{1, 3, 4, 1}, 3, true, []int{1, 3, 4, 1}, []int{})
	test(0, 4, []int{1, 3, 4, 1}, 4, true, []int{1, 4}, []int{1, 3})
	test(0, 4, []int{1, 4, 1, 3}, 2, true, []int{1, 4, 1}, []int{3})
	test(0, 4, []int{1, 4, 1, 3}, 3, true, []int{1, 4, 1, 3}, []int{})
	test(0, 4, []int{1, 4, 1, 3}, 4, true, []int{1, 4, 3}, []int{1})
	test(0, 4, []int{1, 4, 3, 1}, 3, true, []int{1, 4, 3, 1}, []int{})
	test(0, 4, []int{1, 4, 3, 1}, 4, true, []int{3}, []int{1, 1, 4})
	test(0, 4, []int{3, 1, 1, 4}, 1, true, []int{3, 1}, []int{1, 4})
	test(0, 4, []int{3, 1, 1, 4}, 2, true, []int{3, 1, 1}, []int{4})
	test(0, 4, []int{3, 1, 1, 4}, 3, true, []int{3, 1, 1, 4}, []int{})
	test(0, 4, []int{3, 1, 1, 4}, 4, true, []int{3, 1, 4}, []int{1})
	test(0, 4, []int{3, 1, 4, 1}, 3, true, []int{3, 1, 4, 1}, []int{})
	test(0, 4, []int{3, 1, 4, 1}, 4, true, []int{3, 4}, []int{1, 1})
	test(0, 4, []int{3, 4, 1, 1}, 2, true, []int{3, 4, 1}, []int{1})
	test(0, 4, []int{3, 4, 1, 1}, 3, true, []int{3, 4, 1, 1}, []int{})
	test(0, 4, []int{3, 4, 1, 1}, 4, true, []int{4}, []int{1, 1, 3})
	test(0, 4, []int{4, 1, 1, 3}, 1, true, []int{4, 1}, []int{1, 3})
	test(0, 4, []int{4, 1, 1, 3}, 2, true, []int{4, 1, 1}, []int{3})
	test(0, 4, []int{4, 1, 1, 3}, 3, true, []int{4, 1, 1, 3}, []int{})
	test(0, 4, []int{4, 1, 1, 3}, 4, true, []int{4, 1, 3}, []int{1})
	test(0, 4, []int{4, 1, 3, 1}, 3, true, []int{4, 1, 3, 1}, []int{})
	test(0, 4, []int{4, 1, 3, 1}, 4, true, []int{4, 3}, []int{1, 1})
	test(0, 4, []int{4, 3, 1, 1}, 2, true, []int{4, 3, 1}, []int{1})
	test(0, 4, []int{4, 3, 1, 1}, 3, true, []int{4, 3, 1, 1}, []int{})
	test(0, 4, []int{4, 3, 1, 1}, 4, false, []int{}, []int{1, 1, 3, 4})

	test(2, 4, []int{4, 3, 1, 1}, 2, true, []int{4, 3, 1}, []int{1})
	test(3, 4, []int{4, 1, 3, 1}, 4, true, []int{4, 3, 1}, []int{1})
	test(3, 3, []int{4, 1, 3, 1}, 3, true, []int{4, 3, 1}, []int{1})
	test(1, 4, []int{1, 4, 3, 1}, 4, true, []int{3}, []int{1, 1, 4})
	test(2, 4, []int{1, 4, 3, 1}, 4, true, []int{3, 1}, []int{1, 4})
	test(0, 3, []int{4, 1, 1, 3}, 3, true, []int{4, 1, 3}, []int{1})
	test(0, 3, []int{4, 3, 1, 1}, 3, false, []int{}, []int{1, 1, 3, 4})
	test(1, 3, []int{4, 3, 1, 1}, 3, false, []int{1}, []int{1, 3, 4})
	test(2, 3, []int{4, 3, 1, 1}, 3, false, []int{1, 1}, []int{3, 4})
	test(3, 3, []int{4, 3, 1, 1}, 3, false, []int{1, 1, 3}, []int{4})
	test(2, 2, []int{4, 3, 1, 1}, 2, false, []int{1, 1}, []int{3, 4})

	test(0, 4, []int{1, 1, 4, 4}, 0, true, []int{1}, []int{1, 4, 4})
	test(0, 4, []int{1, 1, 4, 4}, 1, true, []int{1, 1}, []int{4, 4})
	test(0, 4, []int{1, 1, 4, 4}, 2, true, []int{1, 1, 4}, []int{4})
	test(0, 4, []int{1, 1, 4, 4}, 3, true, []int{1, 1, 4, 4}, []int{})
	test(0, 4, []int{1, 1, 4, 4}, 4, true, []int{1, 4}, []int{1, 4})
	test(0, 4, []int{1, 4, 1, 4}, 2, true, []int{1, 4, 1}, []int{4})
	test(0, 4, []int{1, 4, 1, 4}, 3, true, []int{1, 4, 1, 4}, []int{})
	test(0, 4, []int{1, 4, 1, 4}, 4, true, []int{1, 4, 4}, []int{1})
	test(0, 4, []int{1, 4, 4, 1}, 3, true, []int{1, 4, 4, 1}, []int{})
	test(0, 4, []int{1, 4, 4, 1}, 4, true, []int{4}, []int{1, 1, 4})
	test(0, 4, []int{4, 1, 1, 4}, 1, true, []int{4, 1}, []int{1, 4})
	test(0, 4, []int{4, 1, 1, 4}, 2, true, []int{4, 1, 1}, []int{4})
	test(0, 4, []int{4, 1, 1, 4}, 3, true, []int{4, 1, 1, 4}, []int{})
	test(0, 4, []int{4, 1, 1, 4}, 4, true, []int{4, 1, 4}, []int{1})
	test(0, 4, []int{4, 1, 4, 1}, 3, true, []int{4, 1, 4, 1}, []int{})
	test(0, 4, []int{4, 1, 4, 1}, 4, true, []int{4, 4}, []int{1, 1})
	test(0, 4, []int{4, 4, 1, 1}, 2, true, []int{4, 4, 1}, []int{1})
	test(0, 4, []int{4, 4, 1, 1}, 3, true, []int{4, 4, 1, 1}, []int{})
	test(0, 4, []int{4, 4, 1, 1}, 4, false, []int{}, []int{1, 1, 4, 4})

	test(0, 4, []int{1, 1, 1, 4}, 0, true, []int{1}, []int{1, 1, 4})
	test(0, 4, []int{1, 1, 1, 4}, 1, true, []int{1, 1}, []int{1, 4})
	test(0, 4, []int{1, 1, 1, 4}, 2, true, []int{1, 1, 1}, []int{4})
	test(0, 4, []int{1, 1, 1, 4}, 3, true, []int{1, 1, 1, 4}, []int{})
	test(0, 4, []int{1, 1, 1, 4}, 4, true, []int{1, 1, 4}, []int{1})
	test(0, 4, []int{1, 1, 4, 1}, 3, true, []int{1, 1, 4, 1}, []int{})
	test(0, 4, []int{1, 1, 4, 1}, 4, true, []int{1, 4}, []int{1, 1})
	test(0, 4, []int{1, 4, 1, 1}, 2, true, []int{1, 4, 1}, []int{1})
	test(0, 4, []int{1, 4, 1, 1}, 3, true, []int{1, 4, 1, 1}, []int{})
	test(0, 4, []int{1, 4, 1, 1}, 4, true, []int{4}, []int{1, 1, 1})
	test(0, 4, []int{4, 1, 1, 1}, 1, true, []int{4, 1}, []int{1, 1})
	test(0, 4, []int{4, 1, 1, 1}, 2, true, []int{4, 1, 1}, []int{1})
	test(0, 4, []int{4, 1, 1, 1}, 3, true, []int{4, 1, 1, 1}, []int{})
	test(0, 4, []int{4, 1, 1, 1}, 4, false, []int{}, []int{1, 1, 1, 4})

	test(0, 4, []int{1, 1, 1, 1}, 0, true, []int{1}, []int{1, 1, 1})
	test(0, 4, []int{1, 1, 1, 1}, 1, true, []int{1, 1}, []int{1, 1})
	test(0, 4, []int{1, 1, 1, 1}, 2, true, []int{1, 1, 1}, []int{1})
	test(0, 4, []int{1, 1, 1, 1}, 3, true, []int{1, 1, 1, 1}, []int{})
	test(0, 4, []int{1, 1, 1, 1}, 4, false, []int{}, []int{1, 1, 1, 1})
}
