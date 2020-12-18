package intvector

import (
	"errors"
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

	var testNum = 1
	s.Push(testNum)

	want := testNum
	got, err := s.Pop()

	if want != got {
		t.Errorf("Pop Test failed. Want %d got %d", want, got)
	}

	if err != nil {
		t.Errorf("Pop test failed with Error : %s", err)
	}

	//check if appropriate error is thrown in situations where the vector is empty
	got, err = s.Pop()
	want = 0

	wantError := errors.New("Empty Vector")
	gotError := err

	if got != want {
		t.Errorf("Pop Test failed. Want %d got %d", want, got)
	}

	if wantError.Error() != gotError.Error() {
		t.Errorf("Pop Test failed, Want Error : %s, got Error : %s", wantError, gotError)
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

func TestUnshift(t *testing.T) {
	var s Intvector
	s.Insert([]int{2, 3, 4}...)
	elem := 1

	curLength := len(s.vec)
	s.Unshift(elem)

	wantElem := elem
	gotElem := s.At(0)

	if wantElem != gotElem {
		t.Errorf("Unshift Test Failed for incoming element, want %d got %d ", wantElem, gotElem)
	}

	wantLength := curLength + 1
	gotLength := len(s.vec)

	if wantLength != gotLength {
		t.Errorf("Unshift Length Test failed, want %d got %d", wantLength, gotLength)
	}
}

func TestUniquePush(t *testing.T) {

	var s Intvector
	length := 10
	//try pushing multiple items
	for i := 0; i < length; i++ {
		//Deliberately insert duplicate elements
		s.UniquePush(i)
		s.UniquePush(i)
		s.UniquePush(i)
	}

	wantLength := length //since duplicate elements are being inserted, the length shouldn't exceed loop iterations.
	gotLength := len(s.vec)

	if wantLength != gotLength {
		t.Errorf("UniquePush Test failed : got length %d, want length %d", gotLength, wantLength)
	}

	m := make(map[int]int, 0)
	//also need to ensure that each element in this array is unique
	for _, v := range s.vec {
		if _, ok := m[v]; !ok {
			m[v] = 1
		} else {
			//this means that an element is duplicated
			t.Errorf("UniquePush Test failed : Duplicate element %d found", v)
		}
	}
}

func TestSize(t *testing.T) {
	var s Intvector
	wantSize := 10

	for i := 0; i < wantSize; i++ {
		s.Push(i)
	}
	gotSize := s.Size()

	if gotSize != wantSize {
		t.Errorf("Size Test Failed : got size %d, want %d", gotSize, wantSize)
	}
}

func TestClear(t *testing.T) {
	var s Intvector
	//populate the vector with some elementes
	for i := 0; i < 10; i++ {
		s.Push(i)
	}
	s.Clear()
	wantLength := 0
	gotLength := len(s.vec)
	if wantLength != gotLength {
		t.Errorf("Clear Test Failed : got vector length %d, want %d", gotLength, wantLength)
	}
}
