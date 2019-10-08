package main

import "fmt"

func main() {
	// card := newCard()
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of spades")
	// cards[1] = "Rahul"
	// fmt.Println(cards[0])

	// for i := 0; i < len(cards); i++ {
	// 	fmt.Println(cards[i])
	// }

	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "Five of Diamonds"
}
