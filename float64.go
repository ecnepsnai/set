package set

import "sync"

// Float64 represents a set of float64s
type Float64 struct {
	data map[float64]int
	lock sync.RWMutex
	lidx int
}

// NewFloat64 create a new set for float64s
func NewFloat64() *Float64 {
	return &Float64{
		data: map[float64]int{},
		lock: sync.RWMutex{},
	}
}

// Add a new float64
func (s *Float64) Add(data float64) {
	s.lock.Lock()
	defer s.lock.Unlock()

	_, ok := s.data[data]
	if !ok {
		s.data[data] = s.lidx
		s.lidx++
	}
}

// Length return the number of items in the set
func (s *Float64) Length() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.lidx
}

// Values return an array of values for the set
func (s *Float64) Values() []float64 {
	s.lock.RLock()
	defer s.lock.RUnlock()

	values := make([]float64, s.lidx)
	for value, index := range s.data {
		values[index] = value
	}
	return values
}
