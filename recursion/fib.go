package recursion

func FibRecursion(n uint) uint {
	if n < 2 {
		return n
	}

	return FibRecursion(n - 1) + FibRecursion(n - 2)
}
