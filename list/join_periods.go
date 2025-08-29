package list

import (
	"fmt"
	"sort"
)

func JoinPeriods(src []int) string {
	sort.Ints(src)
	if len(src) == 0 {
		return ""
	}

	if len(src) == 1 {
		return fmt.Sprintf("%d", src[0])
	}

	if len(src) == 2 {
		return fmt.Sprintf("%d-%d", src[0], src[1])
	}

	var periods [][]int

	period := []int{src[0], src[1]}

	for _, i := range src[2:] {
		if i-period[1] > 1 {
			periods = append(periods, period)
			period = []int{i, i}
		} else {
			period[1] = i
		}
	}

	periods = append(periods, period)

	result := ""
	for _, i := range periods {
		if i[0] == i[1] {
			result += fmt.Sprintf("%d,", i[0])
		} else {
			result += fmt.Sprintf("%d-%d,", i[0], i[1])
		}
	}

	return result[:len(result)-1]
}
