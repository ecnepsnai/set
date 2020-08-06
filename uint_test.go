package set_test

import (
	"testing"

	"github.com/ecnepsnai/set"
)

func TestSetUInt(t *testing.T) {
	s := set.NewUInt()

	CompareInt(t, s.Length(), 0, "Length of empty set must be 0")
	CompareInt(t, len(s.Values()), 0, "Length of empty set must be 0")

	s.Add(uint(1))
	s.Add(uint(2))
	s.Add(uint(3))
	s.Add(uint(3))
	s.Add(uint(2))
	s.Add(uint(1))

	values := s.Values()
	CompareInt(t, s.Length(), 3, "Set length should be expected")
	CompareInt(t, len(values), 3, "Set length should be expected")

	CompareUInt(t, values[0], 1, "Value at index 1 must be correct")
	CompareUInt(t, values[1], 2, "Value at index 2 must be correct")
	CompareUInt(t, values[2], 3, "Value at index 3 must be correct")
}
