package main

import (
	"fmt"
	"strconv"
)

func main() {
	cards := newDeck()
	hand, remainingDeck := deal(cards, 5)
	hand.print()
	remainingDeck.print()
	cards.saveToFile("myCards")
	cards.shuffle()
	cards.print()
	assignment()
}

func assignment() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, value := range numbers {
		if value%2 == 0 {
			fmt.Println(strconv.Itoa(value) + " is even")
		} else {
			fmt.Println(strconv.Itoa(value) + " is odd")
		}
	}
}
