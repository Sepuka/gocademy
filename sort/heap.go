package sort

func heap(in []uint32) []uint32 {
	var l, r int

	l = len(in) / 2
	r = len(in) - 1

	for l > 0 {
		l--
		sift(in, l, len(in) - 1)
	}

	for r > 0 {
		in[0], in[r] = in[r], in[0]
		r--
		sift(in, l, r)
	}

	return in
}

func sift(in []uint32, l, r int) {
	var i, j int
	var x uint32

	i = l
	j = 2*i + 1
	x = in[i]

	if j < r && in[j + 1] > in[j] {
		j++
	}

	for j <= r && in[j] > x {
		in[i] = in[j]
		i = j
		j = 2*j + 1

		if j < r && in[j + 1] > in[j] {
			j++
		}
	}

	in[i] = x
}
