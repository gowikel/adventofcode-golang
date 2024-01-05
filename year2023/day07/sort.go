package day07

// This function defines how to sort a list of ranks by its strength
func SortRanksByStrength(a, b Rank) int {
	// The handtype alone may define the order
	if a.HandType > b.HandType {
		return -1
	} else if b.HandType > a.HandType {
		return 1
	}

	// If two hands have the same hand type, then the
	// cards are compared one by one. The first one
	// that is higher wins
	for i := 0; i < len(a.cards); i++ {
		aCard := a.cards[i]
		bCard := b.cards[i]

		if aCard.Strength() > bCard.Strength() {
			return -1
		} else if bCard.Strength() > aCard.Strength() {
			return 1
		}
	}

	// All other checks have failed, so I try to
	// sort by bid
	if a.Bid > b.Bid {
		return -1
	} else if b.Bid > a.Bid {
		return 1
	}

	// All checks failed, a & b are considered equal.
	return 0
}
