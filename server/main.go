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
			value, exists := s.Get("info")

			if !exists {
				continue
			}

			info := value.(*Player)
			err := session.Write([]byte(fmt.Sprintf("welcome back player%s", info.Uuid)))
			if err != nil {
				return
			}
		}

		id := uuid.NewString()
		player := NewPlayer("")
		player.Uuid = id
		session.Set("info", player)
		game.AddPlayer()

		err := session.Write([]byte("waitingToJoin"))
		if err != nil {
			return
		}
	})

	melodyRouter.HandleDisconnect(func(s *melody.Session) {
		value, exists := s.Get("info")

		if !exists {
			return
		}

		info := value.(*Player)

		err := melodyRouter.BroadcastOthers([]byte(fmt.Sprintf("disconnecting player%s", info.Uuid)), s)
		if err != nil {
			return
		}
	})

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
	game = NewGame()
	game.AddPlayer()
	game.AddPlayer()
	game.AddPlayer()
	game.Start()

	initWebsocket()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
