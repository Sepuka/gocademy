package string

func reverseWords(s string) string {
	var stack string
	var result string

	for i := 0; i < len(s); i++ {
		if s[i] != 32 {
			stack = string(s[i]) + stack
		} else {
			result = result + stack + " "
			stack = ""
		}
	}

	return result + stack
}
