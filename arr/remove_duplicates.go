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

func removeDuplicates2(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}

	var pos = 0

	for i, v := range nums {
		if i < 2 {
			continue
		}
		if v == nums[i-1] && v == nums[i-2] {
			pos = i
			break
		}
	}

	if pos == 0 {
		return len(nums)
	}

	var double = true

	for _, v := range nums[pos:] {
		if v != nums[pos-1] {
			double = false
			nums[pos] = v
			pos++
			continue
		}
		if v == nums[pos-1] && double == false {
			double = true
			nums[pos] = v
			pos++
		}
	}

	return pos
}
