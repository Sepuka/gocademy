package arr

func canJump(nums []int) bool {
	return jump(nums, make(map[int]bool))
}

func jump(nums []int, positions map[int]bool) bool {
	if len(nums) < 2 {
		return true
	}

	var pointer = 0
	var jumpSize int
	var lastIndex = len(nums) - 1
	var testingSlice []int

	jumpSize = nums[pointer]
	if pointer+jumpSize >= lastIndex {
		return true
	}

	for jumpSize > 0 {
		if nums[pointer+jumpSize] > 0 {
			testingSlice = nums[pointer+jumpSize:]
			if _, ok := positions[len(testingSlice)]; ok {
				jumpSize--
			} else {
				if jump(testingSlice, positions) == false {
					positions[len(testingSlice)] = true
					jumpSize--
				} else {
					return true
				}
			}
		} else {
			jumpSize--
		}
	}

	return false
}
