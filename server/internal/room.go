package internal

import (
	"log"
)

type Room struct {
	Name     string
	Sessions map[Session]bool

	Broadcast  chan *Message
	Register   chan *Session
	Unregister chan *Session
}

func NewRoom(m *Message, CreatorSession *Session) *Room {
	r := Room{
		Name:       m.Room,
		Sessions:   make(map[Session]bool),
		Broadcast:  make(chan *Message),
		Register:   make(chan *Session),
		Unregister: make(chan *Session),
	}
	r.Sessions[*CreatorSession] = true
	log.Printf("%v created %v room.", m.User, r.Name)
	go r.Start()
	return &r
}

func (r *Room) Start() {
	for {
		select {
		case session := <-r.Register:
			r.Sessions[*session] = true
		case session := <-r.Unregister:
			r.Sessions[*session] = false
		case message := <-r.Broadcast:
			for session, status := range r.Sessions {
				if status {
					session.Ws.WriteJSON(message)
				}
			}
		}
	}

}
