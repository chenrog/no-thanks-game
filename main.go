package main

import "fmt"

type Player struct {
	Cards  []int
	Tokens int
}

func main() {
	var cards []int
	cards = append(cards, 1)

	fmt.Println(cards)
}
