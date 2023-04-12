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
})
