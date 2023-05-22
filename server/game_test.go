package main_test

import (
	. "github.com/chenrog/no-thanks-game"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Game", func() {
	var game *Game
	BeforeEach(func() {
		game = NewGameWithSeed(1234)
		game.AddPlayer()
		game.AddPlayer()
		game.AddPlayer()
		game.Start()
	})

	Describe("Playing a game", func() {
		Context("with a seed", func() {
			It("Should always have the same first 3 starting cards in the deck", func() {
				Expect(game.Deck.TakeCurrentCard()).To(Equal(14))
				Expect(game.Deck.TakeCurrentCard()).To(Equal(16))
				Expect(game.Deck.TakeCurrentCard()).To(Equal(27))
			})
		})

		Context("Passing a turn", func() {
			It("Should move onto the next player", func() {
				turn := (game.PlayerTurn + 1) % len(game.Players)
				game.Action(Pass)
				Expect(game.PlayerTurn).To(Equal(turn))
			})

			It("Should increase the tokens on the card by 1 and take from the player who passed", func() {
				tokensExpected := game.Players[0].GetTokens() - 1
				game.Action(Pass)
				Expect(game.TokensOnCard).To(Equal(1))
				Expect(game.Players[0].GetTokens()).To(Equal(tokensExpected))
			})
		})

		Context("Taking a card", func() {
			It("Should not pass the turn", func() {
				turn := game.PlayerTurn
				game.Action(Take)
				Expect(game.PlayerTurn).To(Equal(turn))
			})

			It("Should take all tokens on the card", func() {
				tokensExpected := game.Players[0].GetTokens() + 2
				game.Action(Pass)
				game.Action(Pass)
				game.Action(Take)
				Expect(game.TokensOnCard).To(Equal(0))
				Expect(game.Players[2].GetTokens()).To(Equal(tokensExpected))
			})

			It("Should add the card to the player's hand from the deck", func() {
				card := game.Deck.CurrentCard()
				game.Action(Take)

				Expect(game.Deck.CurrentCard()).ToNot(Equal(card))
				Expect(game.Players[0].GetCards()[0]).To(Equal(card))
			})
		})
	})

	Describe("Finishing a game", func() {
		Context("with a seed", func() {
			It("Should always finish the same", func() {
				for !game.Deck.IsEmpty() {
					game.Action(Take)
				}

				Expect(game.IsOver()).To(Equal(true))
				Expect(game.Players[0].GetScore()).To(Equal(122))
				Expect(game.Players[1].GetScore()).To(Equal(-11))
				Expect(game.Players[2].GetScore()).To(Equal(-11))
				Expect(game.GetWinners()).To(Equal([]Player{game.Players[1], game.Players[2]}))
			})
		})
	})
})
