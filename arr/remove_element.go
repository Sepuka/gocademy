package arr

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return len(nums)
	}

	var i = 0
	var rightPointer = len(nums) - 1

	for i <= rightPointer {
		if nums[i] == val {
			for nums[rightPointer] == val {
				if rightPointer == i {
					return i
				}
				rightPointer--
				if rightPointer == i {
					return rightPointer
				}
			}

			nums[i] = nums[rightPointer]
			rightPointer--
		}
		i++
	}

	return rightPointer + 1
}
