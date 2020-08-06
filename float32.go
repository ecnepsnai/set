package set

import "sync"

// Float32 represents a set of float32s
type Float32 struct {
	data map[float32]int
	lock sync.RWMutex
	lidx int
}

// NewFloat32 create a new set for float32s
func NewFloat32() *Float32 {
	return &Float32{
		data: map[float32]int{},
		lock: sync.RWMutex{},
	}
}

// Add a new float32
func (s *Float32) Add(data float32) {
	s.lock.Lock()
	defer s.lock.Unlock()

	_, ok := s.data[data]
	if !ok {
		s.data[data] = s.lidx
		s.lidx++
	}
}

// Length return the number of items in the set
func (s *Float32) Length() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.lidx
}

// Values return an array of values for the set
func (s *Float32) Values() []float32 {
	s.lock.RLock()
	defer s.lock.RUnlock()

	values := make([]float32, s.lidx)
	for value, index := range s.data {
		values[index] = value
	}
	return values
}
