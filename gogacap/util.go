package gogacap

func reverseInts(ints []int) {
	e := len(ints)
	b := 0
	for b+1 < e {
		e--
		ints[b], ints[e] = ints[e], ints[b]
		b++
	}
}
