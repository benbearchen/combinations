package gogacap

func combinationNextInts(ss, us []int) bool {
	if len(ss) == 0 || len(us) == 0 {
		return false
	}

	s := lowerBoundInts(ss, us[len(us)-1])
	u := 0
	if s > 0 {
		c := s - 1
		b := upperBoundInts(us, ss[c])
		ss[c], us[b] = us[b], ss[c]
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
