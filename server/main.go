package main

import (
	"fmt"
)

func main() {
	game := NewGame(7)

	game.Action(0)
	game.Action(0)
	game.Action(0)
	game.Action(1)
	fmt.Println(game)
}
