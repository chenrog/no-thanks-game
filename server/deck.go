package main

import "math/rand"

type Deck struct {
	cards []int
}

func NewDeck(seed int64) *Deck {
	var cards []int
	for i := 3; i <= 35; i++ {
		cards = append(cards, i)
	}

	r := rand.New(rand.NewSource(seed))
	r.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })

	cards = cards[9:]

	return &Deck{cards: cards}
}

func (d *Deck) TakeCurrentCard() int {
	card := d.CurrentCard()
	d.cards = d.cards[1:]
	return card
}

func (d *Deck) CurrentCard() int {
	return d.cards[0]
}

func (d *Deck) CardsLeft() int {
	return len(d.cards)
}

func (d *Deck) IsEmpty() bool {
	return d.CardsLeft() == 0
}
