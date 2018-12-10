package sort

func dummy(in []int) []int {
	for pos, el := range in {
		j := pos
		for j > 0 && in[j-1] > el {
			in[j] = in[j-1]
			j--
		}
		in[j] = el
	}

	return in
}
