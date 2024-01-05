package day07

import (
	"errors"
	"fmt"
	"strings"
)

// Type representing a Card in the Camel Cards game
// the valid values are:
// A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2
type Card struct {
	value    rune
	strength int
}

var validCardSet = map[rune]struct{}{
	'A': {},
	'K': {},
	'Q': {},
	'J': {},
	'T': {},
	'9': {},
	'8': {},
	'7': {},
	'6': {},
	'5': {},
	'4': {},
	'3': {},
	'2': {},
}

// String representation of a Card
func (c Card) String() string {
	return string(c.value)
}

// GoString implementation
func (c Card) GoString() string {
	return fmt.Sprintf("%c(%d)", c.value, c.Strength())
}

// Strength of the card
func (c Card) Strength() int {
	return c.strength
}

// Card constructor. It validates the rune is valid. If not,
// it will return an error.
func NewCard(
	c rune,
	strengthDeterminer func(rune) int,
) (Card, error) {
	if _, ok := validCardSet[c]; !ok {
		return Card{}, fmt.Errorf("%q is an invalid Card", c)
	}

	return Card{
		value:    c,
		strength: strengthDeterminer(c),
	}, nil
}

// Represents the Hand type, based on the given cards
type HandType int

// Ordered from less to more value
const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfKind
	FullHouse
	FourOfKind
	FiveOfKind
)

// String representation of HandType
func (h HandType) String() string {
	names := []string{
		"High Card",
		"One Pair",
		"Two Pair",
		"Three of Kind",
		"Full House",
		"Four of Kind",
		"Five of Kind",
	}

	if int(h) < len(names) {
		return names[int(h)]
	}

	return fmt.Sprintf("HandType(%d)", int(h))
}

// Represents a hand in the Camel Cards game
type Hand struct {
	cards    [5]Card
	HandType HandType
}

func (h Hand) String() string {
	return fmt.Sprintf("%s - %v", h.HandType, h.cards)
}

// Hand constructor from a string, like "32T3K"
// If something goes wrong while parsing the Cards
// the error is returned
func NewHand(
	cards string,
	handTypeDeterminer func([5]Card) HandType,
	cardStrengthDeterminer func(rune) int,
) (Hand, error) {
	result := Hand{}

	if len(cards) != 5 {
		return result, fmt.Errorf(
			"a hand needs 5 cards, but %d were passed",
			len(cards),
		)
	}

	cards = strings.ToUpper(cards)

	for i, cardValue := range cards {
		card, err := NewCard(cardValue, cardStrengthDeterminer)

		if err != nil {
			return result, errors.Join(
				fmt.Errorf("unable to process card %d", i+1),
				err,
			)
		}

		result.cards[i] = card
	}

	result.HandType = handTypeDeterminer(result.cards)

	return result, nil
}

// Represents a Rank, aka, a hand in the Camel Cards
// with its corresponding bid.
type Rank struct {
	Hand
	Bid int
}

// Generates a new rank, with score equal to zero
func NewRank(hand Hand, bid int) Rank {
	return Rank{
		Hand: hand,
		Bid:  bid,
	}
}
