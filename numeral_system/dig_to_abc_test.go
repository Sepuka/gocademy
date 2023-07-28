package numeral_system

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvert(t *testing.T) {
  var dataSet = []struct {
    k int
    expected string
  }{
    {
      k: 1,
      expected: `A`,
    },
    {
      k: 2,
      expected: `B`,
    },
    {
      k: 26,
      expected: `Z`,
    },
    {
      k: 27,
      expected: `AA`,
    },
    {
      k: 53,
      expected: `BA`,
    },
    {
      k: 5000,
      expected: `GJH`,
    },
  }

  for _, caseVal := range dataSet {
    assert.Equal(t, convert(caseVal.k), caseVal.expected)
  }
}
