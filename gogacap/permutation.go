package gogacap

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

	reverseInts(ints[b:])
	if b == 0 {
		return false
	}

	i := b + upperBoundInts(ints[b:], ints[b-1])
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
		i := s + lowerBoundInts(ints[s:], ints[s-1])
		ints[s-1], ints[i-1] = ints[i-1], ints[s-1]
	}

	reverseInts(ints[s:])
	return s > 0
}

func PartPermuNextInt(ints []int, c int) bool {
	reverseInts(ints[c:])
	return PermuNextInt(ints)
}

func PartPermuPrevInt(ints []int, c int) bool {
	r := PermuPrevInt(ints)
	reverseInts(ints[c:])
	return r
}

func SlidePermuNextInt(min, max int, ints []int, c int) (int, bool) {
	if len(ints) == 0 {
		return 0, false
	}

	if min < 0 {
		min = 0
	}

	if max > len(ints) {
		max = len(ints)
	}

	if min > max {
		min, max = max, min
	}

	if c < max {
		return c + 1, true
	}

	reverseInts(ints[max:])

	b := max
	if b == len(ints) {
		b--
	}

	for b > 0 {
		p := b - 1
		if ints[p] < ints[b] {
			break
		} else {
			b = p
		}
	}

	reverseInts(ints[b:])
	if b == 0 {
		return min, false
	}

	i := b + upperBoundInts(ints[b:], ints[b-1])
	ints[b-1], ints[i] = ints[i], ints[b-1]

	if b < min {
		return min, true
	} else {
		return b, true
	}
}

func SlidePermuPrevInt(min, max int, ints []int, c int) (int, bool) {
	if len(ints) == 0 {
		return 0, false
	}

	if min < 0 {
		min = 0
	}

	if max > len(ints) {
		max = len(ints)
	}

	if min > max {
		min, max = max, min
	}

	s := max
	if s == len(ints) {
		s--
	}

	for s > 0 {
		p := s - 1
		if ints[s] < ints[p] {
			break
		} else {
			s = p
		}
	}

	if min <= s && s < c {
		return c - 1, true
	}

	if s > 0 {
		i := s + lowerBoundInts(ints[s:], ints[s-1])
		ints[s-1], ints[i-1] = ints[i-1], ints[s-1]
	}

	reverseInts(ints[s:])
	reverseInts(ints[max:])
	return max, s > 0
}
