package cache

import (
	"fmt"
	"github.com/sepuka/gocademy/linkedList"
)

const (
	hashSize = 3
)

var (
	queue   = &linkedList.LinkedList{}
	hashTbl  map[int]interface{}
)

func getFromCache(key int) interface{} {
	var (
		value interface{}
		ok    bool
	)

	if value, ok = hashTbl[key]; !ok {
		value = calculate(key)
		hashTbl[key] = value
	}

	key = queue.Set(key, value, hashSize)
	if key > 0 {
		delete(hashTbl, key)
	}

	return value
}

func calculate(key int) string {
	return fmt.Sprint(key)
}
