package utils

import (
	"math"

	"github.com/rs/zerolog/log"
	"golang.org/x/exp/constraints"
)

// Returns the absolute calc from the given val
func Abs[T constraints.Signed](a T) T {
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

	log.Debug().Msgf("Calculating GCD of %v", n)

	result = n[0]
	for i := 1; i < len(n); i++ {
		a := result
		b := n[i]

		result = gcd(a, b)

		log.Debug().
			Str("func", "GCD").
			Any("A", a).
			Any("B", b).
			Any("Result", result).
			Msg("")
	}

	log.Debug().Any("Result", result).Msgf("GCD of %v calculated", n)

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

	log.Debug().Msgf("Calculating LCM of %v", n)

	result = n[0]
	for i := 1; i < len(n); i++ {
		a := result
		b := n[i]

		log.Debug().
			Str("func", "LCM").
			Any("A", a).
			Any("B", b).
			Any("Result", result).
			Msg("")

		result = lcm(a, b)
	}

	log.Debug().Any("Result", result).Msgf("LCM of %v calculated", n)

	return result
}

func lcm[T constraints.Integer](a, b T) T {
	return T(math.Abs(float64(a*b)) / float64(GCD[T](a, b)))
}
