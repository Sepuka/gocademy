package sort

func shell(in []uint32) []uint32 {
	var x uint32
	var length, step, i, j int
	length = len(in)

	for step = length/2; step > 0; step /= 2 {
		for i = step; i < length; i++ {
			x = in[i]
			j = i - step

			for j >= step && x < in[j] {
				in[j+step] = in[j]
				j -= step
			}

			if j >= step || x >= in[j] {
				in[j+step] = x
			} else {
				in[j+step] = in[j]
				in[j] = x
			}
		}
	}

	return in
}
