package gogacap

import (
	"testing"
)

import (
	"fmt"
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
