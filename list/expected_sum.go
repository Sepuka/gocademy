package list

func ContainsSum(l []int, s int) bool {
	if len(l) == 0 {
		return false
	}

	acc := make([]int, len(l))
	acc = append(acc, l[0])

	for _, e := range l[1:] {
		for _, i := range acc {
			if e+i == s {
				return true
			}

			acc = append(acc, e)
		}
	}

	return false
}
