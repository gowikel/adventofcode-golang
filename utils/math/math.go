package algo

import "golang.org/x/exp/constraints"

// Returns the absolute calc from the given val
func Abs[T constraints.Signed](a T) T {
	if a < 0 {
		return -a
	}
	return a
}
