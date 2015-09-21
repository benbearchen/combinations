package gogacap

func ChanZeroToNumber(n int) chan []int {
	zero := make(chan []int)
	go func() {
		zero <- []int{}
		close(zero)
	}()

	input := zero
	for i := 0; i <= n; i++ {
		next := make(chan []int)
		go func(i, n int, input, output chan []int) {
			defer close(output)
			for a := range input {
				output <- a

				for j := len(a); j <= n; j++ {
					s := make([]int, len(a)+1)
					copy(s, a)
					s[len(a)] = i

					a = s
					output <- a
				}
			}
		}(i, n, input, next)

		input = next
	}

	return input
}
