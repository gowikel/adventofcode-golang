package algo

import mapset "github.com/deckarep/golang-set/v2"

func UnorderedEqualSlices[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	aSet := mapset.NewSet[T](a...)
	bSet := mapset.NewSet[T](b...)

	if aSet.Cardinality() != bSet.Cardinality() {
		return false
	}

	sm := aSet.SymmetricDifference(bSet)

	return sm.Cardinality() == 0
}
