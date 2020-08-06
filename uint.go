package set

import "sync"

// UInt represents a set of uints
type UInt struct {
	data map[uint]int
	lock sync.RWMutex
	lidx int
}

// NewUInt create a new set for uints
func NewUInt() *UInt {
	return &UInt{
		data: map[uint]int{},
		lock: sync.RWMutex{},
	}
}

// Add a new uint
func (s *UInt) Add(data uint) {
	s.lock.Lock()
	defer s.lock.Unlock()

	_, ok := s.data[data]
	if !ok {
		s.data[data] = s.lidx
		s.lidx++
	}
}

// Length return the number of items in the set
func (s *UInt) Length() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.lidx
}

// Values return an array of values for the set
func (s *UInt) Values() []uint {
	s.lock.RLock()
	defer s.lock.RUnlock()

	values := make([]uint, s.lidx)
	for value, index := range s.data {
		values[index] = value
	}
	return values
}
