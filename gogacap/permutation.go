package gogacap

import (
	"sort"
)

func PermuNextInt(ints []int) bool {
	if len(ints) <= 1 {
		return false
	}

	// find the biggest item from end
	b := len(ints) - 1
	for b > 0 {
		p := b - 1
		if ints[p] < ints[b] {
			break
		} else {
			b = p
		}
	}

	sort.Ints(ints[b:]) // TODO: real reverse
	if b == 0 {
		return false
	}

	v := ints[b-1]
	i := b + sort.SearchInts(ints[b:], v)
	for ints[i] == v {
		i++
	}

	ints[b-1], ints[i] = ints[i], ints[b-1]
	return true
}

func PermuPrevInt(ints []int) bool {
	if len(ints) <= 1 {
		return false
	}

	s := len(ints) - 1
	for s > 0 {
		p := s - 1
		if ints[s] < ints[p] {
			break
		} else {
			s = p
		}
	}

	if s > 0 {
		v := ints[s-1]
		i := s + sort.SearchInts(ints[s:], v)
		for ints[i-1] == v {
			i--
		}

		ints[s-1], ints[i-1] = ints[i-1], ints[s-1]
	}

	sort.Sort(sort.Reverse(sort.IntSlice(ints[s:]))) // TODO: real reverse
	return s > 0
}
