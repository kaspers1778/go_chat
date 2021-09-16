package server

import (
	"encoding/json"
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
	log.Println("Client connected")
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
		message := models.Message{}
		err = json.Unmarshal(p, &message)
		if err != nil {
			log.Println("Cannot unmarshal message")
		}

		newSession := models.Session{
			User: &models.User{Name: message.User},
			Ws:   con,
			Hub:  Hub,
		}
		Sessions[&newSession] = true

		switch message.Kind {
		case "Join":
			Hub.JoinRoom(message, newSession)
		default:
			log.Println("Unknown message kind")
		}

		if err := con.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}
