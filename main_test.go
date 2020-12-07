package intvector

import "testing"

func TestPush(t *testing.T) {
	var v Intvector
	v.Push(5)
	v.Push(6)
	v.Push(7)
	curLength := len(v.vec)
	v.Push(4)
	newLength := len(v.vec)

	if !(newLength == curLength+1 && v.vec[len(v.vec)-1] == 4) {
		t.Errorf("Push Test failed")
	}
}
