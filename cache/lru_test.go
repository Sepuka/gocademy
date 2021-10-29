package cache

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestLru(t *testing.T) {
	hashTbl = map[int]interface{}{}
	var (
	)

	getFromCache(1)
	assert.Equal(t, hashTbl, map[int]interface{}{1: "1"})

	getFromCache(1)
	assert.Equal(t, hashTbl, map[int]interface{}{1: "1"})

	getFromCache(2)
	assert.Equal(t, hashTbl, map[int]interface{}{1: "1", 2: "2"})

	getFromCache(3)
	assert.Equal(t, hashTbl, map[int]interface{}{1: "1", 2: "2", 3: "3"})

	getFromCache(4)
	assert.Equal(t, hashTbl, map[int]interface{}{4: "4", 2: "2", 3: "3"})

	getFromCache(5)
	assert.Equal(t, hashTbl, map[int]interface{}{4: "4", 5: "5", 3: "3"})
}
