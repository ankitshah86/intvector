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

func BenchmarkPush(b *testing.B) {
	var s Intvector
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1; j++ {
			s.Push(j)
		}
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
			t.Errorf("Insert element Test failed")
		}
	}
}

func TestPop(t *testing.T) {

	var s Intvector
	s.Push(1)
	num, err := s.Pop()
	if err != nil || num != 1 {
		t.Errorf("Pop Test failed")
	}

	num, err = s.Pop()
	if err == nil || num != 0 {
		t.Errorf("Pop Test failed")
	}

}

func TestShift(t *testing.T) {
	var s Intvector
	s.Insert([]int{1, 2, 3}...)

	got, err := s.Shift()
	want := 1
	if err != nil {
		t.Errorf("Shift Test failed")
	}

	if got != want {
		t.Errorf("Shift Test Failed got %d, want %d", got, want)
	}
}
