package recursion

import (
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type fibDataSet [2]uint

type FibSuite struct {
	suite.Suite
	dataSet []fibDataSet
}

func (suite *FibSuite) SetupTest() {
	suite.dataSet = []fibDataSet{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
	}
}

func (suite *FibSuite) TestRecursive() {
	for _, set := range suite.dataSet {
		assert.Equal(suite.T(), FibRecursion(set[0]), set[1])
	}
}

func TestFib(t *testing.T)  {
	suite.Run(t, new(FibSuite))
}
