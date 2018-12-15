package test

import (
	"math/rand"
)

func GenerateIntSequence(len int) []int {
	var result []int
	for i := 0; i < len; i++ {
		result = append(result, rand.Int())
	}

	return result
}
