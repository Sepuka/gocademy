package sort

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

type dataSet struct {
	data []int
	expected []int
}

var dataProvider = []dataSet{
	{[]int{0}, []int{0}},
	{[]int{0,1}, []int{0,1}},
	{[]int{1,0}, []int{0,1}},
	{[]int{1,0,3,2,5,4}, []int{0,1,2,3,4,5}},
}

func TestDummy(t *testing.T) {
	for _, set := range dataProvider {
		assert.Equal(t, dummy(set.data), set.expected)
	}
}
