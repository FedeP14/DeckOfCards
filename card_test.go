package DeckOfCards

import (
	"fmt"
	"math/rand"
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

	// TODO: CHECK IF ALL THE SUITS ARE IN THE DECK (FIRST 13 SHOULD BE OF A KIND AND SO ON)
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("Expected: ", exp.String(), "Received: ", cards[0])
	}

}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("Expected: ", exp.String(), "Received: ", cards[0])
	}
}

func TestShuffle(t *testing.T) {
	// making the function deterministic
	shuffleRand = rand.New(rand.NewSource(0))
	// Perm with that seed [40 35 ...]
	origin := New()
	first := origin[40]
	second := origin[35]
	cards := New(Shuffle)

	if cards[0] != first {
		t.Errorf("Expected %s, Received: %s", first, cards[0])
	} else if cards[1] != second {
		t.Errorf("Expected %s, Received: %s", second, cards[1])
	}
}

// Possible adding test of no more than 4 jokers
func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	count := 0
	for _, card := range cards {
		if card.Suit == Joker {
			count++
		}
	}

	if count != 3 {
		t.Error("Expected 3 Jokers, Received: ", count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(filter))

	for _, card := range cards {
		if card.Rank == Two || card.Rank == Three {
			t.Error("Expected all Twos and Threes filtered out")
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	// 13 ranks, 4 suits, 3 decks
	if len(cards) != 13*4*3 {
		t.Errorf("Expected %d cards Received: %d ", 13*4*3, len(cards))
	}
}
