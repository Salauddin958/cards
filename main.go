package main

import (
	"fmt"
	"sort"
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
	sortingSlices()
	copySliceElements()
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

func sortingSlices() {
	array := []int{3, 2, 5, 8, 6}
	sort.Ints(array)
	fmt.Println(array)

	fruits := []string{"peach", "kiwi", "banana", "pineapple"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}

type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}

func copySliceElements() {
	slc1 := []int{58, 69, 40, 45, 11, 56, 67, 21, 65}
	var slc2 []int
	slc3 := make([]int, 5)
	slc4 := []int{78, 50, 67, 77}

	copy_1 := copy(slc2, slc1)
	fmt.Println("\nSlice:", slc2)
	fmt.Println("Total number of elements copied:", copy_1)

	copy_2 := copy(slc3, slc1)
	fmt.Println("\nSlice:", slc3)
	fmt.Println("Total number of elements copied:", copy_2)

	copy_4 := copy(slc1, slc4)
	fmt.Println("\nSlice:", slc1)
	fmt.Println("Total number of elements copied:", copy_4)
}
