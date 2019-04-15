package concurrency

import (
	"github.com/magiconair/properties/assert"
	"sync"
	"testing"
)

const repeatNumber = 1000

func TestConcurrency(t *testing.T) {
	var (
		someStore = concurrencySafeStore{}
		wg = sync.WaitGroup{}
		i int
	)
	for i < repeatNumber {
		wg.Add(1)
		go func(){
			defer wg.Done()
			someStore.Inc()
		}()
		i += 1
	}
	wg.Wait()

	assert.Equal(t, someStore.cnt, repeatNumber)
}
