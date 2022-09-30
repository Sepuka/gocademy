package linkedList

import (
	"github.com/magiconair/properties/assert"
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestAddTwoLists(t *testing.T) {
	var cases = map[string]struct {
		lst1        []int
		lst2        []int
		expectedLst []int
	}{
		`long sequence`: {
			lst1:        []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			lst2:        []int{5, 6, 4},
			expectedLst: []int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		},
		`short sequence`: {
			lst1:        []int{2, 4, 3},
			lst2:        []int{5, 6, 4},
			expectedLst: []int{7, 0, 8},
		},
		`big digits`: {
			lst1:        []int{9, 9, 9, 9, 9, 9, 9},
			lst2:        []int{9, 9, 9, 9},
			expectedLst: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
		`two hundred`: {
			lst1:        []int{9, 9, 1},
			lst2:        []int{1},
			expectedLst: []int{0, 0, 2},
		},
		`simple`: {
			lst1:        []int{5},
			lst2:        []int{5},
			expectedLst: []int{0, 1},
		},
		`middle overflow`: {
			lst1:        []int{2, 8, 9, 9, 9, 9, 8, 9, 9, 9},
			lst2:        []int{8, 1, 2},
			expectedLst: []int{0, 0, 2, 0, 0, 0, 9, 9, 9, 9},
		},
	}

	for _, testCase := range cases {
		var list1 = buildNode(testCase.lst1)
		var list2 = buildNode(testCase.lst2)
		var list3 = buildNode(testCase.expectedLst)

		var actual = addTwoNumbers(list1, list2)

		for list3 != nil {
			assert2.NotNil(t, actual)
			assert.Equal(t, actual.key, list3.key)
			actual = actual.next
			list3 = list3.next
		}
	}
}

func buildNode(lst []int) *Node {
	if len(lst) == 0 {
		return nil
	}

	return &Node{
		key:  lst[0],
		next: buildNode(lst[1:]),
	}
}
