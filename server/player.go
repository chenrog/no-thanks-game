package main

import (
	"fmt"
	"sort"
)

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

func (p *Player) GetCards() []int {
	return p.cards
}

func (p *Player) RemoveToken() {
	p.tokens -= 1
}

func (p *Player) AddTokens(tokens int) {
	p.tokens += tokens
}

func (p *Player) GetTokens() int {
	return p.tokens
}

func (p *Player) GetScore() int {
	sort.Ints(p.cards)
	cards := p.cards

	reverse := func(numbers []int) []int {
		var reversed []int
		for range numbers {
			reversed = append(reversed, 0)
		}

		for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
			reversed[i], reversed[j] = numbers[j], numbers[i]
		}

		return reversed
	}

	stackCards := func(cards []int) []int {
		var stackedCards []int
		for i := range cards {
			if i == (len(cards) - 1) {
				stackedCards = append(stackedCards, cards[i])
				break
			}

			if cards[i] != (cards[i+1] + 1) {
				stackedCards = append(stackedCards, cards[i])
				cards[i] = 0
			}
		}

		return stackedCards
	}

	sum := func(numbers []int) int {
		sum := 0
		for _, i := range numbers {
			sum += i
		}
		return sum
	}

	return sum(stackCards(reverse(cards))) - p.tokens
}

func (p *Player) String() string {
	return fmt.Sprintf("%s | cards: %d, tokens: %d", p.Name, p.cards, p.tokens)
}
