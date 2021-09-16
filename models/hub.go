package models

type Hub struct {
	Rooms map[*Room]bool

	broadcast  chan []byte
	register   chan *Session
	unregister chan *Session
}
