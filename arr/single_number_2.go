package arr

import "math"

func singleNumber(nums []int) int {
	var length = math.Ceil(float64(len(nums)) / float64(3))
	var store = make(map[int]int, int(length))

	for _, n := range nums {
		value, ok := store[n]
		if ok {
			store[n] = value - 1
		} else {
			store[n] = 2
		}
	}

	for k, v := range store {
		if v > 0 {
			return k
		}
	}

	return 0
}
