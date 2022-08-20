//go:generate stringer -type=Suit,Rank
//Command stringer --> utils to print the name of cards ex:"Ace of Hearth"

package DeckOfCards

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
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

// Shuffle Using permutation for shuffling the deck
// r.Perm(len(cards)) return a permutation of 52 element (in the base case)
// ex: Perm() returns [6, 4, 19, 52, 34, ..., 49]
// that means that the first card in cards Slice will be in 6th position, the second card will be in 4th position and so on
// NB: NOT the most efficient way to do it
// TODO: read https://www.calhoun.io/how-to-shuffle-arrays-and-slices-in-go/ and change this

func Shuffle(cards []Card) []Card {
	ret := make([]Card, len(cards))
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(cards))
	// i is the index of the cards Slice, j is the index of the permutation
	for i, j := range perm {
		ret[i] = cards[j]
	}
	return ret
}

// Jokers n is the number of Joker that will be put in the deck
// Jokers usually don't have a proper rank but in case some game uses it the rank will be from 0 to n
func Jokers(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{Suit: Joker, Rank: Rank(i)})
		}
		return cards
	}
}

func Filter(f func(card Card) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card
		for _, card := range cards {
			if !f(card) {
				ret = append(ret, card)
			}
		}
		return ret
	}
}
