package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/olahol/melody"
	"log"
	"net/http"
)

var game *Game
var debug = false

func initWebsocket() {
	melodyRouter := melody.New()

	// The default maximum message size is 512 bytes,
	// but this is not long enough to send game objects
	// Thus, we have to manually increase it
	//melodyRouter.Config.MaxMessageSize = 8192

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		err := melodyRouter.HandleRequest(w, r)
		if err != nil {
			return
		}
	})

	melodyRouter.HandleConnect(func(session *melody.Session) {
		sessions, _ := melodyRouter.Sessions()

		for _, s := range sessions {
			_, exists := s.Get("info")

			if !exists {
				continue
			}

			err := session.Write([]byte("you are an existing session"))
			if err != nil {
				return
			}
		}

		id := uuid.NewString()
		session.Set("info", id)

		err := session.Write([]byte("iam " + id))
		if err != nil {
			return
		}
	})
	//melodyRouter.HandleDisconnect()

	melodyRouter.HandleMessage(func(s *melody.Session, msg []byte) {
		if string(msg) == "pass" {
			game.Action(Pass)
		} else if string(msg) == "take" {
			game.Action(Take)
		}

		if debug {
			log.Println(string(msg))
		}

		outMsg := fmt.Sprint(game)
		err := melodyRouter.Broadcast([]byte(outMsg))
		if err != nil {
			return
		}
	})
}

func main() {
	game = NewGameBuilder().Build()

	initWebsocket()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
