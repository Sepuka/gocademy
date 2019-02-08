package recursion

func FibRecursion(n uint) uint {
	if n < 2 {
		return n
	}

	return FibRecursion(n - 1) + FibRecursion(n - 2)
}

func FibIter(n uint) uint {
	var x, i, p uint

	if n < 2 {
		return n
	}

	x = 1
	p = 0
	for i = 1; i < n; i++ {
		x, p = x+p, x
	}

	return x
}