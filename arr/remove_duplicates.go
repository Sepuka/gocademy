package arr

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	var pos = 0
	for i, v := range nums {
		if i == 0 {
			continue
		}
		if v == nums[i-1] {
			pos = i
			break
		}
	}

	if pos == 0 {
		return len(nums)
	}

	for _, v := range nums[pos+1:] {
		if v != nums[pos-1] {
			nums[pos] = v
			pos++
		}
	}

	return pos
}
