package arr

func rotate(nums []int, k int) {
	var length = len(nums)
	var i int

	k = k % len(nums)

	var buff = make([]int, k)

	for i < k {
		buff[k-i-1] = nums[length-i-1]
		i++
	}

	i = 0
	var pos = length - k - 1
	for pos > -1 {
		nums[length-i-1] = nums[pos]
		pos--
		i++
	}

	copy(nums, buff)
}
