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
	Deck           []int
	FloatingTokens int
	Players        []Player
	PlayerTurn     int
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
		Deck:           deck,
		FloatingTokens: 0,
		Players:        players,
		PlayerTurn:     0,
	}
}

func (g *Game) Action(action Action) {
	if action == Pass {
		g.Players[g.PlayerTurn].RemoveToken()
		g.FloatingTokens += 1
		g.PlayerTurn += 1
	}

	if action == Take {
		g.Players[g.PlayerTurn].AddTokens(g.FloatingTokens)
		g.FloatingTokens = 0
		g.Players[g.PlayerTurn].AddCard(g.Deck[0])
		g.Deck = g.Deck[1:]
	}
}

func (g *Game) String() string {
	output := ""
	output += fmt.Sprintf("deck(%d): %d", len(g.Deck), g.Deck)
	output += fmt.Sprintf("\n%s's turn", g.Players[g.PlayerTurn].Name)
	output += fmt.Sprintf("\n%d token(s) on %d", g.FloatingTokens, g.Deck[0])
	for i := range g.Players {
		output += fmt.Sprintf("\n%s", g.Players[i].String())
	}
	return output
}
