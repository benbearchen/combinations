package gogacap

import (
	"sort"
)

func reverseInts(ints []int) {
	e := len(ints)
	b := 0
	for b+1 < e {
		e--
		ints[b], ints[e] = ints[e], ints[b]
		b++
	}
}

func upperBoundInts(ints []int, x int) int {
	i := sort.SearchInts(ints, x)
	for i < len(ints) && ints[i] == x {
		i++
	}

	return i
}

func lowerBoundInts(ints []int, x int) int {
	i := sort.SearchInts(ints, x)
	for i > 0 && ints[i-1] == x {
		i--
	}

	return i
}
