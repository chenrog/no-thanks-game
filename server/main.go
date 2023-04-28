package main

import (
	"fmt"
)

func main() {
	game := NewGameBuilder().Build()

	game.Action(Pass)
	game.Action(Pass)
	game.Action(Pass)
	game.Action(Pass)
	game.Action(Take)
	fmt.Println(game)
}
