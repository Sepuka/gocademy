package sort

func simpleChoice(in []uint32) []uint32 {
	var minValue uint32
	var minPos int
	for pos, el := range in {
		minValue = el
		minPos = pos
		for cur := pos + 1; cur <= len(in)-1; cur++ {
			if minValue > in[cur] {
				minValue = in[cur]
				minPos = cur
			}
		}
		in[pos] = minValue
		in[minPos] = el
	}

	return in
}
