package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	Cards  []int
	Tokens int
}

func NewPlayer(tokens int) *Player {
	return &Player{Tokens: tokens}
}

func main() {
	var deck []int
	var players []Player

	// DECK SETUP
	for i := 3; i <= 35; i++ {
		deck = append(deck, i)
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })

	deck = deck[9:]

	// PLAYER SETUP
	for i := 0; i < 3; i++ {
		players = append(players, *NewPlayer(11))
	}

	fmt.Println("deck: ", deck)
	fmt.Println("deck count: ", len(deck))
	fmt.Println("players: ", players)
}
