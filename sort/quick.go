package sort

func quick(in []uint32) []uint32 {
	return sort(in, 0, len(in)-1)
}

func sort(in []uint32, l, r int) []uint32 {
	var i, j, pos int
	var x uint32

	i, j = l, r
	pos = (l + r) >> 1
	x = in[pos]

	do := true
	for do {
		for in[i] < x {
			i++
		}
		for x < in[j] {
			j--
		}
		if i <= j {
			in[j], in[i] = in[i], in[j]
			i++
			j--
		}
		do = j > i
	}

	if l < j {
		sort(in, l, j)
	}

	if i < r {
		sort(in, i, r)
	}

	return in
}
