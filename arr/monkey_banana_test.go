package arr

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestMinEatingSpeedFlattening(t *testing.T) {
  var dataSet = []struct {
    k int
    hours int
    piles []int
  }{
    {
      k: 0,
      hours: 0,
      piles: []int{},
    },
    {
      k: 5,
      hours: 8,
      piles: []int{3,6,7,11},
    },
  }

  for _, set := range dataSet {
    assert.Equal(t, minEatingSpeedFlattening(set.piles, set.hours), set.k)
  }
}

func TestMinEatingSpeedGrowing(t *testing.T) {
  var dataSet = []struct {
    k int
    hours int
    piles []int
  }{
    {
      k: 0,
      hours: 0,
      piles: []int{},
    },
    {
      k: 4,
      hours: 8,
      piles: []int{3,6,7,11},
    },
    {
      k: 30,
      hours: 5,
      piles: []int{30,11,23,4,20},
    },
    {
      k: 23,
      hours: 6,
      piles: []int{30,11,23,4,20},
    },
    {
      k: 2,
      hours: 312884469,
      piles: []int{312884470},
    },
    {
      k: 1000000000,
      hours: 3,
      piles: []int{1000000000,1000000000},
    },
  }

  for _, set := range dataSet {
    assert.Equal(t, minEatingSpeedGrowing(set.piles, set.hours), set.k)
  }
}
