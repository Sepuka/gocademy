package arr

func getIntersection(input [][]int) [][]int {
	var result [][]int
	var start, end int
	var i int

	for i < len(input)-1 {
		if hasIntersect(input[i], input[i+1]) {
			if input[0][0] < input[1][0] {
				start = input[0][0]
			} else {
				start = input[1][0]
			}

			if input[0][1] > input[1][1] {
				end = input[0][1]
			} else {
				end = input[1][1]
			}

			result = append(result, []int{start, end})
			i++
		} else {
			result = append(result, input[i])
			if i == len(input)-2 {
				result = append(result, input[i+1])
			}
		}
		i++
	}

	return result
}

func hasIntersect(a []int, b []int) bool {
	var aStart = a[0]
	var aEnd = a[1]
	var bStart = b[0]
	var bEnd = b[1]

	return aStart <= bEnd && aEnd >= bStart
}
