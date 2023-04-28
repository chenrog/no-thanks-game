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

// actual game class past this point

type Game struct {
	Deck           []int
	FloatingTokens int
	Players        []Player
	PlayerTurn     int
	rand           *rand.Rand
}

func NewGame(seed int64, playerCount int) *Game {
	game := Game{
		Deck:           []int{},
		FloatingTokens: 0,
		Players:        []Player{},
		PlayerTurn:     0,
	}

	// DECK SETUP
	for i := 3; i <= 35; i++ {
		game.Deck = append(game.Deck, i)
	}

	game.rand = rand.New(rand.NewSource(seed))
	game.rand.Shuffle(len(game.Deck), func(i, j int) { game.Deck[i], game.Deck[j] = game.Deck[j], game.Deck[i] })

	game.Deck = game.Deck[9:]

	// PLAYER SETUP
	for i := 0; i < playerCount; i++ {
		game.Players = append(game.Players, *NewPlayer(fmt.Sprintf("player %d", i), 11))
	}

	game.rand.Shuffle(len(game.Players), func(i, j int) { game.Players[i], game.Players[j] = game.Players[j], game.Players[i] })

	return &game
}

func (g *Game) Action(action Action) {
	if action == Pass {
		g.Players[g.PlayerTurn].RemoveToken()
		g.FloatingTokens += 1
		g.PlayerTurn += 1
		g.PlayerTurn %= len(g.Players)
	}

	if action == Take {
		g.Players[g.PlayerTurn].AddTokens(g.FloatingTokens)
		g.FloatingTokens = 0
		g.Players[g.PlayerTurn].AddCard(g.Deck[0])
		g.Deck = g.Deck[1:]
	}
}

func (g *Game) IsOver() bool {
	return len(g.Deck) == 0
}

func (g *Game) GetWinners() []Player {
	minScore := 9999
	var winningPlayers []Player

	for p := range g.Players {
		if g.Players[p].GetScore() == minScore {
			winningPlayers = append(winningPlayers, g.Players[p])
		}
		if g.Players[p].GetScore() < minScore {
			winningPlayers = []Player{g.Players[p]}
			minScore = g.Players[p].GetScore()
		}
	}

	return winningPlayers
}

func (g *Game) String() string {
	output := ""
	output += fmt.Sprintf("deck(%d): %d", len(g.Deck), g.Deck)
	output += fmt.Sprintf("\n%s's turn", g.Players[g.PlayerTurn].Name)
	if len(g.Deck) > 0 {
		output += fmt.Sprintf("\n%d token(s) on %d", g.FloatingTokens, g.Deck[0])
	}
	for i := range g.Players {
		output += fmt.Sprintf("\n%s", g.Players[i].String())
	}
	return output
}
