package internal

import (
	"context"
	"database/sql"
	"log"
)

type Hub struct {
	Rooms   map[string]*Room
	roomRep *RoomRepository
	userRep *UserRepository
}

func NewHub(db *sql.DB) *Hub {
	h := Hub{
		Rooms:   make(map[string]*Room),
		roomRep: NewRoomRepository(db),
		userRep: NewUserRepository(db),
	}
	return &h
}

func (h *Hub) RouteMessage(message *Message, session *Session) {
	switch message.Kind {
	case "Join":
		h.JoinRoom(message, session)
	case "Text":
		h.SendMessage(message)
	case "Leave":
		h.LeaveRoom(message, session)
	default:
		log.Println("Unknown message kind")
	}
}

func (h *Hub) JoinRoom(message *Message, session *Session) {
	if h.Rooms[message.Room] != nil {
		h.Rooms[message.Room].Register <- session
	} else {
		h.Rooms[message.Room] = NewRoom(message, session)
		if err := h.roomRep.Create(context.Background(), h.Rooms[message.Room]); err != nil {
			log.Println(err)
		}
	}
	if err := h.userRep.Create(context.Background(), message); err != nil {
		log.Println(err)
	}
	h.Rooms[message.Room].Broadcast <- message
}

func (h *Hub) SendMessage(message *Message) {
	h.Rooms[message.Room].Broadcast <- message
}

func (h *Hub) LeaveRoom(message *Message, session *Session) {
	h.Rooms[message.Room].Unregister <- session
	h.Rooms[message.Room].Broadcast <- message
	if err := h.userRep.Delete(context.Background(), message); err != nil {
		log.Println(err)
	}
}
