package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Game struct {
	Deck           *Deck
	FloatingTokens int
	Players        []Player
	PlayerTurn     int
	seed           int64
}

func NewGame() *Game {
	return NewGameWithSeed(time.Now().UnixNano())
}

func NewGameWithSeed(seed int64) *Game {
	game := Game{
		FloatingTokens: 0,
		Players:        []Player{},
		PlayerTurn:     0,
		seed:           seed,
	}

	return &game
}

func (g *Game) Start() {
	r := rand.New(rand.NewSource(g.seed))
	g.Deck = NewDeck(g.seed)
	r.Shuffle(len(g.Players), func(i, j int) { g.Players[i], g.Players[j] = g.Players[j], g.Players[i] })
}

func (g *Game) AddPlayer() {
	g.Players = append(g.Players, *NewPlayer(fmt.Sprintf("player %d", len(g.Players)), 11))
}

func (g *Game) CurrentPlayer() *Player {
	return &g.Players[g.PlayerTurn]
}

func (g *Game) PlayerCount() int {
	return len(g.Players)
}

func (g *Game) Action(action Action) {
	if action == Pass && g.CurrentPlayer().GetTokens() > 0 {
		g.CurrentPlayer().RemoveToken()
		g.FloatingTokens += 1
		g.PlayerTurn += 1
		g.PlayerTurn %= g.PlayerCount()
	}

	if action == Take {
		g.CurrentPlayer().AddTokens(g.FloatingTokens)
		g.FloatingTokens = 0
		g.CurrentPlayer().AddCard(g.Deck.Pop())
	}
}

func (g *Game) IsOver() bool {
	return g.Deck.Empty()
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
	output += fmt.Sprintf("%d cards remain in the deck", g.Deck.CardsLeft())
	output += fmt.Sprintf("\n%s's turn", g.CurrentPlayer().Name)
	if !g.Deck.Empty() {
		output += fmt.Sprintf("\n%d token(s) on %d", g.FloatingTokens, g.Deck.CurrentCard())
	}
	for i := range g.Players {
		output += fmt.Sprintf("\n%s", g.Players[i].String())
	}
	return output
}

// TODO: fix this up if we need a different string view
//func (g *Game) FullViewString() string {
//	output := ""
//	output += fmt.Sprintf("deck(%d): %d", len(g.Deck.cards), g.Deck)
//	output += fmt.Sprintf("\n%s's turn", g.CurrentPlayer().Name)
//	if len(g.Deck.cards) > 0 {
//		output += fmt.Sprintf("\n%d token(s) on %d", g.FloatingTokens, g.Deck.cards[0])
//	}
//	for i := range g.Players {
//		output += fmt.Sprintf("\n%s", g.Players[i].String())
//	}
//	return output
//}
