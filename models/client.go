package models

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"sync"
	"time"
)

type Client struct {
	Name    string
	wsconn  *websocket.Conn
	mutex   sync.Mutex
	sendBuf chan []byte
	link    string
}

func NewChatUser(name string, link string) (*Client, error) {
	newUser := Client{
		Name:    name,
		link:    link,
		sendBuf: make(chan []byte, 1),
	}
	go newUser.listen()
	go newUser.listenWrite()
	return &newUser, nil

}

func (cl *Client) Connect() *websocket.Conn {
	cl.mutex.Lock()
	defer cl.mutex.Unlock()
	if cl.wsconn != nil {
		return cl.wsconn
	}

	ticker := time.NewTicker(time.Second * 1)
	defer ticker.Stop()
	for ; ; <-ticker.C {
		ws, _, err := websocket.DefaultDialer.Dial(cl.link, nil)
		if err != nil {
			continue
		}
		cl.wsconn = ws
		log.Println("Client connected.")
		return cl.wsconn
	}
}

func (cl *Client) listen() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for range ticker.C {
		ws := cl.Connect()
		if ws == nil {
			continue
		}
		for {
			_, byteMsg, err := ws.ReadMessage()
			if err != nil {
				cl.Stop()
				break
			}
			fmt.Println(string(byteMsg))
		}
	}
}

func (cl *Client) Write(payload interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	cl.sendBuf <- data
	return nil

}

func (cl *Client) listenWrite() {
	for {
		ws := cl.Connect()
		if ws == nil {
			log.Println("No ws connection")
			continue
		}
		data := <-cl.sendBuf
		err := ws.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			log.Println("Write message error")
		}
		log.Printf("Sended: %s", data)
	}
}

func (cl *Client) Stop() {
	if cl.wsconn != nil {
		cl.wsconn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		cl.wsconn.Close()
		cl.wsconn = nil
	}
}
