package set_test

import (
	"testing"

	"github.com/ecnepsnai/set"
)

func TestSetFloat64(t *testing.T) {
	s := set.NewFloat64()

	CompareInt(t, s.Length(), 0, "Length of empty set must be 0")
	CompareInt(t, len(s.Values()), 0, "Length of empty set must be 0")

	s.Add(float64(1.1))
	s.Add(float64(2.1))
	s.Add(float64(3.1))
	s.Add(float64(3.1))
	s.Add(float64(2.1))
	s.Add(float64(1.1))

	values := s.Values()
	CompareInt(t, s.Length(), 3, "Set length should be expected")
	CompareInt(t, len(values), 3, "Set length should be expected")

	CompareFloat64(t, values[0], 1.1, "Value at index 1 must match")
	CompareFloat64(t, values[1], 2.1, "Value at index 2 must match")
	CompareFloat64(t, values[2], 3.1, "Value at index 3 must match")
}
