package sort

import (
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type bubbleSortSuite struct {
	suite.Suite
	dataSet []dataSet
}

func (suite *bubbleSortSuite) SetupTest() {
	suite.dataSet = []dataSet{
		{[]uint32{0}, []uint32{0}},
		{[]uint32{0, 1}, []uint32{0, 1}},
		{[]uint32{1, 0}, []uint32{0, 1}},
		{[]uint32{1, 0, 3, 2, 5, 4}, []uint32{0, 1, 2, 3, 4, 5}},
		{[]uint32{9, 7, 5, 3, 1, 0}, []uint32{0, 1, 3, 5, 7, 9}},
	}
}

func (suite bubbleSortSuite) Test() {
	for _, set := range suite.dataSet {
		assert.Equal(suite.T(), bubble(set.data), set.expected)
	}
}

func TestBubbleSort(t *testing.T) {
	suite.Run(t, new(bubbleSortSuite))
}
