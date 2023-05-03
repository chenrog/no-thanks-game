package main

import (
	"fmt"
	"github.com/olahol/melody"
	"log"
	"net/http"
)

func main() {
	game := NewGameBuilder().Build()

	game.Action(Pass)
	game.Action(Pass)
	game.Action(Pass)
	game.Action(Pass)
	game.Action(Take)
	fmt.Println(game)

	m := melody.New()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		err := m.HandleRequest(w, r)
		if err != nil {
			return
		}
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		log.Println(string(msg))
		err := m.Broadcast(msg)
		if err != nil {
			return
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
