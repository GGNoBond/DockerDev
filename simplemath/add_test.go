package simplemath

import "testing"

func TestAdd(t *testing.T) {
	r := Add(123, 123)
	if r != 246 {
		t.Errorf("Add(123, 123) failed, Got %d, expected 246.", r)
	}
}
