package tree

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Возвращает среднее арифметическое заданного уровня глубины
func TestWalk(t *testing.T) {
	var (
		emptyTree = &Tree{}
		tree0     = &Tree{
			value: 123,
		}
		tree1 = &Tree{
			value: 123,
			left: &Tree{
				value: 1,
			},
			right: &Tree{
				value: 1,
			},
		}
		tree3 = &Tree{ // 0 level
			value: 1,
			left: &Tree{ // 1 level
				value: 2,
				left: &Tree{ // 2 level
					value: 3,
					left: &Tree{ // 3 level
						value: 4,
					},
					right: &Tree{
						value: 4,
					},
				},
				right: &Tree{ // 2 level
					value: 3,
					left: &Tree{ // 3 level
						value: 4,
					},
					right: &Tree{
						value: 4,
					},
				},
			},
			right: &Tree{
				value: 2,
				left: &Tree{ // 2 level
					value: 3,
					left: &Tree{ // 3 level
						value: 4,
					},
					right: &Tree{
						value: 4,
					},
				},
				right: &Tree{ // 2 level
					value: 3,
					left: &Tree{ // 3 level
						value: 4,
					},
					right: &Tree{
						value: 4,
					},
				},
			},
		}
	)

	var testCases = []struct {
		tree     *Tree
		level    uint
		expected float32
		err      error
	}{
		{
			tree:     emptyTree,
			level:    0,
			expected: 0,
		},
		{
			tree:     emptyTree,
			level:    1,
			expected: 0,
			err:      errors.New(`too deep`),
		},
		{
			tree:     tree0,
			level:    0,
			expected: 123,
		},
		{
			tree:     tree1,
			level:    0,
			expected: 123,
		},
		{
			tree:     tree1,
			level:    1,
			expected: 1,
		},
		{
			tree:     tree3,
			level:    3,
			expected: 4,
		},
	}

	// решение с рекурсией
	for _, testCase := range testCases {
		actual, err := WalkAvgLeafRecursive(testCase.tree, testCase.level)
		assert.Equal(t, testCase.expected, actual)
		if testCase.err != nil {
			assert.EqualError(t, err, testCase.err.Error())
		}
	}

	// решение итеративное
	for _, testCase := range testCases {
		actual, err := WalkAvgLeafIter(testCase.tree, testCase.level)
		assert.Equal(t, testCase.expected, actual)
		if testCase.err != nil {
			assert.EqualError(t, err, testCase.err.Error())
		}
	}
}

// Возвращает среднюю глубину дерева
func TestWalkAvgDeep(t *testing.T) {
	var (
		emptyTree = &Tree{}
		// deep is 1
		tree1 = &Tree{
			left:  &Tree{},
			right: &Tree{},
		}
		tree3 = &Tree{ // 0 level
			left: &Tree{ // 1 level
				left: &Tree{ // 2 level
					left:  &Tree{}, // 3 level
					right: &Tree{},
				},
				right: &Tree{ // 2 level
					left:  &Tree{}, // 3 level
					right: &Tree{},
				},
			},
			right: &Tree{
				left: &Tree{ // 2 level
					left:  &Tree{}, // 3 level
					right: &Tree{},
				},
				right: &Tree{ // 2 level
					left:  &Tree{}, // 3 level
					right: &Tree{},
				},
			},
		}
		leftOrientedTree = &Tree{
			left: &Tree{ // 1 level
				left: &Tree{ // 2 level
					left: &Tree{}, // 3 level
				},
			},
		}
		unbalancedTree = &Tree{
			left: &Tree{ // 1 level
				left: &Tree{ // 2 level
					left: &Tree{}, // 3 level
				},
			},
			right: &Tree{},
		}
	)

	var testCases = []struct {
		tree     *Tree
		expected float32
	}{
		{
			tree:     emptyTree,
			expected: 0,
		},
		{
			tree:     tree1,
			expected: 1,
		},
		{
			tree:     tree3,
			expected: 3,
		},
		{
			tree:     leftOrientedTree,
			expected: 1.5,
		},
		{
			tree:     unbalancedTree,
			expected: 2,
		},
	}

	for _, testCase := range testCases {
		actual := WalkAvgDeep(testCase.tree)
		assert.Equal(t, testCase.expected, actual)
	}
}
