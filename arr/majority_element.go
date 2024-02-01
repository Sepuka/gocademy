package arr

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	var cnt = map[int]int{}
	for _, el := range nums {
		v, ok := cnt[el]
		if ok {
			if v+1 > len(nums)/2 {
				return el
			}
			cnt[el]++
		} else {
			cnt[el] = 1
		}
	}

	return 0
}
