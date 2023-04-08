package main

import "fmt"

type Player struct {
	cards  []int
	tokens int
}

func NewPlayer(tokens int) *Player {
	return &Player{tokens: tokens}
}

func (p Player) AddCard(card int) {
	p.cards = append(p.cards, card)
}

func (p Player) String() string {
	return fmt.Sprintf("cards: %d, tokens: %d", p.cards, p.tokens)
}
