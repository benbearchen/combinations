package gogacap

func combinationNextInts(ss, us []int) bool {
	if len(ss) == 0 || len(us) == 0 {
		return false
	}

	s := lowerBoundInts(ss, us[len(us)-1])
	u := 0
	if s > 0 {
		p := s - 1
		b := upperBoundInts(us, ss[p])
		ss[p], us[b] = us[b], ss[p]
		u = b + 1
	}

	inplaceMergeInts(ss[s:], us[u:])
	return s > 0
}

func PartCombiNextInts(ints []int, c int) bool {
	return combinationNextInts(ints[:c], ints[c:])
}

func PartCombiPrevInts(ints []int, c int) bool {
	return combinationNextInts(ints[c:], ints[:c])
}

func SlideCombiNextInts(min, max int, ints []int, c int) (int, bool) {
	length := len(ints)
	if length == 0 {
		return 0, false
	}

	min, max = minMax(min, max, length)
	if c == 0 {
		if c < max {
			return c + 1, true
		} else {
			return c, false
		}
	}

	if c < max {
		b := c + lowerBoundInts(ints[c:], ints[c-1])
		if b != length {
			rotateShiftRightOneInts(ints[c : b+1])
			return c + 1, true
		}
	}

	if min < c {
		e := c - 1
		p := lowerBoundInts(ints[:e], ints[e])
		if p > 0 {
			u := p - 1
			b := c + upperBoundInts(ints[c:], ints[u])
			if b == length {
				ints[u], ints[e] = ints[e], ints[u]
				m := p
				if m < min {
					m = min
				}

				reverseInts(ints[c:])
				reverseInts(ints[m:])
				return m, m > 0
			}
		}
	}

	s := 0
	if c < length {
		s = lowerBoundInts(ints[:c], ints[length-1])
	}

	if s > 0 {
		p := s - 1
		b := c + upperBoundInts(ints[c:], ints[p])
		ints[p], ints[b] = ints[b], ints[p]
		reverseInts(ints[s:c])
		reverseInts(ints[c:])
		e := length - (b + 1 - c)
		reverseInts(ints[s:e])
		m := s
		if m < min {
			m = min
		}

		reverseInts(ints[m:e])
		reverseInts(ints[m:])
		return m, true
	}

	inplaceMergeInts(ints[:c], ints[c:])
	return min, false
}

func SlideCombiPrevInts(min, max int, ints []int, c int) (int, bool) {
	length := len(ints)
	if length == 0 {
		return 0, false
	}

	min, max = minMax(min, max, length)
	e := 0
	if c == 0 {
	} else if c == length {
		if min < c {
			return c - 1, true
		} else {
			return c, false
		}
	} else {
		b := c + lowerBoundInts(ints[c:], ints[c-1])
		if b != c {
			if c == 1 || ints[c-2] <= ints[b-1] {
				ints[c-1], ints[b-1] = ints[b-1], ints[c-1]
				e = c
			} else if min < c {
				rotateShiftLeftOneInts(ints[c-1 : b])
				return c - 1, true
			} else {
				p := lowerBoundInts(ints[:c-1], ints[b-1])
				ints[p], ints[b-1] = ints[b-1], ints[p]
				e = p + 1
				rotateInts(ints[e:b], c-e)
			}
		} else if min < c {
			return c - 1, true
		}
	}

	m := e + (length - lowerBoundInts(ints, ints[length-1]))
	if m > max {
		m = max
	} else if m < min {
		m = min
	}

	rotateBackInts(ints[e:], m-e)
	return m, e > 0
}
