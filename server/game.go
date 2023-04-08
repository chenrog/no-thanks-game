package main

import (
	"fmt"
	"math/rand"
	"time"
)

type GameBuilder struct {
	seed        int64
	playerCount int
}

func NewGameBuilder() *GameBuilder {
	return &GameBuilder{
		seed:        time.Now().UnixNano(),
		playerCount: 3,
	}
}

func (gb *GameBuilder) SetSeed(seed int64) *GameBuilder {
	gb.seed = seed
	return gb
}

func (gb *GameBuilder) SetPlayerCount(count int) *GameBuilder {
	if count < 3 || count > 7 {
		panic("player count invalid, must be 3-7")
	}

	gb.playerCount = count
	return gb
}

func (gb *GameBuilder) Build() *Game {
	return NewGame(gb.seed, gb.playerCount)
}

type Game struct {
	deck           []int
	floatingTokens int
	players        []Player
	playerTurn     int
}

func NewGame(seed int64, playerCount int) *Game {
	var deck []int
	var players []Player

	// DECK SETUP
	for i := 3; i <= 35; i++ {
		deck = append(deck, i)
	}

	rand.New(rand.NewSource(seed))
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })

	deck = deck[9:]

	// PLAYER SETUP
	for i := 0; i < playerCount; i++ {
		players = append(players, *NewPlayer(fmt.Sprintf("player %d", i), 11))
	}

	rand.Shuffle(len(players), func(i, j int) { players[i], players[j] = players[j], players[i] })

	return &Game{
		deck:           deck,
		floatingTokens: 0,
		players:        players,
		playerTurn:     0,
	}
}

func (g *Game) Action(action int) {
	// Pass
	if action == 0 {
		g.players[g.playerTurn].RemoveToken()
		g.floatingTokens += 1
		g.playerTurn += 1
		fmt.Println("here")
	}

	// Take
	if action == 1 {
		g.players[g.playerTurn].AddTokens(g.floatingTokens)
		g.floatingTokens = 0
		g.players[g.playerTurn].AddCard(g.deck[0])
		g.deck = g.deck[1:]
	}
}

func (g *Game) String() string {
	output := ""
	output += fmt.Sprintf("deck(%d): %d", len(g.deck), g.deck)
	output += fmt.Sprintf("\n%s's turn", g.players[g.playerTurn].Name)
	output += fmt.Sprintf("\n%d token(s) on %d", g.floatingTokens, g.deck[0])
	for i := range g.players {
		output += fmt.Sprintf("\n%s", g.players[i].String())
	}
	return output
}
