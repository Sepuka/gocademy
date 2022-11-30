package combinatorics

func climbStairs(n int) int {
	if n < 2 {
		return n
	}

	var ways = 1
	var prev = 1
	var temp = 0
	for i := 1; i < n; i++ {
		temp = ways
		ways = prev + ways
		prev = temp
	}

	return ways
}
