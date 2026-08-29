package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

type RoomEvent struct {
	roomID    string
	videoID   string
	timestamp float64
	event     string
	clientId  string
	conn      *websocket.Conn
}

type Room struct {
	roomID    string
	clientIDs map[string]*websocket.Conn
	videoID   string
	timestamp float64
	playing   bool
	mu        sync.Mutex
}

func (r *Room) setPlaying(playing bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.playing = playing
}

func (r *Room) getState() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.playing
}

func (r *Room) broadcastToAll(msg []byte, excludeClientID string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for id, client := range r.clientIDs {
		if id == excludeClientID {
			continue
		}

		err := client.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Printf("broadcast failed for user %s: %v", id, err)
		}

	}

}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func serveWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

}
