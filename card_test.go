package DeckOfCards

import "fmt"

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
