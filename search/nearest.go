package search

func nearest(l []int, i int, n int) []int {
	startIdx := i - n + 1
	if startIdx < 0 {
		startIdx = 0
	}
	bestIdx := startIdx

	diff := func(s []int) (r int) {
		r = s[len(s)-1]
		for _, v := range s[:len(s)-1] {
			r -= v
		}

		return
	}

	minimum := diff(l[startIdx : startIdx+n])

	for startIdx <= i && len(l) > startIdx+n {
		startIdx += 1
		c := diff(l[startIdx : startIdx+n])
		if c < minimum {
			minimum = c
			bestIdx = startIdx
		}
	}

	return l[bestIdx : bestIdx+n]
}
