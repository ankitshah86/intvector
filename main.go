package intvector

import (
	"errors"
	"runtime"
	"strconv"
)

//Intvector is a vector implementation in golang
type Intvector struct {
	vec []int
}

//Push inserts/pushes a new integer into the int slice
func (v *Intvector) Push(s int) {
	v.vec = append(v.vec, s)
}

//Insert appends a new slice to an existing slice
func (v *Intvector) Insert(s ...int) {
	v.vec = append(v.vec, s...)
}

//Pop removes the first element from the slice and retruns it
func (v *Intvector) Pop() (int, error) {
	var s int

	if len(v.vec) > 0 {
		s = v.vec[len(v.vec)-1]
		v.vec = v.vec[:len(v.vec)-1]
	} else {
		//add better handling here
		return 0, errors.New("Empty Vector")
	}

	return s, nil
}

//Size returns the current size of the vector
func (v *Intvector) Size() int {
	return len(v.vec)
}

//Clear clears out the slice and invokes the garbage collector to reclaim the freed memory.
func (v *Intvector) Clear() {
	v.vec = nil
	runtime.GC()
}

//Reverse function can be used to reverse the vector
func (v *Intvector) Reverse() {
	for i := 0; i < len(v.vec)/2; i++ {
		v.vec[i], v.vec[len(v.vec)-1-i] = v.vec[len(v.vec)-i-1], v.vec[i]
	}
}

//At allows for accesing any element of the vector
func (v *Intvector) At(i int) int {
	return v.vec[i]
}

//Swap function swaps two elements of the vector
func (v *Intvector) Swap(idx1 int, idx2 int) error {

	if idx1 == idx2 {
		return errors.New("idx1 and idx2 are the same number, no swap was performed")
	}

	if idx1 >= len(v.vec) || idx1 < 0 {
		return errors.New("idx1 out of range for vector of length " + strconv.Itoa(len(v.vec)))
	}

	if idx2 >= len(v.vec) || idx2 < 0 {
		return errors.New("idx2 out of range for vector of length " + strconv.Itoa(len(v.vec)))
	}

	v.vec[idx1], v.vec[idx2] = v.vec[idx2], v.vec[idx1]

	return nil
}

//Set function can be used to set the value at a specific index in the vector
func (v *Intvector) Set(idx int, value int) error {

	if idx < 0 {
		return errors.New("idx must be a positive number")
	}

	if idx >= len(v.vec) {
		return errors.New("idx out of range for vector of length " + strconv.Itoa(len(v.vec)))
	}

	v.vec[idx] = value
	return nil
}

//add sorting functionality

//SortedPush pushes the incoming element into the vector in a sorted way
//it is assumed that the Vector is already sorted
func (v *Intvector) SortedPush(n int) {
	v.vec = append(v.vec, n)
	//see if this can be optimized
	if len(v.vec) > 1 && v.vec[len(v.vec)-2] > n {
		idx := len(v.vec) - 1
		for i := len(v.vec) - 2; i >= 0; i-- {
			if n < v.vec[i] {
				v.Swap(idx, i)
				idx = i
			} else {
				break
			}
		}
	}
}
