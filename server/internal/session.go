package internal

import (
	"github.com/gorilla/websocket"
	"log"
)

type Session struct {
	//	User *User
	Ws  *websocket.Conn
	Hub *Hub
}

func (s *Session) Run() {
	message := &Message{}
	var err error
	for {
		err = s.Ws.ReadJSON(message)
		log.Printf("(%v)%v:%v", message.Room, message.User, message.Text)
		if err != nil {
			log.Println(err)
			return
		}
		s.Hub.RouteMessage(message, s)
		message = &Message{}
	}
}
