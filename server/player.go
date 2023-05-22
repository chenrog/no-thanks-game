package main

import (
	"fmt"
	"math"
	"sort"
)

type Player struct {
	Name   string
	cards  []int
	tokens int
	Uuid   string
}

func NewPlayer(name string, tokens int) *Player {
	return &Player{Name: name, tokens: tokens}
}

func (p *Player) TakeCard(card int, tokensOnCard int) {
	p.cards = append(p.cards, card)
	p.tokens += tokensOnCard
}

func (p *Player) GetCards() []int {
	return p.cards
}

func (p *Player) BetToken() {
	p.tokens -= 1
}

func (p *Player) GetTokens() int {
	return p.tokens
}

func (p *Player) GetScore() int {
	sort.Ints(p.cards)
	var scoredCardsSum int

	for i := range p.cards {
		areSequentialCards := func(c1 int, c2 int) bool {
			return math.Abs(float64(c1-c2)) == 1
		}

		if firstCard := i == 0; firstCard {
			scoredCardsSum += p.cards[i]
		} else if !areSequentialCards(p.cards[i], p.cards[i-1]) {
			scoredCardsSum += p.cards[i]
		}
	}

	return scoredCardsSum - p.tokens
}

func (p *Player) String() string {
	return fmt.Sprintf("%s | cards: %d, tokens: %d", p.Name, p.cards, p.tokens)
}
