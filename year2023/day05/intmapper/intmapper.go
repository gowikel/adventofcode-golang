// intconv defines IntConversor structure
// which can be used to map integers between two ranges
package intmapper

// IntConversor is a type that allows to map numbers
// between two or more ranges
type IntMapper struct {
	rs map[int][2]int
}

// Returns a new initialized conversor
func New() IntMapper {
	return IntMapper{
		// map[from] = [2]int{to, length}
		rs: map[int][2]int{},
	}
}

// Adds a new mapping. The behaviour is not defined
// for multiple overlapping ranges.
func (ic *IntMapper) AddMapping(from, to, length int) {
	ic.rs[from] = [2]int{to, length - 1}
}

// IsMapped returns true if the given number
// is inside a mapped rule
func (ic IntMapper) IsMapped(n int) bool {
	_, ok := ic.rs[n]

	if ok {
		return true
	}

	for start, rs := range ic.rs {
		length := rs[1]

		if n >= start && n <= (start+length) {
			return true
		}
	}

	return false
}

// Transform applies the mappings to the
// given integer n. If no mappings are found
// then n is returned
func (ic IntMapper) Transform(n int) int {
	r, ok := ic.rs[n]

	if ok {
		to := r[0]
		return to
	}

	for start, rs := range ic.rs {
		to := rs[0]
		length := rs[1]

		if n >= start && n <= (start+length) {
			return to + (n - start)
		}
	}

	return n
}
