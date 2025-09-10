package list

import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	var ranges []string

	pair := make([]int, 0, 2)
	for _, v := range nums {
		if len(pair) == 0 {
			pair = append(pair, v)
			continue
		}

		if v-pair[len(pair)-1] == 1 {
			if len(pair) == 1 {
				pair = append(pair, v)
			} else {
				pair[len(pair)-1] = v
			}
		} else {
			if len(pair) == 2 {
				ranges = append(ranges, fmt.Sprintf("%d->%d", pair[0], pair[1]))
			} else {
				ranges = append(ranges, fmt.Sprintf("%d", pair[0]))
			}
			pair = pair[:0]
			pair = append(pair, v)
		}
	}

	if len(pair) == 2 {
		ranges = append(ranges, fmt.Sprintf("%d->%d", pair[0], pair[1]))
	} else {
		ranges = append(ranges, fmt.Sprintf("%d", pair[0]))
	}

	return ranges
}
