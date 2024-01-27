package utils_test

import (
	"testing"

	"github.com/gowikel/adventofcode-golang/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestAbs_Zero_Int(t *testing.T) {
	input := 0
	want := 0
	got := utils.Abs[int](input)

	assert.Equal(t, want, got)
}

func TestAbs_Positive_Int(t *testing.T) {
	input := 12
	want := 12
	got := utils.Abs[int](input)

	assert.Equal(t, want, got)
}

func TestAbs_Negative_Int(t *testing.T) {
	input := -34
	want := 34
	got := utils.Abs[int](input)

	assert.Equal(t, want, got)
}

func TestAbs_Zero_Float(t *testing.T) {
	input := 0.0
	want := 0.0
	got := utils.Abs[float64](input)

	assert.Equal(t, want, got)
}

func TestAbs_Positive_Float(t *testing.T) {
	input := 12.34
	want := 12.34
	got := utils.Abs[float64](input)

	assert.Equal(t, want, got)
}

func TestAbs_Negative_Float(t *testing.T) {
	input := -34.12
	want := 34.12
	got := utils.Abs[float64](input)

	assert.Equal(t, want, got)
}

func TestGCD_48_18(t *testing.T) {
	a := 48
	b := 18
	want := 6
	got := utils.GCD(a, b)

	assert.Equal(t, want, got)
}

func TestGCD_101_103(t *testing.T) {
	a := 101
	b := 103
	want := 1
	got := utils.GCD(a, b)

	assert.Equal(t, want, got)
}

func TestGCD_0_20(t *testing.T) {
	a := 0
	b := 20
	want := 20
	got := utils.GCD(a, b)

	assert.Equal(t, want, got)
}

func TestLCM_5_7(t *testing.T) {
	a := 5
	b := 7
	want := 35
	got := utils.LCM(a, b)

	assert.Equal(t, want, got)
}

func TestLCM_10_15(t *testing.T) {
	a := 10
	b := 15
	want := 30
	got := utils.LCM(a, b)

	assert.Equal(t, want, got)
}

func TestLCM_0_20(t *testing.T) {
	a := 0
	b := 20
	want := 0
	got := utils.LCM(a, b)

	assert.Equal(t, want, got)
}
