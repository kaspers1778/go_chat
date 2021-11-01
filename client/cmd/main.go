package main

import (
	"fmt"
	"go_chat/models"
	"time"
)

func main() {

	me, _ := models.NewChatUser("Petro", "ws://127.0.0.1:8081/v1/ws")
	defer me.Stop()
	me.Write(models.Message{
		User: "Petro",
		Kind: "Join",
		Text: "has joined the room.",
		Room: "First",
	})
	time.Sleep(1 * time.Second)
	not_me, _ := models.NewChatUser("Vasyan", "ws://127.0.0.1:8081/v1/ws")
	defer not_me.Stop()
	not_me.Write(models.Message{
		User: "Vasya",
		Kind: "Join",
		Text: "has joined the room.",
		Room: "Second",
	})
	time.Sleep(1 * time.Second)
	not_me.Write(models.Message{
		User: "Vasya",
		Kind: "Text",
		Text: "Hey, Pete)",
		Room: "Second",
	})

	fmt.Scanln()

}
