package gogacap

import (
	"testing"
)

import (
	"fmt"
	"sync"
)

func TestCombinationNextInts(t *testing.T) {
	one := func(a, b, c, d []int, r bool) {
		input := fmt.Sprintf("%v-%v")
		q := combinationNextInts(a, b)
		if q != r || !sliceEq(a, c) || !sliceEq(b, d) {
			t.Errorf("%s => %v-%v/%v != %v-%v/%v", input, a, b, q, c, d, r)
		}
	}

	one([]int{}, []int{}, []int{}, []int{}, false)
	one([]int{0}, []int{}, []int{0}, []int{}, false)
	one([]int{}, []int{0}, []int{}, []int{0}, false)
	one([]int{0}, []int{0}, []int{0}, []int{0}, false)
	one([]int{0}, []int{1}, []int{1}, []int{0}, true)
	one([]int{1}, []int{0}, []int{0}, []int{1}, false)
	one([]int{0, 1}, []int{0}, []int{0, 0}, []int{1}, false)
	one([]int{0, 0}, []int{1}, []int{0, 1}, []int{0}, true)
	one([]int{0, 0}, []int{}, []int{0, 0}, []int{}, false)
	one([]int{0, 1}, []int{2}, []int{0, 2}, []int{1}, true)
	one([]int{0, 2}, []int{1}, []int{1, 2}, []int{0}, true)
	one([]int{1, 2}, []int{0}, []int{0, 1}, []int{2}, false)
	one([]int{1, 2}, []int{0, 3}, []int{1, 3}, []int{0, 2}, true)
	one([]int{1, 2, 4, 4}, []int{8, 10}, []int{1, 2, 4, 8}, []int{4, 10}, true)
	one([]int{1, 2, 4, 8}, []int{4, 10}, []int{1, 2, 4, 10}, []int{4, 8}, true)
	one([]int{1, 2, 4, 10}, []int{4, 8}, []int{1, 2, 8, 10}, []int{4, 4}, true)
	one([]int{1, 2, 8, 10}, []int{4, 4}, []int{1, 4, 4, 8}, []int{2, 10}, true)
	one([]int{1, 4, 4, 8}, []int{2, 10}, []int{1, 4, 4, 10}, []int{2, 8}, true)
	one([]int{1, 4, 4, 10}, []int{2, 8}, []int{1, 4, 8, 10}, []int{2, 4}, true)
	one([]int{1, 4, 8, 10}, []int{2, 4}, []int{2, 4, 4, 8}, []int{1, 10}, true)
	one([]int{2, 4, 4, 8}, []int{1, 10}, []int{2, 4, 4, 10}, []int{1, 8}, true)
	one([]int{2, 4, 4, 10}, []int{1, 8}, []int{2, 4, 8, 10}, []int{1, 4}, true)
	one([]int{2, 4, 8, 10}, []int{1, 4}, []int{4, 4, 8, 10}, []int{1, 2}, true)
	one([]int{4, 4, 8, 10}, []int{1, 2}, []int{1, 2, 4, 4}, []int{8, 10}, false)
	one([]int{8, 10}, []int{1, 2, 4, 4}, []int{1, 2}, []int{4, 4, 8, 10}, false)
	one([]int{1, 2}, []int{4, 4, 8, 10}, []int{1, 4}, []int{2, 4, 8, 10}, true)
}

func TestSlideCombiInts(t *testing.T) {
	pc := 0
	nc := 0
	test := func(min, max int, ints []int, c int, b bool, s, r []int) {
		prev := make([]int, 0, len(ints))
		prev = append(prev, s...)
		prev = append(prev, r...)
		rn, rb := SlideCombiPrevInts(min, max, prev, len(s))
		if rn != c || rb != b || !sliceEq(ints, prev) {
			pc++
			t.Errorf("SlideCombiPrevInts(%d, %d, %v-%v, %d) => (%d, %v)/%v-%v != (%d, %v)/%v-%v", min, max, s, r, len(s), rn, rb, prev[:rn], prev[rn:], c, b, ints[:c], ints[c:])
		}

		input := fmt.Sprintf("%v-%v", ints[:c], ints[c:])
		rn, rb = SlideCombiNextInts(min, max, ints, c)
		if rn != len(s) || rb != b || !sliceEq(ints[:rn], s) || !sliceEq(ints[rn:], r) {
			nc++
			t.Errorf("SlideCombiNextInts(%d, %d, %s, %d) => (%d, %v)/%v-%v != (%d, %v)/%v-%v", min, max, input, c, rn, rb, ints[:rn], ints[rn:], len(s), b, s, r)
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
	test(0, 2, []int{2, 1}, 1, false, []int{}, []int{1, 2})

	test(1, 1, []int{1, 2}, 1, true, []int{2}, []int{1})
	test(1, 1, []int{2, 1}, 1, false, []int{1}, []int{2})

	test(1, 2, []int{1, 2}, 1, true, []int{1, 2}, []int{})
	test(1, 2, []int{1, 2}, 2, true, []int{2}, []int{1})
	test(1, 2, []int{2, 1}, 1, false, []int{1}, []int{2})

	test(2, 2, []int{1, 2}, 2, false, []int{1, 2}, []int{})

	test(0, 0, []int{1, 1}, 0, false, []int{}, []int{1, 1})

	test(0, 1, []int{1, 1}, 0, true, []int{1}, []int{1})
	test(0, 1, []int{1, 1}, 1, false, []int{}, []int{1, 1})

	test(0, 2, []int{1, 1}, 0, true, []int{1}, []int{1})
	test(0, 2, []int{1, 1}, 1, true, []int{1, 1}, []int{})
	test(0, 2, []int{1, 1}, 2, false, []int{}, []int{1, 1})

	test(1, 1, []int{1, 1}, 1, false, []int{1}, []int{1})

	test(1, 2, []int{1, 1}, 1, true, []int{1, 1}, []int{})
	test(1, 2, []int{1, 1}, 2, false, []int{1}, []int{1})

	test(0, 3, []int{1, 2, 3, 4}, 0, true, []int{1}, []int{2, 3, 4})
	test(0, 3, []int{1, 2, 3, 4}, 1, true, []int{1, 2}, []int{3, 4})
	test(0, 3, []int{1, 2, 3, 4}, 2, true, []int{1, 2, 3}, []int{4})
	test(0, 3, []int{1, 2, 3, 4}, 3, true, []int{1, 2, 4}, []int{3})
	test(0, 3, []int{1, 2, 4, 3}, 3, true, []int{1, 3}, []int{2, 4})
	test(0, 3, []int{1, 3, 2, 4}, 2, true, []int{1, 3, 4}, []int{2})
	test(0, 3, []int{1, 3, 4, 2}, 3, true, []int{1, 4}, []int{2, 3})
	test(0, 3, []int{1, 4, 2, 3}, 2, true, []int{2}, []int{1, 3, 4})
	test(0, 3, []int{2, 1, 3, 4}, 1, true, []int{2, 3}, []int{1, 4})
	test(0, 3, []int{2, 3, 1, 4}, 2, true, []int{2, 3, 4}, []int{1})
	test(0, 3, []int{2, 3, 4, 1}, 3, true, []int{2, 4}, []int{1, 3})
	test(0, 3, []int{2, 4, 1, 3}, 2, true, []int{3}, []int{1, 2, 4})
	test(0, 3, []int{3, 1, 2, 4}, 1, true, []int{3, 4}, []int{1, 2})
	test(0, 3, []int{3, 4, 1, 2}, 2, true, []int{4}, []int{1, 2, 3})
	test(0, 3, []int{4, 1, 2, 3}, 1, false, []int{}, []int{1, 2, 3, 4})

	test(0, 4, []int{1, 2, 3, 4}, 0, true, []int{1}, []int{2, 3, 4})
	test(0, 4, []int{1, 2, 3, 4}, 1, true, []int{1, 2}, []int{3, 4})
	test(0, 4, []int{1, 2, 3, 4}, 2, true, []int{1, 2, 3}, []int{4})
	test(0, 4, []int{1, 2, 3, 4}, 3, true, []int{1, 2, 3, 4}, []int{})
	test(0, 4, []int{1, 2, 3, 4}, 4, true, []int{1, 2, 4}, []int{3})
	test(0, 4, []int{1, 2, 4, 3}, 3, true, []int{1, 3}, []int{2, 4})
	test(0, 4, []int{1, 3, 2, 4}, 2, true, []int{1, 3, 4}, []int{2})
	test(0, 4, []int{1, 3, 4, 2}, 3, true, []int{1, 4}, []int{2, 3})
	test(0, 4, []int{1, 4, 2, 3}, 2, true, []int{2}, []int{1, 3, 4})
	test(0, 4, []int{2, 1, 3, 4}, 1, true, []int{2, 3}, []int{1, 4})
	test(0, 4, []int{2, 3, 1, 4}, 2, true, []int{2, 3, 4}, []int{1})
	test(0, 4, []int{2, 3, 4, 1}, 3, true, []int{2, 4}, []int{1, 3})
	test(0, 4, []int{2, 4, 1, 3}, 2, true, []int{3}, []int{1, 2, 4})
	test(0, 4, []int{3, 1, 2, 4}, 1, true, []int{3, 4}, []int{1, 2})
	test(0, 4, []int{3, 4, 1, 2}, 2, true, []int{4}, []int{1, 2, 3})
	test(0, 4, []int{4, 1, 2, 3}, 1, false, []int{}, []int{1, 2, 3, 4})

	test(1, 3, []int{1, 2, 3, 4}, 1, true, []int{1, 2}, []int{3, 4})
	test(1, 3, []int{1, 2, 3, 4}, 2, true, []int{1, 2, 3}, []int{4})
	test(1, 3, []int{1, 2, 3, 4}, 3, true, []int{1, 2, 4}, []int{3})
	test(1, 3, []int{1, 2, 4, 3}, 3, true, []int{1, 3}, []int{2, 4})
	test(1, 3, []int{1, 3, 2, 4}, 2, true, []int{1, 3, 4}, []int{2})
	test(1, 3, []int{1, 3, 4, 2}, 3, true, []int{1, 4}, []int{2, 3})
	test(1, 3, []int{1, 4, 2, 3}, 2, true, []int{2}, []int{1, 3, 4})
	test(1, 3, []int{2, 1, 3, 4}, 1, true, []int{2, 3}, []int{1, 4})
	test(1, 3, []int{2, 3, 1, 4}, 2, true, []int{2, 3, 4}, []int{1})
	test(1, 3, []int{2, 3, 4, 1}, 3, true, []int{2, 4}, []int{1, 3})
	test(1, 3, []int{2, 4, 1, 3}, 2, true, []int{3}, []int{1, 2, 4})
	test(1, 3, []int{3, 1, 2, 4}, 1, true, []int{3, 4}, []int{1, 2})
	test(1, 3, []int{3, 4, 1, 2}, 2, true, []int{4}, []int{1, 2, 3})
	test(1, 3, []int{4, 1, 2, 3}, 1, false, []int{1}, []int{2, 3, 4})

	test(1, 4, []int{1, 2, 3, 4}, 1, true, []int{1, 2}, []int{3, 4})
	test(1, 4, []int{1, 2, 3, 4}, 2, true, []int{1, 2, 3}, []int{4})
	test(1, 4, []int{1, 2, 3, 4}, 3, true, []int{1, 2, 3, 4}, []int{})
	test(1, 4, []int{1, 2, 3, 4}, 4, true, []int{1, 2, 4}, []int{3})
	test(1, 4, []int{1, 2, 4, 3}, 3, true, []int{1, 3}, []int{2, 4})
	test(1, 4, []int{1, 3, 2, 4}, 2, true, []int{1, 3, 4}, []int{2})
	test(1, 4, []int{1, 3, 4, 2}, 3, true, []int{1, 4}, []int{2, 3})
	test(1, 4, []int{1, 4, 2, 3}, 2, true, []int{2}, []int{1, 3, 4})
	test(1, 4, []int{2, 1, 3, 4}, 1, true, []int{2, 3}, []int{1, 4})
	test(1, 4, []int{2, 3, 1, 4}, 2, true, []int{2, 3, 4}, []int{1})
	test(1, 4, []int{2, 3, 4, 1}, 3, true, []int{2, 4}, []int{1, 3})
	test(1, 4, []int{2, 4, 1, 3}, 2, true, []int{3}, []int{1, 2, 4})
	test(1, 4, []int{3, 1, 2, 4}, 1, true, []int{3, 4}, []int{1, 2})
	test(1, 4, []int{3, 4, 1, 2}, 2, true, []int{4}, []int{1, 2, 3})
	test(1, 4, []int{4, 1, 2, 3}, 1, false, []int{1}, []int{2, 3, 4})

	test(2, 2, []int{1, 2, 3, 4}, 2, true, []int{1, 3}, []int{2, 4})
	test(2, 2, []int{1, 3, 2, 4}, 2, true, []int{1, 4}, []int{2, 3})
	test(2, 2, []int{1, 4, 2, 3}, 2, true, []int{2, 3}, []int{1, 4})
	test(2, 2, []int{2, 3, 1, 4}, 2, true, []int{2, 4}, []int{1, 3})
	test(2, 2, []int{2, 4, 1, 3}, 2, true, []int{3, 4}, []int{1, 2})
	test(2, 2, []int{3, 4, 1, 2}, 2, false, []int{1, 2}, []int{3, 4})

	test(2, 3, []int{1, 2, 3, 4}, 2, true, []int{1, 2, 3}, []int{4})
	test(2, 3, []int{1, 2, 3, 4}, 3, true, []int{1, 2, 4}, []int{3})
	test(2, 3, []int{1, 2, 4, 3}, 3, true, []int{1, 3}, []int{2, 4})
	test(2, 3, []int{1, 3, 2, 4}, 2, true, []int{1, 3, 4}, []int{2})
	test(2, 3, []int{1, 3, 4, 2}, 3, true, []int{1, 4}, []int{2, 3})
	test(2, 3, []int{1, 4, 2, 3}, 2, true, []int{2, 3}, []int{1, 4})
	test(2, 3, []int{2, 3, 1, 4}, 2, true, []int{2, 3, 4}, []int{1})
	test(2, 3, []int{2, 3, 4, 1}, 3, true, []int{2, 4}, []int{1, 3})
	test(2, 3, []int{2, 4, 1, 3}, 2, true, []int{3, 4}, []int{1, 2})
	test(2, 3, []int{3, 4, 1, 2}, 2, false, []int{1, 2}, []int{3, 4})

	test(2, 4, []int{1, 2, 3, 4}, 2, true, []int{1, 2, 3}, []int{4})
	test(2, 4, []int{1, 2, 3, 4}, 3, true, []int{1, 2, 3, 4}, []int{})
	test(2, 4, []int{1, 2, 3, 4}, 4, true, []int{1, 2, 4}, []int{3})
	test(2, 4, []int{1, 2, 4, 3}, 3, true, []int{1, 3}, []int{2, 4})
	test(2, 4, []int{1, 3, 2, 4}, 2, true, []int{1, 3, 4}, []int{2})
	test(2, 4, []int{1, 3, 4, 2}, 3, true, []int{1, 4}, []int{2, 3})
	test(2, 4, []int{1, 4, 2, 3}, 2, true, []int{2, 3}, []int{1, 4})
	test(2, 4, []int{2, 3, 1, 4}, 2, true, []int{2, 3, 4}, []int{1})
	test(2, 4, []int{2, 3, 4, 1}, 3, true, []int{2, 4}, []int{1, 3})
	test(2, 4, []int{2, 4, 1, 3}, 2, true, []int{3, 4}, []int{1, 2})
	test(2, 4, []int{3, 4, 1, 2}, 2, false, []int{1, 2}, []int{3, 4})

	test(3, 3, []int{1, 2, 3, 4}, 3, true, []int{1, 2, 4}, []int{3})
	test(3, 3, []int{1, 2, 4, 3}, 3, true, []int{1, 3, 4}, []int{2})
	test(3, 3, []int{1, 3, 4, 2}, 3, true, []int{2, 3, 4}, []int{1})
	test(3, 3, []int{2, 3, 4, 1}, 3, false, []int{1, 2, 3}, []int{4})

	test(3, 4, []int{1, 2, 3, 4}, 3, true, []int{1, 2, 3, 4}, []int{})
	test(3, 4, []int{1, 2, 3, 4}, 4, true, []int{1, 2, 4}, []int{3})
	test(3, 4, []int{1, 2, 4, 3}, 3, true, []int{1, 3, 4}, []int{2})
	test(3, 4, []int{1, 3, 4, 2}, 3, true, []int{2, 3, 4}, []int{1})
	test(3, 4, []int{2, 3, 4, 1}, 3, false, []int{1, 2, 3}, []int{4})

	test(4, 4, []int{1, 2, 3, 4}, 4, false, []int{1, 2, 3, 4}, []int{})

	test(0, 4, []int{1, 1, 3, 4}, 0, true, []int{1}, []int{1, 3, 4})
	test(0, 4, []int{1, 1, 3, 4}, 1, true, []int{1, 1}, []int{3, 4})
	test(0, 4, []int{1, 1, 3, 4}, 2, true, []int{1, 1, 3}, []int{4})
	test(0, 4, []int{1, 1, 3, 4}, 3, true, []int{1, 1, 3, 4}, []int{})
	test(0, 4, []int{1, 1, 3, 4}, 4, true, []int{1, 1, 4}, []int{3})
	test(0, 4, []int{1, 1, 4, 3}, 3, true, []int{1, 3}, []int{1, 4})
	test(0, 4, []int{1, 3, 1, 4}, 2, true, []int{1, 3, 4}, []int{1})
	test(0, 4, []int{1, 3, 4, 1}, 3, true, []int{1, 4}, []int{1, 3})
	test(0, 4, []int{1, 4, 1, 3}, 2, true, []int{3}, []int{1, 1, 4})
	test(0, 4, []int{3, 1, 1, 4}, 1, true, []int{3, 4}, []int{1, 1})
	test(0, 4, []int{3, 4, 1, 1}, 2, true, []int{4}, []int{1, 1, 3})
	test(0, 4, []int{4, 1, 1, 3}, 1, false, []int{}, []int{1, 1, 3, 4})

	test(2, 4, []int{1, 1, 3, 4}, 2, true, []int{1, 1, 3}, []int{4})
	test(2, 4, []int{1, 1, 3, 4}, 3, true, []int{1, 1, 3, 4}, []int{})
	test(2, 4, []int{1, 1, 3, 4}, 4, true, []int{1, 1, 4}, []int{3})
	test(2, 4, []int{1, 1, 4, 3}, 3, true, []int{1, 3}, []int{1, 4})
	test(2, 4, []int{1, 3, 1, 4}, 2, true, []int{1, 3, 4}, []int{1})
	test(2, 4, []int{1, 3, 4, 1}, 3, true, []int{1, 4}, []int{1, 3})
	test(2, 4, []int{1, 4, 1, 3}, 2, true, []int{3, 4}, []int{1, 1})
	test(2, 4, []int{3, 4, 1, 1}, 2, false, []int{1, 1}, []int{3, 4})

	test(0, 4, []int{1, 1, 4, 4}, 0, true, []int{1}, []int{1, 4, 4})
	test(0, 4, []int{1, 1, 4, 4}, 1, true, []int{1, 1}, []int{4, 4})
	test(0, 4, []int{1, 1, 4, 4}, 2, true, []int{1, 1, 4}, []int{4})
	test(0, 4, []int{1, 1, 4, 4}, 3, true, []int{1, 1, 4, 4}, []int{})
	test(0, 4, []int{1, 1, 4, 4}, 4, true, []int{1, 4}, []int{1, 4})
	test(0, 4, []int{1, 4, 1, 4}, 2, true, []int{1, 4, 4}, []int{1})
	test(0, 4, []int{1, 4, 4, 1}, 3, true, []int{4}, []int{1, 1, 4})
	test(0, 4, []int{4, 1, 1, 4}, 1, true, []int{4, 4}, []int{1, 1})
	test(0, 4, []int{4, 4, 1, 1}, 2, false, []int{}, []int{1, 1, 4, 4})

	test(0, 4, []int{1, 1, 1, 4}, 0, true, []int{1}, []int{1, 1, 4})
	test(0, 4, []int{1, 1, 1, 4}, 1, true, []int{1, 1}, []int{1, 4})
	test(0, 4, []int{1, 1, 1, 4}, 2, true, []int{1, 1, 1}, []int{4})
	test(0, 4, []int{1, 1, 1, 4}, 3, true, []int{1, 1, 1, 4}, []int{})
	test(0, 4, []int{1, 1, 1, 4}, 4, true, []int{1, 1, 4}, []int{1})
	test(0, 4, []int{1, 1, 4, 1}, 3, true, []int{1, 4}, []int{1, 1})
	test(0, 4, []int{1, 4, 1, 1}, 2, true, []int{4}, []int{1, 1, 1})
	test(0, 4, []int{4, 1, 1, 1}, 1, false, []int{}, []int{1, 1, 1, 4})

	test(0, 4, []int{1, 1, 1, 1}, 0, true, []int{1}, []int{1, 1, 1})
	test(0, 4, []int{1, 1, 1, 1}, 1, true, []int{1, 1}, []int{1, 1})
	test(0, 4, []int{1, 1, 1, 1}, 2, true, []int{1, 1, 1}, []int{1})
	test(0, 4, []int{1, 1, 1, 1}, 3, true, []int{1, 1, 1, 1}, []int{})
	test(0, 4, []int{1, 1, 1, 1}, 4, false, []int{}, []int{1, 1, 1, 1})

	test(1, 1, []int{1, 1, 1, 1}, 1, false, []int{1}, []int{1, 1, 1})

	test(1, 2, []int{1, 1, 1, 1}, 1, true, []int{1, 1}, []int{1, 1})
	test(1, 2, []int{1, 1, 1, 1}, 2, false, []int{1}, []int{1, 1, 1})

	test(1, 3, []int{1, 1, 1, 1}, 1, true, []int{1, 1}, []int{1, 1})
	test(1, 3, []int{1, 1, 1, 1}, 2, true, []int{1, 1, 1}, []int{1})
	test(1, 3, []int{1, 1, 1, 1}, 3, false, []int{1}, []int{1, 1, 1})

	test(1, 4, []int{1, 1, 1, 1}, 1, true, []int{1, 1}, []int{1, 1})
	test(1, 4, []int{1, 1, 1, 1}, 2, true, []int{1, 1, 1}, []int{1})
	test(1, 4, []int{1, 1, 1, 1}, 3, true, []int{1, 1, 1, 1}, []int{})
	test(1, 4, []int{1, 1, 1, 1}, 4, false, []int{1}, []int{1, 1, 1})

	if pc != 0 || nc != 0 {
		t.Errorf("fail count: prev(%d), next(%d)", pc, nc)
	}
}

func TestSlideCombiInts2(t *testing.T) {
	test := func(a []int) {
		ps := make([]*set, 0, len(a)+1)
		for i := 0; i <= len(a); i++ {
			c := sliceCp(a)
			//fmt.Println("gen PartCombiNextInts: ", a, i)
			s := genSetInts(func() ([]int, []int, bool) {
				f := sliceCp(c[:i])
				b := sliceCp(c[i:])
				return f, b, PartCombiNextInts(c, i)
			}, fmt.Sprintf("PartCombiNextInts(%v, %d)", c, i))

			checkOrder(s, false, t)
			ps = append(ps, s)
		}

		for i := 0; i <= len(a); i++ {
			for j := i; j <= len(a); j++ {
				c := sliceCp(a)
				p := i
				//fmt.Println("gen SliceCombiNextInts: ", a, i, j)
				s := genSetInts(func() ([]int, []int, bool) {
					f := sliceCp(c[:p])
					b := sliceCp(c[p:])
					r := false
					p, r = SlideCombiNextInts(i, j, c, p)
					return f, b, r
				}, fmt.Sprintf("SlideCombiNextInts(%d, %d, %v, %d)", i, j, c, i))

				checkOrder(s, false, t)
				m := make([][]int, 0, len(s.selected))
				for k := i; k <= j; k++ {
					m = append(m, ps[k].selected...)
				}

				checkMatch(s, m, t)

				c = sliceCp(a)
				p, _ = SlideCombiPrevInts(i, j, c, i)
				//fmt.Println("gen SliceCombiPrevInts: ", c, i, j)
				sr := genSetInts(func() ([]int, []int, bool) {
					f := sliceCp(c[:p])
					b := sliceCp(c[p:])
					r := false
					p, r = SlideCombiPrevInts(i, j, c, p)
					//fmt.Println("gen SliceCombiPrevInts: ", f, b, p, r)
					return f, b, r
				}, fmt.Sprintf("SlideCombiPrevInts(%d, %d, %v, %d)", i, j, c, i))

				checkReverse(sr, s, t)
			}
		}
	}

	quicktest, longtest, forevertest := 5, 7, 9
	d := []int{quicktest, longtest, forevertest}
	cases := ChanZeroToNumber(d[0])

	var wg sync.WaitGroup
	pass, until := true, []int{}
	for c := range cases {
		if sliceEq(c, until) {
			pass = false
		}

		if pass {
			continue
		}

		if len(c) == len(until)+1 {
			fmt.Println(c)
		}

		wg.Add(1)
		go func(c []int) {
			defer wg.Done()
			test(c)
		}(c)
	}

	wg.Wait()
}
