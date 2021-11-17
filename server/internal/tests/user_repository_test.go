package tests

import (
	"context"
	"go_chat/server"
	"go_chat/server/db"
	"go_chat/server/internal"
	"testing"
)

func TestUserRepository(t *testing.T) {
	db := db.NewDB(server.DB_CONNECTION_STRING)
	userRep := internal.NewUserRepository(db)
	message := internal.Message{
		User: "TestUser",
		Kind: "Join",
		Text: "New User",
		Room: "TestRoom",
	}
	t.Run("Add new user", func(t *testing.T) {
		err := userRep.Create(context.Background(), &message)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
	})
	t.Run("Remove user", func(t *testing.T) {
		err := userRep.Delete(context.Background(), &message)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
	})

}
