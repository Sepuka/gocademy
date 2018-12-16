package sort

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type choiceSortSuite struct {
	suite.Suite
	dataSet []dataSet
}

func (suite *choiceSortSuite) SetupTest() {
	suite.dataSet = []dataSet{
		{[]uint32{0}, []uint32{0}},
		{[]uint32{0, 1}, []uint32{0, 1}},
		{[]uint32{1, 0}, []uint32{0, 1}},
		{[]uint32{1, 0, 3, 2, 5, 4}, []uint32{0, 1, 2, 3, 4, 5}},
	}
}

func (suite choiceSortSuite) TestDummy() {
	for _, set := range suite.dataSet {
		assert.Equal(suite.T(), set.expected, simpleChoice(set.data))
	}
}

func TestChoiceSort(t *testing.T) {
	suite.Run(t, new(choiceSortSuite))
}
