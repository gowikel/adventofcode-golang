package utils

import (
	"math"

	"golang.org/x/exp/constraints"
)

type Real interface {
	constraints.Signed | constraints.Float
}

// Returns the absolute calc from the given val
func Abs[T Real](a T) T {
	if a < 0 {
		return -a
	}
	return a
}

// Calculates the Greater Common Divisor of a set of integers
// panics if there are less than one argument
func GCD[T constraints.Integer](n ...T) T {
	var result T

	switch len(n) {
	case 0:
		panic("GCD needs at least one integer")
	case 1:
		return n[0]
	}

	result = n[0]
	for i := 1; i < len(n); i++ {
		a := result
		b := n[i]

		result = gcd(a, b)
	}

	return result
}

func gcd[T constraints.Integer](a, b T) T {
	if b > a {
		a, b = b, a
	}

	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// Calculates the Lowest Common Multiple of a set of integers
func LCM[T constraints.Integer](n ...T) T {
	var result T

	switch len(n) {
	case 0:
		panic("LCM needs at least one integer")
	case 1:
		return n[0]
	}

	result = n[0]
	for i := 1; i < len(n); i++ {
		a := result
		b := n[i]

		result = lcm(a, b)
	}

	return result
}

func lcm[T constraints.Integer](a, b T) T {
	return T(math.Abs(float64(a*b)) / float64(GCD[T](a, b)))
}
