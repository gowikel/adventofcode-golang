package search_test

import (
	"testing"

	"github.com/gowikel/adventofcode-golang/year2024/day07/parser"
	"github.com/gowikel/adventofcode-golang/year2024/day07/search"
	"github.com/stretchr/testify/assert"
)

func TestSearch__one_operand(t *testing.T) {
	eq := parser.Equation{
		Result:   12,
		Operands: []int{12},
	}
	result := search.Part1(eq)
	assert.True(t, result)
}

func TestSearch__one_operand_no_solution(t *testing.T) {
	eq := parser.Equation{
		Result:   12,
		Operands: []int{13},
	}
	result := search.Part1(eq)
	assert.False(t, result)
}

func TestSearch__zero_special_case(t *testing.T) {
	eq := parser.Equation{
		Result:   0,
		Operands: []int{},
	}
	result := search.Part1(eq)
	assert.True(t, result)
}

func TestSearch__zero_one_operand(t *testing.T) {
	eq := parser.Equation{
		Result:   0,
		Operands: []int{0},
	}
	result := search.Part1(eq)
	assert.True(t, result)
}

func TestSearch__n408__simple_multiplication(t *testing.T) {
	eq := parser.Equation{
		Result:   408,
		Operands: []int{12, 34},
	}

	result := search.Part1(eq)
	assert.True(t, result)
}

func TestSearch__n46__simple_addition(t *testing.T) {
	eq := parser.Equation{
		Result:   46,
		Operands: []int{12, 34},
	}

	result := search.Part1(eq)
	assert.True(t, result)
}

func TestSearch__n409__no_solution(t *testing.T) {
	eq := parser.Equation{
		Result:   409,
		Operands: []int{12, 34},
	}
	result := search.Part1(eq)
	assert.False(t, result)
}

func TestSearch__n2576_addition_first(t *testing.T) {
	eq := parser.Equation{
		Result:   2576,
		Operands: []int{12, 34, 56},
	}
	result := search.Part1(eq)
	assert.True(t, result)
}

func TestSearch__n464_multiplication_first(t *testing.T) {
	eq := parser.Equation{
		Result:   464,
		Operands: []int{12, 34, 56},
	}
	result := search.Part1(eq)
	assert.True(t, result)
}

func TestSearch__n100_no_solution(t *testing.T) {
	eq := parser.Equation{
		Result:   100,
		Operands: []int{12, 34, 56},
	}
	result := search.Part1(eq)
	assert.False(t, result)
}

func TestSearch__n180_valid(t *testing.T) {
	eq := parser.Equation{
		Result:   180,
		Operands: []int{12, 34, 56, 78},
	}
	result := search.Part1(eq)
	assert.True(t, result)
}

func TestSearch__n1782144_valid(t *testing.T) {
	eq := parser.Equation{
		Result:   1_782_144,
		Operands: []int{12, 34, 56, 78},
	}
	result := search.Part1(eq)
	assert.True(t, result)
}

func TestSearch__n7956_valid(t *testing.T) {
	eq := parser.Equation{
		Result:   7956,
		Operands: []int{12, 34, 56, 78},
	}
	result := search.Part1(eq)
	assert.True(t, result)
}

func TestSearch__n2654_valid(t *testing.T) {
	eq := parser.Equation{
		Result:   2654,
		Operands: []int{12, 34, 56, 78},
	}
	result := search.Part1(eq)
	assert.True(t, result)
}

func TestSearch__n542_valid(t *testing.T) {
	eq := parser.Equation{
		Result:   542,
		Operands: []int{12, 34, 56, 78},
	}
	result := search.Part1(eq)
	assert.True(t, result)
}

func TestSearch__valid(t *testing.T) {
	eq := parser.Equation{
		Result:   22926,
		Operands: []int{12, 34, 56, 78},
	}
	result := search.Part1(eq)
	assert.True(t, result)
}

func TestSearch__n185_no_solution(t *testing.T) {
	eq := parser.Equation{
		Result:   185,
		Operands: []int{12, 34, 56, 78},
	}
	result := search.Part1(eq)
	assert.False(t, result)
}

func TestSearch__n3267(t *testing.T) {
	eq := parser.Equation{
		Result:   3267,
		Operands: []int{81, 40, 27},
	}
	result := search.Part1(eq)
	assert.True(t, result)
}
