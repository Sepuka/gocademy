package search

func isSubsequence(s, t []rune) bool {
	var e int
	for i, v := range s {
		for e < len(t) {
			if v == t[e] {
				e += 1
				break
			}
			e += 1
		}
		if e == len(t) {
			return i == len(s)-1
		}
	}

	return true
}
