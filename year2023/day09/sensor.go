package day09

import (
	"fmt"
)

type SensorRead struct {
	reads        []int
	computations [][]int
}

func NewSensorRead(reads []int) *SensorRead {
	return &SensorRead{
		reads: reads,
	}
}

// Performs the substractions unti a sequence is all zero
func (sr *SensorRead) Compute() error {

	lastSeq := sr.reads

	for !isAllZeroes(lastSeq) && len(lastSeq) >= 2 {
		var newSeq = make([]int, 0, len(lastSeq)-1)

		for i := 1; i < len(lastSeq); i++ {
			a := lastSeq[i-1]
			b := lastSeq[i]

			r := b - a
			newSeq = append(newSeq, r)
		}

		sr.computations = append(sr.computations, newSeq)
		lastSeq = newSeq
	}

	if !isAllZeroes(lastSeq) {
		return fmt.Errorf(
			"unable to compute the sequence for %v",
			sr.reads,
		)
	}

	return nil
}

func (sr SensorRead) ExtrapolateForward() int {
	var result int
	if len(sr.computations) == 0 {
		return result
	}

	for _, l := range sr.computations {
		result += l[len(l)-1]
	}

	result += sr.reads[len(sr.reads)-1]

	return result
}

func (sr SensorRead) ExtrapolateBackward() int {
	if len(sr.computations) == 0 {
		return 0
	}

	var result int

	for i := len(sr.computations) - 2; i >= 0; i-- {
		v := sr.computations[i][0]

		result = v - result
	}

	f := sr.reads[0]
	result = f - result

	return result
}

func isAllZeroes(l []int) bool {
	for _, n := range l {
		if n != 0 {
			return false
		}
	}

	return true
}
