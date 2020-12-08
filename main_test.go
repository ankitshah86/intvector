package intvector

import (
	"testing"
)

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

func TestInsert(t *testing.T) {

	var s Intvector

	sliceToInsert := []int{1, 2, 3}
	vecLength := len(s.vec)

	s.Insert(sliceToInsert...)
	newVecLength := len(s.vec)

	if newVecLength-vecLength != len(sliceToInsert) {
		t.Errorf("Insert Length Test failed")
	}

	//the last elements must be th ones inserted
	for i := newVecLength - 1; i > vecLength-1; i-- {
		if s.vec[i] != sliceToInsert[i-vecLength] {
			t.Errorf("Insert element Test failed`")
		}
	}
}
