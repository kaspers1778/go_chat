package main

import (
	"go_chat/server"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/v1/status", server.StatusHandler)
	http.HandleFunc("/v1/ws", server.WsHandler)
	log.Println("Server is on air")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
