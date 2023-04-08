package main

import "fmt"

type Player struct {
	Name   string
	cards  []int
	tokens int
}

func NewPlayer(name string, tokens int) *Player {
	return &Player{Name: name, tokens: tokens}
}

func (p *Player) AddCard(card int) {
	p.cards = append(p.cards, card)
}

func (p *Player) RemoveToken() {
	p.tokens -= 1
}

func (p *Player) AddTokens(tokens int) {
	p.tokens += tokens
}

func (p *Player) String() string {
	return fmt.Sprintf("%s | cards: %d, tokens: %d", p.Name, p.cards, p.tokens)
}
