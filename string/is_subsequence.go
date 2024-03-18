package string

func isSubsequence(s string, t string) bool {
	var pos_t int
	var found = false

	for i := range s {
		for pos_t < len(t) {
			if s[i] == t[pos_t] {
				pos_t++
				found = true
				break
			}
			pos_t++
		}

		if !found {
			return false
		}

		if i == len(s)-1 {
			return true
		}

		found = false
	}

	return true
}
