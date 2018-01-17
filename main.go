package main

import(
    "fmt"
    "strconv"
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

func main() {
    var deck []Card = BuildDeck()
    fmt.Println("Deck Size: " + strconv.Itoa(len(deck)))
}