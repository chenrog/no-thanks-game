package main

import (
	"fmt"
)

func main() {
	game := NewGameBuilder().Build()

	game.Action(0)
	game.Action(0)
	game.Action(1)
	fmt.Println(game)
}
