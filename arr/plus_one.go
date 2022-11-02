package arr

func onePlus(digits []int) []int {
	var ln = len(digits)
	for i, _ := range digits {
		var key = ln - i - 1
		digits[key]++
		if digits[key] < 10 {
			return digits
		}

		digits[key] = 0
	}

	return append([]int{1}, digits...)
}
