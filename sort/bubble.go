package sort

func bubble(in []uint32) []uint32 {
	var j int

	for pos := range in {
		for j = pos; j > 0; j-- {
			if in[j-1] > in[j] {
				in[j-1], in[j] = in[j], in[j-1]
			} else {
				break
			}
		}
	}

	return in
}
