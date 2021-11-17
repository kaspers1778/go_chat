package main

import (
	"go_chat/server/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/v1/status", handlers.StatusHandler)
	http.HandleFunc("/v1/ws", handlers.WsHandler)
	log.Println("Server is on air")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
