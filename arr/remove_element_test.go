package arr

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestRemoveElement(t *testing.T) {
	var cases = []struct {
		nums         []int
		val          int
		expectedLen  int
		expectedNums []int
	}{
		{
			nums:         []int{},
			val:          0,
			expectedLen:  0,
			expectedNums: []int{},
		},
		{
			nums:         []int{4, 5},
			val:          5,
			expectedLen:  1,
			expectedNums: []int{4},
		},
		{
			nums:         []int{3, 2, 2, 3},
			val:          3,
			expectedLen:  2,
			expectedNums: []int{2, 2},
		},
		{
			nums:         []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:          2,
			expectedLen:  5,
			expectedNums: []int{0, 1, 4, 0, 3},
		},
	}

	for _, testCase := range cases {
		assert.Equal(t, removeElement(testCase.nums, testCase.val), testCase.expectedLen)
		assert.Equal(t, testCase.nums[0:testCase.expectedLen], testCase.expectedNums)
	}
}
