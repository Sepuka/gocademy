package arr

func merge(nums1 []int, m int, nums2 []int, n int) {
	var pos1 int
	var pos2 int
	var result []int

	for pos2 < n {
		if nums1[pos1] > nums2[pos2] || pos1 == m {
			result = append(result, nums2[pos2])
			pos2++
		} else {
			result = append(result, nums1[pos1])
			pos1++
		}
	}

	for pos1 < m {
		result = append(result, nums1[pos1])
		pos1++
	}

	for i, r := range result {
		nums1[i] = r
	}
}
