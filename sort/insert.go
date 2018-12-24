package sort

// Пример наивной реализации устойчивой сортировки
func dummy(in []uint32) []uint32 {
	var j int
	for pos, el := range in {
		for j = pos; j > 0 && in[j-1] > el; j-- {
			in[j] = in[j-1]
		}
		in[j] = el
	}

	return in
}

// Улучшение состоит в бинарном поиске подходящей позиции для нового элемента
func improvedDummy(in []uint32) []uint32 {
	var left, right, med, j int

	for pos, el := range in {
		left = 0
		right = pos

		// Поиск делением пополам нужно использовать начиная с определенной
		// длины левой, сортированной, части
		// У меня сейчас он выигрывает начиная с +- 85к элементов
		for left < right {
			med = (left + right) >> 1
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
