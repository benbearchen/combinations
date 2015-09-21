package gogacap

import (
	"testing"
)

func TestChanZeroToNumber(t *testing.T) {
	aa := [][]int{
		[]int{},
	}

	c := ChanZeroToNumber(-1)
	for _, a := range aa {
		b := <-c
		if !sliceEq(a, b) {
			t.Errorf("%v != %v", a, b)
		}
	}

	aa = [][]int{
		[]int{},
		[]int{0},
	}

	c = ChanZeroToNumber(0)
	for _, a := range aa {
		b := <-c
		if !sliceEq(a, b) {
			t.Errorf("%v != %v", a, b)
		}
	}

	aa = [][]int{
		[]int{},
		[]int{1},
		[]int{1, 1},
		[]int{0},
		[]int{0, 1},
		[]int{0, 0},
	}

	c = ChanZeroToNumber(1)
	for _, a := range aa {
		b := <-c
		if !sliceEq(a, b) {
			t.Errorf("%v != %v", a, b)
		}
	}

	aa = [][]int{
		[]int{},
		[]int{2},
		[]int{2, 2},
		[]int{2, 2, 2},
		[]int{1},
		[]int{1, 2},
		[]int{1, 2, 2},
		[]int{1, 1},
		[]int{1, 1, 2},
		[]int{1, 1, 1},
		[]int{0},
		[]int{0, 2},
		[]int{0, 2, 2},
		[]int{0, 1},
		[]int{0, 1, 2},
		[]int{0, 1, 1},
		[]int{0, 0},
		[]int{0, 0, 2},
		[]int{0, 0, 1},
		[]int{0, 0, 0},
	}

	c = ChanZeroToNumber(2)
	for _, a := range aa {
		b := <-c
		if !sliceEq(a, b) {
			t.Errorf("%v != %v", a, b)
		}
	}

	c = ChanZeroToNumber(10)
	for range c {
	}
}
