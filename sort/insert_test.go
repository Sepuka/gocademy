package sort

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type insertSortSuite struct {
	suite.Suite
	dataSet []dataSet
}

func (suite *insertSortSuite) SetupTest() {
	suite.dataSet = []dataSet{
		{[]int{0}, []int{0}},
		{[]int{0, 1}, []int{0, 1}},
		{[]int{1, 0}, []int{0, 1}},
		{[]int{1, 0, 3, 2, 5, 4}, []int{0, 1, 2, 3, 4, 5}},
	}
}

func (suite insertSortSuite) TestDummy() {
	for _, set := range suite.dataSet {
		assert.Equal(suite.T(), dummy(set.data), set.expected)
	}
}

func (suite insertSortSuite) TestImprovedDummy() {
	for _, set := range suite.dataSet {
		assert.Equal(suite.T(), set.expected, improvedDummy(set.data))
	}
}

func TestInsertSort(t *testing.T) {
	suite.Run(t, new(insertSortSuite))
}
