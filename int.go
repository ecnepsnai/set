package set

import "sync"

// Int represents a set of ints
type Int struct {
	data map[int]int
	lock sync.RWMutex
	lidx int
}

// NewInt create a new set for ints
func NewInt() *Int {
	return &Int{
		data: map[int]int{},
		lock: sync.RWMutex{},
	}
}

// Add a new int
func (s *Int) Add(data int) {
	s.lock.Lock()
	defer s.lock.Unlock()

	_, ok := s.data[data]
	if !ok {
		s.data[data] = s.lidx
		s.lidx++
	}
}

// Length return the number of items in the set
func (s *Int) Length() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.lidx
}

// Values return an array of values for the set
func (s *Int) Values() []int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	values := make([]int, s.lidx)
	for value, index := range s.data {
		values[index] = value
	}
	return values
}
