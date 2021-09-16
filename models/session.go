package models

import "github.com/gorilla/websocket"

type Session struct {
	User *User
	Ws   *websocket.Conn
	Hub  *Hub
}
