package simplemath

import "testing"

func TestSqrt(t *testing.T) {
	r := Sqrt(64)
	if r != 8 {
		t.Errorf("Sqrt(64) failed, Got %d, excepted 8.", r)
	}
}
