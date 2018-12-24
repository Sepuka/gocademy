package sort

import (
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type insertSortSuite struct {
	suite.Suite
	dataSet []dataSet
}

func (suite *insertSortSuite) SetupTest() {
	suite.dataSet = []dataSet{
		{[]uint32{0}, []uint32{0}},
		{[]uint32{0, 1}, []uint32{0, 1}},
		{[]uint32{1, 0}, []uint32{0, 1}},
		{[]uint32{1, 0, 3, 2, 5, 4}, []uint32{0, 1, 2, 3, 4, 5}},
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
