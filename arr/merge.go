package arr

func merge(nums1 []int, m int, nums2 []int, n int) {
	var pos0 int
	var pos1 int
	var pos2 int
	var length = 0
	if len(nums1) > len(nums2) {
		length = len(nums1)
	} else {
		length = len(nums2)
	}
	var result = make([]int, length)

	for pos2 < n {
		if nums1[pos1] > nums2[pos2] || pos1 == m {
			result[pos0] = nums2[pos2]
			pos0++
			pos2++
		} else {
			result[pos0] = nums1[pos1]
			pos0++
			pos1++
		}
	}

	for pos1 < m {
		result[pos0] = nums1[pos1]
		pos0++
		pos1++
	}

	for i, r := range result {
		nums1[i] = r
	}
}
