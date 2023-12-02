package arr

// Ищет в последовательности чисел наиболее выгодные позиции покупки и продажи, с максимальным профитом
func maxProfit(prices []int) int {
	var min = prices[0]
	var max = 0
	var minTmp = min
	var diff = max - min

	for _, el := range prices[1:] {
		if el < minTmp {
			minTmp = el
			continue
		}

		if el-minTmp > diff {
			min = minTmp
			max = el
			diff = max - min
		}
	}

	if diff > 0 {
		return diff
	} else {
		return 0
	}
}
