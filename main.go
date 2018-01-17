package main

import(
    "fmt"
    "strconv"
    "golang-blackjack/pkg/card"
)

func main() {
    var deck []card.Card = card.BuildDeck()
    card.Shuffle(deck)

	deck, player_hand := card.DealHand(deck)
	deck, dealer_hand := card.DealHand(deck)
	
	fmt.Println("---------PLAYER---------")
	for x := 0; x < len(player_hand); x++ {
		fmt.Println(player_hand[x].Rank + " Of " + player_hand[x].Suit)
	}
	
	fmt.Println()
	
	fmt.Println("---------DEALER---------")
	for x := 0; x < len(dealer_hand); x++ {
		fmt.Println(dealer_hand[x].Rank + " Of " + dealer_hand[x].Suit)
	}
	
	fmt.Println()
    fmt.Println("Deck Size: " + strconv.Itoa(len(deck)))
}
