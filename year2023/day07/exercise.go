package day07

import (
	"slices"

	"github.com/rs/zerolog/log"
)

type Exercise struct{}

func (e Exercise) Part1(data string) int {
	var result int
	ranks, err := Parse(
		data,
		BasicHandDeterminer,
		BasicCardStrengthDeterminer,
	)
	if err != nil {
		log.Fatal().Err(err).Msg("Error while parsing the file")
	}

	slices.SortFunc[[]Rank](ranks, SortRanksByStrength)

	for i, rank := range ranks {
		result += (len(ranks) - i) * rank.Bid
	}

	return result
}

func (e Exercise) Part2(data string) int {
	var result int
	ranks, err := Parse(
		data,
		WildcardHandDeterminer,
		WildcardCardStrengthDeterminer,
	)
	if err != nil {
		log.Fatal().Err(err).Msg("Error while parsing the file")
	}

	slices.SortFunc[[]Rank](ranks, SortRanksByStrength)

	for i, rank := range ranks {
		result += (len(ranks) - i) * rank.Bid
	}

	return result
}

var basicCardStrengths = map[rune]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'J': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
}

// Determines the card strength based on the order
// of the given card, AKQJT98765432.
func BasicCardStrengthDeterminer(r rune) int {
	return basicCardStrengths[r]
}

// Determines the strength of the card exactly like in
// BasicCardStrengthDeterminer
// but moving 'J' to the bottom
func WildcardCardStrengthDeterminer(r rune) int {
	strength := basicCardStrengths[r]

	switch r {
	case 'T', '9', '8', '7', '6', '5', '4', '3', '2':
		strength += 1
	case 'J':
		strength = 1
	}

	return strength
}

// - Five of a kind, where all five cards have the same label: AAAAA
// - Four of a kind, where four cards have the same label and one card
// has a different label: AA8AA
// - Full house, where three cards have the same label, and the
// remaining two cards share a different label: 23332
// - Three of a kind, where three cards have the same label, and
// the remaining two cards are each different from any other card in
// the hand: TTT98
// - Two pair, where two cards share one label, two other cards share
// a second label, and the remaining card has a third label: 23432
// - One pair, where two cards share one label, and the other three
// cards have a different label from the pair and each other: A23A4
// - High card, where all cards' labels are distinct: 23456

// Calculates which HandType should be used of a given set of cards.
// It does not return any error, as it does not check that the cards
// are valid.
func BasicHandDeterminer(cards [5]Card) HandType {
	counter := map[rune]int{}

	for _, card := range cards {
		counter[card.value] += 1
	}

	// All cards are the same
	if len(counter) == 1 {
		return FiveOfKind
	}

	groups := make([]int, 0, len(counter))

	for _, group := range counter {
		groups = append(groups, group)
	}

	slices.Sort[[]int](groups)
	slices.Reverse[[]int](groups)

	if groups[0] == 4 {
		return FourOfKind
	} else if groups[0] == 3 && groups[1] == 2 {
		return FullHouse
	} else if groups[0] == 3 {
		return ThreeOfKind
	} else if groups[0] == 2 && groups[1] == 2 {
		return TwoPair
	} else if groups[0] == 2 {
		return OnePair
	}

	return HighCard
}

// The function takes into account that 'J' are comodin
// cards, and therefore, can count towards other types
func WildcardHandDeterminer(cards [5]Card) HandType {
	counter := map[rune]int{}

	for _, card := range cards {
		counter[card.value] += 1
	}

	// All cards are equal
	if len(counter) == 1 {
		return FiveOfKind
	}

	// Fetch comodin cards first, and pop them from the groups
	wildcards := counter['J']
	delete(counter, 'J')

	groups := make([]int, 0, len(counter))

	for _, group := range counter {
		groups = append(groups, group)
	}

	slices.Sort[[]int](groups)
	slices.Reverse[[]int](groups)

	// May happen, 'JJJJA'
	if groups[0]+wildcards == 5 {
		return FiveOfKind
	} else if groups[0]+wildcards == 4 {
		return FourOfKind
	} else if (groups[0]+wildcards == 3 && groups[1] == 2) ||
		(groups[0] == 3 && groups[1]+wildcards == 2) {

		return FullHouse
	} else if groups[0]+wildcards == 3 {
		return ThreeOfKind
	} else if (groups[0]+wildcards == 2 && groups[1] == 2) ||
		(groups[0] == 2 && groups[1]+wildcards == 2) {

		return TwoPair
	} else if groups[0]+wildcards == 2 {
		return OnePair
	}

	return HighCard
}
