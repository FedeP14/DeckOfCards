//go:generate stringer -type=Suit,Rank
//Command stringer --> utils to print the name of cards ex:"Ace of Hearth"

package DeckOfCards

import (
	"fmt"
	"sort"
)

type Suit uint8

const (
	Spade Suit = iota // Spade equal 0, Diamond 1, Club 2, Heart 3
	Diamond
	Club
	Heart
	Joker
)

var suits = [...]Suit{Spade, Diamond, Club, Heart} // needed for New() function cause Joker is a special Suit

type Rank uint8

const (
	_ Rank = iota // The value of every rank is the rank itself, so we skip 0 with _
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace // iterator for managing Ranks
	maxRank = King
)

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String() // return Joker
	}
	//return the name of the card
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func New(opts ...func([]Card) []Card) []Card { // ... means optional
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}

	// adding options
	for _, opt := range opts {
		cards = opt(cards)
	}

	return cards
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

func Sort(less func(cards []Card) func(i, j int) bool) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, Less(cards))
		return cards
	}
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRanks(cards[i]) < absRanks(cards[j])
	}
}

// Cards of Spade kind will have value between 14 and 26
// Cards of Diamond kind will have value between 27 and 39 and so on
// absoluteRanks are indispensable for sorting the deck
func absRanks(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}
