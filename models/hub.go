package models

import (
	"log"
)

type Hub struct {
	Rooms map[string]*Room
}

func NewHub() *Hub {
	h := Hub{
		Rooms: make(map[string]*Room),
	}
	return &h
}

func (h *Hub) RouteMessage(message *Message, session *Session) {
	switch message.Kind {
	case "Join":
		h.JoinRoom(message, session)
	case "Text":
		h.SendMessage(message)
	default:
		log.Println("Unknown message kind")
	}
}

func (h *Hub) JoinRoom(message *Message, session *Session) {
	if h.Rooms[message.Room] != nil {
		h.Rooms[message.Room].register <- session
	} else {
		h.Rooms[message.Room] = NewRoom(message, session)
	}
	h.Rooms[message.Room].broadcast <- message
}

func (h *Hub) SendMessage(message *Message) {
	h.Rooms[message.Room].broadcast <- message
}
