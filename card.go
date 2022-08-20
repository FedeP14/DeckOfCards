//go:generate stringer -type=Suit,Rank
//Command stringer --> utils to print the name of cards ex:"Ace of Hearth"

package DeckOfCards

import "fmt"

type Suit uint8

const (
	Spade Suit = iota // Spade equal 0, Diamond 1, Club 2, Heart 3
	Diamond
	Club
	Heart
	Joker
)

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
