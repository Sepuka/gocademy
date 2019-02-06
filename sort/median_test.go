package sort

import (
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type quickSearchSuite struct {
	suite.Suite
	dataSet []dataSelect
}

func (suite *quickSearchSuite) SetupTest() {
	suite.dataSet = []dataSelect{
		{[]uint32{0}, 0},
		{[]uint32{0, 1}, 0},
		{[]uint32{1, 0}, 1},
		{[]uint32{16, 12, 99, 95, 18, 87, 10}, 18},
	}
}

func (suite quickSearchSuite) TestSearch() {
	for _, set := range suite.dataSet {
		assert.Equal(suite.T(), median(set.data), set.expected)
	}
}

func TestQuickSearch(t *testing.T) {
	suite.Run(t, new(quickSearchSuite))
}
