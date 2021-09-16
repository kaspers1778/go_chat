package models

type Hub struct {
	Rooms map[string]*Room

	broadcast  chan []byte
	register   chan *Session
	unregister chan *Session
}

func NewHub() *Hub {
	h := Hub{
		Rooms:      make(map[string]*Room),
		broadcast:  make(chan []byte),
		register:   make(chan *Session),
		unregister: make(chan *Session),
	}
	return &h
}

func (h *Hub) JoinRoom(message Message, session Session) {
	if h.Rooms[message.Room] != nil {
		h.Rooms[message.Room].AddSession(session)
	} else {
		h.Rooms[message.Room] = NewRoom(message.Room, session)
	}
}
