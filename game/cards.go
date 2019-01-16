package game

import (
	"math/rand"
	"time"
)

// Suits ..
type Suits string

const (
	clubs    = "♠"
	diamonds = "♥"
	hearts   = "♣"
	spades   = "♦"
)

// Ranks ..
type Ranks string

const (
	ace   = "A"
	two   = "2"
	three = "3"
	four  = "4"
	five  = "5"
	six   = "6"
	seven = "7"
	eight = "8"
	nine  = "9"
	ten   = "10"
	jack  = "J"
	queen = "Q"
	king  = "K"
)

// Card ..
type Card struct {
	Suit Suits
	Rank Ranks
}

// Deck ..
type Deck []Card

// NewDeck ..
func NewDeck() Deck {
	suits := []Suits{clubs, diamonds, hearts, spades}
	ranks := []Ranks{ace, two, three, four, five, six, seven, eight, nine, ten, jack, queen, king}
	deck := Deck{}
	i := 0
	for _, suit := range suits {
		for _, rank := range ranks {
			deck = append(deck, Card{suit, rank})
			i++
		}
	}
	return deck
}

// StandardShuffle ..
func StandardShuffle(deck Deck) {
	s := rand.NewSource(time.Now().UnixNano())
	deckSize := len(deck)
	cuttingTimes := 4 + rand.New(s).Intn(11)
	for i := 0; i < cuttingTimes; i++ {
		cutIndex := rand.New(s).Intn(deckSize)
		splitIndex := rand.New(s).Intn(cutIndex)
		subDeck := append(deck[cutIndex:], deck[splitIndex:cutIndex]...)
		deck = append(deck[:splitIndex], subDeck...)
	}
}

// CutShuffle ..
func CutShuffle(deck Deck) {
	s := rand.NewSource(time.Now().UnixNano())
	deckSize := len(deck)
	cuttingTimes := 1 + rand.New(s).Intn(3)
	for i := 0; i < cuttingTimes; i++ {
		cutIndex := rand.New(s).Intn(deckSize)
		leftDeck := deck[:cutIndex]
		rightDeck := deck[cutIndex:]
		deck = Deck{}
		leftDeckSize := len(leftDeck)
		rightDeckSize := len(rightDeck)
		lh := 0
		rh := 0
		// TODO
		for lh < leftDeckSize && rh < rightDeckSize {
			ll := lh
			lh += rand.New(s).Intn(4)
			if lh > leftDeckSize {
				lh = leftDeckSize
			}
			deck = append(deck, leftDeck[ll:lh]...)
			rl := rh
			rh := rand.New(s).Intn(4)
			if rh > rightDeckSize {
				rh = rightDeckSize
			}
			deck = append(deck, rightDeck[rl:rh]...)
		}
		deck = append(deck, append(leftDeck[lh:leftDeckSize], rightDeck[rh:rightDeckSize]...)...)
	}
}

// ShuffleCards ..
func ShuffleCards(deck Deck) {
	StandardShuffle(deck)
	CutShuffle(deck)
}
