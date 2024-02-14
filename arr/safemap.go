// Package join contains an interview task.
//
// Есть интерфейс SafeMap, предоставляющий конкурентную безопасную работу с map[int]int.
// Нужно реализовать его имплементацию, не используя sync.Map.
//
// Замечание: количество чтений во много раз больше числа записей.
//
// Как проверять:
// go test -race -v
package arr

import "sync"

// SafeMap is concurrent safe map interface.
type SafeMap interface {
	GetOrCreate(key, value int) int
}

// safeMap is an implementation of SafeMap interface.
type safeMap struct {
	sync.RWMutex
	data map[int]int
}

// GetOrCreate implements SafeMap interface.
func (s *safeMap) GetOrCreate(key, value int) int {
	s.RLock()

	if v, ok := s.data[key]; ok {
		s.RUnlock()

		return v
	}

	s.RUnlock()
	s.Lock()
	defer s.Unlock()

	if v, ok := s.data[key]; ok {
		return v
	}

	s.data[key] = value

	return value
}
