package models

import "log"

type Room struct {
	Name     string
	sessions map[Session]bool

	broadcast  chan *Message
	register   chan *Session
	unregister chan *Session
}

func NewRoom(m *Message, CreatorSession *Session) *Room {
	r := Room{
		Name:       m.Room,
		sessions:   make(map[Session]bool),
		broadcast:  make(chan *Message),
		register:   make(chan *Session),
		unregister: make(chan *Session),
	}
	r.sessions[*CreatorSession] = true
	log.Printf("%v created %v room.", m.User, r.Name)
	go r.Start()
	return &r
}

func (r *Room) Start() {
	for {
		select {
		case session := <-r.register:
			r.sessions[*session] = true
		case session := <-r.unregister:
			r.sessions[*session] = false
		case message := <-r.broadcast:
			for session, _ := range r.sessions {
				session.Ws.WriteJSON(message)
			}
		}
	}

}
