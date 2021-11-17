package internal

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"sync"
	"time"
)

type Client struct {
	Name    string
	Wsconn  *websocket.Conn
	mutex   sync.Mutex
	sendBuf chan Message
	getBuf  chan *Message
	link    string
}

func NewChatUser(name string, link string, getBuf *chan *Message) (*Client, error) {
	newUser := Client{
		Name:    name,
		link:    link,
		sendBuf: make(chan Message, 1),
		getBuf:  *getBuf,
	}
	go newUser.listen()
	go newUser.listenWrite()
	return &newUser, nil
}

func (cl *Client) Connect() *websocket.Conn {
	cl.mutex.Lock()
	defer cl.mutex.Unlock()
	if cl.Wsconn != nil {
		return cl.Wsconn
	}

	ticker := time.NewTicker(time.Second * 1)
	defer ticker.Stop()
	for ; ; <-ticker.C {
		ws, _, err := websocket.DefaultDialer.Dial(cl.link, nil)
		if err != nil {
			continue
		}
		cl.Wsconn = ws
		//log.Println("Client connected.")
		return cl.Wsconn
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
		message := &Message{}
		for {
			err := ws.ReadJSON(message)
			if err != nil {
				cl.Stop()
				break
			}
			cl.getBuf <- message
			fmt.Println(message)
		}
	}
}

func (cl *Client) Write(payload Message) error {
	cl.sendBuf <- payload
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
		err := ws.WriteJSON(data)
		if err != nil {
			log.Println("Write message error")
		}
		//log.Printf("Sended: %s", data)
	}
}

func (cl *Client) Stop() {
	if cl.Wsconn != nil {
		cl.Wsconn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		cl.Wsconn.Close()
		cl.Wsconn = nil
	}
}

func (cl *Client) WriteTextMessage(text string, room string) {
	if err := cl.Write(Message{
		User: cl.Name,
		Kind: "Text",
		Text: text,
		Room: room,
	}); err != nil {
		log.Print("Message has not been sent")
	}
}

func (cl *Client) WriteJoinMessage(room string) {
	if err := cl.Write(Message{
		User: cl.Name,
		Kind: "Join",
		Text: "New user",
		Room: room,
	}); err != nil {
		log.Print("Message has not been sent")
	}
}

func (cl *Client) WriteLeaveMessage(room string) {
	if err := cl.Write(Message{
		User: cl.Name,
		Kind: "Leave",
		Text: "is gone",
		Room: room,
	}); err != nil {
		log.Print("Message has not been sent")
	}
}
