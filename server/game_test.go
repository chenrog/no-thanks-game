package main_test

import (
	. "github.com/chenrog/no-thanks-game"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Game", func() {
	var game *Game
	BeforeEach(func() {
		game = NewGameBuilder().SetSeed(1234).SetPlayerCount(3).Build()
	})

	Describe("Creating a game", func() {
		Context("with too few or too many players", func() {
			It("should panic", func() {
				Expect(func() { NewGameBuilder().SetPlayerCount(2).Build() }).To(Panic())
				Expect(func() { NewGameBuilder().SetPlayerCount(8).Build() }).To(Panic())
			})
		})

		Context("with default parameters", func() {
			It("should have 3 players and not panic", func() {
				Expect(func() { NewGameBuilder().Build() }).ToNot(Panic())
				game := NewGameBuilder().Build()
				Expect(len(game.Players)).To(Equal(3))
			})
		})
	})

	Describe("Playing a game", func() {
		Context("with a seed", func() {
			It("Should always have the same first 3 starting cards in the deck", func() {
				Expect(game.Deck[0:3]).To(Equal([]int{14, 16, 27}))
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
					Expect(game.FloatingTokens).To(Equal(1))
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
					Expect(game.FloatingTokens).To(Equal(0))
					Expect(game.Players[2].GetTokens()).To(Equal(tokensExpected))
				})

				It("Should add the card to the player's hand from the deck", func() {
					card := game.Deck[0]
					game.Action(Take)

					Expect(game.Deck[0]).ToNot(Equal(card))
					Expect(game.Players[0].GetCards()[0]).To(Equal(card))
				})
			})
		})
	})

	Describe("Finishing a game", func() {
		Context("with a seed", func() {
			It("Should always finish the same", func() {
				for len(game.Deck) > 0 {
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
