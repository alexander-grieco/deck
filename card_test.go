package deck

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Diamond})
	fmt.Println(Card{Rank: Three, Suit: Spade})
	fmt.Println(Card{Rank: Four, Suit: Club})
	fmt.Println(Card{Rank: Queen, Suit: Club})
	fmt.Println(Card{Rank: King, Suit: Heart})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Two of Diamonds
	// Three of Spades
	// Four of Clubs
	// Queen of Clubs
	// King of Hearts
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	// 13 * 4
	if len(cards) != 52 {
		t.Error("Wrong number of cards in a new deck.")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("Expected", exp, "as first card. Received:", cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := New(CustomSort(Less))
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Errorf("Expected %s as first card. Received %s", exp, cards[0])
	}
}

func TestJokers(t *testing.T) {
	numJokers := 3
	cards := New(Jokers(numJokers))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != numJokers {
		t.Errorf("Expected %d Jokers, but there are %d Jokers.", numJokers, count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(filter))
	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Error("Expected all 2s and 3s to be filtered out, but found a", c.Rank)
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	// 13 Ranks, 4 Suits, 3 Decks
	if len(cards) != 13*4*3 {
		t.Errorf("Expected %d cards, received %d cards", 13*4*3, len(cards))
	}
}

func TestShuffle(t *testing.T) {
	// make shuffleRand deterministic
	// First call to shuffleRand.Perm(52) should be:
	// [40, 35 ...]
	shuffleRand = rand.New(rand.NewSource(0))
	orig := New()
	first := orig[40]
	second := orig[35]
	cards := New(Shuffle)
	if cards[0] != first {
		t.Errorf("Expected the first card to be %s, received %s", first, cards[0])
	}
	if cards[1] != second {
		t.Errorf("Expected the second card to be%s, received %s", second, cards[1])
	}
}
