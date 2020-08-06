package set

import "sync"

// String represents a set of strings
type String struct {
	data map[string]int
	lock sync.RWMutex
	lidx int
}

// NewString create a new set for strings
func NewString() *String {
	return &String{
		data: map[string]int{},
		lock: sync.RWMutex{},
	}
}

// Add a new string
func (s *String) Add(data string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	_, ok := s.data[data]
	if !ok {
		s.data[data] = s.lidx
		s.lidx++
	}
}

// Length return the number of items in the set
func (s *String) Length() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.lidx
}

// Values return an array of values for the set
func (s *String) Values() []string {
	s.lock.RLock()
	defer s.lock.RUnlock()

	values := make([]string, s.lidx)
	for value, index := range s.data {
		values[index] = value
	}
	return values
}
