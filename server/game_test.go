package main_test

import (
	. "github.com/chenrog/no-thanks-game"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Game", func() {
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
		Context("with a seed of 1234", func() {
			It("Should always have the same first 3 starting cards in the deck", func() {
				gb := NewGameBuilder().SetSeed(1234).SetPlayerCount(3)
				game := gb.Build()
				Expect(game.Deck[0:3]).To(Equal([]int{14, 16, 27}))
			})
		})
	})

	Describe("Finishing a game", func() {
		Context("with a seed of 1234", func() {
			It("Should always finish the same", func() {
				game := NewGameBuilder().SetSeed(1234).SetPlayerCount(3).Build()
				for len(game.Deck) > 0 {
					game.Action(Take)
				}

				Expect(game.IsOver()).To(Equal(true))
				Expect(game.Players[0].GetScore()).To(Equal(133))
			})
		})
	})
})
