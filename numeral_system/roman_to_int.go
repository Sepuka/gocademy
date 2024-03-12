package numeral_system

var romanMap = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	var r = 0
	var i = len(s) - 1
	var prev = 0

	for i >= 0 {
		var v = romanMap[s[i]]
		if r == 0 {
			r = v
		} else {
			if v < prev {
				r -= v
			} else {
				r += v
			}
		}
		prev = v
		i--
	}

	return r
}
