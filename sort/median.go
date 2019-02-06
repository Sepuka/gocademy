package sort

import "github.com/sepuka/gocademy/sort/interal"

func median(in []uint32) uint32 {
	var (
		l, r, i, j, pivot int
		x                 uint32
	)

	l = 0
	r = len(in) - 1
	pivot = (l + r) >> 1
	do := true

	for do {
		i, j = l, r
		x = in[pivot]
		interal.Partition(in, x, &i, &j)
		if j < pivot {
			l = i
		}
		if pivot < i {
			r = j
		}
		do = l < r
	}

	return x
}
