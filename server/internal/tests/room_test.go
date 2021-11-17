package tests

import (
	"go_chat/server/internal"
	"testing"
	"time"
)

func TestRoom(t *testing.T) {
	createMessage := internal.Message{
		User: "Creator",
		Kind: "Join",
		Text: "New user",
		Room: "Room1",
	}
	session := internal.Session{
		Ws:  nil,
		Hub: nil,
	}
	room := internal.NewRoom(&createMessage, &session)
	t.Run("room creates with right name", func(t *testing.T) {
		if room.Name != createMessage.Room {
			t.Fail()
		}
	})
	t.Run("room add session on join", func(t *testing.T) {
		if len(room.Sessions) != 1 {
			t.Fail()
		}
	})
	t.Run("room deactivate session on quit", func(t *testing.T) {
		room.Unregister <- &session
		time.Sleep(time.Second)
		if room.Sessions[session] != false {
			t.Fail()
		}
	})
}
