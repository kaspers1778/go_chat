package models

type Room struct {
	Name  string
	users map[Session]bool

	broadcast  chan []byte
	register   chan *Session
	unregister chan *Session
}
