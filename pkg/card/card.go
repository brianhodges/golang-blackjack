package card

import (
	"time"
	"math/rand"
)

type Card struct {
    Suit string
    Rank string
}

func BuildDeck() []Card {
    suits := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
    ranks := []string{"Ace", "King", "Queen", "Jack", "10", "9", "8", "7", "6", "5", "4", "3", "2"}
    var deck []Card
    
    for x := 0; x < len(suits); x++ {
        for y := 0; y < len(ranks); y++ {
            card := Card{
                Suit: suits[x],
                Rank: ranks[y],
            }
            deck = append(deck, card)
        }
    }
    return deck
}

func Shuffle(deck []Card) {
	for x := 0; x < 10; x++ {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for len(deck) > 0 {
			n := len(deck)
			randIndex := r.Intn(n)
			deck[n-1], deck[randIndex] = deck[randIndex], deck[n-1]
			deck = deck[:n-1]
		}
	}
}

func DealHand(deck []Card) ([]Card, []Card) {
	var hand []Card
	for x := 0; x < 2; x++ {
		hand = append(hand, deck[x])
		deck[x] = deck[len(deck)-1] 
		deck = deck[:len(deck)-1]
	}
	return deck, hand
}
