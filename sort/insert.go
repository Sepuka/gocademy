package sort

// Пример наивной реализации устойчивой сортировки
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

// Улучшение состоит в бинарном поиске подходящей позиции для нового элемента
func improvedDummy(in []int) []int {
	var left, right, med, j int

	for pos, el := range in {
		if pos == 0 {
			continue
		}
		left = 0
		right = pos

		for left < right {
			med = (left + right) / 2
			if in[med] <= el {
				left = med + 1
			} else {
				right = med
			}
		}

		for j = pos; j >= right+1; j-- {
			in[j] = in[j-1]
		}

		in[right] = el
	}

	return in
}
