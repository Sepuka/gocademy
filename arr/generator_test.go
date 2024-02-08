package arr

import (
	"errors"
	"sort"
	"sync"
	"testing"
)

func TestNew(t *testing.T) {
	cases := []*struct {
		sync.Mutex
		name       string
		start      int
		stop       int
		step       int
		concurrent int
		expected   []int
		err        error
	}{
		{name: "empty", step: 10, concurrent: 1},
		{name: "simple", start: 1, stop: 5, step: 1, concurrent: 1, expected: []int{1, 2, 3, 4}},
		{name: "simple_threads", start: 1, stop: 5, step: 1, concurrent: 2, expected: []int{1, 2, 3, 4}},
		{name: "big_step", start: 1, stop: 10, step: 10, concurrent: 1, expected: []int{1}},
		{name: "one_step_one_thread", start: 1, stop: 2, step: 1, concurrent: 1, expected: []int{1}},
		{name: "one_step_threads", start: 1, stop: 2, step: 1, concurrent: 2, expected: []int{1}},
		{name: "two_step_one_thread", start: 1, stop: 10, step: 2, concurrent: 1, expected: []int{1, 3, 5, 7, 9}},
		{name: "two_step_threads", start: 1, stop: 10, step: 2, concurrent: 4, expected: []int{1, 3, 5, 7, 9}},
		{name: "failed_negative_step", start: 1, stop: 10, step: -2, err: errors.New("failed step")},
		{name: "failed_zero_step", start: 1, stop: 10, step: 0, err: errors.New("failed step")},
		{name: "failed_start_stop", start: 10, stop: 9, step: 1, err: errors.New("failed start/stop")},
		{
			name:       "long_sequence_one_thread",
			start:      10,
			stop:       50,
			step:       2,
			concurrent: 1,
			expected:   []int{10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48},
		},
		{
			name:       "long_sequence_threads",
			start:      10,
			stop:       50,
			step:       2,
			concurrent: 5,
			expected:   []int{10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			var wg sync.WaitGroup
			result := make([]int, 0, len(c.expected))
			g := New(c.start, c.stop, c.step)

			wg.Add(c.concurrent)
			for i := 0; i < c.concurrent; i++ {
				go func() {
					for v, ok := g.Next(); ok; v, ok = g.Next() {
						c.Lock()
						result = append(result, v)
						c.Unlock()
					}
					wg.Done()
				}()
			}
			wg.Wait()
			sort.Ints(result)

			if err := g.Err(); c.err != nil {
				// expected error
				if err == nil {
					tt.Errorf("expected err=%v, but gotten nil", c.err)
				}
			} else {
				// expected successful
				if err != nil {
					tt.Errorf("unexpected error=%v", err)
					return
				}
				if n, m := len(c.expected), len(result); n != m {
					tt.Errorf("failed result len, expected=%d, gotten=%d", n, m)
				} else {
					for i, e := range c.expected {
						if e != result[i] {
							tt.Errorf("failed result[%d], expected=%d, gotten=%d", i, e, result[i])
						}
					}
				}
			}
		})
	}
}
