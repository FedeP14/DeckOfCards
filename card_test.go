package DeckOfCards

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Suit: Spade, Rank: Ace})
	fmt.Println(Card{Suit: Diamond, Rank: Two})
	fmt.Println(Card{Suit: Club, Rank: Three})
	fmt.Println(Card{Suit: Heart, Rank: Four})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Spades
	// Two of Diamonds
	// Three of Clubs
	// Four of Hearts
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	// 13 ranks, 4 suits
	if len(cards) != 13*4 {
		t.Error("Incorrect number of cards in a deck")
	}
	
	// TODO: CHECK IF ALL THE SUITS ARE IN THE DECK (FIRST 13 SHOULD BE OF A KIND AND SO GO ON)
}
