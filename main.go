package intvector

import (
	"errors"
	"runtime"
)

//Intvector is a vector implementation in golang
type Intvector struct {
	vec []int
}

//Push inserts/pushes a new integer into the int slice
func (v *Intvector) Push(s int) {
	v.vec = append(v.vec, s)
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
