package main

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	id   string
	conn *websocket.Conn
	send chan []byte
	room *Room
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (h *Hub) handler(w http.ResponseWriter, r *http.Request) (c *Client) {
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	roomID := r.URL.Query().Get("room")
	room := h.ccRoom(roomID, newClient)

	newClient := &Client{
		id:   uuid.NewString(),
		conn: wsConn,
		send: make(chan []byte),
		room: room,
	}

	return newClient

}
