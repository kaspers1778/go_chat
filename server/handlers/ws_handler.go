package handlers

import (
	"github.com/gorilla/websocket"
	"go_chat/server"
	"log"
	"net/http"

	"go_chat/server/db"
	"go_chat/server/internal"
)

var Hub = internal.NewHub(db.NewDB(server.DB_CONNECTION_STRING))

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WsHandler(w http.ResponseWriter, req *http.Request) {
	ws, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
	newSession := internal.Session{
		Ws:  ws,
		Hub: Hub,
	}
	go newSession.Run()
}
