package server

import (
	"fmt"
	"github.com/gorilla/websocket"
	"go_chat/models"
	"log"
	"net/http"
)

const StatusOK = 200

var Sessions = make(map[*models.Session]bool)

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
	log.Println("Client connected")
	newSession := models.Session{
		User: nil,
		Ws:   ws,
		Hub:  nil,
	}
	Sessions[&newSession] = true
	reader(ws)
}

func reader(con *websocket.Conn) {
	for {
		messageType, p, err := con.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(string(p))

		if err := con.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}
