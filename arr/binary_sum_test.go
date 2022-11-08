package arr

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestBinarySum(t *testing.T) {
	var dataSet = []struct {
		first    string
		second   string
		expected string
	}{
		{first: `0`, second: `0`, expected: `0`},
		{first: `0`, second: `1`, expected: `1`},
		{first: `1`, second: `0`, expected: `1`},
		{first: `1`, second: `1`, expected: `10`},
		{first: `10`, second: `0`, expected: `10`},
		{first: `10`, second: `1`, expected: `11`},
		{first: `11`, second: `1`, expected: `100`},
	}

	for _, i := range dataSet {
		assert.Equal(t, binarySum(i.first, i.second), i.expected)
	}
}
