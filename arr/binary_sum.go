package arr

func binarySum(a string, b string) string {
	var firstLn = int32(len(a))
	var secondLn = int32(len(b))
	var result []byte
	var over bool
	var longest, shortest int32
	var long, short string
	if firstLn > secondLn {
		result = make([]byte, firstLn+1)
		longest = firstLn
		shortest = secondLn
		long = a
		short = b
	} else {
		result = make([]byte, secondLn+1)
		longest = secondLn
		shortest = firstLn
		long = b
		short = a
	}

	for i := int32(0); i < longest; i++ {
		var longKey = longest - i - 1
		var shortKey = shortest - i - 1
		var resultKey = longKey + 1
		if shortKey > -1 {
			switch long[longKey] {
			case '0':
				if short[shortKey] == '0' {
					if over {
						result[resultKey] = '1'
						over = false
					} else {
						result[resultKey] = '0'
					}
				} else {
					if over {
						result[resultKey] = '0'
					} else {
						result[resultKey] = '1'
					}
				}
			case '1':
				if short[shortKey] == '0' {
					if over {
						result[resultKey] = '0'
					} else {
						result[resultKey] = '1'
					}
				} else {
					if over {
						result[resultKey] = '1'
					} else {
						result[resultKey] = '0'
						over = true
					}
				}
			}
		} else {
			if long[longKey] == '1' {
				if over {
					result[resultKey] = '0'
				} else {
					result[resultKey] = '1'
				}
			} else {
				if over {
					result[resultKey] = '1'
					over = false
				} else {
					result[resultKey] = '0'
				}
			}
		}
	}

	if over {
		result[0] = '1'
	} else {
		return string(result[1:])
	}

	return string(result)
}
