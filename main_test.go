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
	gotElem, _ := s.At(0)

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

func TestReverse(t *testing.T) {
	var s Intvector

	wantLength := 99

	for i := 0; i < wantLength; i++ {
		s.Push(i)
	}

	s.Reverse()

	//make sure the length hasn't changed
	gotLength := s.Size()

	if gotLength != wantLength {
		t.Errorf("Reverse Test Failed : got length %d, want length %d", gotLength, wantLength)
	}

	for i := 0; i < wantLength; i++ {
		wantElem := 99 - i - 1
		gotElem, _ := s.At(i)
		if wantElem != gotElem {
			t.Errorf("Reverse Test Failed : got element %d, want %d", gotElem, wantElem)
		}
	}

}

func TestAt(t *testing.T) {
	var s Intvector

	length := 10
	for i := 0; i < length; i++ {
		s.Push(i * 2)
	}

	for i := 0; i < length; i++ {
		want := i * 2
		got, _ := s.At(i)

		if want != got {
			t.Errorf("At Test Failed : at index %d got element %d, want %d ", i, got, want)
		}
	}

	//test for invalid indices
	idx := -1

	_, err := s.At(idx)

	if err == nil {
		t.Errorf("At Test Failed : should throw error for invalid index %d", idx)
	}

	idx = length
	_, err = s.At(idx)

	if err == nil {
		t.Errorf("At Test Failed : should throw error for invalid index %d", idx)
	}

}

func TestSwap(t *testing.T) {
	var s Intvector

	s.Insert([]int{1, 2, 3}...)

	s.Swap(0, 1)

	want := 2
	got, _ := s.At(0)

	if want != got {
		t.Errorf("Swap Test Failed : want element %d at 0, got %d", want, got)
	}

	want = 1
	got, _ = s.At(1)

	if want != got {
		t.Errorf("Swap Test Failed : want element %d at 0, got %d", want, got)
	}

	err := s.Swap(0, 0)
	if err == nil {
		t.Errorf("Swap Test Failed : should throw error for same index swap.")
	}

	err = s.Swap(0, s.Size()+1)
	if err == nil {
		t.Errorf("Swap Test Failed : should throw error for out of range index.")
	}

	err = s.Swap(-1, s.Size()-1)
	if err == nil {
		t.Errorf("Swap Test Failed : should throw error for out of range index.")
	}
}

func TestSet(t *testing.T) {
	var s Intvector

	//init 10 positions with zeros
	for i := 0; i < 10; i++ {
		s.Push(0)

		//set the value
		s.Set(i, i*2)

		want := i * 2
		got, _ := s.At(i)

		if want != got {
			t.Errorf("Set Test Failed : want element %d at index %d, got %d", want, i, got)
		}
	}

	err := s.Set(s.Size()+1, 0)
	if err == nil {
		t.Errorf("Set Test Failed : should throw error for out of range index.")
	}

	err = s.Set(-1, 0)
	if err == nil {
		t.Errorf("Set Test Failed : should throw error for out of range index.")
	}

}

func TestSortedPush(t *testing.T) {
	//Test this function with random elements
	var s Intvector
	testSlice := []int{4, 23, 62, 7, 2, 7, 8, 1, 1, 5, 3, 8, 93, 34, 38}

	for _, v := range testSlice {
		s.SortedPush(v)
	}

	//test length of vector to ensure no element was missed
	wantLength := len(testSlice)
	gotLength := s.Size()
	if wantLength != gotLength {
		t.Errorf("SortedPush Length Test Failed : want %d got %d", wantLength, gotLength)
	}

	//check if the array is sorted
	isSoreted := true
	for i, v := range s.vec {
		if i > 1 {
			if v < s.vec[i-1] {
				isSoreted = false
				break
			}
		}
	}

	if !isSoreted {
		t.Errorf("SortedPush Test Failed : The vector is unsorted")
	}
}

func TestSort(t *testing.T) {
	//Test this function with unsorted random elements
	var s Intvector
	testSlice := []int{4, 23, 62, 7, 2, 7, 8, 1, 1, 5, 3, 8, 93, 34, 38}

	for _, v := range testSlice {
		s.Push(v)
	}

	//sort vector here
	s.Sort()

	//check if the vector is sorted
	isSoreted := true
	for i, v := range s.vec {
		if i > 1 {
			if v < s.vec[i-1] {
				isSoreted = false
				break
			}
		}
	}

	if !isSoreted {
		t.Errorf("Sort Test Failed : The vector is unsorted")
	}
}

func TestIsSorted(t *testing.T) {
	var s Intvector
	s.Push(1)
	want := true
	got := s.IsSorted()
	if want != got {
		t.Errorf("IsSorted Test failed : want %t got %t", want, got)
	}

	//test with sorted vector
	s.Insert([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}...)
	want = true
	got = s.IsSorted()

	if want != got {
		t.Errorf("IsSorted Test failed : want %t got %t", want, got)
	}

	//apppend some random unsorted elements
	s.Insert([]int{3, 5, 2, 7, 2, 5, 4, 6, 4, 23, 56, 9, 3, 3}...)
	want = false
	got = s.IsSorted()

	if want != got {
		t.Errorf("IsSorted Test failed : want %t got %t", want, got)
	}
}

func TestFirst(t *testing.T) {
	var s Intvector

	res, err := s.First()

	want := 0
	got := res

	if err == nil {
		t.Errorf("First method Test failed : should throw error for empty vector")
	}

	if want != got {
		t.Errorf("First method Test failed : want %d, got %d", want, got)
	}

	s.Insert([]int{1, 2, 3, 4, 5}...)

	want = 1
	got, err = s.First()

	if err != nil {
		t.Errorf("First method Test failed : should not throw error for empty vector")
	}

	if want != got {
		t.Errorf("First method Test failed : want %d, got %d", want, got)
	}
}

func TestLast(t *testing.T) {
	var s Intvector

	res, err := s.Last()

	want := 0
	got := res

	if err == nil {
		t.Errorf("Last method Test failed : should throw error for empty vector")
	}

	if want != got {
		t.Errorf("Last method Test failed : want %d, got %d", want, got)
	}

	s.Insert([]int{1, 2, 3, 4, 5}...)

	want = 5
	got, err = s.Last()

	if err != nil {
		t.Errorf("Last method Test failed : should not throw error for empty vector")
	}

	if want != got {
		t.Errorf("Last method Test failed : want %d, got %d", want, got)
	}
}

func TestSearch(t *testing.T) {
	var s Intvector
	n := 10
	for i := 0; i <= n; i++ {
		s.Push(i * 2)
	}
	want := 10
	got := s.Search(10 * 2)
	if want != got {
		t.Errorf("Search test failed : want %d, got %d", want, got)
	}
	//also test for non - existent element
	want = -1
	got = s.Search(11 * 2)

	if want != got {
		t.Errorf("Search test failed : should return %d instead of %d for non-existent element", want, got)
	}
}

func TestSearchAll(t *testing.T) {
	var s Intvector
	n := 10
	for i := 0; i < n; i++ {
		//this would ensure that each number will be be inserted the number of times equal to its value
		for j := 0; j < i; j++ {
			s.Push(i)
		}
	}

	start := 0

	for i := 0; i < n; i++ {
		want := i
		got := len(s.SearchAll(i))
		if want != got {
			t.Errorf("SearchAll Test Failed : want total count of %d to be %d, got %d", i, want, got)
		}
		//also check indexes  may need elaborate testing here
		if i > 1 {
			start += (i - 1)
		}

		wantSlice := []int{}
		for j := 0; j < got; j++ {
			wantSlice = append(wantSlice, start+j)
		}
		gotSlice := s.SearchAll(i)
		//fmt.Println(wantSlice, gotSlice)

		isSliceEqual := true

		for j := 0; j < len(wantSlice); j++ {
			if wantSlice[j] != gotSlice[j] {
				isSliceEqual = false
				break
			}
		}

		if !isSliceEqual {
			t.Errorf("SearchAll Test Failed : slice returned by SearchAll is inaccurate, want slice %d, got %d", wantSlice, gotSlice)
		}
	}
}

func TestMin(t *testing.T) {
	var s Intvector
	wantNum, wantIdx := 0, -1
	gotNum, gotIdx := s.Min()

	if wantNum != gotNum {
		t.Errorf("Min Test Failed : for empty vector, want num %d, got %d", wantNum, gotNum)
	}

	if wantIdx != gotIdx {
		t.Errorf("Min Test Failed : for empty vector, want index %d, got %d", wantIdx, gotIdx)
	}

	s.Insert([]int{1, 2, -4, 34, 2788, 24, 2, 4, 4, 0, 223, 6453, 234677, 234, 89, 76, -778, 345, 22, 4, 66, 4, 3, 7, 9, 8, 1, 2}...)

	wantNum, wantIdx = -778, 16
	gotNum, gotIdx = s.Min()

	if wantNum != gotNum {
		t.Errorf("Min Test Failed : want num %d, got %d", wantNum, gotNum)
	}

	if wantIdx != gotIdx {
		t.Errorf("Min Test Failed : want index %d, got %d", wantIdx, gotIdx)
	}
}

func TestMax(t *testing.T) {
	var s Intvector
	wantNum, wantIdx := 0, -1
	gotNum, gotIdx := s.Max()

	if wantNum != gotNum {
		t.Errorf("Min Test Failed : for empty vector, want num %d, got %d", wantNum, gotNum)
	}

	if wantIdx != gotIdx {
		t.Errorf("Min Test Failed : for empty vector, want index %d, got %d", wantIdx, gotIdx)
	}

	s.Insert([]int{1, 2, -4, 34, 2788, 24, 2, 4, 4, 0, 223, 6453, 234677, 234, 89, 76, -778, 345, 22, 4, 66, 4, 3, 7, 9, 8, 1, 2}...)

	wantNum, wantIdx = 234677, 12
	gotNum, gotIdx = s.Max()

	if wantNum != gotNum {
		t.Errorf("Min Test Failed : want num %d, got %d", wantNum, gotNum)
	}

	if wantIdx != gotIdx {
		t.Errorf("Min Test Failed : want index %d, got %d", wantIdx, gotIdx)
	}
}
