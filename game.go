package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Game struct {
	deck    []int
	players []Player
}

func NewGame(playerCount int) *Game {
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
	for i := 0; i < playerCount; i++ {
		players = append(players, *NewPlayer(11))
	}

	return &Game{
		deck:    deck,
		players: players,
	}
}

func (g Game) String() string {
	output := ""
	output += fmt.Sprintf("deck(%d): %d", len(g.deck), g.deck)
	for i := range g.players {
		output += fmt.Sprintf("\nplayer %d: %s", i, g.players[i].String())
	}
	return output
}
