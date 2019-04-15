package concurrency

import "sync"

type concurrencySafeStore struct {
	sync.Mutex
	cnt int
}

func (s *concurrencySafeStore) Inc() {
	s.Lock()
	defer s.Unlock()
	s.cnt += 1
}
