package main

import(
    "fmt"
    "strconv"
    "golang-blackjack/pkg/card"
)

func main() {
    var deck []card.Card = card.BuildDeck()
    fmt.Println("Deck Size: " + strconv.Itoa(len(deck)))
}
