package set_test

import "testing"

func CompareString(t *testing.T, value, expected string, must string) {
	if value != expected {
		t.Errorf("%s. Expected %s got %s", must, expected, value)
	}
}

func CompareInt(t *testing.T, value, expected int, must string) {
	if value != expected {
		t.Errorf("%s. Expected %d got %d", must, expected, value)
	}
}

func CompareUInt(t *testing.T, value, expected uint, must string) {
	if value != expected {
		t.Errorf("%s. Expected %d got %d", must, expected, value)
	}
}

func CompareFloat32(t *testing.T, value, expected float32, must string) {
	if value != expected {
		t.Errorf("%s. Expected %f got %f", must, expected, value)
	}
}

func CompareFloat64(t *testing.T, value, expected float64, must string) {
	if value != expected {
		t.Errorf("%s. Expected %f got %f", must, expected, value)
	}
}
