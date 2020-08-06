package set_test

import (
	"testing"

	"github.com/ecnepsnai/set"
)

func TestSetInt(t *testing.T) {
	s := set.NewInt()

	CompareInt(t, s.Length(), 0, "Length of empty set must be 0")
	CompareInt(t, len(s.Values()), 0, "Length of empty set must be 0")

	s.Add(int(1))
	s.Add(int(2))
	s.Add(int(3))
	s.Add(int(3))
	s.Add(int(2))
	s.Add(int(1))

	values := s.Values()
	CompareInt(t, s.Length(), 3, "Set length should be expected")
	CompareInt(t, len(values), 3, "Set length should be expected")

	CompareInt(t, values[0], 1, "Value at index 1 must be correct")
	CompareInt(t, values[1], 2, "Value at index 2 must be correct")
	CompareInt(t, values[2], 3, "Value at index 3 must be correct")
}
