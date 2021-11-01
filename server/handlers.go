package server

import (
	"fmt"
	"github.com/gorilla/websocket"
	"go_chat/models"
	"log"
	"net/http"
)

const StatusOK = 200

var Hub = models.NewHub()

var Sessions = make(map[*models.Session]bool)
var Rooms = make(map[string]*models.Room)

func StatusHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "{\nStatus: %v\n}", StatusOK)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WsHandler(w http.ResponseWriter, req *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	ws, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println(err)
	}
	newSession := models.Session{
		Ws:  ws,
		Hub: Hub,
	}
	Sessions[&newSession] = true
	go newSession.Run()
}
