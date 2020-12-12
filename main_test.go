package intvector

import (
	"testing"
)

func TestPush(t *testing.T) {
	var v Intvector
	curLength := len(v.vec)
	v.Push(4)

	lengthGot := len(v.vec)
	lengthWant := curLength + 1

	if lengthGot != lengthWant {
		t.Errorf("Incorrect Vector length : Want %d, got %d", lengthWant, lengthGot)
	}

	elemGot := v.vec[len(v.vec)-1]
	elemWant := 4

	if elemGot != elemWant {
		t.Errorf("Incorrect last element : want %d,got %d", elemWant, elemGot)
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

	gotLengthDifference := newVecLength - vecLength
	wantLenghtDifference := len(sliceToInsert)

	if gotLengthDifference != wantLenghtDifference {
		t.Errorf("Incorrect Vector length Difference : Want %d, got %d", wantLenghtDifference, gotLengthDifference)
	}

	//the last elements must be the ones inserted
	for i := newVecLength - 1; i > vecLength-1; i-- {
		gotElem := s.vec[i]
		wantElem := sliceToInsert[i-vecLength]
		if gotElem != wantElem {
			t.Errorf("Incorrect insert element at index %d, want %d, got %d", i, wantElem, gotElem)
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
