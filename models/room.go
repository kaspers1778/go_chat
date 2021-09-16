package models

import (
	"fmt"
)

type Room struct {
	Name     string
	sessions map[Session]bool

	broadcast chan []byte
}

func NewRoom(Name string, CreatorSession Session) *Room {

	r := Room{
		Name:      Name,
		sessions:  make(map[Session]bool),
		broadcast: make(chan []byte),
	}
	r.AddSession(CreatorSession)
	return &r
}

func (r *Room) AddSession(session Session) {
	r.sessions[session] = true
	r.broadcast <- []byte(fmt.Sprintf("%v connected to %v", session.User.Name, r.Name))
}
