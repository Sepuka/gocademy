package arr

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

type testCase struct {
	readPart   int // 0-100
	timeout    time.Duration
	concurrent int
	wg         sync.WaitGroup
}

// genData returns map[int]int and fills it by values 0-100.
func genData() map[int]int {
	d := make(map[int]int, 1000)
	for i := 0; i < 100; i++ {
		d[i] = i
	}
	return d
}

func process(s SafeMap, readPart int) error {
	const (
		maxRead  = 100
		maxWrite = 1_000_000_000
	)
	var key, value int
	if key = rand.Intn(maxRead); key >= readPart {
		// try write
		key = rand.Intn(maxWrite)
		value = s.GetOrCreate(key, key)
	} else {
		// read
		value = s.GetOrCreate(key, key)
	}
	if key != value {
		return fmt.Errorf("failed result key/value %d != %d", key, value)
	}
	return nil
}

func testGetOrCreate(t *testing.T, c *testCase, s SafeMap) {
	deadline := time.Now().Add(c.timeout)
	c.wg.Add(c.concurrent)
	for j := 0; j < c.concurrent; j++ {
		go func(k int) {
			for time.Now().Before(deadline) {
				if err := process(s, c.readPart); err != nil {
					t.Error(err)
				}
			}
			c.wg.Done()
		}(j)
	}
	c.wg.Wait()
}

func TestSafeMap_GetOrCreate(t *testing.T) {
	cases := []testCase{
		{readPart: 90, timeout: time.Second, concurrent: 2},
		{readPart: 95, timeout: time.Second, concurrent: 100},
		{readPart: 10, timeout: time.Second, concurrent: 100},
		{readPart: 50, timeout: time.Second, concurrent: 1},
	}
	rand.Seed(time.Now().UnixNano())
	for i := range cases {
		s := &safeMap{data: genData()}
		testGetOrCreate(t, &cases[i], s)
	}
}
