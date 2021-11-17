package main

import "go_chat/client/internal"

func main() {
	cli := internal.NewCLI("ws://127.0.0.1:8081/v1/ws")
	cli.Run()
}
