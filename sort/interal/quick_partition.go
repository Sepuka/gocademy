package interal

func Partition(in []uint32, x uint32, i, j *int) {
	for *i < *j {
		for in[*i] < x {
			*i++
		}
		for x < in[*j] {
			*j--
		}
		if *i <= *j {
			in[*j], in[*i] = in[*i], in[*j]
			*i++
			*j--
		}
	}
}
