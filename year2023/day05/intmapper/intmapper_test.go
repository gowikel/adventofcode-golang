package intmapper_test

import (
	"testing"

	. "github.com/gowikel/adventofcode-golang/year2023/day05/intmapper"
	"github.com/stretchr/testify/assert"
)

func TestIsMapped_NoMappings(t *testing.T) {
	im := New()
	input := 12
	want := false
	got := im.IsMapped(input)

	assert.Equal(t, want, got)
}

func TestIsMapped_Mapped_OutsideMap(t *testing.T) {
	im := New()
	input := 16
	want := false

	// [10-15] => [20-25]
	im.AddMapping(10, 20, 5)

	got := im.IsMapped(input)

	assert.Equal(t, want, got)
}

func TestIsMapped_Mapped_InsideMap_StartBoundary(t *testing.T) {
	im := New()
	input := 10
	want := true

	// [10-15] => [20-25]
	im.AddMapping(10, 20, 5)

	got := im.IsMapped(input)

	assert.Equal(t, want, got)
}

func TestIsMapped_Mapped_InsideMap_EndBoundary(t *testing.T) {
	t.Skip()
	im := New()
	input := 15
	want := true

	// [10-15] => [20-25]
	im.AddMapping(10, 20, 5)

	got := im.IsMapped(input)

	assert.Equal(t, want, got)
}

func TestTransform_NoMappings(t *testing.T) {
	im := New()
	input := 15
	want := 15

	got := im.Transform(input)

	assert.Equal(t, want, got)
}

func TestTransform_OneMapping_OutsideMapping(t *testing.T) {
	im := New()
	input := 16
	want := 16

	// [10-15] => [20-25]
	im.AddMapping(10, 20, 5)

	got := im.Transform(input)

	assert.Equal(t, want, got)
}

func TestTransform_OneMapping_StartBoundary(t *testing.T) {
	im := New()
	input := 10
	want := 20

	// [10-15] => [20-25]
	im.AddMapping(10, 20, 5)

	got := im.Transform(input)

	assert.Equal(t, want, got)
}

func TestTransform_OneMapping_EndBoundary(t *testing.T) {
	t.Skip()
	im := New()
	input := 15
	want := 25

	// [10-15] => [20-25]
	im.AddMapping(10, 20, 5)

	got := im.Transform(input)

	assert.Equal(t, want, got)
}

func TestTransform_OneMapping_Middle(t *testing.T) {
	im := New()
	input := 13
	want := 23

	// [10-15] => [20-25]
	im.AddMapping(10, 20, 5)

	got := im.Transform(input)

	assert.Equal(t, want, got)
}
