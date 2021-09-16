package main

import (
	"fmt"
	"go_chat/models"
)

func main() {
	me, _ := models.NewChatUser("Petro", "ws://127.0.0.1:8081/v1/ws")
	defer me.Stop()
	me.Write(models.Message{
		User: "Petro",
		Kind: "Join",
		Text: "",
		Room: "First",
	})
	not_me, _ := models.NewChatUser("Vasyan", "ws://127.0.0.1:8081/v1/ws")
	defer not_me.Stop()
	not_me.Write(models.Message{
		User: "Vasya",
		Kind: "Join",
		Text: "",
		Room: "First",
	})

	fmt.Scanln()

}
