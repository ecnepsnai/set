package set_test

import (
	"testing"

	"github.com/ecnepsnai/set"
)

func TestSetString(t *testing.T) {
	s := set.NewString()

	CompareInt(t, s.Length(), 0, "Length of empty set must be 0")
	CompareInt(t, len(s.Values()), 0, "Length of empty set must be 0")

	s.Add("1")
	s.Add("2")
	s.Add("3")
	s.Add("3")
	s.Add("2")
	s.Add("1")

	values := s.Values()
	CompareInt(t, s.Length(), 3, "Set length should be expected")
	CompareInt(t, len(values), 3, "Set length should be expected")

	CompareString(t, values[0], "1", "Value at index 1 must be correct")
	CompareString(t, values[1], "2", "Value at index 2 must be correct")
	CompareString(t, values[2], "3", "Value at index 3 must be correct")
}
