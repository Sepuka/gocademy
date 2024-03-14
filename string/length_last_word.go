package string

func lengthOfLastWord(s string) int {
	var p = len(s) - 1
	var space = ' '
	var offset = 0

	for p >= 0 {
		if rune(s[p]) == space {
			p--
		} else {
			offset = len(s) - p - 1
			break
		}
	}

	for p >= 0 {
		if rune(s[p]) != space {
			p--
			continue
		}

		return len(s) - p - 1 - offset
	}

	return len(s) - offset
}
