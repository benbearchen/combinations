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

func inplaceMergeInts(a, b []int) {
	c := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			c = append(c, a[i])
			i++
		} else {
			c = append(c, b[j])
			j++
		}
	}

	if i < len(a) {
		c = append(c, a[i:]...)
	}

	if j < len(b) {
		c = append(c, b[j:]...)
	}

	copy(a, c[:len(a)])
	copy(b, c[len(a):])
}

func minMax(min, max, size int) (int, int) {
	if min < 0 {
		min = 0
	}

	if max > size {
		max = size
	}

	if min > max {
		min, max = max, min
	}

	return min, max
}

func rotateShiftRightOneInts(ints []int) {
	if len(ints) <= 1 {
		return
	}

	c := len(ints)
	v := ints[c-1]
	for i := c - 1; i > 0; i-- {
		ints[i] = ints[i-1]
	}

	ints[0] = v
}

func rotateShiftLeftOneInts(ints []int) {
	if len(ints) <= 1 {
		return
	}

	c := len(ints)
	v := ints[0]
	for i := 1; i < c; i++ {
		ints[i-1] = ints[i]
	}

	ints[c-1] = v
}

func rotateInts(ints []int, c int) {
	reverseInts(ints[:c])
	reverseInts(ints[c:])
	reverseInts(ints)
}

func rotateBackInts(ints []int, c int) {
	reverseInts(ints)
	reverseInts(ints[:c])
	reverseInts(ints[c:])
}
